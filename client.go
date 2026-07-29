package binance_go_fix

import (
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

	"github.com/quantaverse-org/binance-go-fix/internal/fixsbe"
	"github.com/quantaverse-org/binance-go-fix/message"

	"github.com/gogf/gf/v2/frame/g"
)

const (
	targetCompId = "SPOT"
	orderHost    = "fix-oe.binance.com"
	marketHost   = "fix-md.binance.com"
	// subscription channel 需要吸收同步请求等待期间穿插到达的主动推送，避免阻塞 FIX 读循环。
	subscriptionChannelCapacity = 1024
)

var (
	ErrReconnecting    = errors.New("reconnecting, retry later")
	ErrResponseTimeout = errors.New("response timeout")
)

// responseResult 将正常响应和 Reject 统一送入同一个请求等待通道。
type responseResult struct {
	response message.Response
	err      error
}

// responseWaiter 表示一个正在等待单条同步响应的请求。
// 同一个请求可能通过多个业务 ID 关联响应。
type responseWaiter struct {
	ch  chan responseResult // 接收正常响应或 Reject。
	ids []string            // 注册到 respChannels 的全部业务 ID，用于统一清理别名。
}

type sbeMarketRoute struct {
	symbol     string
	templateID uint16
}

type sbeBookLevel struct {
	symbol    string
	entryType message.MDEntryType
	price     float64
}

// MarketSubscription 暴露市场数据服务器的推送消息。
type MarketSubscription struct {
	// MarketData 中只会出现 *MarketDataSnapshot 和 *MarketDataIncrementalRefresh。
	MarketData <-chan message.Response
}

// OrderSubscription 暴露订单服务器的账户级推送消息。
type OrderSubscription struct {
	OrderExecution  <-chan *message.ExecutionReport
	OrderListStatus <-chan *message.ListStatus
}

// subscriptionSenders 只由 Client 的消息分发协程持有，防止调用方误写更新通道。
type subscriptionSenders struct {
	marketData      chan<- message.Response
	orderExecution  chan<- *message.ExecutionReport
	orderListStatus chan<- *message.ListStatus
}

func initMarketSubscription(cap int) (*subscriptionSenders, *MarketSubscription) {
	marketData := make(chan message.Response, cap)
	return &subscriptionSenders{marketData: marketData}, &MarketSubscription{MarketData: marketData}
}

func initOrderSubscription(cap int) (*subscriptionSenders, *OrderSubscription) {
	orderExecution := make(chan *message.ExecutionReport, cap)
	orderListStatus := make(chan *message.ListStatus, cap)

	return &subscriptionSenders{
			orderExecution:  orderExecution,
			orderListStatus: orderListStatus,
		}, &OrderSubscription{
			OrderExecution:  orderExecution,
			OrderListStatus: orderListStatus,
		}
}

// sendSubscription 在客户端关闭时放弃发送，否则允许缓冲区提供有限背压。
func sendSubscription[T any](ctx context.Context, ch chan<- T, value T) {
	select {
	case <-ctx.Done():
	case ch <- value:
	}
}

type ApiKey struct {
	UserName   string
	PrivateKey ed25519.PrivateKey
}

type EncodingMode uint8

const (
	// EncodingModeFIX uses text FIX requests and responses on port 9000.
	EncodingModeFIX EncodingMode = iota
	// EncodingModeFIXRequestSBEResponse uses text FIX requests and SBE responses on port 9001.
	EncodingModeFIXRequestSBEResponse
	// EncodingModeSBE uses SBE requests and responses on port 9002.
	EncodingModeSBE
)

type ClientConfig struct {
	EnableNotify      bool
	ClientName        string
	ChannelCapacity   int
	HeartbeatInterval time.Duration
	ReconnectInterval time.Duration
	ResponseTimeout   time.Duration
	WriteTimeout      time.Duration
	ResponseMode      message.ResponseMode
	EncodingMode      EncodingMode
	ApiKey            *ApiKey
}

func NewClientConfig(apiKey *ApiKey, clientName string) *ClientConfig {
	return &ClientConfig{
		EnableNotify:      false,
		ClientName:        clientName,
		ChannelCapacity:   subscriptionChannelCapacity,
		HeartbeatInterval: time.Second * 30,
		ReconnectInterval: time.Second * 1,
		ResponseTimeout:   time.Second * 10,
		WriteTimeout:      time.Second * 1,
		ResponseMode:      message.ResponseModeEverything,
		EncodingMode:      EncodingModeFIX,
		ApiKey:            apiKey,
	}
}

func (c *ClientConfig) WithEnableNotify() *ClientConfig {
	c.EnableNotify = true
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

func (c *ClientConfig) WithEncodingMode(mode EncodingMode) *ClientConfig {
	c.EncodingMode = mode
	return c
}

type MarketClient struct {
	*Client
}

// NewMarketClient 建立市场数据会话；仅在 EnableNotify 开启时创建 MarketSubscription。
func NewMarketClient(config *ClientConfig) (*MarketClient, *MarketSubscription, error) {
	var senders *subscriptionSenders
	var subscription *MarketSubscription
	if config.EnableNotify {
		senders, subscription = initMarketSubscription(config.ChannelCapacity)
	}

	client, err := newClient(marketHost, config, senders)
	if err != nil {
		return nil, nil, err
	}
	return &MarketClient{Client: client}, subscription, nil
}

func (c *MarketClient) InstrumentList(ctx context.Context, req *message.InstrumentListRequest) (*message.InstrumentList, error) {
	resp, err := c.requestAndWait(ctx, req, req.InstrumentReqID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.InstrumentList)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeInstrumentList)
	}
	return result, nil
}

func (c *MarketClient) MarketData(ctx context.Context, req *message.MarketDataRequest) error {
	// 取消订阅没有成功响应，消息写入连接后即可返回。
	if req.SubscriptionRequestType == message.SubscriptionRequestTypeUnsubscribe {
		c.removeResubRequest(req.MDReqID)
		c.removeSBEMarketRoutes(req.MDReqID)
		return c.request(req, false)
	}

	// SBE 行情响应不携带 MDReqID，必须在请求写入前登记 symbol/template 路由。
	if encodingModeUsesSBEResponses(c.config.EncodingMode) {
		if err := c.registerSBEMarketRoutes(req); err != nil {
			return err
		}
	}

	// 第一条 Snapshot 或 IncrementalRefresh 表示订阅成功，Reject 表示订阅失败。
	_, err := c.requestAndWait(ctx, req, req.MDReqID)
	if err != nil {
		c.removeResubRequest(req.MDReqID)
		c.removeSBEMarketRoutes(req.MDReqID)
		return err
	}
	// 确认订阅成功后保存请求副本，重连完成时用于恢复订阅。
	c.setResubRequest(req.MDReqID, cloneMarketDataRequest(req))
	return nil
}

type OrderClient struct {
	*Client
}

// NewOrderClient 建立订单会话；OrderSubscription 用于接收账户级 ExecutionReport 和 ListStatus 推送。
func NewOrderClient(config *ClientConfig) (*OrderClient, *OrderSubscription, error) {
	var senders *subscriptionSenders
	var subscription *OrderSubscription
	if config.EnableNotify {
		senders, subscription = initOrderSubscription(config.ChannelCapacity)
	}

	client, err := newClient(orderHost, config, senders)
	if err != nil {
		return nil, nil, err
	}
	return &OrderClient{Client: client}, subscription, nil
}

func (c *OrderClient) NewOrderSingle(ctx context.Context, req *message.NewOrderSingle) (*message.ExecutionReport, error) {
	// 首条匹配 ClOrdID 的 ExecutionReport 是同步 ACK，后续状态变化进入 Subscription。
	resp, err := c.requestAndWait(ctx, req, req.ClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) OrderCancel(ctx context.Context, req *message.OrderCancelRequest) (*message.ExecutionReport, error) {
	resp, err := c.requestAndWait(ctx, req, req.ClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) OrderMassCancel(ctx context.Context, req *message.OrderMassCancelRequest) (*message.OrderMassCancelReport, error) {
	// 这里只等待汇总报告；每个被取消订单的 ExecutionReport 属于账户级订阅消息。
	resp, err := c.requestAndWait(ctx, req, req.ClOrdID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.OrderMassCancelReport)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeOrderMassCancelReport)
	}
	return result, nil
}

func (c *OrderClient) OrderReplace(ctx context.Context, req *message.OrderCancelRequestAndNewOrderSingle) (*message.ExecutionReport, error) {
	// 新订单 ID 和被取消订单 ID 都可能出现在 ACK 中，因此注册为同一 waiter 的别名。
	resp, err := c.requestAndWait(ctx, req, req.ClOrdID, req.CancelClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) ListStatus(ctx context.Context, req *message.NewOrderList) (*message.ListStatus, error) {
	// ListStatus 是该请求的同步结果，各子订单的 ExecutionReport 通过 Subscription 接收。
	resp, err := c.requestAndWait(ctx, req, req.ClListID)
	if err != nil {
		return nil, err
	}
	result, ok := resp.(*message.ListStatus)
	if !ok {
		return nil, unexpectedResponseError(resp, message.MsgTypeListStatus)
	}
	return result, nil
}

func (c *OrderClient) OrderAmendKeepPriority(ctx context.Context, req *message.OrderAmendKeepPriorityRequest) (*message.ExecutionReport, error) {
	// 修改结果由 ExecutionReport 确认；若订单属于列表，额外 ListStatus 会进入 Subscription。
	resp, err := c.requestAndWait(ctx, req, req.ClOrdID)
	return executionReportResponse(resp, err)
}

func (c *OrderClient) Limit(ctx context.Context, req *message.LimitQuery) (*message.LimitResponse, error) {
	resp, err := c.requestAndWait(ctx, req, req.ReqID)
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
	host   string
	config *ClientConfig

	// connLk 串行化写入、重连和序列号递增，reader 只由消息处理协程使用。
	connLk sync.Mutex
	conn   *tls.Conn
	reader *bufio.Reader
	// SbeGoMarshaller 会复用内部缓冲；读响应和写请求必须使用不同实例。
	sbeMarshaller        *fixsbe.SbeGoMarshaller
	sbeRequestMarshaller *fixsbe.SbeGoMarshaller
	id                   uint32
	closed               chan struct{}
	hbChannel            chan *message.TestRequest

	// respChannels 用业务 ID 将应用层响应路由给正在等待的请求。
	respChanLk   sync.Mutex
	respChannels map[string]*responseWaiter

	// rejectChannels 用请求的 MsgSeqNum 将会话级 Reject 路由回原请求。
	rejectChanLk   sync.Mutex
	rejectChannels map[uint32]chan responseResult

	// resubReqs 保存成功建立的市场数据订阅，用于断线重连后恢复。
	resubReqLk sync.Mutex
	resubReqs  map[string]message.Request

	// SBE 行情省略 MDReqID，使用订阅时登记的 symbol/template 找回原请求。
	sbeMarketLk     sync.Mutex
	sbeMarketRoutes map[sbeMarketRoute]string
	sbeBookLevels   map[string]map[sbeBookLevel]struct{}

	// subscription 为 nil 时忽略所有未匹配请求的主动推送。
	subscription *subscriptionSenders
}

func newClient(host string, config *ClientConfig, subscription *subscriptionSenders) (*Client, error) {
	address, err := clientAddress(host, config.EncodingMode)
	if err != nil {
		return nil, err
	}
	// ServerName 用于 TLS SNI 和证书主机名校验。
	conn, err := tls.Dial("tcp", address, &tls.Config{
		ServerName: host,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return &Client{
		host:                 host,
		config:               config,
		conn:                 conn,
		reader:               bufio.NewReader(conn),
		sbeMarshaller:        fixsbe.NewSbeGoMarshaller(),
		sbeRequestMarshaller: fixsbe.NewSbeGoMarshaller(),
		closed:               make(chan struct{}),
		hbChannel:            make(chan *message.TestRequest, 1),
		respChannels:         make(map[string]*responseWaiter),
		rejectChannels:       make(map[uint32]chan responseResult),
		resubReqs:            make(map[string]message.Request),
		sbeMarketRoutes:      make(map[sbeMarketRoute]string),
		sbeBookLevels:        make(map[string]map[sbeBookLevel]struct{}),
		subscription:         subscription,
	}, nil
}

func clientAddress(host string, mode EncodingMode) (string, error) {
	switch mode {
	case EncodingModeFIX:
		return host + ":9000", nil
	case EncodingModeFIXRequestSBEResponse:
		return host + ":9001", nil
	case EncodingModeSBE:
		return host + ":9002", nil
	default:
		return "", fmt.Errorf("unsupported encoding mode: %d", mode)
	}
}

func encodingModeUsesSBEResponses(mode EncodingMode) bool {
	return mode == EncodingModeFIXRequestSBEResponse || mode == EncodingModeSBE
}

func encodingModeUsesSBERequests(mode EncodingMode) bool {
	return mode == EncodingModeSBE
}

func (c *Client) Run(ctx context.Context) error {
	// 后台协程启动前先同步完成 Logon，确保调用方拿到的是可用会话。
	if err := c.logon(); err != nil {
		_ = c.conn.Close()
		return fmt.Errorf("failed to logon: %w", err)
	}
	go c.handlingMessage(ctx)   // 唯一的网络读取和消息分发协程。
	go c.handlingHeartbeat(ctx) // 定时心跳以及 TestRequest 响应协程。
	return nil
}

func (c *Client) UtilClosed(ctx context.Context) {
	select {
	case <-ctx.Done():
	case <-c.closed:
	}
}

// handlingMessage 持续读取完整 FIX 消息，并负责断线检测、分发和重连。
func (c *Client) handlingMessage(ctx context.Context) {
	defer close(c.closed)

	readTimeout := 0
	for {
		// readIncomingMessage 根据连接模式读取文本 FIX 或一个完整 SBE frame。
		inbound, err := c.readIncomingMessage()
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
				// SBE 的 SOFH 提供 frame 边界；读取或解码失败后无法保证 reader
				// 仍对齐下一条 frame，因此重新建立会话并恢复行情订阅。
				if encodingModeUsesSBEResponses(c.config.EncodingMode) {
					if c.reconnecting(ctx) {
						return
					}
					if c.resubscribing(ctx) {
						return
					}
				}
			}
			continue
		} else {
			readTimeout = 0
		}

		// dispatchInboundMessage 返回 reconnect=true 表示服务器要求结束当前会话。
		reconnect, err := c.dispatchInboundMessage(ctx, inbound)
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

	address, err := clientAddress(c.host, c.config.EncodingMode)
	if err != nil {
		return err
	}
	conn, err := tls.Dial("tcp", address, &tls.Config{
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
	c.clearSBEBookLevels()

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

// dispatchMessage 保留文本 FIX 的测试和内部入口，实际路由与 SBE 共用 dispatchInboundMessage。
func (c *Client) dispatchMessage(ctx context.Context, msg *message.Message) (bool, error) {
	inbound, err := decodeFIXInboundMessage(msg)
	if err != nil {
		return false, err
	}
	return c.dispatchInboundMessage(ctx, inbound)
}

func decodeFIXInboundMessage(msg *message.Message) (*inboundMessage, error) {
	msgType, err := msg.MsgType()
	if err != nil {
		return nil, err
	}
	sendingTime, err := msg.SendingTime()
	if err != nil {
		return nil, err
	}

	var response message.Response
	switch msgType {
	case message.MsgTypeHeartbeat:
		response = new(message.Heartbeat)
	case message.MsgTypeTestRequest:
		response = new(message.TestRequest)
	case message.MsgTypeReject:
		response = new(message.Reject)
	case message.MsgTypeLogout:
		response = new(message.Logout)
	case message.MsgTypeNews:
		response = new(message.News)
	case message.MsgTypeLogon:
		response = new(message.LogonResponse)
	case message.MsgTypeLimitResponse:
		response = new(message.LimitResponse)
	case message.MsgTypeInstrumentList:
		response = new(message.InstrumentList)
	case message.MsgTypeMarketDataSnapshot:
		response = new(message.MarketDataSnapshot)
	case message.MsgTypeMarketDataRequestReject:
		response = new(message.MarketDataRequestReject)
	case message.MsgTypeMarketDataIncrementalRefresh:
		response = new(message.MarketDataIncrementalRefresh)
	case message.MsgTypeExecutionReport:
		response = new(message.ExecutionReport)
	case message.MsgTypeOrderCancelReject:
		response = new(message.OrderCancelReject)
	case message.MsgTypeOrderMassCancelReport:
		response = new(message.OrderMassCancelReport)
	case message.MsgTypeOrderAmendReject:
		response = new(message.OrderAmendReject)
	case message.MsgTypeListStatus:
		response = new(message.ListStatus)
	default:
		return nil, fmt.Errorf("unexpected message type %s", msgType)
	}

	if err := response.FromMessage(msg); err != nil {
		return nil, err
	}
	inbound := &inboundMessage{
		msgType:     msgType,
		sendingTime: sendingTime,
		response:    response,
	}
	if seqNumValue, ok := msg.GetField(message.TagMsgSeqNum); ok {
		seqNum, err := message.ParseUint(seqNumValue)
		if err != nil {
			return nil, err
		}
		inbound.seqNum = uint32(seqNum)
	}
	return inbound, nil
}

// dispatchInboundMessage routes an already decoded response to heartbeat,
// synchronous request waiters, or the appropriate Subscription.
func (c *Client) dispatchInboundMessage(ctx context.Context, inbound *inboundMessage) (bool, error) {
	switch resp := inbound.response.(type) {
	case *message.Heartbeat:
		g.Log().Debugf(ctx, "BinanceFixClient %s: received heartbeat", c.config.ClientName)
	case *message.TestRequest:
		// TestRequest 交给心跳协程回复，避免消息读取协程直接执行网络写入。
		select {
		case <-ctx.Done():
		case c.hbChannel <- resp:
		default:
		}
	case *message.Reject:
		// 会话级 Reject 使用 RefSeqNum 对应原始请求的 MsgSeqNum。
		if resp.RefSeqNum != nil {
			c.deliverRejectError(*resp.RefSeqNum, resp)
		}
	case *message.Logout, *message.News:
		// 服务端 Logout 或 News 都表示当前会话不可继续使用。
		return true, nil
	case *message.LogonResponse:
		return false, fmt.Errorf("unexpected logon response after session establishment")
	case *message.LimitResponse:
		c.deliverResponse(resp.ReqID, resp)
	case *message.InstrumentList:
		c.deliverResponse(resp.InstrumentReqID, resp)
	case *message.MarketDataSnapshot:
		if err := c.resolveSBEMarketData(inbound, resp); err != nil {
			return false, err
		}
		// 第一条行情完成订阅 waiter；行情本身始终保留在 Subscription 中供调用方消费。
		c.deliverResponse(resp.MDReqID, resp)
		if c.subscription != nil {
			sendSubscription[message.Response](ctx, c.subscription.marketData, resp)
		}
	case *message.MarketDataRequestReject:
		// 订阅被拒绝后不能继续参与重订阅，并立即终止该请求的 waiter。
		c.removeResubRequest(resp.MDReqID)
		c.removeSBEMarketRoutes(resp.MDReqID)
		c.deliverResponseError(resp.MDReqID, resp)
	case *message.MarketDataIncrementalRefresh:
		if err := c.resolveSBEMarketData(inbound, resp); err != nil {
			return false, err
		}
		// Trade 订阅以第一条增量行情确认成功，后续增量行情只进入 Subscription。
		c.deliverResponse(resp.MDReqID, resp)
		if c.subscription != nil {
			sendSubscription[message.Response](ctx, c.subscription.marketData, resp)
		}
	case *message.ExecutionReport:
		// Rejected 是请求错误；正常报告优先作为 ACK，未匹配 waiter 时作为账户级推送。
		if resp.ExecType == message.ExecTypeRejected {
			c.deliverResponseError(resp.ClOrdID, resp)
		} else if !c.deliverResponse(resp.ClOrdID, resp) && c.subscription != nil {
			sendSubscription(ctx, c.subscription.orderExecution, resp)
		}
	case *message.OrderCancelReject:
		c.deliverResponseError(resp.ClOrdID, resp)
	case *message.OrderMassCancelReport:
		if resp.MassCancelResponse == message.MassCancelResponseCancelRequestRejected {
			c.deliverResponseError(resp.ClOrdID, resp)
		} else {
			c.deliverResponse(resp.ClOrdID, resp)
		}
	case *message.OrderAmendReject:
		c.deliverResponseError(resp.ClOrdID, resp)
	case *message.ListStatus:
		// ListStatus 与 ExecutionReport 一样，既可能是请求响应，也可能是账户级主动推送。
		if resp.ListOrderStatus == message.ListOrderStatusReject {
			c.deliverResponseError(resp.ClListID, resp)
		} else if !c.deliverResponse(resp.ClListID, resp) && c.subscription != nil {
			sendSubscription(ctx, c.subscription.orderListStatus, resp)
		}
	default:
		return false, fmt.Errorf("unexpected response type %T", inbound.response)
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
	if err != nil && errors.Is(err, ErrReconnecting) {
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
	// ResponseMode 仅适用于订单会话，market data 会话携带该字段会被拒绝。
	if c.host == orderHost {
		req.WithResponseMode(c.config.ResponseMode)
	}
	if c.config.EncodingMode == EncodingModeFIXRequestSBEResponse {
		req.WithSbeSchema(sbeSchemaID, sbeSchemaVersion)
	}

	// 每个新会话的本地发送序列号从 1 开始，Logon 占用第一号。
	c.id = 1
	if err := c.writeRequest(req, c.id, time.Now()); err != nil {
		return err
	}

	// 后台读取协程尚未启动，因此由当前协程同步读取 Logon 响应。
	inbound, err := c.readIncomingMessage()
	if err != nil {
		return err
	}
	if inbound.msgType == message.MsgTypeReject {
		if reject, ok := inbound.response.(*message.Reject); ok {
			return reject
		}
	}
	if inbound.msgType != message.MsgTypeLogon {
		return fmt.Errorf("unexpected logon response type: %s", inbound.msgType)
	}
	if response, ok := inbound.response.(*message.LogonResponse); ok && response.SbeSchemaIdVersionDeprecated {
		g.Log().Warningf(context.Background(), "BinanceFixClient %s: FIX SBE schema %d:%d is deprecated",
			c.config.ClientName, sbeSchemaID, sbeSchemaVersion)
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

// requestAndWait 注册业务 ID，发送请求，并等待一条正常响应或任意错误。
func (c *Client) requestAndWait(ctx context.Context, req message.Request, ids ...string) (message.Response, error) {
	// 单响应 channel 预留一个位置，消息读取协程无需等待调用方消费。
	ch := make(chan responseResult, 1)
	// 一个请求可能使用多个等价业务 ID；先去重再绑定到同一个 waiter。
	ids = compactResponseIDs(ids)
	c.registerRespWaiter(ids, ch)

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

	timer := time.NewTimer(c.config.ResponseTimeout)
	defer timer.Stop()

	select {
	case result := <-ch:
		if result.err != nil {
			return nil, result.err
		}
		return result.response, nil
	case <-timer.C:
		return nil, ErrResponseTimeout
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// sendRequest 串行生成 MsgSeqNum、构造消息并写入当前 TLS 连接。
func (c *Client) sendRequest(req message.Request, block bool, rejectCh chan responseResult) (uint32, error) {
	if block {
		// 内部恢复流程允许等待重连锁释放。
		c.connLk.Lock()
	} else {
		// 外部请求不等待重连，立即返回连接不可用错误。
		if !c.connLk.TryLock() {
			return 0, ErrReconnecting
		}
	}
	defer c.connLk.Unlock()

	// 构造消息和注册 Reject 必须使用同一个尚未递增的序列号。
	seqNum := c.id
	// 必须先注册再写入，避免响应过快而找不到等待者。
	if rejectCh != nil {
		c.registerRejectChannel(seqNum, rejectCh)
	}
	if err := c.writeRequest(req, seqNum, time.Now()); err != nil {
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
func (c *Client) registerRespWaiter(keys []string, ch chan responseResult) {
	if len(keys) == 0 {
		return
	}

	c.respChanLk.Lock()
	defer c.respChanLk.Unlock()

	// 复制 keys，确保调用方后续复用原 slice 不会破坏 waiter 的清理列表。
	waiter := &responseWaiter{
		ch:  ch,
		ids: append([]string(nil), keys...),
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
// 返回 false 表示没有同步请求在等待，该消息可继续作为 Subscription 推送。
func (c *Client) deliverResponse(id string, resp message.Response) bool {
	return c.deliverResponseResult(id, responseResult{response: resp})
}

// deliverResponseError 将业务级 Reject 作为 error 投递，并立即结束整个 waiter。
func (c *Client) deliverResponseError(id string, err error) {
	c.deliverResponseResult(id, responseResult{err: err})
}

// deliverResponseResult 完成 waiter 并投递一条结果。
func (c *Client) deliverResponseResult(id string, result responseResult) bool {
	c.respChanLk.Lock()
	waiter, ok := c.respChannels[id]
	if !ok {
		c.respChanLk.Unlock()
		return false
	}

	// 第一条正常响应或错误都会完成请求；后续同 ID 消息进入 Subscription。
	for _, key := range waiter.ids {
		if c.respChannels[key] == waiter {
			delete(c.respChannels, key)
		}
	}
	c.respChanLk.Unlock()

	// channel 已预留一个位置；default 只用于防御异常重复或竞态投递。
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

func (c *Client) registerSBEMarketRoutes(req *message.MarketDataRequest) error {
	templates := sbeMarketTemplates(req)
	routes := make([]sbeMarketRoute, 0, len(req.Symbols)*len(templates))
	for _, symbol := range req.Symbols {
		for _, templateID := range templates {
			routes = append(routes, sbeMarketRoute{symbol: symbol, templateID: templateID})
		}
	}

	c.sbeMarketLk.Lock()
	defer c.sbeMarketLk.Unlock()
	if c.sbeMarketRoutes == nil {
		c.sbeMarketRoutes = make(map[sbeMarketRoute]string)
	}
	if c.sbeBookLevels == nil {
		c.sbeBookLevels = make(map[string]map[sbeBookLevel]struct{})
	}
	for _, route := range routes {
		if existingID, ok := c.sbeMarketRoutes[route]; ok && existingID != req.MDReqID {
			return fmt.Errorf("SBE market route already registered: symbol=%s template=%d", route.symbol, route.templateID)
		}
	}
	for _, route := range routes {
		c.sbeMarketRoutes[route] = req.MDReqID
	}
	return nil
}

func sbeMarketTemplates(req *message.MarketDataRequest) []uint16 {
	var hasBook, hasTrade bool
	for _, entryType := range req.MDEntryTypes {
		switch entryType {
		case message.MDEntryTypeBid, message.MDEntryTypeOffer:
			hasBook = true
		case message.MDEntryTypeTrade:
			hasTrade = true
		}
	}

	templates := make([]uint16, 0, 3)
	if hasBook {
		templates = append(templates, sbeTemplateMarketDataSnapshot)
		if req.MarketDepth == 1 {
			templates = append(templates, sbeTemplateMarketDataIncrementalBookTicker)
		} else {
			templates = append(templates, sbeTemplateMarketDataIncrementalDepth)
		}
	}
	if hasTrade {
		templates = append(templates, sbeTemplateMarketDataIncrementalTrade)
	}
	return templates
}

func (c *Client) removeSBEMarketRoutes(id string) {
	if id == "" {
		return
	}
	c.sbeMarketLk.Lock()
	defer c.sbeMarketLk.Unlock()
	for route, routeID := range c.sbeMarketRoutes {
		if routeID == id {
			delete(c.sbeMarketRoutes, route)
		}
	}
	delete(c.sbeBookLevels, id)
}

func (c *Client) clearSBEBookLevels() {
	c.sbeMarketLk.Lock()
	c.sbeBookLevels = make(map[string]map[sbeBookLevel]struct{})
	c.sbeMarketLk.Unlock()
}

// resolveSBEMarketData restores the MDReqID omitted from SBE market responses.
// It also derives NEW/CHANGE/DELETE using the locally tracked price levels.
func (c *Client) resolveSBEMarketData(inbound *inboundMessage, response message.Response) error {
	if inbound.templateID == 0 {
		return nil
	}

	symbol := inbound.marketSymbol
	switch resp := response.(type) {
	case *message.MarketDataSnapshot:
		if symbol == "" {
			symbol = resp.Symbol
		}
	case *message.MarketDataIncrementalRefresh:
		if symbol == "" && len(resp.Entries) > 0 {
			symbol = resp.Entries[0].Symbol
		}
	default:
		return nil
	}
	if symbol == "" {
		return fmt.Errorf("SBE market response template %d is missing Symbol", inbound.templateID)
	}

	c.sbeMarketLk.Lock()
	defer c.sbeMarketLk.Unlock()
	id, ok := c.sbeMarketRoutes[sbeMarketRoute{symbol: symbol, templateID: inbound.templateID}]
	if !ok {
		return fmt.Errorf("no SBE market route for symbol=%s template=%d", symbol, inbound.templateID)
	}

	switch resp := response.(type) {
	case *message.MarketDataSnapshot:
		resp.MDReqID = id
		resp.SendingTime = inbound.sendingTime
		levels := c.sbeBookLevels[id]
		if levels == nil {
			levels = make(map[sbeBookLevel]struct{}, len(resp.Entries))
			c.sbeBookLevels[id] = levels
		}
		// 一个 MDReqID 可以包含多个交易对；新快照只替换当前 Symbol 的档位。
		for level := range levels {
			if level.symbol == symbol {
				delete(levels, level)
			}
		}
		for _, entry := range resp.Entries {
			levels[sbeBookLevel{symbol: symbol, entryType: entry.MDEntryType, price: entry.MDEntryPx}] = struct{}{}
		}
	case *message.MarketDataIncrementalRefresh:
		resp.MDReqID = id
		resp.SendingTime = inbound.sendingTime
		if inbound.templateID == sbeTemplateMarketDataIncrementalTrade {
			return nil
		}
		levels := c.sbeBookLevels[id]
		if levels == nil {
			levels = make(map[sbeBookLevel]struct{})
			c.sbeBookLevels[id] = levels
		}
		for i := range resp.Entries {
			entry := &resp.Entries[i]
			level := sbeBookLevel{symbol: symbol, entryType: entry.MDEntryType, price: entry.MDEntryPx}
			if entry.MDUpdateAction == message.MDUpdateActionDelete {
				delete(levels, level)
				continue
			}
			if _, exists := levels[level]; exists {
				entry.MDUpdateAction = message.MDUpdateActionChange
			} else {
				entry.MDUpdateAction = message.MDUpdateActionNew
				levels[level] = struct{}{}
			}
		}
	}
	return nil
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

// writeRequest 根据连接模式编码并写入一条完整请求。
func (c *Client) writeRequest(req message.Request, seqNum uint32, sendingTime time.Time) error {
	if encodingModeUsesSBERequests(c.config.EncodingMode) {
		if c.sbeRequestMarshaller == nil {
			c.sbeRequestMarshaller = fixsbe.NewSbeGoMarshaller()
		}
		frame, err := encodeSBERequest(
			c.sbeRequestMarshaller,
			req,
			c.config.ClientName,
			targetCompId,
			seqNum,
			sendingTime,
		)
		if err != nil {
			return err
		}
		return c.writeBytes(frame)
	}

	msg, err := req.ToMessage(c.config.ClientName, targetCompId, seqNum, sendingTime)
	if err != nil {
		return err
	}
	return c.writeMessage(msg)
}

// writeMessage 设置单次写超时并一次性写入完整文本 FIX 消息。
func (c *Client) writeMessage(msg *message.Message) error {
	return c.writeBytes([]byte(msg.RawMessage()))
}

func (c *Client) writeBytes(data []byte) error {
	if err := c.conn.SetWriteDeadline(time.Now().Add(c.config.WriteTimeout)); err != nil {
		return err
	}
	if _, err := c.conn.Write(data); err != nil {
		return err
	}
	return nil
}

func (c *Client) readIncomingMessage() (*inboundMessage, error) {
	if encodingModeUsesSBEResponses(c.config.EncodingMode) {
		if err := c.conn.SetReadDeadline(time.Now().Add(c.config.HeartbeatInterval)); err != nil {
			return nil, fmt.Errorf("failed to set read deadline: %w", err)
		}
		if c.sbeMarshaller == nil {
			c.sbeMarshaller = fixsbe.NewSbeGoMarshaller()
		}
		return readSBEMessage(c.reader, c.sbeMarshaller)
	}

	msg, err := c.readMessage()
	if err != nil {
		return nil, err
	}
	return decodeFIXInboundMessage(msg)
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
