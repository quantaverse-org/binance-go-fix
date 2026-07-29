package binance_go_fix

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/quantaverse-org/binance-go-fix/internal/fixsbe"
	"github.com/quantaverse-org/binance-go-fix/message"
)

type sbeTestMessage interface {
	Encode(*fixsbe.SbeGoMarshaller, io.Writer, bool) error
	SbeBlockLength() uint16
	SbeTemplateId() uint16
	SbeSchemaId() uint16
	SbeSchemaVersion() uint16
}

func TestReadSBELogonAck(t *testing.T) {
	value := &fixsbe.LogonAck{
		EncryptMethod:                0,
		HeartBtInt:                   30,
		SbeSchemaIdVersionDeprecated: fixsbe.BoolEnum.False,
		UUID:                         []byte("session-uuid"),
	}
	sendingTime := time.Date(2026, time.July, 29, 12, 30, 33, 725732000, time.UTC)

	inbound := decodeSBEFrameForTest(t, encodeSBEFrameForTest(t, value, 7, sendingTime))
	if inbound.msgType != message.MsgTypeLogon {
		t.Fatalf("MsgType = %s, want %s", inbound.msgType, message.MsgTypeLogon)
	}
	if inbound.seqNum != 7 {
		t.Fatalf("SeqNum = %d, want 7", inbound.seqNum)
	}
	if !inbound.sendingTime.Equal(sendingTime) {
		t.Fatalf("SendingTime = %v, want %v", inbound.sendingTime, sendingTime)
	}
	response, ok := inbound.response.(*message.LogonResponse)
	if !ok {
		t.Fatalf("response type = %T, want *message.LogonResponse", inbound.response)
	}
	if response.EncryptMethod != message.EncryptMethodNone || response.HeartBtInt != 30 || response.UUID != "session-uuid" {
		t.Fatalf("response = %+v", response)
	}
	if response.SbeSchemaIdVersionDeprecated {
		t.Fatal("SbeSchemaIdVersionDeprecated = true, want false")
	}
}

func TestReadSBEReject(t *testing.T) {
	var value fixsbe.Reject
	fixsbe.RejectInit(&value)
	value.RefSeqNum = 9
	value.RefTagID = uint32(message.TagMsgType)
	value.SessionRejectReason = fixsbe.SessionRejectReason.RequiredTagMissing
	value.ErrorCode = -1102
	value.RefMsgType = []byte(message.MsgTypeNewOrderSingle)
	value.Text = []byte("Missing field")

	inbound := decodeSBEFrameForTest(t, encodeSBEFrameForTest(t, &value, 10, time.Now()))
	response, ok := inbound.response.(*message.Reject)
	if !ok {
		t.Fatalf("response type = %T, want *message.Reject", inbound.response)
	}
	if response.RefSeqNum == nil || *response.RefSeqNum != 9 {
		t.Fatalf("RefSeqNum = %v, want 9", response.RefSeqNum)
	}
	if response.RefTagID == nil || *response.RefTagID != message.TagMsgType {
		t.Fatalf("RefTagID = %v, want %d", response.RefTagID, message.TagMsgType)
	}
	if response.SessionRejectReason == nil || *response.SessionRejectReason != message.SessionRejectReasonRequiredTagMissing {
		t.Fatalf("SessionRejectReason = %v", response.SessionRejectReason)
	}
	if response.ErrorCode == nil || *response.ErrorCode != -1102 {
		t.Fatalf("ErrorCode = %v, want -1102", response.ErrorCode)
	}
	if response.Text != "Missing field" {
		t.Fatalf("Text = %q, want %q", response.Text, "Missing field")
	}
}

func TestReadSBEMarketDataSnapshot(t *testing.T) {
	value := &fixsbe.MarketDataSnapshot{
		LastBookUpdateID: 42,
		PriceExponent:    -2,
		QtyExponent:      -3,
		MDEntriesBids: []fixsbe.MarketDataSnapshotMDEntriesBids{
			{MDEntryPx: 12345, MDEntrySize: 2500},
		},
		MDEntriesAsks: []fixsbe.MarketDataSnapshotMDEntriesAsks{
			{MDEntryPx: 12346, MDEntrySize: 1750},
		},
		Symbol: []byte("BTCUSDT"),
	}

	inbound := decodeSBEFrameForTest(t, encodeSBEFrameForTest(t, value, 11, time.Now()))
	response, ok := inbound.response.(*message.MarketDataSnapshot)
	if !ok {
		t.Fatalf("response type = %T, want *message.MarketDataSnapshot", inbound.response)
	}
	if response.Symbol != "BTCUSDT" || response.LastBookUpdateID != 42 || response.NoMDEntries != 2 {
		t.Fatalf("response = %+v", response)
	}
	if response.Entries[0].MDEntryType != message.MDEntryTypeBid ||
		!floatClose(response.Entries[0].MDEntryPx, 123.45) ||
		!floatClose(response.Entries[0].MDEntrySize, 2.5) {
		t.Fatalf("bid = %+v", response.Entries[0])
	}
	if response.Entries[1].MDEntryType != message.MDEntryTypeOffer ||
		!floatClose(response.Entries[1].MDEntryPx, 123.46) ||
		!floatClose(response.Entries[1].MDEntrySize, 1.75) {
		t.Fatalf("ask = %+v", response.Entries[1])
	}
}

func TestReadSBEMarketDataTradeGroup(t *testing.T) {
	transactTime := time.Date(2026, time.July, 22, 2, 25, 50, 143785000, time.UTC)
	value := &fixsbe.MarketDataIncrementalTrade{
		TransactTime:  transactTime.UnixMicro(),
		PriceExponent: -2,
		QtyExponent:   -3,
		MDEntries: []fixsbe.MarketDataIncrementalTradeMDEntries{
			{
				TradeID:       101,
				MDEntryPx:     20001,
				MDEntrySize:   1250,
				AggressorSide: fixsbe.Side.Buy,
			},
			{
				TradeID:       102,
				MDEntryPx:     20002,
				MDEntrySize:   500,
				AggressorSide: fixsbe.Side.Sell,
			},
		},
		Symbol: []byte("ETHUSDT"),
	}

	inbound := decodeSBEFrameForTest(t, encodeSBEFrameForTest(t, value, 12, time.Now()))
	response, ok := inbound.response.(*message.MarketDataIncrementalRefresh)
	if !ok {
		t.Fatalf("response type = %T, want *message.MarketDataIncrementalRefresh", inbound.response)
	}
	if response.NoMDEntries != 2 || len(response.Entries) != 2 {
		t.Fatalf("entries = %+v", response.Entries)
	}
	if response.Entries[0].TradeID != 101 ||
		!floatClose(response.Entries[0].MDEntryPx, 200.01) ||
		!floatClose(response.Entries[0].MDEntrySize, 1.25) ||
		response.Entries[0].AggressorSide != message.AggressorSideBuy ||
		!response.Entries[0].TransactTime.Equal(transactTime) {
		t.Fatalf("first trade = %+v", response.Entries[0])
	}
}

func TestReadSBEMessageRejectsInvalidSOFH(t *testing.T) {
	frame := make([]byte, sbeSOFHLength)
	binary.LittleEndian.PutUint32(frame[:4], sbeSOFHLength+sbeMessageHeaderLength)
	binary.LittleEndian.PutUint16(frame[4:], 0x1234)

	_, err := readSBEMessage(
		bufio.NewReader(bytes.NewReader(frame)),
		fixsbe.NewSbeGoMarshaller(),
	)
	if err == nil || !strings.Contains(err.Error(), "encoding type") {
		t.Fatalf("readSBEMessage() error = %v, want encoding type error", err)
	}
}

func TestReadSBEResponseTemplates(t *testing.T) {
	heartbeat := &fixsbe.Heartbeat{}
	testRequest := &fixsbe.TestRequest{TestReqID: []byte("test")}
	logout := &fixsbe.Logout{Text: []byte("bye")}
	news := &fixsbe.News{Headline: []byte("reconnect")}

	executionAck := new(fixsbe.ExecutionReportAck)
	fixsbe.ExecutionReportAckInit(executionAck)
	executionAck.ExecType = fixsbe.ExecType.New
	executionAck.OrdStatus = fixsbe.OrdStatus.New
	executionAck.OrdRejReason = fixsbe.OrdRejReason.NullValue
	executionAck.Symbol = []byte("BTCUSDT")

	executionReport := new(fixsbe.ExecutionReport)
	fixsbe.ExecutionReportInit(executionReport)
	executionReport.PriceExponent = -2
	executionReport.QtyExponent = -3
	executionReport.OrdType = fixsbe.OrdType.Limit
	executionReport.Side = fixsbe.Side.Buy
	executionReport.PegMoveType = fixsbe.PegMoveType.NullValue
	executionReport.ExecType = fixsbe.ExecType.New
	executionReport.CumQty = 0
	executionReport.AggressorIndicator = fixsbe.BoolEnum.NullValue
	executionReport.LastQty = 0
	executionReport.OrdStatus = fixsbe.OrdStatus.New
	executionReport.WorkingFloor = fixsbe.WorkingFloor.NullValue
	executionReport.WorkingIndicator = fixsbe.BoolEnum.NullValue
	executionReport.SOR = fixsbe.BoolEnum.NullValue
	executionReport.OrdRejReason = fixsbe.OrdRejReason.NullValue
	executionReport.ExpiryReason = fixsbe.ExpiryReason.NullValue
	executionReport.Symbol = []byte("BTCUSDT")

	cancelReject := new(fixsbe.OrderCancelReject)
	fixsbe.OrderCancelRejectInit(cancelReject)
	cancelReject.CancelRestrictions = fixsbe.CancelRestrictions.NullValue
	cancelReject.CxlRejResponseTo = fixsbe.CxlRejResponseTo.OrderCancelRequest
	cancelReject.ErrorCode = -1
	cancelReject.ClOrdID = []byte("cancel-1")
	cancelReject.Symbol = []byte("BTCUSDT")
	cancelReject.ErrorText = []byte("rejected")

	listStatus := new(fixsbe.ListStatus)
	fixsbe.ListStatusInit(listStatus)
	listStatus.ContingencyType = fixsbe.ContingencyType.NullValue
	listStatus.ListStatusType = fixsbe.ListStatusType.Response
	listStatus.ListOrderStatus = fixsbe.ListOrderStatus.AllDone
	listStatus.ListRejectReason = fixsbe.ListRejectReason.NullValue

	massCancel := new(fixsbe.OrderMassCancelReport)
	fixsbe.OrderMassCancelReportInit(massCancel)
	massCancel.MassCancelRequestType = fixsbe.MassCancelRequestType.CancelSymbolOrders
	massCancel.MassCancelResponse = fixsbe.MassCancelResponse.CancelSymbolOrders
	massCancel.MassCancelRejectReason = fixsbe.MassCancelRejectReason.NullValue
	massCancel.Symbol = []byte("BTCUSDT")
	massCancel.ClOrdID = []byte("mass-1")

	amendReject := new(fixsbe.OrderAmendReject)
	fixsbe.OrderAmendRejectInit(amendReject)
	amendReject.QtyExponent = -3
	amendReject.OrderQty = 1000
	amendReject.ErrorCode = -1
	amendReject.ClOrdID = []byte("amend-1")
	amendReject.Symbol = []byte("BTCUSDT")
	amendReject.ErrorText = []byte("rejected")

	limitResponse := &fixsbe.LimitResponse{ReqID: []byte("limit-1")}
	instrumentList := &fixsbe.InstrumentList{InstrumentReqID: []byte("instrument-1")}

	marketReject := new(fixsbe.MarketDataRequestReject)
	fixsbe.MarketDataRequestRejectInit(marketReject)
	marketReject.MDReqID = []byte("md-1")

	bookTicker := &fixsbe.MarketDataIncrementalBookTicker{
		LastBookUpdateID: 1,
		PriceExponent:    -2,
		QtyExponent:      -3,
		Symbol:           []byte("BTCUSDT"),
	}
	depth := &fixsbe.MarketDataIncrementalDepth{
		FirstBookUpdateID: 1,
		LastBookUpdateID:  2,
		PriceExponent:     -2,
		QtyExponent:       -3,
		Symbol:            []byte("BTCUSDT"),
	}

	tests := []struct {
		name    string
		value   sbeTestMessage
		msgType message.MsgType
	}{
		{name: "heartbeat", value: heartbeat, msgType: message.MsgTypeHeartbeat},
		{name: "test request", value: testRequest, msgType: message.MsgTypeTestRequest},
		{name: "logout", value: logout, msgType: message.MsgTypeLogout},
		{name: "news", value: news, msgType: message.MsgTypeNews},
		{name: "execution report ack", value: executionAck, msgType: message.MsgTypeExecutionReport},
		{name: "execution report", value: executionReport, msgType: message.MsgTypeExecutionReport},
		{name: "order cancel reject", value: cancelReject, msgType: message.MsgTypeOrderCancelReject},
		{name: "list status", value: listStatus, msgType: message.MsgTypeListStatus},
		{name: "mass cancel report", value: massCancel, msgType: message.MsgTypeOrderMassCancelReport},
		{name: "amend reject", value: amendReject, msgType: message.MsgTypeOrderAmendReject},
		{name: "limit response", value: limitResponse, msgType: message.MsgTypeLimitResponse},
		{name: "instrument list", value: instrumentList, msgType: message.MsgTypeInstrumentList},
		{name: "market reject", value: marketReject, msgType: message.MsgTypeMarketDataRequestReject},
		{name: "book ticker", value: bookTicker, msgType: message.MsgTypeMarketDataIncrementalRefresh},
		{name: "depth", value: depth, msgType: message.MsgTypeMarketDataIncrementalRefresh},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inbound := decodeSBEFrameForTest(t, encodeSBEFrameForTest(t, test.value, 1, time.Now()))
			if inbound.msgType != test.msgType {
				t.Fatalf("MsgType = %s, want %s", inbound.msgType, test.msgType)
			}
			if inbound.response == nil {
				t.Fatal("response = nil")
			}
		})
	}
}

func TestResolveSBEMarketData(t *testing.T) {
	client := &Client{
		sbeMarketRoutes: make(map[sbeMarketRoute]string),
		sbeBookLevels:   make(map[string]map[sbeBookLevel]struct{}),
	}
	request := message.NewMarketDataRequest("md-1", message.SubscriptionRequestTypeSubscribe)
	request.Symbols = []string{"BTCUSDT", "ETHUSDT"}
	request.MDEntryTypes = []message.MDEntryType{message.MDEntryTypeBid, message.MDEntryTypeOffer}
	request.MarketDepth = 2
	if err := client.registerSBEMarketRoutes(request); err != nil {
		t.Fatalf("registerSBEMarketRoutes() error = %v", err)
	}

	snapshotTime := time.Now().UTC()
	snapshot := &message.MarketDataSnapshot{
		Symbol: "BTCUSDT",
		Entries: []message.MarketDataSnapshotEntry{
			{MDEntryType: message.MDEntryTypeBid, MDEntryPx: 100, MDEntrySize: 2},
		},
	}
	if err := client.resolveSBEMarketData(&inboundMessage{
		templateID:  sbeTemplateMarketDataSnapshot,
		sendingTime: snapshotTime,
	}, snapshot); err != nil {
		t.Fatalf("resolve snapshot error = %v", err)
	}
	if snapshot.MDReqID != "md-1" || !snapshot.SendingTime.Equal(snapshotTime) {
		t.Fatalf("snapshot = %+v", snapshot)
	}

	secondSnapshot := &message.MarketDataSnapshot{
		Symbol: "ETHUSDT",
		Entries: []message.MarketDataSnapshotEntry{
			{MDEntryType: message.MDEntryTypeBid, MDEntryPx: 100, MDEntrySize: 5},
		},
	}
	if err := client.resolveSBEMarketData(&inboundMessage{
		templateID:  sbeTemplateMarketDataSnapshot,
		sendingTime: snapshotTime,
	}, secondSnapshot); err != nil {
		t.Fatalf("resolve second snapshot error = %v", err)
	}

	refresh := &message.MarketDataIncrementalRefresh{
		Entries: []message.MarketDataIncrementalRefreshEntry{
			{Symbol: "BTCUSDT", MDEntryType: message.MDEntryTypeBid, MDEntryPx: 100, MDEntrySize: 3, MDUpdateAction: message.MDUpdateActionNew},
			{Symbol: "BTCUSDT", MDEntryType: message.MDEntryTypeBid, MDEntryPx: 99, MDEntrySize: 1, MDUpdateAction: message.MDUpdateActionNew},
			{Symbol: "BTCUSDT", MDEntryType: message.MDEntryTypeBid, MDEntryPx: 100, MDUpdateAction: message.MDUpdateActionDelete},
		},
	}
	if err := client.resolveSBEMarketData(&inboundMessage{
		templateID:  sbeTemplateMarketDataIncrementalDepth,
		sendingTime: snapshotTime.Add(time.Millisecond),
	}, refresh); err != nil {
		t.Fatalf("resolve refresh error = %v", err)
	}
	if refresh.MDReqID != "md-1" {
		t.Fatalf("MDReqID = %q, want md-1", refresh.MDReqID)
	}
	wantActions := []message.MDUpdateAction{
		message.MDUpdateActionChange,
		message.MDUpdateActionNew,
		message.MDUpdateActionDelete,
	}
	for i, want := range wantActions {
		if refresh.Entries[i].MDUpdateAction != want {
			t.Fatalf("entry %d action = %s, want %s", i, refresh.Entries[i].MDUpdateAction, want)
		}
	}

	secondRefresh := &message.MarketDataIncrementalRefresh{
		Entries: []message.MarketDataIncrementalRefreshEntry{
			{Symbol: "ETHUSDT", MDEntryType: message.MDEntryTypeBid, MDEntryPx: 100, MDEntrySize: 6, MDUpdateAction: message.MDUpdateActionNew},
		},
	}
	if err := client.resolveSBEMarketData(&inboundMessage{
		templateID:  sbeTemplateMarketDataIncrementalDepth,
		sendingTime: snapshotTime.Add(time.Millisecond),
	}, secondRefresh); err != nil {
		t.Fatalf("resolve second refresh error = %v", err)
	}
	if got := secondRefresh.Entries[0].MDUpdateAction; got != message.MDUpdateActionChange {
		t.Fatalf("second symbol action = %s, want %s", got, message.MDUpdateActionChange)
	}
}

func encodeSBEFrameForTest(t *testing.T, value sbeTestMessage, seqNum uint32, sendingTime time.Time) []byte {
	t.Helper()

	marshaller := fixsbe.NewSbeGoMarshaller()
	var payload bytes.Buffer
	header := fixsbe.MessageHeader{
		BlockLength: value.SbeBlockLength(),
		TemplateId:  value.SbeTemplateId(),
		SchemaId:    value.SbeSchemaId(),
		Version:     value.SbeSchemaVersion(),
		SeqNum:      seqNum,
		SendingTime: sendingTime.UnixMicro(),
	}
	if err := header.Encode(marshaller, &payload); err != nil {
		t.Fatalf("encode header error = %v", err)
	}
	if err := value.Encode(marshaller, &payload, true); err != nil {
		t.Fatalf("encode body error = %v", err)
	}

	frame := make([]byte, sbeSOFHLength+payload.Len())
	binary.LittleEndian.PutUint32(frame[:4], uint32(len(frame)))
	binary.LittleEndian.PutUint16(frame[4:6], sbeEncodingType)
	copy(frame[sbeSOFHLength:], payload.Bytes())
	return frame
}

func decodeSBEFrameForTest(t *testing.T, frame []byte) *inboundMessage {
	t.Helper()

	inbound, err := readSBEMessage(
		bufio.NewReader(bytes.NewReader(frame)),
		fixsbe.NewSbeGoMarshaller(),
	)
	if err != nil {
		t.Fatalf("readSBEMessage() error = %v", err)
	}
	return inbound
}

func floatClose(got, want float64) bool {
	return math.Abs(got-want) < 1e-12
}
