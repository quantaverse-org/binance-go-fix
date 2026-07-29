package binance_go_fix

import (
	"bytes"
	"crypto/ed25519"
	"encoding/binary"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/quantaverse-org/binance-go-fix/internal/fixsbe"
	"github.com/quantaverse-org/binance-go-fix/message"
)

func TestEncodeSBELogon(t *testing.T) {
	privateKey := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{1}, ed25519.SeedSize))
	request := message.NewLogonRequest(
		"api-key",
		privateKey,
		30,
		message.MessageHandlingSequential,
	).WithResponseMode(message.ResponseModeEverything)
	sendingTime := time.Date(2026, time.July, 29, 12, 30, 33, 725123456, time.UTC)

	frame, err := encodeSBERequest(
		fixsbe.NewSbeGoMarshaller(),
		request,
		"CLIENT",
		"SPOT",
		1,
		sendingTime,
	)
	if err != nil {
		t.Fatalf("encodeSBERequest() error = %v", err)
	}

	var decoded fixsbe.Logon
	header := decodeSBERequestFrameForTest(t, frame, &decoded)
	if header.TemplateId != decoded.SbeTemplateId() ||
		header.SchemaId != sbeSchemaID ||
		header.Version != sbeSchemaVersion ||
		header.SeqNum != 1 ||
		header.SendingTime != sendingTime.UnixMicro() {
		t.Fatalf("header = %+v", header)
	}
	if decoded.EncryptMethod != 0 ||
		decoded.HeartBtInt != 30 ||
		decoded.ResetSeqNumFlag != fixsbe.BoolEnum.True ||
		decoded.MessageHandling != fixsbe.MessageHandling.Sequential ||
		decoded.ResponseMode != fixsbe.ResponseMode.Everything {
		t.Fatalf("decoded logon = %+v", decoded)
	}
	if decoded.ExecutionReportType != fixsbe.ExecutionReportType.NullValue ||
		decoded.DropCopyFlag != fixsbe.BoolEnum.NullValue ||
		decoded.RecvWindow != math.MaxUint32 {
		t.Fatalf("optional logon fields = %+v", decoded)
	}
	if string(decoded.SenderCompId) != "CLIENT" ||
		string(decoded.TargetCompId) != "SPOT" ||
		string(decoded.Username) != "api-key" {
		t.Fatalf("logon strings = %+v", decoded)
	}

	wantSignature, err := message.SignLogonRawData(
		privateKey,
		"CLIENT",
		"SPOT",
		1,
		"20260729-12:30:33.725123",
	)
	if err != nil {
		t.Fatalf("SignLogonRawData() error = %v", err)
	}
	if string(decoded.RawData) != wantSignature {
		t.Fatalf("RawData = %q, want %q", decoded.RawData, wantSignature)
	}
}

func TestEncodeSBENewOrderSingleDecimals(t *testing.T) {
	sor := true
	request := message.NewNewOrderSingle("order-1", message.OrdTypeLimit, message.SideBuy, "BTCUSDT")
	request.OrderQty = 1.25
	request.Price = 12.345
	request.TriggerPrice = 10.5
	request.CashOrderQty = 20
	request.MaxFloor = 0.5
	request.TimeInForce = message.TimeInForceGoodTillCancel
	request.PegOffsetValue = 2
	request.PegMoveType = message.PegMoveTypeFixed
	request.SOR = &sor

	frame, err := encodeSBERequest(
		fixsbe.NewSbeGoMarshaller(),
		request,
		"CLIENT",
		"SPOT",
		2,
		time.Now(),
	)
	if err != nil {
		t.Fatalf("encodeSBERequest() error = %v", err)
	}

	var decoded fixsbe.NewOrderSingle
	decodeSBERequestFrameForTest(t, frame, &decoded)
	if decoded.PriceExponent != -3 ||
		decoded.Price != 12345 ||
		decoded.TriggerPrice != 10500 ||
		decoded.CashOrderQty != 20000 {
		t.Fatalf("price encoding = exponent %d, price %d, trigger %d, cash %d",
			decoded.PriceExponent,
			decoded.Price,
			decoded.TriggerPrice,
			decoded.CashOrderQty,
		)
	}
	if decoded.QtyExponent != -2 || decoded.OrderQty != 125 || decoded.MaxFloor != 50 {
		t.Fatalf("quantity encoding = exponent %d, qty %d, maxFloor %d",
			decoded.QtyExponent,
			decoded.OrderQty,
			decoded.MaxFloor,
		)
	}
	if decoded.OrdType != fixsbe.OrdType.Limit ||
		decoded.Side != fixsbe.Side.Buy ||
		decoded.TimeInForce != fixsbe.TimeInForce.GoodTillCancel ||
		decoded.PegMoveType != fixsbe.PegMoveType.Fixed ||
		decoded.PegOffsetValue != 2 ||
		decoded.SOR != fixsbe.BoolEnum.True {
		t.Fatalf("order enums = %+v", decoded)
	}
	if decoded.ExecInst != fixsbe.ExecInst.NullValue ||
		decoded.StrategyID != math.MinInt64 ||
		decoded.TargetStrategy != math.MinInt32 {
		t.Fatalf("optional fields were not encoded as null: %+v", decoded)
	}
}

func TestEncodeSBENewOrderListGroups(t *testing.T) {
	request := message.NewNewOrderList("list-1", []message.NewOrderListOrder{
		{
			OrderFields: message.OrderFields{
				ClOrdID:     "order-1",
				OrderQty:    1.5,
				OrdType:     message.OrdTypeLimit,
				Price:       10.25,
				Side:        message.SideSell,
				Symbol:      "BTCUSDT",
				TimeInForce: message.TimeInForceGoodTillCancel,
			},
			ListTriggeringInstructions: []message.ListTriggeringInstruction{
				{
					ListTriggerType:         message.ListTriggerTypeFilled,
					ListTriggerTriggerIndex: 0,
					ListTriggerAction:       message.ListTriggerActionRelease,
				},
			},
		},
	})
	request.ContingencyType = message.ContingencyTypeOneTriggersTheOther

	frame, err := encodeSBERequest(
		fixsbe.NewSbeGoMarshaller(),
		request,
		"CLIENT",
		"SPOT",
		3,
		time.Now(),
	)
	if err != nil {
		t.Fatalf("encodeSBERequest() error = %v", err)
	}

	var decoded fixsbe.NewOrderList
	decodeSBERequestFrameForTest(t, frame, &decoded)
	if decoded.ContingencyType != fixsbe.ContingencyType.OneTriggersTheOther ||
		decoded.OPO != fixsbe.BoolEnum.NullValue ||
		string(decoded.ClListID) != "list-1" ||
		len(decoded.Orders) != 1 {
		t.Fatalf("decoded list = %+v", decoded)
	}
	order := decoded.Orders[0]
	if order.PriceExponent != -2 || order.Price != 1025 ||
		order.QtyExponent != -1 || order.OrderQty != 15 ||
		len(order.ListTriggeringInstructions) != 1 {
		t.Fatalf("decoded order = %+v", order)
	}
	instruction := order.ListTriggeringInstructions[0]
	if instruction.ListTriggerType != fixsbe.ListTriggerType.Filled ||
		instruction.ListTriggerTriggerIndex != 0 ||
		instruction.ListTriggerAction != fixsbe.ListTriggerAction.Release {
		t.Fatalf("decoded instruction = %+v", instruction)
	}
}

func TestEncodeSBERequestTemplates(t *testing.T) {
	privateKey := ed25519.NewKeyFromSeed(bytes.Repeat([]byte{2}, ed25519.SeedSize))
	orderFields := message.OrderFields{
		ClOrdID:     "order-1",
		OrderQty:    1,
		OrdType:     message.OrdTypeLimit,
		Price:       10,
		Side:        message.SideBuy,
		Symbol:      "BTCUSDT",
		TimeInForce: message.TimeInForceGoodTillCancel,
	}

	newOrderList := message.NewNewOrderList("list-1", []message.NewOrderListOrder{
		{OrderFields: orderFields},
	})
	newOrderList.ContingencyType = message.ContingencyTypeOneCancelsTheOther

	cancelAndNew := message.NewOrderCancelRequestAndNewOrderSingle(
		message.OrderCancelRequestAndNewOrderModeStopOnFailure,
		"replacement-1",
		message.OrdTypeLimit,
		message.SideBuy,
		"BTCUSDT",
	)
	cancelAndNew.OrderQty = 1
	cancelAndNew.Price = 10

	amend := message.NewOrderAmendKeepPriorityRequest("amend-1", "BTCUSDT")
	amend.OrderQty = 0.5

	marketData := message.NewMarketDataRequest("md-1", message.SubscriptionRequestTypeSubscribe)
	marketData.MarketDepth = 2
	marketData.Symbols = []string{"BTCUSDT"}
	marketData.MDEntryTypes = []message.MDEntryType{message.MDEntryTypeBid, message.MDEntryTypeOffer}

	tests := []struct {
		name       string
		request    message.Request
		templateID uint16
		body       sbeBodyDecoder
	}{
		{
			name:       "heartbeat",
			request:    message.NewHeartbeat("test"),
			templateID: 20001,
			body:       new(fixsbe.Heartbeat),
		},
		{
			name:       "test request",
			request:    message.NewTestRequest("test"),
			templateID: 20002,
			body:       new(fixsbe.TestRequest),
		},
		{
			name:       "logout",
			request:    message.NewLogout("bye"),
			templateID: 20004,
			body:       new(fixsbe.Logout),
		},
		{
			name: "logon",
			request: message.NewLogonRequest(
				"api-key",
				privateKey,
				30,
				message.MessageHandlingSequential,
			),
			templateID: 20008,
			body:       new(fixsbe.Logon),
		},
		{
			name:       "new order single",
			request:    &message.NewOrderSingle{OrderFields: orderFields},
			templateID: 99,
			body:       new(fixsbe.NewOrderSingle),
		},
		{
			name:       "new order list",
			request:    newOrderList,
			templateID: 100,
			body:       new(fixsbe.NewOrderList),
		},
		{
			name:       "order cancel",
			request:    message.NewOrderCancelRequest("cancel-1", "BTCUSDT"),
			templateID: 101,
			body:       new(fixsbe.OrderCancelRequest),
		},
		{
			name:       "cancel and new",
			request:    cancelAndNew,
			templateID: 97,
			body:       new(fixsbe.OrderCancelRequestAndNewOrderSingle),
		},
		{
			name:       "mass cancel",
			request:    message.NewOrderMassCancelRequest("mass-1", "BTCUSDT"),
			templateID: 103,
			body:       new(fixsbe.OrderMassCancelRequest),
		},
		{
			name:       "amend",
			request:    amend,
			templateID: 105,
			body:       new(fixsbe.OrderAmendKeepPriorityRequest),
		},
		{
			name:       "limit query",
			request:    message.NewLimitQuery("limit-1"),
			templateID: 120,
			body:       new(fixsbe.LimitQuery),
		},
		{
			name: "instrument list",
			request: message.NewInstrumentListRequest(
				"instrument-1",
				message.InstrumentListRequestTypeAllInstruments,
			),
			templateID: 200,
			body:       new(fixsbe.InstrumentListRequest),
		},
		{
			name:       "market data",
			request:    marketData,
			templateID: 202,
			body:       new(fixsbe.MarketDataRequest),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			frame, err := encodeSBERequest(
				fixsbe.NewSbeGoMarshaller(),
				test.request,
				"CLIENT",
				"SPOT",
				10,
				time.Now(),
			)
			if err != nil {
				t.Fatalf("encodeSBERequest() error = %v", err)
			}
			header := decodeSBERequestFrameForTest(t, frame, test.body)
			if header.TemplateId != test.templateID {
				t.Fatalf("TemplateId = %d, want %d", header.TemplateId, test.templateID)
			}
		})
	}
}

func TestEncodeSBERequestValidation(t *testing.T) {
	order := message.NewNewOrderSingle("order-1", message.OrdTypePegged, message.SideBuy, "BTCUSDT")
	order.PegOffsetValue = 1.5
	_, err := encodeSBERequest(
		fixsbe.NewSbeGoMarshaller(),
		order,
		"CLIENT",
		"SPOT",
		1,
		time.Now(),
	)
	if err == nil || !strings.Contains(err.Error(), "PegOffsetValue") {
		t.Fatalf("encodeSBERequest() error = %v, want PegOffsetValue error", err)
	}

	cancel := message.NewOrderCancelRequest("cancel-1", "BTCUSDT")
	cancel.ListID = "not-an-int"
	_, err = encodeSBERequest(
		fixsbe.NewSbeGoMarshaller(),
		cancel,
		"CLIENT",
		"SPOT",
		1,
		time.Now(),
	)
	if err == nil || !strings.Contains(err.Error(), "ListID") {
		t.Fatalf("encodeSBERequest() error = %v, want ListID error", err)
	}
}

func decodeSBERequestFrameForTest(
	t *testing.T,
	frame []byte,
	body sbeBodyDecoder,
) fixsbe.MessageHeader {
	t.Helper()
	header, reader := readSBERequestHeaderForTest(t, frame)
	if err := body.Decode(
		fixsbe.NewSbeGoMarshaller(),
		reader,
		header.Version,
		header.BlockLength,
		true,
	); err != nil {
		t.Fatalf("decode SBE request body error = %v", err)
	}
	if reader.Len() != 0 {
		t.Fatalf("request frame has %d unread bytes", reader.Len())
	}
	return header
}

func readSBERequestHeaderForTest(
	t *testing.T,
	frame []byte,
) (fixsbe.MessageHeader, *bytes.Reader) {
	t.Helper()
	if len(frame) < sbeSOFHLength+sbeMessageHeaderLength {
		t.Fatalf("frame length = %d", len(frame))
	}
	if got := binary.LittleEndian.Uint32(frame[:4]); got != uint32(len(frame)) {
		t.Fatalf("SOFH messageLength = %d, want %d", got, len(frame))
	}
	if got := binary.LittleEndian.Uint16(frame[4:sbeSOFHLength]); got != sbeEncodingType {
		t.Fatalf("SOFH encodingType = 0x%04X, want 0x%04X", got, sbeEncodingType)
	}

	reader := bytes.NewReader(frame[sbeSOFHLength:])
	var header fixsbe.MessageHeader
	if err := header.Decode(fixsbe.NewSbeGoMarshaller(), reader, 0); err != nil {
		t.Fatalf("decode SBE request header error = %v", err)
	}
	return header, reader
}
