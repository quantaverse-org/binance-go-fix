package message

import (
	"strings"
	"testing"
)

func TestNewOrderSingleToMessage(t *testing.T) {
	sor := true
	request := NewNewOrderSingle("order-1", OrdTypeLimit, SideBuy, "LTCBNB")
	request.OrderQty = "5"
	request.Price = "10"
	request.TimeInForce = TimeInForceFillOrKill
	request.ExecInst = ExecInstParticipateDontInitiate
	request.MaxFloor = "1"
	request.CashOrderQty = "50"
	request.TargetStrategy = "1000000"
	request.StrategyID = "42"
	request.SelfTradePreventionMode = SelfTradePreventionModeExpireTaker
	request.PegOffsetValue = "2"
	request.PegPriceType = PegPriceTypeMarketPeg
	request.PegMoveType = PegMoveTypeFixed
	request.PegOffsetType = PegOffsetTypePriceTier
	request.TriggerType = TriggerTypePriceMovement
	request.TriggerAction = TriggerActionActivate
	request.TriggerPrice = "9"
	request.TriggerPriceType = TriggerPriceTypeLastTrade
	request.TriggerPriceDirection = TriggerPriceDirectionUp
	request.TriggerTrailingDeltaBips = "100"
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
	message, err := ParseMessage(withSOH("8=FIX.4.4|9=1|35=8|34=2|49=SPOT|52=20240611-09:01:46.228950|56=EXAMPLE|11=order-1|14=1.00000000|17=144|32=0.50000000|37=76|38=5.00000000|39=1|40=2|44=10.00000000|54=1|55=LTCBNB|59=4|60=20240611-09:01:46.228000|150=F|151=4.00000000|636=Y|1057=Y|25001=1|25017=5.00000000|25018=20240611-09:01:46.228000|25023=20240611-09:01:46.228000|25032=Y|10=000|"))
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
	if report.OrderID != "76" {
		t.Fatalf("OrderID = %q, want %q", report.OrderID, "76")
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
	if report.CumQty != "1.00000000" {
		t.Fatalf("CumQty = %q, want %q", report.CumQty, "1.00000000")
	}
	if report.LastQty != "0.50000000" {
		t.Fatalf("LastQty = %q, want %q", report.LastQty, "0.50000000")
	}
	if report.OrdStatus != OrdStatusPartiallyFilled {
		t.Fatalf("OrdStatus = %q, want %q", report.OrdStatus, OrdStatusPartiallyFilled)
	}
	if report.WorkingIndicator != "Y" {
		t.Fatalf("WorkingIndicator = %q, want %q", report.WorkingIndicator, "Y")
	}
	if report.SOR != "Y" {
		t.Fatalf("SOR = %q, want %q", report.SOR, "Y")
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
	if reject.OrderID != "2" {
		t.Fatalf("OrderID = %q, want %q", reject.OrderID, "2")
	}
	if reject.Symbol != "LTCBNB" {
		t.Fatalf("Symbol = %q, want %q", reject.Symbol, "LTCBNB")
	}
	if reject.CxlRejResponseTo != CxlRejResponseToOrderCancelRequest {
		t.Fatalf("CxlRejResponseTo = %q, want %q", reject.CxlRejResponseTo, CxlRejResponseToOrderCancelRequest)
	}
	if reject.ErrorCode != "-1013" {
		t.Fatalf("ErrorCode = %q, want %q", reject.ErrorCode, "-1013")
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
				OrderQty:    "1",
				OrdType:     OrdTypeLimit,
				Price:       "0.25",
				Side:        SideSell,
				Symbol:      "LTCBNB",
				TimeInForce: TimeInForceGoodTillCancel,
			},
		},
		{
			OrderFields: OrderFields{
				ClOrdID:  "pending-1",
				OrderQty: "1",
				OrdType:  OrdTypeMarket,
				Side:     SideSell,
				Symbol:   "LTCBNB",
			},
			ListTriggeringInstructions: []ListTriggeringInstruction{
				{
					ListTriggerType:         ListTriggerTypeFilled,
					ListTriggerTriggerIndex: "0",
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

func TestMarketDataRequestToMessageRepeatingGroups(t *testing.T) {
	aggregatedBook := true
	request := NewMarketDataRequest("BOOK_TICKER_STREAM", SubscriptionRequestTypeSubscribe)
	request.MarketDepth = "1"
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
	if response.NoLimitIndicators != "2" {
		t.Fatalf("NoLimitIndicators = %q, want %q", response.NoLimitIndicators, "2")
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

func assertDisplayContains(t *testing.T, display string, want string) {
	t.Helper()

	if !strings.Contains(display, want) {
		t.Fatalf("Display() = %q, want substring %q", display, want)
	}
}
