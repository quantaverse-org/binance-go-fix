package binance_go_fix

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"

	"github.com/quantaverse-org/binance-go-fix/internal/fixsbe"
	"github.com/quantaverse-org/binance-go-fix/message"
)

const (
	sbeSchemaID            uint16 = 1
	sbeSchemaVersion       uint16 = 1
	sbeEncodingType        uint16 = 0xEB50
	sbeSOFHLength                 = 6
	sbeMessageHeaderLength        = 20
	maxSBEMessageLength           = 16 << 20

	sbeTemplateExecutionReportAck              uint16 = 198
	sbeTemplateExecutionReport                 uint16 = 98
	sbeTemplateOrderCancelReject               uint16 = 96
	sbeTemplateListStatus                      uint16 = 102
	sbeTemplateOrderMassCancelReport           uint16 = 104
	sbeTemplateOrderAmendReject                uint16 = 106
	sbeTemplateLimitResponse                   uint16 = 121
	sbeTemplateInstrumentList                  uint16 = 201
	sbeTemplateMarketDataRequestReject         uint16 = 203
	sbeTemplateMarketDataSnapshot              uint16 = 204
	sbeTemplateMarketDataIncrementalTrade      uint16 = 205
	sbeTemplateMarketDataIncrementalBookTicker uint16 = 206
	sbeTemplateMarketDataIncrementalDepth      uint16 = 207
	sbeTemplateHeartbeat                       uint16 = 20001
	sbeTemplateTestRequest                     uint16 = 20002
	sbeTemplateReject                          uint16 = 20003
	sbeTemplateLogout                          uint16 = 20004
	sbeTemplateLogonAck                        uint16 = 20009
	sbeTemplateNews                            uint16 = 20100
)

type inboundMessage struct {
	msgType      message.MsgType
	templateID   uint16
	seqNum       uint32
	sendingTime  time.Time
	marketSymbol string
	response     message.Response
}

type sbeBodyDecoder interface {
	Decode(*fixsbe.SbeGoMarshaller, io.Reader, uint16, uint16, bool) error
}

// readSBEMessage reads one complete SOFH frame and decodes it with the generated
// FIX SBE 1.1 codecs. The frame boundary prevents one malformed body from
// consuming bytes belonging to the next response.
func readSBEMessage(reader *bufio.Reader, marshaller *fixsbe.SbeGoMarshaller) (*inboundMessage, error) {
	var sofh [sbeSOFHLength]byte
	if _, err := io.ReadFull(reader, sofh[:]); err != nil {
		return nil, fmt.Errorf("read SBE SOFH: %w", err)
	}

	messageLength := binary.LittleEndian.Uint32(sofh[:4])
	encodingType := binary.LittleEndian.Uint16(sofh[4:])
	if encodingType != sbeEncodingType {
		return nil, fmt.Errorf("invalid SBE encoding type: 0x%04X", encodingType)
	}
	minLength := uint32(sbeSOFHLength + sbeMessageHeaderLength)
	if messageLength < minLength {
		return nil, fmt.Errorf("invalid SBE message length: %d", messageLength)
	}
	if messageLength > maxSBEMessageLength {
		return nil, fmt.Errorf("SBE message length %d exceeds limit %d", messageLength, maxSBEMessageLength)
	}

	payload := make([]byte, int(messageLength)-sbeSOFHLength)
	if _, err := io.ReadFull(reader, payload); err != nil {
		return nil, fmt.Errorf("read SBE payload: %w", err)
	}
	frame := bytes.NewReader(payload)

	var header fixsbe.MessageHeader
	if err := header.Decode(marshaller, frame, 0); err != nil {
		return nil, fmt.Errorf("decode SBE message header: %w", err)
	}
	if header.SchemaId != sbeSchemaID {
		return nil, fmt.Errorf("unsupported SBE schema ID: %d", header.SchemaId)
	}

	inbound, err := decodeSBEBody(marshaller, frame, header)
	if err != nil {
		return nil, err
	}
	// Version 1 is fully known, so unread bytes indicate an invalid frame. For a
	// later compatible version, generated codecs may intentionally leave newly
	// appended variable data unread.
	if header.Version <= sbeSchemaVersion && frame.Len() != 0 {
		return nil, fmt.Errorf("SBE template %d has %d unread bytes", header.TemplateId, frame.Len())
	}
	inbound.templateID = header.TemplateId
	inbound.seqNum = header.SeqNum
	inbound.sendingTime = time.UnixMicro(header.SendingTime).UTC()
	return inbound, nil
}

func decodeSBEBody(marshaller *fixsbe.SbeGoMarshaller, reader io.Reader, header fixsbe.MessageHeader) (*inboundMessage, error) {
	var (
		msgType      message.MsgType
		decoder      sbeBodyDecoder
		response     message.Response
		marketSymbol string
		adapt        func()
	)

	switch header.TemplateId {
	case sbeTemplateHeartbeat:
		value := new(fixsbe.Heartbeat)
		decoder = value
		msgType = message.MsgTypeHeartbeat
		adapt = func() { response = &message.Heartbeat{TestReqID: string(value.TestReqID)} }
	case sbeTemplateTestRequest:
		value := new(fixsbe.TestRequest)
		decoder = value
		msgType = message.MsgTypeTestRequest
		adapt = func() { response = &message.TestRequest{TestReqID: string(value.TestReqID)} }
	case sbeTemplateReject:
		value := new(fixsbe.Reject)
		decoder = value
		msgType = message.MsgTypeReject
		adapt = func() { response = adaptSBEReject(value) }
	case sbeTemplateLogout:
		value := new(fixsbe.Logout)
		decoder = value
		msgType = message.MsgTypeLogout
		adapt = func() { response = &message.Logout{Text: string(value.Text)} }
	case sbeTemplateLogonAck:
		value := new(fixsbe.LogonAck)
		decoder = value
		msgType = message.MsgTypeLogon
		adapt = func() {
			response = &message.LogonResponse{
				EncryptMethod:                message.EncryptMethod(sbeUintEnum(value.EncryptMethod)),
				HeartBtInt:                   int64(value.HeartBtInt),
				UUID:                         string(value.UUID),
				SbeSchemaIdVersionDeprecated: value.SbeSchemaIdVersionDeprecated == fixsbe.BoolEnum.True,
			}
		}
	case sbeTemplateNews:
		value := new(fixsbe.News)
		decoder = value
		msgType = message.MsgTypeNews
		adapt = func() { response = &message.News{Headline: string(value.Headline)} }
	case sbeTemplateExecutionReportAck:
		value := new(fixsbe.ExecutionReportAck)
		decoder = value
		msgType = message.MsgTypeExecutionReport
		adapt = func() { response = adaptSBEExecutionReportAck(value) }
	case sbeTemplateExecutionReport:
		value := new(fixsbe.ExecutionReport)
		decoder = value
		msgType = message.MsgTypeExecutionReport
		adapt = func() { response = adaptSBEExecutionReport(value) }
	case sbeTemplateOrderCancelReject:
		value := new(fixsbe.OrderCancelReject)
		decoder = value
		msgType = message.MsgTypeOrderCancelReject
		adapt = func() { response = adaptSBEOrderCancelReject(value) }
	case sbeTemplateListStatus:
		value := new(fixsbe.ListStatus)
		decoder = value
		msgType = message.MsgTypeListStatus
		adapt = func() { response = adaptSBEListStatus(value) }
	case sbeTemplateOrderMassCancelReport:
		value := new(fixsbe.OrderMassCancelReport)
		decoder = value
		msgType = message.MsgTypeOrderMassCancelReport
		adapt = func() { response = adaptSBEOrderMassCancelReport(value) }
	case sbeTemplateOrderAmendReject:
		value := new(fixsbe.OrderAmendReject)
		decoder = value
		msgType = message.MsgTypeOrderAmendReject
		adapt = func() { response = adaptSBEOrderAmendReject(value) }
	case sbeTemplateLimitResponse:
		value := new(fixsbe.LimitResponse)
		decoder = value
		msgType = message.MsgTypeLimitResponse
		adapt = func() { response = adaptSBELimitResponse(value) }
	case sbeTemplateInstrumentList:
		value := new(fixsbe.InstrumentList)
		decoder = value
		msgType = message.MsgTypeInstrumentList
		adapt = func() { response = adaptSBEInstrumentList(value) }
	case sbeTemplateMarketDataRequestReject:
		value := new(fixsbe.MarketDataRequestReject)
		decoder = value
		msgType = message.MsgTypeMarketDataRequestReject
		adapt = func() { response = adaptSBEMarketDataRequestReject(value) }
	case sbeTemplateMarketDataSnapshot:
		value := new(fixsbe.MarketDataSnapshot)
		decoder = value
		msgType = message.MsgTypeMarketDataSnapshot
		adapt = func() {
			marketSymbol = string(value.Symbol)
			response = adaptSBEMarketDataSnapshot(value)
		}
	case sbeTemplateMarketDataIncrementalTrade:
		value := new(fixsbe.MarketDataIncrementalTrade)
		decoder = value
		msgType = message.MsgTypeMarketDataIncrementalRefresh
		adapt = func() {
			marketSymbol = string(value.Symbol)
			response = adaptSBEMarketDataTrade(value)
		}
	case sbeTemplateMarketDataIncrementalBookTicker:
		value := new(fixsbe.MarketDataIncrementalBookTicker)
		decoder = value
		msgType = message.MsgTypeMarketDataIncrementalRefresh
		adapt = func() {
			marketSymbol = string(value.Symbol)
			response = adaptSBEMarketDataBookTicker(value)
		}
	case sbeTemplateMarketDataIncrementalDepth:
		value := new(fixsbe.MarketDataIncrementalDepth)
		decoder = value
		msgType = message.MsgTypeMarketDataIncrementalRefresh
		adapt = func() {
			marketSymbol = string(value.Symbol)
			response = adaptSBEMarketDataDepth(value)
		}
	default:
		return nil, fmt.Errorf("unsupported SBE template ID: %d", header.TemplateId)
	}

	if err := decoder.Decode(marshaller, reader, header.Version, header.BlockLength, true); err != nil {
		return nil, fmt.Errorf("decode SBE template %d: %w", header.TemplateId, err)
	}
	adapt()
	return &inboundMessage{msgType: msgType, marketSymbol: marketSymbol, response: response}, nil
}

func adaptSBEReject(value *fixsbe.Reject) *message.Reject {
	response := &message.Reject{
		RefMsgType: string(value.RefMsgType),
		Text:       string(value.Text),
	}
	if value.RefSeqNum != value.RefSeqNumNullValue() {
		refSeqNum := value.RefSeqNum
		response.RefSeqNum = &refSeqNum
	}
	if value.RefTagID != value.RefTagIDNullValue() {
		refTagID := message.Tag(value.RefTagID)
		response.RefTagID = &refTagID
	}
	if value.SessionRejectReason != fixsbe.SessionRejectReason.NullValue {
		reason := message.SessionRejectReason(sbeUintEnum(uint8(value.SessionRejectReason)))
		response.SessionRejectReason = &reason
	}
	if value.ErrorCode != value.ErrorCodeNullValue() {
		errorCode := int64(value.ErrorCode)
		response.ErrorCode = &errorCode
	}
	return response
}

func adaptSBEExecutionReportAck(value *fixsbe.ExecutionReportAck) *message.ExecutionReport {
	response := &message.ExecutionReport{
		ClOrdID:      string(value.ClOrdID),
		Symbol:       string(value.Symbol),
		ExecType:     message.ExecType(sbeCharEnum(byte(value.ExecType))),
		OrdStatus:    message.OrdStatus(sbeCharEnum(byte(value.OrdStatus))),
		OrdRejReason: message.OrdRejReason(sbeUintEnum(uint8(value.OrdRejReason))),
		Text:         string(value.ErrorText),
	}
	if value.OrderID != value.OrderIDNullValue() {
		response.OrderID = value.OrderID
	}
	if value.ListID != value.ListIDNullValue() {
		response.ListID = strconv.FormatInt(value.ListID, 10)
	}
	response.TransactTime = sbeTimestamp(value.TransactTime, value.TransactTimeNullValue())
	if value.ErrorCode != value.ErrorCodeNullValue() {
		response.ErrorCode = int64(value.ErrorCode)
	}
	return response
}

func adaptSBEExecutionReport(value *fixsbe.ExecutionReport) *message.ExecutionReport {
	response := &message.ExecutionReport{
		ClOrdID:                 string(value.ClOrdID),
		OrigClOrdID:             string(value.OrigClOrdID),
		Symbol:                  string(value.Symbol),
		SecondarySymbol:         string(value.SecondarySymbol),
		CounterSymbol:           string(value.CounterSymbol),
		OrdType:                 message.OrdType(sbeCharEnum(byte(value.OrdType))),
		Side:                    message.Side(sbeCharEnum(byte(value.Side))),
		ExecInst:                message.ExecInst(sbeCharEnum(byte(value.ExecInst))),
		TriggerType:             message.TriggerType(sbeCharEnum(byte(value.TriggerType))),
		TriggerAction:           message.TriggerAction(sbeCharEnum(byte(value.TriggerAction))),
		TriggerPriceType:        message.TriggerPriceType(sbeCharEnum(byte(value.TriggerPriceType))),
		TriggerPriceDirection:   message.TriggerPriceDirection(sbeCharEnum(byte(value.TriggerPriceDirection))),
		PegPriceType:            message.PegPriceType(sbeCharEnum(byte(value.PegPriceType))),
		PegMoveType:             message.PegMoveType(sbeUintEnum(uint8(value.PegMoveType))),
		PegOffsetType:           message.PegOffsetType(sbeCharEnum(byte(value.PegOffsetType))),
		TimeInForce:             message.TimeInForce(sbeCharEnum(byte(value.TimeInForce))),
		SelfTradePreventionMode: message.SelfTradePreventionMode(sbeCharEnum(byte(value.SelfTradePreventionMode))),
		ExecType:                message.ExecType(sbeCharEnum(byte(value.ExecType))),
		OrdStatus:               message.OrdStatus(sbeCharEnum(byte(value.OrdStatus))),
		MatchType:               message.MatchType(sbeCharEnum(byte(value.MatchType))),
		OrderCapacity:           message.OrderCapacity(sbeCharEnum(byte(value.OrderCapacity))),
		OrdRejReason:            message.OrdRejReason(sbeUintEnum(uint8(value.OrdRejReason))),
		ExpiryReason:            sbeUintEnum(uint8(value.ExpiryReason)),
		Text:                    string(value.ErrorText),
	}

	response.ExecID = sbeOptionalIntString(value.ExecID, value.ExecIDNullValue())
	response.OrderID = sbeOptionalInt64(value.OrderID, value.OrderIDNullValue())
	response.OrderQty = sbeDecimal(value.OrderQty, value.QtyExponent, value.OrderQtyNullValue())
	response.Price = sbeDecimal(value.Price, value.PriceExponent, value.PriceNullValue())
	response.TriggerPrice = sbeDecimal(value.TriggerPrice, value.PriceExponent, value.TriggerPriceNullValue())
	if value.TriggerTrailingDeltaBips != value.TriggerTrailingDeltaBipsNullValue() {
		response.TriggerTrailingDeltaBips = int64(value.TriggerTrailingDeltaBips)
	}
	if value.PegOffsetValue != value.PegOffsetValueNullValue() {
		response.PegOffsetValue = float64(value.PegOffsetValue)
	}
	response.PeggedPrice = sbeDecimal(value.PeggedPrice, value.PriceExponent, value.PeggedPriceNullValue())
	response.TransactTime = sbeTimestamp(value.TransactTime, value.TransactTimeNullValue())
	response.OrderCreationTime = sbeOptionalInt64(value.OrderCreationTime, value.OrderCreationTimeNullValue())
	response.MaxFloor = sbeDecimal(value.MaxFloor, value.QtyExponent, value.MaxFloorNullValue())
	response.ListID = sbeOptionalIntString(value.ListID, value.ListIDNullValue())
	response.CashOrderQty = sbeDecimal(value.CashOrderQty, value.PriceExponent, value.CashOrderQtyNullValue())
	if value.TargetStrategy != value.TargetStrategyNullValue() {
		response.TargetStrategy = int64(value.TargetStrategy)
	}
	response.StrategyID = sbeOptionalInt64(value.StrategyID, value.StrategyIDNullValue())
	response.CumQty = sbeDecimal(value.CumQty, value.QtyExponent, value.CumQtyNullValue())
	response.LeavesQty = sbeDecimal(value.LeavesQty, value.QtyExponent, value.LeavesQtyNullValue())
	response.CumQuoteQty = sbeDecimal(value.CumQuoteQty, value.PriceExponent, value.CumQuoteQtyNullValue())
	response.AggressorIndicator = sbeOptionalBool(value.AggressorIndicator)
	response.TradeID = sbeOptionalIntString(value.TradeID, value.TradeIDNullValue())
	response.LastPx = sbeDecimal(value.LastPx, value.PriceExponent, value.LastPxNullValue())
	response.LastQty = sbeDecimal(value.LastQty, value.QtyExponent, value.LastQtyNullValue())
	response.SecondaryOrderID = sbeOptionalInt64(value.SecondaryOrderID, value.SecondaryOrderIDNullValue())
	response.SecondaryExternalAccountID = sbeOptionalInt64(value.SecondaryExternalAccountID, value.SecondaryExternalAccountIDNullValue())
	response.AllocID = sbeOptionalInt64(value.AllocID, value.AllocIDNullValue())
	if value.WorkingFloor != fixsbe.WorkingFloor.NullValue {
		response.WorkingFloor = int64(value.WorkingFloor)
	}
	response.WorkingIndicator = sbeOptionalBool(value.WorkingIndicator)
	response.WorkingTime = sbeTimestamp(value.WorkingTime, value.WorkingTimeNullValue())
	response.TrailingTime = sbeTimestamp(value.TrailingTime, value.TrailingTimeNullValue())
	response.PreventedMatchID = sbeOptionalInt64(value.PreventedMatchID, value.PreventedMatchIDNullValue())
	response.PreventedExecutionPrice = sbeDecimal(value.PreventedExecutionPrice, value.PriceExponent, value.PreventedExecutionPriceNullValue())
	response.PreventedExecutionQty = sbeDecimal(value.PreventedExecutionQty, value.QtyExponent, value.PreventedExecutionQtyNullValue())
	response.TradeGroupID = sbeOptionalInt64(value.TradeGroupID, value.TradeGroupIDNullValue())
	response.CounterOrderID = sbeOptionalInt64(value.CounterOrderID, value.CounterOrderIDNullValue())
	response.PreventedQty = sbeDecimal(value.PreventedQty, value.QtyExponent, value.PreventedQtyNullValue())
	response.LastPreventedQty = sbeDecimal(value.LastPreventedQty, value.QtyExponent, value.LastPreventedQtyNullValue())
	response.SOR = sbeOptionalBool(value.SOR)
	if value.ErrorCode != value.ErrorCodeNullValue() {
		response.ErrorCode = int64(value.ErrorCode)
	}

	response.NoMiscFees = uint64(len(value.MiscFees))
	response.MiscFees = make([]message.MiscFee, len(value.MiscFees))
	for i, fee := range value.MiscFees {
		response.MiscFees[i] = message.MiscFee{
			MiscFeeAmt:  sbeDecimal(fee.MiscFeeAmt, fee.Exponent, fee.MiscFeeAmtNullValue()),
			MiscFeeCurr: string(fee.MiscFeeCurr),
			MiscFeeType: message.MiscFeeType(sbeUintEnum(uint8(fee.MiscFeeType))),
		}
	}
	return response
}

func adaptSBEOrderCancelReject(value *fixsbe.OrderCancelReject) *message.OrderCancelReject {
	response := &message.OrderCancelReject{
		ClOrdID:            string(value.ClOrdID),
		OrigClOrdID:        string(value.OrigClOrdID),
		OrigClListID:       string(value.OrigClListID),
		Symbol:             string(value.Symbol),
		CancelRestrictions: message.CancelRestrictions(sbeUintEnum(uint8(value.CancelRestrictions))),
		CxlRejResponseTo:   message.CxlRejResponseTo(sbeCharEnum(byte(value.CxlRejResponseTo))),
		ErrorCode:          int64(value.ErrorCode),
		Text:               string(value.ErrorText),
	}
	response.OrderID = sbeOptionalInt64(value.OrderID, value.OrderIDNullValue())
	response.ListID = sbeOptionalIntString(value.ListID, value.ListIDNullValue())
	return response
}

func adaptSBEListStatus(value *fixsbe.ListStatus) *message.ListStatus {
	response := &message.ListStatus{
		ListID:           sbeOptionalIntString(value.ListID, value.ListIDNullValue()),
		ClListID:         string(value.ClListID),
		OrigClListID:     string(value.OrigClListID),
		ContingencyType:  message.ContingencyType(sbeUintEnum(uint8(value.ContingencyType))),
		ListStatusType:   message.ListStatusType(sbeUintEnum(uint8(value.ListStatusType))),
		ListOrderStatus:  message.ListOrderStatus(sbeUintEnum(uint8(value.ListOrderStatus))),
		ListRejectReason: message.ListRejectReason(sbeUintEnum(uint8(value.ListRejectReason))),
		TransactTime:     sbeTimestamp(value.TransactTime, value.TransactTimeNullValue()),
		NoOrders:         uint64(len(value.Orders)),
		Orders:           make([]message.ListStatusOrder, len(value.Orders)),
	}
	for i, order := range value.Orders {
		instructions := make([]message.ListTriggeringInstruction, len(order.ListTriggeringInstructions))
		for j, instruction := range order.ListTriggeringInstructions {
			instructions[j] = message.ListTriggeringInstruction{
				ListTriggerType:         message.ListTriggerType(sbeCharEnum(byte(instruction.ListTriggerType))),
				ListTriggerTriggerIndex: int64(instruction.ListTriggerTriggerIndex),
				ListTriggerAction:       message.ListTriggerAction(sbeCharEnum(byte(instruction.ListTriggerAction))),
			}
		}
		response.Orders[i] = message.ListStatusOrder{
			Symbol:                       string(order.Symbol),
			OrderID:                      sbeOptionalInt64(order.OrderID, order.OrderIDNullValue()),
			ClOrdID:                      string(order.ClOrdID),
			OrdRejReason:                 message.OrdRejReason(sbeUintEnum(uint8(order.OrdRejReason))),
			Text:                         string(order.ErrorText),
			NoListTriggeringInstructions: uint64(len(instructions)),
			ListTriggeringInstructions:   instructions,
		}
		if order.ErrorCode != order.ErrorCodeNullValue() {
			response.Orders[i].ErrorCode = int64(order.ErrorCode)
		}
	}
	// Preserve the existing top-level reject fields using the first rejected
	// order while retaining every order's own values in Orders.
	for _, order := range response.Orders {
		if order.ErrorCode != 0 || order.Text != "" || order.OrdRejReason != "" {
			response.Symbol = order.Symbol
			response.OrdRejReason = order.OrdRejReason
			response.ErrorCode = order.ErrorCode
			response.Text = order.Text
			break
		}
	}
	return response
}

func adaptSBEOrderMassCancelReport(value *fixsbe.OrderMassCancelReport) *message.OrderMassCancelReport {
	response := &message.OrderMassCancelReport{
		Symbol:                 string(value.Symbol),
		ClOrdID:                string(value.ClOrdID),
		MassCancelRequestType:  message.MassCancelRequestType(sbeCharEnum(byte(value.MassCancelRequestType))),
		MassCancelResponse:     message.MassCancelResponse(sbeCharEnum(byte(value.MassCancelResponse))),
		MassCancelRejectReason: message.MassCancelRejectReason(sbeUintEnum(uint8(value.MassCancelRejectReason))),
		TotalAffectedOrders:    sbeOptionalInt64(value.TotalAffectedOrders, value.TotalAffectedOrdersNullValue()),
		Text:                   string(value.ErrorText),
	}
	if value.ErrorCode != value.ErrorCodeNullValue() {
		response.ErrorCode = int64(value.ErrorCode)
	}
	return response
}

func adaptSBEOrderAmendReject(value *fixsbe.OrderAmendReject) *message.OrderAmendReject {
	return &message.OrderAmendReject{
		ClOrdID:     string(value.ClOrdID),
		OrigClOrdID: string(value.OrigClOrdID),
		OrderID:     sbeOptionalInt64(value.OrderID, value.OrderIDNullValue()),
		Symbol:      string(value.Symbol),
		OrderQty:    sbeDecimal(value.OrderQty, value.QtyExponent, value.OrderQtyNullValue()),
		ErrorCode:   int64(value.ErrorCode),
		Text:        string(value.ErrorText),
	}
}

func adaptSBELimitResponse(value *fixsbe.LimitResponse) *message.LimitResponse {
	response := &message.LimitResponse{
		ReqID:             string(value.ReqID),
		NoLimitIndicators: uint64(len(value.LimitIndicators)),
		LimitIndicators:   make([]message.LimitIndicator, len(value.LimitIndicators)),
	}
	for i, indicator := range value.LimitIndicators {
		response.LimitIndicators[i] = message.LimitIndicator{
			LimitType:                    message.LimitType(sbeCharEnum(byte(indicator.LimitType))),
			LimitCount:                   int64(indicator.LimitCount),
			LimitMax:                     int64(indicator.LimitMax),
			LimitResetIntervalResolution: message.LimitResetIntervalResolution(sbeCharEnum(byte(indicator.LimitResetIntervalResolution))),
		}
		if indicator.LimitResetInterval != indicator.LimitResetIntervalNullValue() {
			response.LimitIndicators[i].LimitResetInterval = int64(indicator.LimitResetInterval)
		}
	}
	return response
}

func adaptSBEInstrumentList(value *fixsbe.InstrumentList) *message.InstrumentList {
	response := &message.InstrumentList{
		InstrumentReqID: string(value.InstrumentReqID),
		NoRelatedSym:    uint64(len(value.RelatedSym)),
		Instruments:     make([]message.Instrument, len(value.RelatedSym)),
	}
	for i, instrument := range value.RelatedSym {
		response.Instruments[i] = message.Instrument{
			Symbol:                string(instrument.Symbol),
			Currency:              string(instrument.Currency),
			MinTradeVol:           sbeDecimal(instrument.MinTradeVol, instrument.QtyExponent, instrument.MinTradeVolNullValue()),
			MaxTradeVol:           sbeDecimal(instrument.MaxTradeVol, instrument.QtyExponent, instrument.MaxTradeVolNullValue()),
			MinQtyIncrement:       sbeDecimal(instrument.MinQtyIncrement, instrument.QtyExponent, instrument.MinQtyIncrementNullValue()),
			MarketMinTradeVol:     sbeDecimal(instrument.MarketMinTradeVol, instrument.QtyExponent, instrument.MarketMinTradeVolNullValue()),
			MarketMaxTradeVol:     sbeDecimal(instrument.MarketMaxTradeVol, instrument.QtyExponent, instrument.MarketMaxTradeVolNullValue()),
			MarketMinQtyIncrement: sbeDecimal(instrument.MarketMinQtyIncrement, instrument.QtyExponent, instrument.MarketMinQtyIncrementNullValue()),
			StartPriceRange:       sbeDecimal(instrument.StartPriceRange, instrument.PriceExponent, instrument.StartPriceRangeNullValue()),
			EndPriceRange:         sbeDecimal(instrument.EndPriceRange, instrument.PriceExponent, instrument.EndPriceRangeNullValue()),
			MinPriceIncrement:     sbeDecimal(instrument.MinPriceIncrement, instrument.PriceExponent, instrument.MinPriceIncrementNullValue()),
		}
	}
	return response
}

func adaptSBEMarketDataRequestReject(value *fixsbe.MarketDataRequestReject) *message.MarketDataRequestReject {
	response := &message.MarketDataRequestReject{
		MDReqID:        string(value.MDReqID),
		MDReqRejReason: message.MDReqRejReason(sbeCharEnum(byte(value.MDReqRejReason))),
		Text:           string(value.Text),
	}
	if value.ErrorCode != value.ErrorCodeNullValue() {
		response.ErrorCode = int64(value.ErrorCode)
	}
	return response
}

func adaptSBEMarketDataSnapshot(value *fixsbe.MarketDataSnapshot) *message.MarketDataSnapshot {
	entries := make([]message.MarketDataSnapshotEntry, 0, len(value.MDEntriesBids)+len(value.MDEntriesAsks))
	for _, entry := range value.MDEntriesBids {
		entries = append(entries, message.MarketDataSnapshotEntry{
			MDEntryType: message.MDEntryTypeBid,
			MDEntryPx:   sbeDecimal(entry.MDEntryPx, value.PriceExponent, entry.MDEntryPxNullValue()),
			MDEntrySize: sbeDecimal(entry.MDEntrySize, value.QtyExponent, entry.MDEntrySizeNullValue()),
		})
	}
	for _, entry := range value.MDEntriesAsks {
		entries = append(entries, message.MarketDataSnapshotEntry{
			MDEntryType: message.MDEntryTypeOffer,
			MDEntryPx:   sbeDecimal(entry.MDEntryPx, value.PriceExponent, entry.MDEntryPxNullValue()),
			MDEntrySize: sbeDecimal(entry.MDEntrySize, value.QtyExponent, entry.MDEntrySizeNullValue()),
		})
	}
	return &message.MarketDataSnapshot{
		Symbol:           string(value.Symbol),
		LastBookUpdateID: sbeOptionalInt64(value.LastBookUpdateID, value.LastBookUpdateIDNullValue()),
		NoMDEntries:      uint64(len(entries)),
		Entries:          entries,
	}
}

func adaptSBEMarketDataTrade(value *fixsbe.MarketDataIncrementalTrade) *message.MarketDataIncrementalRefresh {
	entries := make([]message.MarketDataIncrementalRefreshEntry, len(value.MDEntries))
	for i, entry := range value.MDEntries {
		entries[i] = message.MarketDataIncrementalRefreshEntry{
			MDUpdateAction: message.MDUpdateActionNew,
			MDEntryPx:      sbeDecimal(entry.MDEntryPx, value.PriceExponent, entry.MDEntryPxNullValue()),
			MDEntrySize:    sbeDecimal(entry.MDEntrySize, value.QtyExponent, entry.MDEntrySizeNullValue()),
			MDEntryType:    message.MDEntryTypeTrade,
			Symbol:         string(value.Symbol),
			TransactTime:   sbeTimestamp(value.TransactTime, value.TransactTimeNullValue()),
			TradeID:        entry.TradeID,
			AggressorSide:  message.AggressorSide(sbeCharEnum(byte(entry.AggressorSide))),
		}
	}
	return &message.MarketDataIncrementalRefresh{
		NoMDEntries: uint64(len(entries)),
		Entries:     entries,
	}
}

func adaptSBEMarketDataBookTicker(value *fixsbe.MarketDataIncrementalBookTicker) *message.MarketDataIncrementalRefresh {
	entries := make([]message.MarketDataIncrementalRefreshEntry, 0, len(value.MDEntriesBids)+len(value.MDEntriesAsks))
	for _, entry := range value.MDEntriesBids {
		entries = append(entries, newSBEBookEntry(
			message.MDEntryTypeBid,
			string(value.Symbol),
			0,
			value.LastBookUpdateID,
			entry.MDEntryPx,
			entry.MDEntrySize,
			entry.MDEntrySizeNullValue(),
			value.PriceExponent,
			value.QtyExponent,
		))
	}
	for _, entry := range value.MDEntriesAsks {
		entries = append(entries, newSBEBookEntry(
			message.MDEntryTypeOffer,
			string(value.Symbol),
			0,
			value.LastBookUpdateID,
			entry.MDEntryPx,
			entry.MDEntrySize,
			entry.MDEntrySizeNullValue(),
			value.PriceExponent,
			value.QtyExponent,
		))
	}
	return &message.MarketDataIncrementalRefresh{
		NoMDEntries: uint64(len(entries)),
		Entries:     entries,
	}
}

func adaptSBEMarketDataDepth(value *fixsbe.MarketDataIncrementalDepth) *message.MarketDataIncrementalRefresh {
	entries := make([]message.MarketDataIncrementalRefreshEntry, 0, len(value.MDEntriesBids)+len(value.MDEntriesAsks))
	for _, entry := range value.MDEntriesBids {
		entries = append(entries, newSBEBookEntry(
			message.MDEntryTypeBid,
			string(value.Symbol),
			value.FirstBookUpdateID,
			value.LastBookUpdateID,
			entry.MDEntryPx,
			entry.MDEntrySize,
			entry.MDEntrySizeNullValue(),
			value.PriceExponent,
			value.QtyExponent,
		))
	}
	for _, entry := range value.MDEntriesAsks {
		entries = append(entries, newSBEBookEntry(
			message.MDEntryTypeOffer,
			string(value.Symbol),
			value.FirstBookUpdateID,
			value.LastBookUpdateID,
			entry.MDEntryPx,
			entry.MDEntrySize,
			entry.MDEntrySizeNullValue(),
			value.PriceExponent,
			value.QtyExponent,
		))
	}
	return &message.MarketDataIncrementalRefresh{
		NoMDEntries: uint64(len(entries)),
		Entries:     entries,
	}
}

func newSBEBookEntry(
	entryType message.MDEntryType,
	symbol string,
	firstUpdateID int64,
	lastUpdateID int64,
	price int64,
	size int64,
	sizeNull int64,
	priceExponent int8,
	qtyExponent int8,
) message.MarketDataIncrementalRefreshEntry {
	action := message.MDUpdateActionNew
	if size == sizeNull {
		action = message.MDUpdateActionDelete
	}
	return message.MarketDataIncrementalRefreshEntry{
		MDUpdateAction:    action,
		MDEntryPx:         sbeDecimal(price, priceExponent, math.MinInt64),
		MDEntrySize:       sbeDecimal(size, qtyExponent, sizeNull),
		MDEntryType:       entryType,
		Symbol:            symbol,
		FirstBookUpdateID: firstUpdateID,
		LastBookUpdateID:  lastUpdateID,
	}
}

func sbeDecimal(mantissa int64, exponent int8, nullValue int64) float64 {
	if mantissa == nullValue {
		return 0
	}
	return float64(mantissa) * math.Pow10(int(exponent))
}

func sbeTimestamp(value, nullValue int64) time.Time {
	if value == nullValue {
		return time.Time{}
	}
	return time.UnixMicro(value).UTC()
}

func sbeOptionalInt64(value, nullValue int64) int64 {
	if value == nullValue {
		return 0
	}
	return value
}

func sbeOptionalIntString(value, nullValue int64) string {
	if value == nullValue {
		return ""
	}
	return strconv.FormatInt(value, 10)
}

func sbeOptionalBool(value fixsbe.BoolEnumEnum) *bool {
	switch value {
	case fixsbe.BoolEnum.True:
		result := true
		return &result
	case fixsbe.BoolEnum.False:
		result := false
		return &result
	default:
		return nil
	}
}

func sbeCharEnum(value byte) string {
	if value == 0 {
		return ""
	}
	return string([]byte{value})
}

func sbeUintEnum(value uint8) string {
	if value == math.MaxUint8 {
		return ""
	}
	return strconv.FormatUint(uint64(value), 10)
}
