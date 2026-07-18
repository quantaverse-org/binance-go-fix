package binance_go_fix

import (
	"binance-go-fix/message"
	"bufio"
	"context"
	"crypto/ed25519"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	targetCompId = "SPOT"
	orderHost    = "fix-oe.binance.com"
	marketHost   = "fix-md.binance.com"
	// updating channel 需要吸收同步请求等待期间穿插到达的主动推送，避免阻塞 FIX 读循环。
	updatingChannelCapacity = 1024
)

var (
	ErrConnectionBroken = errors.New("connection is broken, please try again later")
	ErrResponseTimeout  = errors.New("response timeout")
	LogonFailed         = errors.New("logon failed")
)

// responseResult 将正常响应和 Reject 统一送入同一个请求等待通道。
type responseResult struct {
	response message.Response
	err      error
}

// responseWaiter 表示一个正在等待响应的请求。
// 同一个请求可能通过多个业务 ID 关联响应，也可能需要收集多条响应消息。
type responseWaiter struct {
	ch        chan responseResult // 接收正常响应或 Reject。
	ids       []string            // 注册到 respChannels 的全部业务 ID，用于统一清理别名。
	remaining int                 // 尚未收到的正常响应数量；归零后不再拦截后续主动推送。
}

// Updating 暴露服务器主动推送消息的只读通道。
// 请求对应的同步响应会优先进入 responseWaiter，未匹配的消息才会进入这些通道。
type Updating struct {
	// MarketData 中只会出现 *MarketDataSnapshot 和 *MarketDataIncrementalRefresh。
	MarketData      <-chan message.Response
	OrderExecution  <-chan *message.ExecutionReport
	OrderListStatus <-chan *message.ListStatus
}

// updatingSenders 只由 Client 的消息分发协程持有，防止调用方误写更新通道。
type updatingSenders struct {
	marketData      chan<- message.Response
	orderExecution  chan<- *message.ExecutionReport
	orderListStatus chan<- *message.ListStatus
}

// initUpdating 创建内部发送端和对外只读端，共享同一组带缓冲 channel。
func initUpdating() (*updatingSenders, *Updating) {
	marketData := make(chan message.Response, updatingChannelCapacity)
	orderExecution := make(chan *message.ExecutionReport, updatingChannelCapacity)
	orderListStatus := make(chan *message.ListStatus, updatingChannelCapacity)

	return &updatingSenders{
			marketData:      marketData,
			orderExecution:  orderExecution,
			orderListStatus: orderListStatus,
		}, &Updating{
			MarketData:      marketData,
			OrderExecution:  orderExecution,
			OrderListStatus: orderListStatus,
		}
}

// sendUpdating 在客户端关闭时放弃发送，否则允许缓冲区提供有限背压。
func sendUpdating[T any](ctx context.Context, ch chan<- T, value T) {
	select {
	case <-ctx.Done():
	case ch <- value:
	}
}

type ApiKey struct {
	UserName   string
	PrivateKey ed25519.PrivateKey
}

type ClientConfig struct {
	EnableNotify      bool
	ClientName        string
	HeartbeatInterval time.Duration
	ReconnectInterval time.Duration
	ResponseTimeout   time.Duration
	WriteTimeout      time.Duration
	ResponseMode      message.ResponseMode
	ApiKey            *ApiKey
}

func NewClientConfig(apiKey *ApiKey) *ClientConfig {
	return &ClientConfig{
		EnableNotify:      false,
		ClientName:        "QUANTAVERSE",
		HeartbeatInterval: time.Second * 30,
		ReconnectInterval: time.Second * 1,
		ResponseTimeout:   time.Second * 10,
		WriteTimeout:      time.Second * 1,
		ResponseMode:      message.ResponseModeEverything,
		ApiKey:            apiKey,
	}
}

func (c *ClientConfig) WithEnableNotify() *ClientConfig {
	c.EnableNotify = true
	return c
}

func (c *ClientConfig) WithClientName(name string) *ClientConfig {
	c.ClientName = name
	return c
}

func (c *ClientConfig) WithHeartbeatInterval(interval time.Duration) *ClientConfig {
	c.HeartbeatInterval = interval
	return c
}

func (c *ClientConfig) WithReconnectInterval(interval time.Duration) *ClientConfig {
	c.ReconnectInterval = interval
	return c
}

func (c *ClientConfig) WithResponseTimeout(timeout time.Duration) *ClientConfig {
	c.ResponseTimeout = timeout
	return c
}

func (c *ClientConfig) WithWriteTimeout(timeout time.Duration) *ClientConfig {
	c.WriteTimeout = timeout
	return c
}

func (c *ClientConfig) WithResponseMode(responseMode message.ResponseMode) *ClientConfig {
	c.ResponseMode = responseMode
	return c
}

type MarketClient struct {
	*Client
}

// NewMarketClient 建立市场数据会话；仅在 EnableNotify 开启时创建 Updating。
func NewMarketClient(ctx context.Context, config *ClientConfig) (*MarketClient, *Updating, error) {
	var senders *updatingSenders
	var updating *Updating
	if config.EnableNotify {
		senders, updating = initUpdating()
	}

	client, err := newClient(ctx, marketHost, config, senders)
	if err != nil {
		return nil, nil, err
	}
	return &MarketClient{Client: client}, updating, nil
}

func (c *MarketClient) InstrumentList(req *message.InstrumentListRequest) (*message.InstrumentList, error) {
	resp, err := c.requestAndWait(req, req.InstrumentReqID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.InstrumentList)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeInstrumentList)
	}
	return result, nil
}

func (c *MarketClient) MarketData(req *message.MarketDataRequest) ([]*message.MarketDataSnapshot, error) {
	// 取消订阅不返回 Snapshot，并且必须先从断线重订阅列表中移除。
	if req.SubscriptionRequestType == message.SubscriptionRequestTypeUnsubscribe {
		c.removeResubRequest(req.MDReqID)
		return nil, c.request(req, false)
	}

	// 每个 Symbol 会返回一条初始 Snapshot，全部收齐前保留 response waiter。
	responses, err := c.requestAndWaitMany(req, len(req.Symbols), req.MDReqID)
	if err != nil {
		c.removeResubRequest(req.MDReqID)
		return nil, err
	}
	// 初始订阅成功后保存请求副本，重连完成时用于恢复订阅。
	c.setResubRequest(req.MDReqID, cloneMarketDataRequest(req))

	// 统一的 response channel 使用接口类型，这里校验并转换成具体 Snapshot slice。
	result := make([]*message.MarketDataSnapshot, 0, len(responses))
	for _, resp := range responses {
		snapshot, ok := resp.(*message.MarketDataSnapshot)
		if !ok {
			return nil, unexpectedResponseError(resp, message.MsgTypeMarketDataSnapshot)
		}
		result = append(result, snapshot)
	}
	return result, nil
}

type OrderClient struct {
	*Client
}

// NewOrderClient 建立订单会话；Updating 用于接收账户级 ExecutionReport 和 ListStatus 推送。
func NewOrderClient(ctx context.Context, config *ClientConfig) (*OrderClient, *Updating, error) {
	var senders *updatingSenders
	var updating *Updating
	if config.EnableNotify {
		senders, updating = initUpdating()
	}

	client, err := newClient(ctx, orderHost, config, senders)
	if err != nil {
		return nil, nil, err
	}
	return &OrderClient{Client: client}, updating, nil
}

func (c *OrderClient) NewOrderSingle(req *message.NewOrderSingle) (*message.ExecutionReport, error) {
	// 首条匹配 ClOrdID 的 ExecutionReport 是同步 ACK，后续状态变化进入 Updating。
	resp, err := c.requestAndWait(req, req.ClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) OrderCancel(req *message.OrderCancelRequest) (*message.ExecutionReport, error) {
	resp, err := c.requestAndWait(
		req,
		req.ClOrdID,
	)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) OrderMassCancel(req *message.OrderMassCancelRequest) (*message.OrderMassCancelReport, error) {
	// 这里只等待汇总报告；每个被取消订单的 ExecutionReport 属于账户级更新。
	resp, err := c.requestAndWait(req, req.ClOrdID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.OrderMassCancelReport)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeOrderMassCancelReport)
	}
	return result, nil
}

func (c *OrderClient) OrderReplace(req *message.OrderCancelRequestAndNewOrderSingle) (*message.ExecutionReport, error) {
	// 新订单 ID 和被取消订单 ID 都可能出现在 ACK 中，因此注册为同一 waiter 的别名。
	resp, err := c.requestAndWait(req, req.ClOrdID, req.CancelClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) ListStatus(req *message.NewOrderList) (*message.ListStatus, error) {
	// ListStatus 是该请求的同步结果，各子订单的 ExecutionReport 通过 Updating 接收。
	resp, err := c.requestAndWait(req, req.ClListID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.ListStatus)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeListStatus)
	}
	return result, nil
}

func (c *OrderClient) OrderAmendKeepPriority(req *message.OrderAmendKeepPriorityRequest) (*message.ExecutionReport, error) {
	// 修改结果由 ExecutionReport 确认；若订单属于列表，额外 ListStatus 会进入 Updating。
	resp, err := c.requestAndWait(
		req,
		req.ClOrdID,
	)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) Limit(req *message.LimitQuery) (*message.LimitResponse, error) {
	resp, err := c.requestAndWait(req, req.ReqID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.LimitResponse)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeLimitResponse)
	}
	return result, nil
}

type Client struct {
	ctx    context.Context
	host   string
	config *ClientConfig

	// connLk 串行化写入、重连和序列号递增，reader 只由消息处理协程使用。
	connLk    sync.Mutex
	conn      *tls.Conn
	reader    *bufio.Reader
	id        uint32
	closed    chan struct{}
	hbChannel chan *message.TestRequest

	// respChannels 用业务 ID 将应用层响应路由给正在等待的请求。
	respChanLk   sync.Mutex
	respChannels map[string]*responseWaiter

	// rejectChannels 用请求的 MsgSeqNum 将会话级 Reject 路由回原请求。
	rejectChanLk   sync.Mutex
	rejectChannels map[uint32]chan responseResult

	// resubReqs 保存成功建立的市场数据订阅，用于断线重连后恢复。
	resubReqLk sync.Mutex
	resubReqs  map[string]message.Request

	// updating 为 nil 时忽略所有未匹配请求的主动推送。
	updating *updatingSenders
}

// newClient 完成 TLS 连接、FIX Logon，并启动消息和心跳两个后台协程。
func newClient(ctx context.Context, host string, config *ClientConfig, updating *updatingSenders) (*Client, error) {
	// ServerName 用于 TLS SNI 和证书主机名校验。
	conn, err := tls.Dial("tcp", host+":9000", &tls.Config{
		ServerName: host,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	c := &Client{
		ctx:            ctx,
		host:           host,
		config:         config,
		conn:           conn,
		reader:         bufio.NewReader(conn),
		closed:         make(chan struct{}),
		hbChannel:      make(chan *message.TestRequest, 1),
		respChannels:   make(map[string]*responseWaiter),
		rejectChannels: make(map[uint32]chan responseResult),
		resubReqs:      make(map[string]message.Request),
		updating:       updating,
	}
	// 后台协程启动前先同步完成 Logon，确保调用方拿到的是可用会话。
	if err = c.logon(); err != nil {
		_ = c.conn.Close()
		return nil, fmt.Errorf("failed to logon: %w", err)
	}
	go c.handlingMessage(ctx)   // 唯一的网络读取和消息分发协程。
	go c.handlingHeartbeat(ctx) // 定时心跳以及 TestRequest 响应协程。

	return c, nil
}

// handlingMessage 持续读取完整 FIX 消息，并负责断线检测、分发和重连。
func (c *Client) handlingMessage(ctx context.Context) {
	defer close(c.closed)

	readTimeout := 0
	for {
		// readMessage 使用 HeartbeatInterval 作为读超时，用于发现静默连接。
		msg, err := c.readMessage()
		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				g.Log().Warningf(ctx, "BinanceFixClient %s: read message timeout", c.config.ClientName)
				readTimeout++
				// 第一次超时发送 TestRequest；连续超时说明连接失效，需要重连。
				if readTimeout > 1 {
					if c.reconnecting(ctx) {
						g.Log().Infof(ctx, "BinanceFixClient %s: stop handling message", c.config.ClientName)
						return
					}
					if c.resubscribing(ctx) {
						g.Log().Infof(ctx, "BinanceFixClient %s: stop handling message", c.config.ClientName)
						return
					}
					readTimeout = 0
					continue
				}
				// 第一次读超时主动探测连接，等待下一个周期确认是否失效。
				if err = c.sendTestReq(); err != nil {
					g.Log().Errorf(ctx, "BinanceFixClient %s: send test request error: %v", c.config.ClientName, err)
				}
			} else {
				readTimeout = 0
				g.Log().Warningf(ctx, "BinanceFixClient %s: read message error: %v", c.config.ClientName, err)
			}
			continue
		} else {
			readTimeout = 0
		}

		// dispatchMessage 返回 reconnect=true 表示服务器要求结束当前会话。
		reconnect, err := c.dispatchMessage(ctx, msg)
		if err != nil {
			g.Log().Errorf(ctx, "BinanceFixClient %s: dispatch message error: %v", c.config.ClientName, err)
			continue
		}

		if reconnect {
			// 先重新建立并登录会话，再恢复所有已保存的行情订阅。
			if c.reconnecting(ctx) {
				return
			}
			if c.resubscribing(ctx) {
				return
			}
		}

		// 每轮分发后检查取消信号，尽量发送 Logout 后关闭连接。
		select {
		case <-ctx.Done():
			if err = c.logout(); err != nil {
				g.Log().Warningf(ctx, "BinanceFixClient %s: failed to logout: %v", c.config.ClientName, err)
			}
			_ = c.conn.Close()
			g.Log().Infof(ctx, "BinanceFixClient %s: stop handling message", c.config.ClientName)
			return
		default:
		}
	}
}

// reconnecting 独占连接锁并按 ReconnectInterval 重试，直到成功或 context 取消。
func (c *Client) reconnecting(ctx context.Context) bool {
	c.connLk.Lock()
	defer c.connLk.Unlock()

	// 先关闭旧连接，唤醒可能仍停留在旧 socket 上的操作。
	_ = c.conn.Close()

	ticker := time.NewTicker(c.config.ReconnectInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return true
		case <-ticker.C:
			if err := c.reconnect(ctx); err != nil {
				g.Log().Warningf(ctx, "BinanceFixClient %s: reconnect failed: %v", c.config.ClientName, err)
			} else {
				return false
			}
		}
	}
}

// reconnect 建立新的 TLS 连接、替换 reader，并从序列号 1 开始重新 Logon。
func (c *Client) reconnect(ctx context.Context) error {
	g.Log().Infof(ctx, "BinanceFixClient %s: try reconnecting...", c.config.ClientName)

	conn, err := tls.Dial("tcp", c.host+":9000", &tls.Config{
		ServerName: c.host,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		return err
	}
	// reader 内部缓存与旧连接绑定，重连后必须一并替换。
	c.conn = conn
	c.reader = bufio.NewReader(conn)

	if err = c.logon(); err != nil {
		_ = c.conn.Close()
		return err
	}

	g.Log().Infof(ctx, "BinanceFixClient %s: reconnected successfully", c.config.ClientName)
	return nil
}

// resubscribing 重放断线前的行情订阅，仅重试发送失败的请求。
func (c *Client) resubscribing(ctx context.Context) bool {
	reqs := c.resubRequests()
	for len(reqs) > 0 {
		g.Log().Infof(ctx, "BinanceFixClient %s: resubscribing to %d requests...", c.config.ClientName, len(reqs))

		// 成功项立即移出本轮重试，避免重复订阅。
		failed := make([]message.Request, 0, len(reqs))
		for _, req := range reqs {
			if err := c.request(req, true); err != nil {
				g.Log().Warningf(ctx, "BinanceFixClient %s: failed to resubscribe request: %v", c.config.ClientName, err)
				failed = append(failed, req)
			}
		}
		if len(failed) == 0 {
			return false
		}
		reqs = failed

		select {
		case <-ctx.Done():
			return true
		case <-time.After(100 * time.Millisecond):
		}
	}
	return false
}

// dispatchMessage 解析消息类型，并将消息路由到心跳、请求 waiter 或 Updating。
// 返回 true 表示当前 FIX 会话应重连。
func (c *Client) dispatchMessage(ctx context.Context, msg *message.Message) (bool, error) {
	msgType, err := msg.MsgType()
	if err != nil {
		return false, err
	}
	now := time.Now()
	sendingTime, err := msg.SendingTime()
	if err != nil {
		return false, err
	}
	// TODO: use metrics instead of log
	g.Log().Debugf(ctx, "BinanceFixClient %s: received message type %s, latency %v", c.config.ClientName, msgType, now.Sub(sendingTime))

	switch msgType {
	case message.MsgTypeHeartbeat:
		g.Log().Debugf(ctx, "BinanceFixClient %s: received heartbeat", c.config.ClientName)
	case message.MsgTypeTestRequest:
		// TestRequest 交给心跳协程回复，避免消息读取协程直接执行网络写入。
		resp := new(message.TestRequest)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		select {
		case <-ctx.Done():
		case c.hbChannel <- resp:
		default:
		}
	case message.MsgTypeReject:
		// 会话级 Reject 使用 RefSeqNum 对应原始请求的 MsgSeqNum。
		resp := new(message.Reject)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if resp.RefSeqNum != nil {
			c.deliverRejectError(*resp.RefSeqNum, resp)
		}
	case message.MsgTypeLogout:
		// 服务端 Logout 或 News 都表示当前会话不可继续使用。
		return true, nil
	case message.MsgTypeNews:
		return true, nil
	case message.MsgTypeLimitResponse:
		// 普通查询响应只投递给对应 waiter，不属于服务器主动推送。
		resp := new(message.LimitResponse)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		c.deliverResponse(resp.ReqID, resp)
	case message.MsgTypeInstrumentList:
		resp := new(message.InstrumentList)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		c.deliverResponse(resp.InstrumentReqID, resp)
	case message.MsgTypeMarketDataSnapshot:
		// 订阅多个 Symbol 时需要先收齐全部初始 Snapshot；waiter 完成后的 Snapshot 才是更新。
		resp := new(message.MarketDataSnapshot)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if !c.deliverResponse(resp.MDReqID, resp) && c.updating != nil {
			sendUpdating(ctx, c.updating.marketData, message.Response(resp))
		}
	case message.MsgTypeMarketDataRequestReject:
		// 订阅被拒绝后不能继续参与重订阅，并立即终止该请求的 waiter。
		resp := new(message.MarketDataRequestReject)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		c.removeResubRequest(resp.MDReqID)
		c.deliverResponseError(resp.MDReqID, resp)
	case message.MsgTypeMarketDataIncrementalRefresh:
		// 增量行情没有同步请求等待者，始终属于订阅更新。
		resp := new(message.MarketDataIncrementalRefresh)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if c.updating != nil {
			sendUpdating(ctx, c.updating.marketData, message.Response(resp))
		}
	case message.MsgTypeExecutionReport:
		// Rejected 是请求错误；正常报告优先作为 ACK，未匹配 waiter 时作为账户级推送。
		resp := new(message.ExecutionReport)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if resp.ExecType == message.ExecTypeRejected {
			c.deliverResponseError(resp.ClOrdID, resp)
		} else if !c.deliverResponse(resp.ClOrdID, resp) && c.updating != nil {
			sendUpdating(ctx, c.updating.orderExecution, resp)
		}
	case message.MsgTypeOrderCancelReject:
		resp := new(message.OrderCancelReject)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		c.deliverResponseError(resp.ClOrdID, resp)
	case message.MsgTypeOrderMassCancelReport:
		// MassCancelResponse 决定报告应作为成功响应还是请求错误返回。
		resp := new(message.OrderMassCancelReport)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if resp.MassCancelResponse == message.MassCancelResponseCancelRequestRejected {
			c.deliverResponseError(resp.ClOrdID, resp)
		} else {
			c.deliverResponse(resp.ClOrdID, resp)
		}
	case message.MsgTypeOrderAmendReject:
		resp := new(message.OrderAmendReject)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		c.deliverResponseError(resp.ClOrdID, resp)
	case message.MsgTypeListStatus:
		// ListStatus 与 ExecutionReport 一样，既可能是请求响应，也可能是账户级主动推送。
		resp := new(message.ListStatus)
		if err = resp.FromMessage(msg); err != nil {
			return false, err
		}
		if resp.ListOrderStatus == message.ListOrderStatusReject {
			c.deliverResponseError(resp.ClListID, resp)
		} else if !c.deliverResponse(resp.ClListID, resp) && c.updating != nil {
			sendUpdating(ctx, c.updating.orderListStatus, resp)
		}
	default:
		return false, fmt.Errorf("unexpected message type %s", msgType)
	}

	return false, nil
}

// handlingHeartbeat 定时发送 Heartbeat，并响应服务端 TestRequest 中携带的 TestReqID。
func (c *Client) handlingHeartbeat(ctx context.Context) {
	ticker := time.NewTicker(c.config.HeartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			g.Log().Infof(ctx, "BinanceFixClient %s: stop handling heartbeat", c.config.ClientName)
			return
		case req := <-c.hbChannel:
			// TestRequest 的响应必须原样带回 TestReqID。
			if err := c.sendHeartbeat(req.TestReqID); err != nil {
				g.Log().Errorf(ctx, "BinanceFixClient %s: failed to send heartbeat: %v", c.config.ClientName, err)
			}
		case <-ticker.C:
			// 普通周期心跳不携带 TestReqID。
			if err := c.sendHeartbeat(""); err != nil {
				g.Log().Errorf(ctx, "BinanceFixClient %s: failed to send heartbeat: %v", c.config.ClientName, err)
			}
		}
	}
}

// sendTestReq 在读超时后主动探测连接是否仍然存活。
func (c *Client) sendTestReq() error {
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	req := message.NewTestRequest(id)
	return c.request(req, true)
}

// sendHeartbeat 允许在重连持有连接锁时跳过心跳，避免把正常重连误报为错误。
func (c *Client) sendHeartbeat(reqId string) error {
	req := message.NewHeartbeat(reqId)
	err := c.request(req, false)
	if err != nil && errors.Is(err, ErrConnectionBroken) {
		return nil
	}
	return err
}

// logon 构造签名后的 Logon 请求，并同步验证服务端返回的第一条消息。
func (c *Client) logon() error {
	// Logon 携带 API Key、Ed25519 签名参数和心跳间隔。
	req := message.NewLogonRequest(
		c.config.ApiKey.UserName,
		c.config.ApiKey.PrivateKey,
		int64(c.config.HeartbeatInterval.Seconds()),
		message.MessageHandlingSequential,
	)
	// ResponseMode 仅属于订单接入会话，Market Data schema 未定义 tag 25036。
	if c.host == orderHost {
		req.WithResponseMode(c.config.ResponseMode)
	}

	// 每个新 FIX 会话的本地发送序列号从 1 开始，Logon 占用第一号。
	c.id = 1
	msg, err := req.ToMessage(c.config.ClientName, targetCompId, c.id, time.Now())
	if err != nil {
		return err
	}
	if err = c.writeMessage(msg); err != nil {
		return err
	}

	// 后台读取协程尚未启动，因此由当前协程同步读取 Logon 响应。
	msg, err = c.readMessage()
	if err != nil {
		return err
	}
	msgTy, err := msg.MsgType()
	if err != nil {
		return err
	}
	if msgTy != message.MsgTypeLogon {
		return LogonFailed
	}
	// 下一条客户端消息从序列号 2 开始。
	c.id++

	return nil
}

// logout 发送 Logout，但不等待服务端响应。
func (c *Client) logout() error {
	return c.request(message.NewLogout(""), false)
}

// request 只负责发送消息，适用于无需等待业务响应的请求。
func (c *Client) request(req message.Request, block bool) error {
	_, err := c.sendRequest(req, block, nil)
	return err
}

// requestAndWait 是单响应请求的便捷封装。
func (c *Client) requestAndWait(req message.Request, ids ...string) (message.Response, error) {
	responses, err := c.requestAndWaitMany(req, 1, ids...)
	if err != nil {
		return nil, err
	}
	return responses[0], nil
}

// requestAndWaitMany 注册业务 ID，发送请求，并等待 expected 条正常响应或任意错误。
func (c *Client) requestAndWaitMany(req message.Request, expected int, ids ...string) ([]message.Response, error) {
	// channel 容量与预期响应数一致，消息读取协程无需等待调用方逐条消费。
	ch := make(chan responseResult, expected)
	// 一个请求可能使用多个等价业务 ID；先去重再绑定到同一个 waiter。
	ids = compactResponseIDs(ids)
	c.registerRespWaiter(ids, ch, expected)

	// 同一个 channel 也按 MsgSeqNum 注册，用于接收会话级 Reject。
	seqNum, err := c.sendRequest(req, false, ch)
	if err != nil {
		c.removeRespChannels(ids)
		return nil, err
	}
	// 无论成功、超时、Reject 还是 context 取消，都清理两类路由表。
	defer func() {
		c.removeRespChannels(ids)
		c.removeRejectChannel(seqNum)
	}()

	// 多响应共享同一个总超时，而不是每收到一条就重新计时。
	timer := time.NewTimer(c.config.ResponseTimeout)
	defer timer.Stop()

	// Reject 会通过 result.err 立即结束；正常响应则收集到 expected 条。
	responses := make([]message.Response, 0, expected)
	for len(responses) < expected {
		select {
		case result := <-ch:
			if result.err != nil {
				return nil, result.err
			}
			responses = append(responses, result.response)
		case <-timer.C:
			return nil, ErrResponseTimeout
		case <-c.ctx.Done():
			return nil, c.ctx.Err()
		}
	}
	return responses, nil
}

// sendRequest 串行生成 MsgSeqNum、构造消息并写入当前 TLS 连接。
func (c *Client) sendRequest(req message.Request, block bool, rejectCh chan responseResult) (uint32, error) {
	if block {
		// 内部恢复流程允许等待重连锁释放。
		c.connLk.Lock()
	} else {
		// 外部请求不等待重连，立即返回连接不可用错误。
		if !c.connLk.TryLock() {
			return 0, ErrConnectionBroken
		}
	}
	defer c.connLk.Unlock()

	// 构造消息和注册 Reject 必须使用同一个尚未递增的序列号。
	seqNum := c.id
	msg, err := req.ToMessage(c.config.ClientName, targetCompId, seqNum, time.Now())
	if err != nil {
		return 0, err
	}

	// 必须先注册再写入，避免响应过快而找不到等待者。
	if rejectCh != nil {
		c.registerRejectChannel(seqNum, rejectCh)
	}
	if err = c.writeMessage(msg); err != nil {
		if rejectCh != nil {
			c.removeRejectChannel(seqNum)
		}
		return 0, err
	}
	// 只有消息成功写入后才消费当前序列号。
	c.id++
	return seqNum, nil
}

// compactResponseIDs 去除同一请求中的重复业务 ID，避免重复注册和清理。
func compactResponseIDs(ids []string) []string {
	compacted := ids[:0]
	seen := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		compacted = append(compacted, id)
	}
	return compacted
}

// registerRespWaiter 将所有等价业务 ID 指向同一个 waiter。
// expected 控制同步响应边界，避免多响应请求的第二条消息被误判为主动推送。
func (c *Client) registerRespWaiter(keys []string, ch chan responseResult, expected int) {
	if len(keys) == 0 {
		return
	}

	c.respChanLk.Lock()
	defer c.respChanLk.Unlock()

	// 复制 keys，确保调用方后续复用原 slice 不会破坏 waiter 的清理列表。
	waiter := &responseWaiter{
		ch:        ch,
		ids:       append([]string(nil), keys...),
		remaining: expected,
	}
	for _, key := range keys {
		c.respChannels[key] = waiter
	}
}

// removeRespChannels 清理请求退出时仍然存在的业务 ID 路由。
func (c *Client) removeRespChannels(ids []string) {
	if len(ids) == 0 {
		return
	}
	c.respChanLk.Lock()
	defer c.respChanLk.Unlock()
	for _, id := range ids {
		delete(c.respChannels, id)
	}
}

// registerRejectChannel 使用发送序列号关联会话级 Reject 和原始请求。
func (c *Client) registerRejectChannel(seqNum uint32, ch chan responseResult) {
	c.rejectChanLk.Lock()
	c.rejectChannels[seqNum] = ch
	c.rejectChanLk.Unlock()
}

// removeRejectChannel 清理已完成请求的会话级 Reject 路由。
func (c *Client) removeRejectChannel(seqNum uint32) {
	c.rejectChanLk.Lock()
	delete(c.rejectChannels, seqNum)
	c.rejectChanLk.Unlock()
}

// deliverResponse 尝试将正常响应投递给业务 ID 对应的 waiter。
// 返回 false 表示没有同步请求在等待，该消息可继续作为 Updating 推送。
func (c *Client) deliverResponse(id string, resp message.Response) bool {
	return c.deliverResponseResult(id, responseResult{response: resp})
}

// deliverResponseError 将业务级 Reject 作为 error 投递，并立即结束整个 waiter。
func (c *Client) deliverResponseError(id string, err error) {
	c.deliverResponseResult(id, responseResult{err: err})
}

// deliverResponseResult 更新 waiter 状态并投递一条结果。
func (c *Client) deliverResponseResult(id string, result responseResult) bool {
	c.respChanLk.Lock()
	waiter, ok := c.respChannels[id]
	if !ok {
		c.respChanLk.Unlock()
		return false
	}

	// 任意错误都会终止请求；正常响应则只消费一个预期响应名额。
	if result.err != nil {
		waiter.remaining = 0
	} else {
		waiter.remaining--
	}
	// waiter 完成时一次性删除全部别名，后续同 ID 消息将进入 Updating。
	if waiter.remaining <= 0 {
		for _, key := range waiter.ids {
			if c.respChannels[key] == waiter {
				delete(c.respChannels, key)
			}
		}
	}
	c.respChanLk.Unlock()

	// channel 按预期响应数分配容量；default 只用于防御重复或竞态投递。
	select {
	case waiter.ch <- result:
	default:
	}
	return true
}

// deliverRejectError 按 RefSeqNum 投递会话级 Reject，并保证同一 Reject 只消费一次。
func (c *Client) deliverRejectError(seqNum uint32, err error) {
	c.rejectChanLk.Lock()
	ch, ok := c.rejectChannels[seqNum]
	if ok {
		delete(c.rejectChannels, seqNum)
	}
	c.rejectChanLk.Unlock()
	if !ok {
		return
	}

	// 请求 channel 已预留容量，default 防止异常重复消息阻塞读取协程。
	select {
	case ch <- responseResult{err: err}:
	default:
	}
}

// setResubRequest 保存独立的订阅请求副本，供连接恢复后重放。
func (c *Client) setResubRequest(id string, req message.Request) {
	if id == "" || req == nil {
		return
	}
	c.resubReqLk.Lock()
	c.resubReqs[id] = req
	c.resubReqLk.Unlock()
}

// removeResubRequest 在取消订阅或订阅被拒绝时移除重订阅记录。
func (c *Client) removeResubRequest(id string) {
	if id == "" {
		return
	}
	c.resubReqLk.Lock()
	delete(c.resubReqs, id)
	c.resubReqLk.Unlock()
}

// resubRequests 返回当前订阅快照，避免网络发送期间长期持有 resubReqLk。
func (c *Client) resubRequests() []message.Request {
	c.resubReqLk.Lock()
	defer c.resubReqLk.Unlock()
	reqs := make([]message.Request, 0, len(c.resubReqs))
	for _, req := range c.resubReqs {
		reqs = append(reqs, req)
	}
	return reqs
}

// cloneMarketDataRequest 深拷贝请求中的 slice 和指针字段，隔离调用方后续修改。
func cloneMarketDataRequest(req *message.MarketDataRequest) *message.MarketDataRequest {
	if req == nil {
		return nil
	}
	cloned := *req
	cloned.Symbols = append([]string(nil), req.Symbols...)
	cloned.MDEntryTypes = append([]message.MDEntryType(nil), req.MDEntryTypes...)
	if req.AggregatedBook != nil {
		aggregatedBook := *req.AggregatedBook
		cloned.AggregatedBook = &aggregatedBook
	}
	return &cloned
}

func executionReportResponse(resp message.Response, err error) (*message.ExecutionReport, error) {
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.ExecutionReport)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeExecutionReport)
	}
	return result, nil
}

func unexpectedResponseError(resp message.Response, want message.MsgType) error {
	return fmt.Errorf("unexpected response type %T, want %s", resp, want)
}

// writeMessage 设置单次写超时并一次性写入完整 FIX 消息。
func (c *Client) writeMessage(msg *message.Message) error {
	if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.WriteTimeout)); err != nil {
		return err
	}
	if _, err := c.conn.Write([]byte(msg.RawMessage())); err != nil {
		return err
	}
	return nil
}

// readMessage 先读取 FIX 头部，再根据 BodyLength 精确读取消息体和 CheckSum。
func (c *Client) readMessage() (*message.Message, error) {
	// 一个心跳周期内没有任何入站数据时返回超时，由上层执行连接探测。
	if err := c.conn.SetReadDeadline(time.Now().Add(c.config.HeartbeatInterval)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	// BeginString 是首个 SOH 分隔字段，可用于快速拒绝非 FIX 4.4 数据。
	beginStr, err := c.reader.ReadString(message.SOH)
	if err != nil {
		return nil, fmt.Errorf("failed to read SOH: %w", err)
	}
	if beginStr != "8=FIX.4.4\x01" {
		return nil, fmt.Errorf("invalid begin string: %s", beginStr)
	}

	// BodyLength 是第二个字段，表示从 MsgType 起到 CheckSum 前的字节数。
	bodyLenStr, err := c.reader.ReadString(message.SOH)
	if err != nil {
		return nil, fmt.Errorf("failed to read SOH: %w", err)
	}
	bls := strings.TrimPrefix(bodyLenStr, "9=")
	bls = strings.TrimSuffix(bls, "\x01")
	bodyLen, err := strconv.ParseUint(bls, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid body len field: %s", bodyLenStr)
	}

	// CheckSum 固定为 "10=xxx<SOH>" 共 7 字节，因此在 BodyLength 基础上额外读取 7 字节。
	buf := make([]byte, bodyLen+7)
	n, err := io.ReadFull(c.reader, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to read body and checksum (%d/%d bytes): %w", n, len(buf), err)
	}

	msg, err := message.ParseMessage(beginStr + bodyLenStr + string(buf))
	if err != nil {
		return nil, fmt.Errorf("parse FIX message: %w", err)
	}
	return msg, nil
}
