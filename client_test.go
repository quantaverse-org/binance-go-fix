package binance_go_fix

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"binance-go-fix/message"
)

func TestClientConfigWithMethods(t *testing.T) {
	apiKey := &ApiKey{UserName: "api-key"}

	config := NewClientConfig(apiKey).
		WithEnableNotify().
		WithClientName("TEST").
		WithHeartbeatInterval(10 * time.Second).
		WithReconnectInterval(2 * time.Second).
		WithResponseTimeout(5 * time.Second).
		WithWriteTimeout(3 * time.Second).
		WithResponseMode(message.ResponseModeOnlyAcks)

	if !config.EnableNotify {
		t.Fatal("EnableNotify = false, want true")
	}
	if config.ClientName != "TEST" {
		t.Fatalf("ClientName = %q, want %q", config.ClientName, "TEST")
	}
	if config.HeartbeatInterval != 10*time.Second {
		t.Fatalf("HeartbeatInterval = %v, want %v", config.HeartbeatInterval, 10*time.Second)
	}
	if config.ReconnectInterval != 2*time.Second {
		t.Fatalf("ReconnectInterval = %v, want %v", config.ReconnectInterval, 2*time.Second)
	}
	if config.ResponseTimeout != 5*time.Second {
		t.Fatalf("ResponseTimeout = %v, want %v", config.ResponseTimeout, 5*time.Second)
	}
	if config.WriteTimeout != 3*time.Second {
		t.Fatalf("WriteTimeout = %v, want %v", config.WriteTimeout, 3*time.Second)
	}
	if config.ResponseMode != message.ResponseModeOnlyAcks {
		t.Fatalf("ResponseMode = %v, want %s", config.ResponseMode, message.ResponseModeOnlyAcks)
	}
	if config.ApiKey != apiKey {
		t.Fatal("ApiKey was changed")
	}
}

func TestDispatchMessageSendsMarketDataIncrementalUpdate(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=X|52=20241019-05:40:11.466313|262=md-1|268=1|279=0|270=1|269=2|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	senders, subscription := initMarketSubscription(1)
	client := &Client{
		config:       &ClientConfig{ClientName: "test-client"},
		subscription: senders,
		respChannels: make(map[string]*responseWaiter),
	}
	ch := make(chan responseResult, 1)
	client.registerRespWaiter([]string{"md-1"}, ch)

	dispatched := make(chan error, 1)
	go func() {
		reconnect, dispatchErr := client.dispatchMessage(ctx, msg)
		if dispatchErr == nil && reconnect {
			dispatchErr = errors.New("dispatchMessage() reconnect = true, want false")
		}
		dispatched <- dispatchErr
	}()

	select {
	case update := <-subscription.MarketData:
		resp, ok := update.(*message.MarketDataIncrementalRefresh)
		if !ok {
			t.Fatalf("market data update type = %T, want *message.MarketDataIncrementalRefresh", update)
		}
		if resp.MDReqID != "md-1" {
			t.Fatalf("MDReqID = %q, want %q", resp.MDReqID, "md-1")
		}
		if len(resp.Entries) != 1 || resp.Entries[0].MDEntryType != message.MDEntryTypeTrade {
			t.Fatalf("Entries = %+v", resp.Entries)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for market data incremental update")
	}
	result := readClientResponseResult(t, ch)
	if result.err != nil {
		t.Fatalf("response error = %v", result.err)
	}
	if _, ok := result.response.(*message.MarketDataIncrementalRefresh); !ok {
		t.Fatalf("response type = %T, want *message.MarketDataIncrementalRefresh", result.response)
	}
	if err := <-dispatched; err != nil {
		t.Fatal(err)
	}
}

func TestDispatchMessageSendsOrderExecutionUpdate(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=8|34=2|49=SPOT|52=20240611-09:01:46.228950|56=EXAMPLE|11=order-1|14=0|17=144|32=0|37=76|39=0|40=2|54=1|55=LTCBNB|150=0|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	senders, subscription := initOrderSubscription(1)
	client := &Client{
		config:       &ClientConfig{ClientName: "test-client"},
		subscription: senders,
	}

	dispatched := make(chan error, 1)
	go func() {
		reconnect, dispatchErr := client.dispatchMessage(ctx, msg)
		if dispatchErr == nil && reconnect {
			dispatchErr = errors.New("dispatchMessage() reconnect = true, want false")
		}
		dispatched <- dispatchErr
	}()

	select {
	case resp := <-subscription.OrderExecution:
		if resp.ClOrdID != "order-1" {
			t.Fatalf("ClOrdID = %q, want %q", resp.ClOrdID, "order-1")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for order execution update")
	}
	if err := <-dispatched; err != nil {
		t.Fatal(err)
	}
}

func TestDispatchMessageUnexpectedType(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=UNKNOWN|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	client := new(Client)
	if _, err := client.dispatchMessage(context.Background(), msg); err == nil {
		t.Fatal("dispatchMessage() error = nil, want error")
	}
}

func TestDispatchMessageDeliversResponseChannel(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=8|34=2|49=SPOT|52=20240611-09:01:46.228950|56=EXAMPLE|11=order-1|14=0|17=144|32=0|37=76|39=0|40=2|54=1|55=LTCBNB|150=0|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	client := &Client{
		config:       &ClientConfig{ClientName: "test-client"},
		respChannels: make(map[string]*responseWaiter),
	}
	ch := make(chan responseResult, 1)
	client.registerRespWaiter([]string{"order-1"}, ch)

	reconnect, err := client.dispatchMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("dispatchMessage() error = %v", err)
	}
	if reconnect {
		t.Fatal("dispatchMessage() reconnect = true, want false")
	}

	result := readClientResponseResult(t, ch)
	if result.err != nil {
		t.Fatalf("response error = %v", result.err)
	}
	resp, ok := result.response.(*message.ExecutionReport)
	if !ok {
		t.Fatalf("response type = %T, want *message.ExecutionReport", result.response)
	}
	if resp.ClOrdID != "order-1" {
		t.Fatalf("ClOrdID = %q, want %q", resp.ClOrdID, "order-1")
	}
}

func TestDeliverResponseCompletesAllAliases(t *testing.T) {
	client := &Client{respChannels: make(map[string]*responseWaiter)}
	ch := make(chan responseResult, 1)
	client.registerRespWaiter([]string{"new-order", "cancel-order"}, ch)

	if delivered := client.deliverResponse("cancel-order", new(message.ExecutionReport)); !delivered {
		t.Fatal("first response was not delivered")
	}
	if len(client.respChannels) != 0 {
		t.Fatalf("len(respChannels) = %d, want 0", len(client.respChannels))
	}
	if delivered := client.deliverResponse("new-order", new(message.ExecutionReport)); delivered {
		t.Fatal("second response was delivered to completed waiter")
	}
}

func TestDispatchMessageMarketDataCompletesWaiterAndSendsEveryUpdate(t *testing.T) {
	rawMessages := []string{
		"8=FIX.4.4|9=1|35=W|52=20241019-05:41:52.867164|262=md-1|55=BTCUSDT|268=1|269=0|270=1|271=2|10=000|",
		"8=FIX.4.4|9=1|35=W|52=20241019-05:41:52.867165|262=md-1|55=ETHUSDT|268=1|269=1|270=3|271=4|10=000|",
	}

	senders, subscription := initMarketSubscription(2)
	client := &Client{
		config:       &ClientConfig{ClientName: "test-client"},
		subscription: senders,
		respChannels: make(map[string]*responseWaiter),
	}
	ch := make(chan responseResult, 1)
	client.registerRespWaiter([]string{"md-1"}, ch)

	for i, raw := range rawMessages {
		msg, err := message.ParseMessage(clientTestSOH(raw))
		if err != nil {
			t.Fatalf("ParseMessage(%d) error = %v", i, err)
		}
		if _, err = client.dispatchMessage(context.Background(), msg); err != nil {
			t.Fatalf("dispatchMessage(%d) error = %v", i, err)
		}
	}
	if _, registered := client.respChannels["md-1"]; registered {
		t.Fatal("response waiter was not removed after the first market data message")
	}

	result := readClientResponseResult(t, ch)
	if result.err != nil {
		t.Fatalf("response error = %v", result.err)
	}
	snapshot, ok := result.response.(*message.MarketDataSnapshot)
	if !ok {
		t.Fatalf("response type = %T, want *message.MarketDataSnapshot", result.response)
	}
	if snapshot.Symbol != "BTCUSDT" {
		t.Fatalf("response Symbol = %q, want %q", snapshot.Symbol, "BTCUSDT")
	}

	for _, want := range []string{"BTCUSDT", "ETHUSDT"} {
		update := <-subscription.MarketData
		snapshot, ok := update.(*message.MarketDataSnapshot)
		if !ok {
			t.Fatalf("market data update type = %T, want *message.MarketDataSnapshot", update)
		}
		if snapshot.Symbol != want {
			t.Fatalf("update Symbol = %q, want %q", snapshot.Symbol, want)
		}
		if len(snapshot.Entries) != 1 {
			t.Fatalf("update Entries = %+v", snapshot.Entries)
		}
	}
}

func TestDispatchMessageDeliversSessionRejectAsError(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=3|52=20240611-09:01:46.228950|45=7|58=Bad request|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	client := &Client{
		config:         &ClientConfig{ClientName: "test-client"},
		rejectChannels: make(map[uint32]chan responseResult),
	}
	ch := make(chan responseResult, 1)
	client.registerRejectChannel(7, ch)

	reconnect, err := client.dispatchMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("dispatchMessage() error = %v", err)
	}
	if reconnect {
		t.Fatal("dispatchMessage() reconnect = true, want false")
	}

	result := readClientResponseResult(t, ch)
	reject, ok := result.err.(*message.Reject)
	if !ok {
		t.Fatalf("response error = %T, want *message.Reject", result.err)
	}
	if reject.RefSeqNum == nil || *reject.RefSeqNum != 7 {
		t.Fatalf("RefSeqNum = %v, want 7", reject.RefSeqNum)
	}
}

func TestDispatchMessageMarketDataRejectRemovesResubRequest(t *testing.T) {
	msg, err := message.ParseMessage(clientTestSOH("8=FIX.4.4|9=1|35=Y|52=20240611-09:01:46.228950|262=md-1|58=Bad request|25016=-1102|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	client := &Client{
		config:       &ClientConfig{ClientName: "test-client"},
		respChannels: make(map[string]*responseWaiter),
		resubReqs: map[string]message.Request{
			"md-1": message.NewMarketDataRequest("md-1", message.SubscriptionRequestTypeSubscribe),
		},
	}
	ch := make(chan responseResult, 1)
	client.registerRespWaiter([]string{"md-1"}, ch)

	reconnect, err := client.dispatchMessage(context.Background(), msg)
	if err != nil {
		t.Fatalf("dispatchMessage() error = %v", err)
	}
	if reconnect {
		t.Fatal("dispatchMessage() reconnect = true, want false")
	}

	result := readClientResponseResult(t, ch)
	if _, ok := result.err.(*message.MarketDataRequestReject); !ok {
		t.Fatalf("response error = %T, want *message.MarketDataRequestReject", result.err)
	}
	if _, ok := client.resubReqs["md-1"]; ok {
		t.Fatal("resub request was not removed")
	}
	if _, ok := client.respChannels["md-1"]; ok {
		t.Fatal("response waiter was not removed after reject")
	}
}

func clientTestSOH(msg string) string {
	return strings.ReplaceAll(msg, "|", string(message.SOH))
}

func readClientResponseResult(t *testing.T, ch <-chan responseResult) responseResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	select {
	case result := <-ch:
		return result
	case <-ctx.Done():
		t.Fatal("timed out waiting for response channel")
		return responseResult{}
	}
}
