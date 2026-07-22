package message

import (
	"strings"
	"testing"
	"time"
)

func TestNewOrderSingleToMessage(t *testing.T) {
	sor := true
	request := NewNewOrderSingle("order-1", OrdTypeLimit, SideBuy, "LTCBNB")
	request.OrderQty = 5
	request.Price = 10
	request.TimeInForce = TimeInForceFillOrKill
	request.ExecInst = ExecInstParticipateDontInitiate
	request.MaxFloor = 1
	request.CashOrderQty = 50
	request.TargetStrategy = 1000000
	request.StrategyID = 42
	request.SelfTradePreventionMode = SelfTradePreventionModeExpireTaker
	request.PegOffsetValue = 2
	request.PegPriceType = PegPriceTypeMarketPeg
	request.PegMoveType = PegMoveTypeFixed
	request.PegOffsetType = PegOffsetTypePriceTier
	request.TriggerType = TriggerTypePriceMovement
	request.TriggerAction = TriggerActionActivate
	request.TriggerPrice = 9
	request.TriggerPriceType = TriggerPriceTypeLastTrade
	request.TriggerPriceDirection = TriggerPriceDirectionUp
	request.TriggerTrailingDeltaBips = 100
	request.SOR = &sor

	message, err := request.ToMessage("EXAMPLE", "SPOT", 2, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}

	assertField(t, message, TagMsgType, string(MsgTypeNewOrderSingle))
	assertField(t, message, TagSenderCompID, "EXAMPLE")
	assertField(t, message, TagTargetCompID, "SPOT")
	assertField(t, message, TagMsgSeqNum, "2")
	assertField(t, message, TagSendingTime, "20240627-11:17:25.223")
	assertField(t, message, TagClOrdID, "order-1")
	assertField(t, message, TagOrderQty, "5")
	assertField(t, message, TagOrdType, "2")
	assertField(t, message, TagExecInst, "6")
	assertField(t, message, TagPrice, "10")
	assertField(t, message, TagSide, "1")
	assertField(t, message, TagSymbol, "LTCBNB")
	assertField(t, message, TagTimeInForce, "4")
	assertField(t, message, TagMaxFloor, "1")
	assertField(t, message, TagCashOrderQty, "50")
	assertField(t, message, TagTargetStrategy, "1000000")
	assertField(t, message, TagStrategyID, "42")
	assertField(t, message, TagSelfTradePreventionMode, "2")
	assertField(t, message, TagPegOffsetValue, "2")
	assertField(t, message, TagPegPriceType, "4")
	assertField(t, message, TagPegMoveType, "1")
	assertField(t, message, TagPegOffsetType, "3")
	assertField(t, message, TagTriggerType, "4")
	assertField(t, message, TagTriggerAction, "1")
	assertField(t, message, TagTriggerPrice, "9")
	assertField(t, message, TagTriggerPriceType, "2")
	assertField(t, message, TagTriggerPriceDirection, "U")
	assertField(t, message, TagTriggerTrailingDeltaBips, "100")
	assertField(t, message, TagSOR, "Y")
}

func TestExecutionReportFromMessage(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=8|34=2|49=SPOT|52=20240611-09:01:46.228950|56=EXAMPLE|11=order-1|14=1.00000000|17=144|32=0.50000000|37=76|38=5.00000000|39=1|40=2|44=10.00000000|54=1|55=LTCBNB|59=4|60=20240611-09:01:46.228000|150=F|151=4.00000000|636=Y|1057=Y|25001=1|25017=5.00000000|25018=1718096506228|25023=20240611-09:01:46.228000|25032=Y|136=2|137=0.001|138=BNB|139=4|137=0.002|138=USDT|139=4|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	report := new(ExecutionReport)
	if err := report.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}

	if report.ClOrdID != "order-1" {
		t.Fatalf("ClOrdID = %q, want %q", report.ClOrdID, "order-1")
	}
	if report.OrderID != 76 {
		t.Fatalf("OrderID = %d, want %d", report.OrderID, 76)
	}
	if report.OrdType != OrdTypeLimit {
		t.Fatalf("OrdType = %q, want %q", report.OrdType, OrdTypeLimit)
	}
	if report.Side != SideBuy {
		t.Fatalf("Side = %q, want %q", report.Side, SideBuy)
	}
	if report.Symbol != "LTCBNB" {
		t.Fatalf("Symbol = %q, want %q", report.Symbol, "LTCBNB")
	}
	if report.ExecType != ExecTypeTrade {
		t.Fatalf("ExecType = %q, want %q", report.ExecType, ExecTypeTrade)
	}
	if report.OrderQty != 5 || report.Price != 10 {
		t.Fatalf("OrderQty = %v, Price = %v, want 5 and 10", report.OrderQty, report.Price)
	}
	if got, want := report.TransactTime, mustParseTimestamp(t, "20240611-09:01:46.228000"); !got.Equal(want) {
		t.Fatalf("TransactTime = %v, want %v", got, want)
	}
	if report.OrderCreationTime != 1718096506228 {
		t.Fatalf("OrderCreationTime = %d, want %d", report.OrderCreationTime, int64(1718096506228))
	}
	if report.CumQty != 1 {
		t.Fatalf("CumQty = %v, want %v", report.CumQty, 1.0)
	}
	if report.LastQty != 0.5 {
		t.Fatalf("LastQty = %v, want %v", report.LastQty, 0.5)
	}
	if report.LeavesQty != 4 || report.CumQuoteQty != 5 {
		t.Fatalf("LeavesQty = %v, CumQuoteQty = %v, want 4 and 5", report.LeavesQty, report.CumQuoteQty)
	}
	if !boolPointerEqual(report.AggressorIndicator, true) {
		t.Fatalf("AggressorIndicator = %v, want true", report.AggressorIndicator)
	}
	if report.OrdStatus != OrdStatusPartiallyFilled {
		t.Fatalf("OrdStatus = %q, want %q", report.OrdStatus, OrdStatusPartiallyFilled)
	}
	if !boolPointerEqual(report.WorkingIndicator, true) {
		t.Fatalf("WorkingIndicator = %v, want true", report.WorkingIndicator)
	}
	if !boolPointerEqual(report.SOR, true) {
		t.Fatalf("SOR = %v, want true", report.SOR)
	}
	if report.NoMiscFees != 2 || len(report.MiscFees) != 2 {
		t.Fatalf("NoMiscFees = %d, len(MiscFees) = %d, want 2", report.NoMiscFees, len(report.MiscFees))
	}
	if got := report.MiscFees[0]; got.MiscFeeAmt != 0.001 || got.MiscFeeCurr != "BNB" || got.MiscFeeType != MiscFeeTypeExchangeFees {
		t.Fatalf("MiscFees[0] = %+v", got)
	}
	if got := report.MiscFees[1]; got.MiscFeeAmt != 0.002 || got.MiscFeeCurr != "USDT" || got.MiscFeeType != MiscFeeTypeExchangeFees {
		t.Fatalf("MiscFees[1] = %+v", got)
	}
}

func TestOrderCancelRejectFromMessage(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=9|34=2|49=SPOT|52=20240613-01:12:41.320869|56=EXAMPLE|11=cancel-1|37=2|55=LTCBNB|58=Unknown order sent.|434=1|25016=-1013|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	reject := new(OrderCancelReject)
	if err := reject.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}

	if reject.ClOrdID != "cancel-1" {
		t.Fatalf("ClOrdID = %q, want %q", reject.ClOrdID, "cancel-1")
	}
	if reject.OrderID != 2 {
		t.Fatalf("OrderID = %d, want %d", reject.OrderID, 2)
	}
	if reject.Symbol != "LTCBNB" {
		t.Fatalf("Symbol = %q, want %q", reject.Symbol, "LTCBNB")
	}
	if reject.CxlRejResponseTo != CxlRejResponseToOrderCancelRequest {
		t.Fatalf("CxlRejResponseTo = %q, want %q", reject.CxlRejResponseTo, CxlRejResponseToOrderCancelRequest)
	}
	if reject.ErrorCode != -1013 {
		t.Fatalf("ErrorCode = %d, want %d", reject.ErrorCode, -1013)
	}
	if reject.Text != "Unknown order sent." {
		t.Fatalf("Text = %q, want %q", reject.Text, "Unknown order sent.")
	}
}

func TestNewOrderListToMessageRepeatingOrders(t *testing.T) {
	request := NewNewOrderList("list-1", []NewOrderListOrder{
		{
			OrderFields: OrderFields{
				ClOrdID:     "working-1",
				OrderQty:    1,
				OrdType:     OrdTypeLimit,
				Price:       0.25,
				Side:        SideSell,
				Symbol:      "LTCBNB",
				TimeInForce: TimeInForceGoodTillCancel,
			},
		},
		{
			OrderFields: OrderFields{
				ClOrdID:  "pending-1",
				OrderQty: 1,
				OrdType:  OrdTypeMarket,
				Side:     SideSell,
				Symbol:   "LTCBNB",
			},
			ListTriggeringInstructions: []ListTriggeringInstruction{
				{
					ListTriggerType:         ListTriggerTypeFilled,
					ListTriggerTriggerIndex: 0,
					ListTriggerAction:       ListTriggerActionRelease,
				},
			},
		},
	})
	request.ContingencyType = ContingencyTypeOneTriggersTheOther

	message, err := request.ToMessage("EXAMPLE", "SPOT", 3, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}

	assertField(t, message, TagMsgType, string(MsgTypeNewOrderList))
	assertField(t, message, TagClListID, "list-1")
	assertField(t, message, TagContingencyType, "2")
	assertField(t, message, TagNoOrders, "2")

	display := message.Display()
	assertDisplayContains(t, display, "73=2|11=working-1|38=1|40=2|44=0.25|54=2|55=LTCBNB|59=1|")
	assertDisplayContains(t, display, "|11=pending-1|38=1|40=1|54=2|55=LTCBNB|25010=1|25011=3|25012=0|25013=1|")
	if got := strings.Count(display, "|11="); got != 2 {
		t.Fatalf("ClOrdID count = %d, want %d in %q", got, 2, display)
	}
}

func TestListStatusFromMessageRepeatingOrders(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=N|34=2|49=SPOT|52=20240607-02:19:07.837191|56=EXAMPLE|60=20240607-02:19:07.836000|66=25|73=2|55=BTCUSDT|37=52|11=working-1|55=BTCUSDT|37=53|11=pending-1|25010=2|25011=3|25012=0|25013=1|25011=1|25012=2|25013=2|429=4|431=3|1385=2|25014=list-1|25015=orig-list-1|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	response := new(ListStatus)
	if err = response.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if response.Symbol != "" {
		t.Fatalf("top-level Symbol = %q, want empty", response.Symbol)
	}
	if response.NoOrders != 2 || len(response.Orders) != 2 {
		t.Fatalf("NoOrders = %d, len(Orders) = %d, want 2", response.NoOrders, len(response.Orders))
	}
	if got := response.Orders[0]; got.Symbol != "BTCUSDT" || got.OrderID != 52 || got.ClOrdID != "working-1" {
		t.Fatalf("Orders[0] = %+v", got)
	}
	second := response.Orders[1]
	if second.Symbol != "BTCUSDT" || second.OrderID != 53 || second.ClOrdID != "pending-1" {
		t.Fatalf("Orders[1] = %+v", second)
	}
	if second.NoListTriggeringInstructions != 2 || len(second.ListTriggeringInstructions) != 2 {
		t.Fatalf("Orders[1] nested count = %d, len = %d, want 2", second.NoListTriggeringInstructions, len(second.ListTriggeringInstructions))
	}
	if got := second.ListTriggeringInstructions[0]; got.ListTriggerType != ListTriggerTypeFilled || got.ListTriggerTriggerIndex != 0 || got.ListTriggerAction != ListTriggerActionRelease {
		t.Fatalf("Orders[1].ListTriggeringInstructions[0] = %+v", got)
	}
	if got := second.ListTriggeringInstructions[1]; got.ListTriggerType != ListTriggerTypeActivated || got.ListTriggerTriggerIndex != 2 || got.ListTriggerAction != ListTriggerActionCancel {
		t.Fatalf("Orders[1].ListTriggeringInstructions[1] = %+v", got)
	}
}

func TestMarketDataRequestToMessageRepeatingGroups(t *testing.T) {
	aggregatedBook := true
	request := NewMarketDataRequest("BOOK_TICKER_STREAM", SubscriptionRequestTypeSubscribe)
	request.MarketDepth = 1
	request.AggregatedBook = &aggregatedBook
	request.Symbols = []string{"BTCUSDT", "ETHUSDT"}
	request.MDEntryTypes = []MDEntryType{MDEntryTypeBid, MDEntryTypeOffer}

	message, err := request.ToMessage("EXAMPLE", "SPOT", 4, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}

	assertField(t, message, TagMsgType, string(MsgTypeMarketDataRequest))
	assertField(t, message, TagMDReqID, "BOOK_TICKER_STREAM")
	assertField(t, message, TagSubscriptionRequestType, "1")
	assertField(t, message, TagMarketDepth, "1")
	assertField(t, message, TagAggregatedBook, "Y")
	assertField(t, message, TagNoRelatedSym, "2")
	assertField(t, message, TagNoMDEntryTypes, "2")

	display := message.Display()
	assertDisplayContains(t, display, "146=2|55=BTCUSDT|55=ETHUSDT|")
	assertDisplayContains(t, display, "267=2|269=0|269=1|")
	if got := strings.Count(display, "|55="); got != 2 {
		t.Fatalf("Symbol count = %d, want %d in %q", got, 2, display)
	}
	if got := strings.Count(display, "|269="); got != 2 {
		t.Fatalf("MDEntryType count = %d, want %d in %q", got, 2, display)
	}
}

func TestMarketDataSnapshotFromMessageRepeatingGroup(t *testing.T) {
	msg, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=W|52=20241019-05:41:52.867164|262=DEPTH_1|55=BTCUSDT|25044=100|268=2|269=0|270=65000.10|271=1.25|269=1|270=65000.20|271=2.50|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	response := new(MarketDataSnapshot)
	if err = response.FromMessage(msg); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if response.NoMDEntries != 2 || len(response.Entries) != 2 {
		t.Fatalf("NoMDEntries = %d, len(Entries) = %d, want 2", response.NoMDEntries, len(response.Entries))
	}
	if got, want := response.SendingTime, mustParseTimestamp(t, "20241019-05:41:52.867164"); !got.Equal(want) {
		t.Fatalf("SendingTime = %v, want %v", got, want)
	}
	if got := response.Entries[0]; got.MDEntryType != MDEntryTypeBid || got.MDEntryPx != 65000.10 || got.MDEntrySize != 1.25 {
		t.Fatalf("Entries[0] = %+v", got)
	}
	if got := response.Entries[1]; got.MDEntryType != MDEntryTypeOffer || got.MDEntryPx != 65000.20 || got.MDEntrySize != 2.50 {
		t.Fatalf("Entries[1] = %+v", got)
	}
}

func TestMarketDataIncrementalRefreshFromMessageRepeatingGroup(t *testing.T) {
	msg, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=X|52=20241019-05:40:11.466313|262=TRADE_3|268=3|279=0|269=2|270=10.00000|271=0.01000|55=BNBBUSD|1003=0|60=20241019-05:40:11.464000|2446=1|25043=100|25044=102|279=0|269=2|270=10.00000|271=0.02000|1003=1|60=20241019-05:40:11.465000|279=0|269=2|270=10.00000|271=0.03000|1003=2|60=20241019-05:40:11.466000|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	response := new(MarketDataIncrementalRefresh)
	if err = response.FromMessage(msg); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if response.NoMDEntries != 3 || len(response.Entries) != 3 {
		t.Fatalf("NoMDEntries = %d, len(Entries) = %d, want 3", response.NoMDEntries, len(response.Entries))
	}
	if got, want := response.SendingTime, mustParseTimestamp(t, "20241019-05:40:11.466313"); !got.Equal(want) {
		t.Fatalf("SendingTime = %v, want %v", got, want)
	}
	first := response.Entries[0]
	if first.MDUpdateAction != MDUpdateActionNew || first.MDEntryType != MDEntryTypeTrade || first.AggressorSide != AggressorSideBuy {
		t.Fatalf("Entries[0] = %+v", first)
	}
	if got, want := first.TransactTime, mustParseTimestamp(t, "20241019-05:40:11.464000"); !got.Equal(want) {
		t.Fatalf("Entries[0].TransactTime = %v, want %v", got, want)
	}
	if first.Symbol != "BNBBUSD" || first.FirstBookUpdateID != 100 || first.LastBookUpdateID != 102 {
		t.Fatalf("Entries[0] inherited fields = %+v", first)
	}
	for i, entry := range response.Entries[1:] {
		if entry.Symbol != "BNBBUSD" || entry.FirstBookUpdateID != 100 || entry.LastBookUpdateID != 102 {
			t.Fatalf("Entries[%d] inherited fields = %+v", i+1, entry)
		}
	}
	if response.Entries[2].TradeID != 2 || response.Entries[2].MDEntrySize != 0.03 {
		t.Fatalf("Entries[2] = %+v", response.Entries[2])
	}
	if got, want := response.Entries[2].TransactTime, mustParseTimestamp(t, "20241019-05:40:11.466000"); !got.Equal(want) {
		t.Fatalf("Entries[2].TransactTime = %v, want %v", got, want)
	}
}

func TestMarketDataIncrementalRefreshRejectsInvalidTransactTime(t *testing.T) {
	msg, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=X|52=20260722-02:25:50.143785|262=TRADE|268=1|279=0|269=2|270=1|60=invalid|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}
	if err = new(MarketDataIncrementalRefresh).FromMessage(msg); err == nil {
		t.Fatal("FromMessage() error = nil, want invalid TransactTime error")
	}
}

func TestMarketDataSnapshotFromMessageRejectsGroupCountMismatch(t *testing.T) {
	msg, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=W|52=20241019-05:41:52.867164|262=DEPTH_1|55=BTCUSDT|268=2|269=0|270=65000.10|271=1.25|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}
	if err = new(MarketDataSnapshot).FromMessage(msg); err == nil {
		t.Fatal("FromMessage() error = nil, want group count mismatch")
	}
}

func TestLimitResponseFromMessage(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=XLR|34=2|49=SPOT|52=20240614-05:42:42.724057|56=EXAMPLE|6136=req-1|25003=2|25004=2|25005=1|25006=1000|25007=10|25008=s|25004=1|25005=0|25006=200|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	response := new(LimitResponse)
	if err := response.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}

	if response.ReqID != "req-1" {
		t.Fatalf("ReqID = %q, want %q", response.ReqID, "req-1")
	}
	if response.NoLimitIndicators != 2 {
		t.Fatalf("NoLimitIndicators = %d, want %d", response.NoLimitIndicators, 2)
	}
	if len(response.LimitIndicators) != 2 {
		t.Fatalf("len(LimitIndicators) = %d, want 2", len(response.LimitIndicators))
	}
	if got := response.LimitIndicators[0]; got.LimitType != LimitTypeMessage || got.LimitCount != 1 || got.LimitMax != 1000 || got.LimitResetInterval != 10 || got.LimitResetIntervalResolution != LimitResetIntervalResolutionSecond {
		t.Fatalf("LimitIndicators[0] = %+v", got)
	}
	if got := response.LimitIndicators[1]; got.LimitType != LimitTypeOrder || got.LimitCount != 0 || got.LimitMax != 200 {
		t.Fatalf("LimitIndicators[1] = %+v", got)
	}
}

func TestInstrumentListRequestToMessage(t *testing.T) {
	message, err := NewInstrumentListRequest("ALL_INFO", InstrumentListRequestTypeAllInstruments).
		ToMessage("EXAMPLE", "SPOT", 5, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}

	assertField(t, message, TagMsgType, string(MsgTypeInstrumentListRequest))
	assertField(t, message, TagInstrumentReqID, "ALL_INFO")
	assertField(t, message, TagInstrumentListRequestType, "4")
	if _, ok := message.GetField(TagSymbol); ok {
		t.Fatal("Symbol field exists, want absent for all-instruments request")
	}

	request := NewInstrumentListRequest("BTCUSDT_INFO", InstrumentListRequestTypeSingleInstrument)
	request.Symbol = "BTCUSDT"
	message, err = request.ToMessage("EXAMPLE", "SPOT", 6, exampleSendingTime())
	if err != nil {
		t.Fatalf("ToMessage() error = %v", err)
	}
	assertField(t, message, TagSymbol, "BTCUSDT")
}

func TestInstrumentListFromMessageRepeatingGroup(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=y|49=SPOT|56=EXAMPLE|34=2|52=20250114-08:46:56.100147|320=ALL_INFO|146=2|55=BTCUSDT|15=USDT|562=0.00001000|1140=9000.00000000|25039=0.00001000|25040=0.00000001|25041=76.79001236|25042=0.00000001|969=0.01000000|2551=0.01|2552=1000000|55=ETHBTC|15=BTC|562=0.00010000|1140=100000.00000000|25039=0.00010000|969=0.00000001|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	response := new(InstrumentList)
	if err = response.FromMessage(message); err != nil {
		t.Fatalf("FromMessage() error = %v", err)
	}
	if response.InstrumentReqID != "ALL_INFO" || response.NoRelatedSym != 2 || len(response.Instruments) != 2 {
		t.Fatalf("response = %+v", response)
	}
	first := response.Instruments[0]
	if first.Symbol != "BTCUSDT" || first.Currency != "USDT" || first.MinTradeVol != 0.00001 || first.MaxTradeVol != 9000 {
		t.Fatalf("Instruments[0] = %+v", first)
	}
	if first.MarketMinTradeVol != 0.00000001 || first.MarketMaxTradeVol != 76.79001236 || first.MarketMinQtyIncrement != 0.00000001 {
		t.Fatalf("Instruments[0] market limits = %+v", first)
	}
	if first.MinPriceIncrement != 0.01 || first.StartPriceRange != 0.01 || first.EndPriceRange != 1000000 {
		t.Fatalf("Instruments[0] price limits = %+v", first)
	}
	if got := response.Instruments[1]; got.Symbol != "ETHBTC" || got.Currency != "BTC" || got.MinQtyIncrement != 0.0001 || got.MinPriceIncrement != 0.00000001 {
		t.Fatalf("Instruments[1] = %+v", got)
	}
}

func TestResponseRepeatingGroupsRejectCountMismatch(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		response Response
	}{
		{
			name:     "execution report fees",
			raw:      "8=FIX.4.4|9=1|35=8|40=2|54=1|55=BTCUSDT|150=F|14=1|32=1|39=2|136=2|137=0.1|138=USDT|139=4|10=000|",
			response: new(ExecutionReport),
		},
		{
			name:     "list status orders",
			raw:      "8=FIX.4.4|9=1|35=N|429=4|431=3|73=2|55=BTCUSDT|37=1|11=order-1|10=000|",
			response: new(ListStatus),
		},
		{
			name:     "limit indicators",
			raw:      "8=FIX.4.4|9=1|35=XLR|6136=req-1|25003=2|25004=1|25005=0|25006=200|10=000|",
			response: new(LimitResponse),
		},
		{
			name:     "instruments",
			raw:      "8=FIX.4.4|9=1|35=y|320=req-1|146=2|55=BTCUSDT|15=USDT|10=000|",
			response: new(InstrumentList),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message, err := ParseMessage(withSOH(tt.raw))
			if err != nil {
				t.Fatalf("ParseMessage() error = %v", err)
			}
			if err = tt.response.FromMessage(message); err == nil {
				t.Fatal("FromMessage() error = nil, want group count mismatch")
			}
		})
	}
}

func TestListStatusRejectsNestedGroupCountMismatch(t *testing.T) {
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=N|429=4|431=3|73=1|55=BTCUSDT|37=1|11=order-1|25010=2|25011=3|25012=0|25013=1|10=000|"))
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}
	if err = new(ListStatus).FromMessage(message); err == nil {
		t.Fatal("FromMessage() error = nil, want nested group count mismatch")
	}
}

func TestApplicationResponsesRejectInvalidTypedFields(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		response Response
	}{
		{
			name:     "execution report int",
			raw:      "8=FIX.4.4|9=1|35=8|37=invalid|40=2|54=1|55=BTCUSDT|150=0|14=0|32=0|39=0|10=000|",
			response: new(ExecutionReport),
		},
		{
			name:     "execution report bool",
			raw:      "8=FIX.4.4|9=1|35=8|40=2|54=1|55=BTCUSDT|150=0|14=0|32=0|39=0|636=invalid|10=000|",
			response: new(ExecutionReport),
		},
		{
			name:     "list status timestamp",
			raw:      "8=FIX.4.4|9=1|35=N|429=4|431=3|60=invalid|10=000|",
			response: new(ListStatus),
		},
		{
			name:     "limit response int",
			raw:      "8=FIX.4.4|9=1|35=XLR|6136=req-1|25003=1|25004=1|25005=invalid|25006=200|10=000|",
			response: new(LimitResponse),
		},
		{
			name:     "instrument qty",
			raw:      "8=FIX.4.4|9=1|35=y|320=req-1|146=1|55=BTCUSDT|15=USDT|562=invalid|10=000|",
			response: new(InstrumentList),
		},
		{
			name:     "snapshot price",
			raw:      "8=FIX.4.4|9=1|35=W|52=20241019-05:41:52.867164|262=depth|55=BTCUSDT|268=1|269=0|270=invalid|271=1|10=000|",
			response: new(MarketDataSnapshot),
		},
		{
			name:     "incremental trade id",
			raw:      "8=FIX.4.4|9=1|35=X|52=20241019-05:40:11.466313|262=trade|268=1|279=0|270=1|269=2|1003=invalid|10=000|",
			response: new(MarketDataIncrementalRefresh),
		},
		{
			name:     "market data reject error code",
			raw:      "8=FIX.4.4|9=1|35=Y|262=depth|25016=invalid|10=000|",
			response: new(MarketDataRequestReject),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message, err := ParseMessage(withSOH(tt.raw))
			if err != nil {
				t.Fatalf("ParseMessage() error = %v", err)
			}
			if err = tt.response.FromMessage(message); err == nil {
				t.Fatal("FromMessage() error = nil, want typed field parse error")
			}
		})
	}
}

func TestApplicationRejectErrorsOnlyIncludeErrorCodeAndText(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "execution report",
			err:  &ExecutionReport{ClOrdID: "order-1", ErrorCode: -1013, Text: "Invalid quantity"},
			want: "order rejected, errorCode=-1013, text=Invalid quantity",
		},
		{
			name: "order cancel reject",
			err:  &OrderCancelReject{ClOrdID: "cancel-1", ErrorCode: -2011, Text: "Unknown order"},
			want: "order cancel reject, errorCode=-2011, text=Unknown order",
		},
		{
			name: "order mass cancel reject",
			err:  &OrderMassCancelReport{ClOrdID: "cancel-all-1", ErrorCode: -1102, Text: "Missing symbol"},
			want: "order mass cancel reject, errorCode=-1102, text=Missing symbol",
		},
		{
			name: "list status",
			err:  &ListStatus{ClListID: "list-1", ErrorCode: -1013, Text: "Invalid list"},
			want: "order list reject, errorCode=-1013, text=Invalid list",
		},
		{
			name: "order amend reject",
			err:  &OrderAmendReject{ClOrdID: "amend-1", ErrorCode: -2010, Text: "Amend rejected"},
			want: "order amend reject, errorCode=-2010, text=Amend rejected",
		},
		{
			name: "market data request reject",
			err:  &MarketDataRequestReject{MDReqID: "md-1", ErrorCode: -1191, Text: "Too many subscriptions"},
			want: "market data request reject, errorCode=-1191, text=Too many subscriptions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Fatalf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func assertDisplayContains(t *testing.T, display string, want string) {
	t.Helper()

	if !strings.Contains(display, want) {
		t.Fatalf("Display() = %q, want substring %q", display, want)
	}
}

func mustParseTimestamp(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := ParseTimestamp(value)
	if err != nil {
		t.Fatalf("ParseTimestamp(%q) error = %v", value, err)
	}
	return parsed
}

func boolPointerEqual(got *bool, want bool) bool {
	return got != nil && *got == want
}
