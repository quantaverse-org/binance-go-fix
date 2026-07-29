package binance_go_fix

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/quantaverse-org/binance-go-fix/internal/fixsbe"
	"github.com/quantaverse-org/binance-go-fix/message"
)

type sbeRequestMessage interface {
	Encode(*fixsbe.SbeGoMarshaller, io.Writer, bool) error
	SbeBlockLength() uint16
	SbeTemplateId() uint16
	SbeSchemaId() uint16
	SbeSchemaVersion() uint16
}

// encodeSBERequest builds one complete SOFH + message header + request frame.
func encodeSBERequest(
	marshaller *fixsbe.SbeGoMarshaller,
	req message.Request,
	senderCompID string,
	targetCompID string,
	seqNum uint32,
	sendingTime time.Time,
) ([]byte, error) {
	body, err := adaptSBERequest(req, senderCompID, targetCompID, seqNum, sendingTime)
	if err != nil {
		return nil, err
	}

	var payload bytes.Buffer
	header := fixsbe.MessageHeader{
		BlockLength: body.SbeBlockLength(),
		TemplateId:  body.SbeTemplateId(),
		SchemaId:    body.SbeSchemaId(),
		Version:     body.SbeSchemaVersion(),
		SeqNum:      seqNum,
		SendingTime: sendingTime.UnixMicro(),
	}
	if err := header.Encode(marshaller, &payload); err != nil {
		return nil, fmt.Errorf("encode SBE message header: %w", err)
	}
	if err := body.Encode(marshaller, &payload, true); err != nil {
		return nil, fmt.Errorf("encode SBE template %d: %w", body.SbeTemplateId(), err)
	}

	messageLength := sbeSOFHLength + payload.Len()
	if messageLength > maxSBEMessageLength {
		return nil, fmt.Errorf("SBE message length %d exceeds limit %d", messageLength, maxSBEMessageLength)
	}

	frame := make([]byte, messageLength)
	binary.LittleEndian.PutUint32(frame[:4], uint32(messageLength))
	binary.LittleEndian.PutUint16(frame[4:sbeSOFHLength], sbeEncodingType)
	copy(frame[sbeSOFHLength:], payload.Bytes())
	return frame, nil
}

func adaptSBERequest(
	req message.Request,
	senderCompID string,
	targetCompID string,
	seqNum uint32,
	sendingTime time.Time,
) (sbeRequestMessage, error) {
	switch value := req.(type) {
	case *message.Heartbeat:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.Heartbeat")
		}
		testReqID, err := sbeString8("TestReqID", value.TestReqID, false)
		if err != nil {
			return nil, err
		}
		return &fixsbe.Heartbeat{TestReqID: testReqID}, nil
	case *message.TestRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.TestRequest")
		}
		testReqID, err := sbeString8("TestReqID", value.TestReqID, true)
		if err != nil {
			return nil, err
		}
		return &fixsbe.TestRequest{TestReqID: testReqID}, nil
	case *message.Logout:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.Logout")
		}
		text, err := sbeString16("Text", value.Text, false)
		if err != nil {
			return nil, err
		}
		return &fixsbe.Logout{Text: text}, nil
	case *message.LogonRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.LogonRequest")
		}
		return adaptSBELogon(value, senderCompID, targetCompID, seqNum, sendingTime)
	case *message.NewOrderSingle:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.NewOrderSingle")
		}
		return adaptSBENewOrderSingle(value)
	case *message.NewOrderList:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.NewOrderList")
		}
		return adaptSBENewOrderList(value)
	case *message.OrderCancelRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.OrderCancelRequest")
		}
		return adaptSBEOrderCancelRequest(value)
	case *message.OrderCancelRequestAndNewOrderSingle:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.OrderCancelRequestAndNewOrderSingle")
		}
		return adaptSBEOrderCancelAndNewOrder(value)
	case *message.OrderMassCancelRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.OrderMassCancelRequest")
		}
		return adaptSBEOrderMassCancelRequest(value)
	case *message.OrderAmendKeepPriorityRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.OrderAmendKeepPriorityRequest")
		}
		return adaptSBEOrderAmendKeepPriorityRequest(value)
	case *message.LimitQuery:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.LimitQuery")
		}
		reqID, err := sbeString8("ReqID", value.ReqID, true)
		if err != nil {
			return nil, err
		}
		return &fixsbe.LimitQuery{ReqID: reqID}, nil
	case *message.InstrumentListRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.InstrumentListRequest")
		}
		return adaptSBEInstrumentListRequest(value)
	case *message.MarketDataRequest:
		if value == nil {
			return nil, fmt.Errorf("encode SBE request: nil *message.MarketDataRequest")
		}
		return adaptSBEMarketDataRequest(value)
	default:
		return nil, fmt.Errorf("unsupported SBE request type %T", req)
	}
}

func adaptSBELogon(
	req *message.LogonRequest,
	senderCompID string,
	targetCompID string,
	seqNum uint32,
	sendingTime time.Time,
) (*fixsbe.Logon, error) {
	encryptMethod, err := sbeUint8Enum("EncryptMethod", string(req.EncryptMethod), false)
	if err != nil {
		return nil, err
	}
	heartBtInt, err := sbeRequiredUint32("HeartBtInt", req.HeartBtInt)
	if err != nil {
		return nil, err
	}
	messageHandling, err := sbeUint8Enum("MessageHandling", string(req.MessageHandling), false)
	if err != nil {
		return nil, err
	}

	responseMode := uint8(math.MaxUint8)
	if req.ResponseMode != nil {
		responseMode, err = sbeUint8Enum("ResponseMode", string(*req.ResponseMode), true)
		if err != nil {
			return nil, err
		}
	}

	sender, err := sbeString8("SenderCompID", senderCompID, true)
	if err != nil {
		return nil, err
	}
	target, err := sbeString8("TargetCompID", targetCompID, true)
	if err != nil {
		return nil, err
	}
	username, err := sbeString16("Username", req.Username, true)
	if err != nil {
		return nil, err
	}
	rawDataValue, err := message.SignLogonRawData(
		req.PrivateKey,
		senderCompID,
		targetCompID,
		seqNum,
		message.FormatTimestampMs(sendingTime),
	)
	if err != nil {
		return nil, err
	}
	rawData, err := sbeString16("RawData", rawDataValue, true)
	if err != nil {
		return nil, err
	}

	return &fixsbe.Logon{
		EncryptMethod:       encryptMethod,
		HeartBtInt:          heartBtInt,
		ResetSeqNumFlag:     sbeBool(req.ResetSeqNumFlag),
		MessageHandling:     fixsbe.MessageHandlingEnum(messageHandling),
		ResponseMode:        fixsbe.ResponseModeEnum(responseMode),
		ExecutionReportType: fixsbe.ExecutionReportType.NullValue,
		DropCopyFlag:        encodeSBEOptionalBool(req.DropCopyFlag),
		RecvWindow:          math.MaxUint32,
		SenderCompId:        sender,
		TargetCompId:        target,
		RawData:             rawData,
		Username:            username,
	}, nil
}

type encodedSBEOrderFields struct {
	PriceExponent            int8
	QtyExponent              int8
	OrderQty                 int64
	OrdType                  fixsbe.OrdTypeEnum
	ExecInst                 fixsbe.ExecInstEnum
	Price                    int64
	TriggerType              fixsbe.TriggerTypeEnum
	TriggerAction            fixsbe.TriggerActionEnum
	TriggerPrice             int64
	TriggerPriceType         fixsbe.TriggerPriceTypeEnum
	TriggerPriceDirection    fixsbe.TriggerPriceDirectionEnum
	TriggerTrailingDeltaBips uint64
	PegOffsetValue           uint8
	PegPriceType             fixsbe.PegPriceTypeEnum
	PegMoveType              fixsbe.PegMoveTypeEnum
	PegOffsetType            fixsbe.PegOffsetTypeEnum
	Side                     fixsbe.SideEnum
	TimeInForce              fixsbe.TimeInForceEnum
	MaxFloor                 int64
	CashOrderQty             int64
	TargetStrategy           int32
	StrategyID               int64
	SelfTradePreventionMode  fixsbe.SelfTradePreventionModeEnum
	SOR                      fixsbe.BoolEnumEnum
	ClOrdID                  []byte
	Symbol                   []byte
}

func encodeSBEOrderFields(order message.OrderFields, includeSOR bool) (encodedSBEOrderFields, error) {
	var result encodedSBEOrderFields

	price, err := newSBEFixedPoint(order.Price, order.TriggerPrice, order.CashOrderQty)
	if err != nil {
		return result, fmt.Errorf("encode price fields: %w", err)
	}
	qty, err := newSBEFixedPoint(order.OrderQty, order.MaxFloor)
	if err != nil {
		return result, fmt.Errorf("encode quantity fields: %w", err)
	}

	ordType, err := sbeCharEnumValue("OrdType", string(order.OrdType), true)
	if err != nil {
		return result, err
	}
	execInst, err := sbeCharEnumValue("ExecInst", string(order.ExecInst), false)
	if err != nil {
		return result, err
	}
	triggerType, err := sbeCharEnumValue("TriggerType", string(order.TriggerType), false)
	if err != nil {
		return result, err
	}
	triggerAction, err := sbeCharEnumValue("TriggerAction", string(order.TriggerAction), false)
	if err != nil {
		return result, err
	}
	triggerPriceType, err := sbeCharEnumValue("TriggerPriceType", string(order.TriggerPriceType), false)
	if err != nil {
		return result, err
	}
	triggerDirection, err := sbeCharEnumValue("TriggerPriceDirection", string(order.TriggerPriceDirection), false)
	if err != nil {
		return result, err
	}
	pegPriceType, err := sbeCharEnumValue("PegPriceType", string(order.PegPriceType), false)
	if err != nil {
		return result, err
	}
	pegMoveType, err := sbeUint8Enum("PegMoveType", string(order.PegMoveType), false)
	if err != nil {
		return result, err
	}
	pegOffsetType, err := sbeCharEnumValue("PegOffsetType", string(order.PegOffsetType), false)
	if err != nil {
		return result, err
	}
	side, err := sbeCharEnumValue("Side", string(order.Side), true)
	if err != nil {
		return result, err
	}
	timeInForce, err := sbeCharEnumValue("TimeInForce", string(order.TimeInForce), false)
	if err != nil {
		return result, err
	}
	selfTradePreventionMode, err := sbeCharEnumValue(
		"SelfTradePreventionMode",
		string(order.SelfTradePreventionMode),
		false,
	)
	if err != nil {
		return result, err
	}
	triggerTrailingDeltaBips, err := sbeOptionalUint64("TriggerTrailingDeltaBips", order.TriggerTrailingDeltaBips)
	if err != nil {
		return result, err
	}
	pegOffsetValue, err := sbeOptionalUint8Float("PegOffsetValue", order.PegOffsetValue)
	if err != nil {
		return result, err
	}
	targetStrategy, err := sbeOptionalInt32("TargetStrategy", order.TargetStrategy)
	if err != nil {
		return result, err
	}
	strategyID, err := encodeSBEOptionalInt64("StrategyID", order.StrategyID)
	if err != nil {
		return result, err
	}
	clOrdID, err := sbeString8("ClOrdID", order.ClOrdID, true)
	if err != nil {
		return result, err
	}
	symbol, err := sbeString8("Symbol", order.Symbol, true)
	if err != nil {
		return result, err
	}

	result.PriceExponent = price.exponent
	result.QtyExponent = qty.exponent
	result.OrderQty, err = qty.optionalMantissa(order.OrderQty)
	if err != nil {
		return result, fmt.Errorf("encode OrderQty: %w", err)
	}
	result.OrdType = fixsbe.OrdTypeEnum(ordType)
	result.ExecInst = fixsbe.ExecInstEnum(execInst)
	result.Price, err = price.optionalMantissa(order.Price)
	if err != nil {
		return result, fmt.Errorf("encode Price: %w", err)
	}
	result.TriggerType = fixsbe.TriggerTypeEnum(triggerType)
	result.TriggerAction = fixsbe.TriggerActionEnum(triggerAction)
	result.TriggerPrice, err = price.optionalMantissa(order.TriggerPrice)
	if err != nil {
		return result, fmt.Errorf("encode TriggerPrice: %w", err)
	}
	result.TriggerPriceType = fixsbe.TriggerPriceTypeEnum(triggerPriceType)
	result.TriggerPriceDirection = fixsbe.TriggerPriceDirectionEnum(triggerDirection)
	result.TriggerTrailingDeltaBips = triggerTrailingDeltaBips
	result.PegOffsetValue = pegOffsetValue
	result.PegPriceType = fixsbe.PegPriceTypeEnum(pegPriceType)
	result.PegMoveType = fixsbe.PegMoveTypeEnum(pegMoveType)
	result.PegOffsetType = fixsbe.PegOffsetTypeEnum(pegOffsetType)
	result.Side = fixsbe.SideEnum(side)
	result.TimeInForce = fixsbe.TimeInForceEnum(timeInForce)
	result.MaxFloor, err = qty.optionalMantissa(order.MaxFloor)
	if err != nil {
		return result, fmt.Errorf("encode MaxFloor: %w", err)
	}
	result.CashOrderQty, err = price.optionalMantissa(order.CashOrderQty)
	if err != nil {
		return result, fmt.Errorf("encode CashOrderQty: %w", err)
	}
	result.TargetStrategy = targetStrategy
	result.StrategyID = strategyID
	result.SelfTradePreventionMode = fixsbe.SelfTradePreventionModeEnum(selfTradePreventionMode)
	result.SOR = fixsbe.BoolEnum.NullValue
	if includeSOR {
		result.SOR = encodeSBEOptionalBool(order.SOR)
	}
	result.ClOrdID = clOrdID
	result.Symbol = symbol
	return result, nil
}

func adaptSBENewOrderSingle(req *message.NewOrderSingle) (*fixsbe.NewOrderSingle, error) {
	order, err := encodeSBEOrderFields(req.OrderFields, true)
	if err != nil {
		return nil, err
	}
	return &fixsbe.NewOrderSingle{
		PriceExponent:            order.PriceExponent,
		QtyExponent:              order.QtyExponent,
		OrderQty:                 order.OrderQty,
		OrdType:                  order.OrdType,
		ExecInst:                 order.ExecInst,
		Price:                    order.Price,
		TriggerType:              order.TriggerType,
		TriggerAction:            order.TriggerAction,
		TriggerPrice:             order.TriggerPrice,
		TriggerPriceType:         order.TriggerPriceType,
		TriggerPriceDirection:    order.TriggerPriceDirection,
		TriggerTrailingDeltaBips: order.TriggerTrailingDeltaBips,
		PegOffsetValue:           order.PegOffsetValue,
		PegPriceType:             order.PegPriceType,
		PegMoveType:              order.PegMoveType,
		PegOffsetType:            order.PegOffsetType,
		Side:                     order.Side,
		TimeInForce:              order.TimeInForce,
		MaxFloor:                 order.MaxFloor,
		CashOrderQty:             order.CashOrderQty,
		TargetStrategy:           order.TargetStrategy,
		StrategyID:               order.StrategyID,
		SelfTradePreventionMode:  order.SelfTradePreventionMode,
		SOR:                      order.SOR,
		ClOrdID:                  order.ClOrdID,
		Symbol:                   order.Symbol,
	}, nil
}

func adaptSBEOrderCancelAndNewOrder(
	req *message.OrderCancelRequestAndNewOrderSingle,
) (*fixsbe.OrderCancelRequestAndNewOrderSingle, error) {
	order, err := encodeSBEOrderFields(req.OrderFields, false)
	if err != nil {
		return nil, err
	}
	mode, err := sbeUint8Enum(
		"OrderCancelRequestAndNewOrderSingleMode",
		string(req.Mode),
		true,
	)
	if err != nil {
		return nil, err
	}
	rateLimitMode, err := sbeUint8Enum(
		"OrderRateLimitExceededMode",
		string(req.OrderRateLimitExceededMode),
		false,
	)
	if err != nil {
		return nil, err
	}
	orderID, err := encodeSBEOptionalInt64("OrderID", req.OrderID)
	if err != nil {
		return nil, err
	}
	cancelRestrictions, err := sbeUint8Enum(
		"CancelRestrictions",
		string(req.CancelRestrictions),
		false,
	)
	if err != nil {
		return nil, err
	}
	cancelClOrdID, err := sbeString8("CancelClOrdID", req.CancelClOrdID, false)
	if err != nil {
		return nil, err
	}
	origClOrdID, err := sbeString8("OrigClOrdID", req.OrigClOrdID, false)
	if err != nil {
		return nil, err
	}

	return &fixsbe.OrderCancelRequestAndNewOrderSingle{
		OrderCancelRequestAndNewOrderSingleMode: fixsbe.OrderCancelRequestAndNewOrderSingleModeEnum(mode),
		OrderRateLimitExceededMode:              fixsbe.OrderRateLimitExceededModeEnum(rateLimitMode),
		OrderID:                                 orderID,
		CancelRestrictions:                      fixsbe.CancelRestrictionsEnum(cancelRestrictions),
		PriceExponent:                           order.PriceExponent,
		QtyExponent:                             order.QtyExponent,
		OrderQty:                                order.OrderQty,
		OrdType:                                 order.OrdType,
		ExecInst:                                order.ExecInst,
		Price:                                   order.Price,
		TriggerType:                             order.TriggerType,
		TriggerAction:                           order.TriggerAction,
		TriggerPrice:                            order.TriggerPrice,
		TriggerPriceType:                        order.TriggerPriceType,
		TriggerPriceDirection:                   order.TriggerPriceDirection,
		TriggerTrailingDeltaBips:                order.TriggerTrailingDeltaBips,
		PegOffsetValue:                          order.PegOffsetValue,
		PegPriceType:                            order.PegPriceType,
		PegMoveType:                             order.PegMoveType,
		PegOffsetType:                           order.PegOffsetType,
		Side:                                    order.Side,
		TimeInForce:                             order.TimeInForce,
		MaxFloor:                                order.MaxFloor,
		CashOrderQty:                            order.CashOrderQty,
		TargetStrategy:                          order.TargetStrategy,
		StrategyID:                              order.StrategyID,
		SelfTradePreventionMode:                 order.SelfTradePreventionMode,
		CancelClOrdID:                           cancelClOrdID,
		OrigClOrdID:                             origClOrdID,
		ClOrdID:                                 order.ClOrdID,
		Symbol:                                  order.Symbol,
	}, nil
}

func adaptSBENewOrderList(req *message.NewOrderList) (*fixsbe.NewOrderList, error) {
	contingencyType, err := sbeUint8Enum("ContingencyType", string(req.ContingencyType), true)
	if err != nil {
		return nil, err
	}
	clListID, err := sbeString8("ClListID", req.ClListID, true)
	if err != nil {
		return nil, err
	}
	if len(req.Orders) > math.MaxUint8 {
		return nil, fmt.Errorf("Orders has %d entries, maximum is %d", len(req.Orders), math.MaxUint8)
	}

	orders := make([]fixsbe.NewOrderListOrders, len(req.Orders))
	for i, requestOrder := range req.Orders {
		order, err := encodeSBEOrderFields(requestOrder.OrderFields, false)
		if err != nil {
			return nil, fmt.Errorf("encode Orders[%d]: %w", i, err)
		}
		if len(requestOrder.ListTriggeringInstructions) > math.MaxUint8 {
			return nil, fmt.Errorf(
				"Orders[%d].ListTriggeringInstructions has %d entries, maximum is %d",
				i,
				len(requestOrder.ListTriggeringInstructions),
				math.MaxUint8,
			)
		}

		instructions := make(
			[]fixsbe.NewOrderListOrdersListTriggeringInstructions,
			len(requestOrder.ListTriggeringInstructions),
		)
		for j, instruction := range requestOrder.ListTriggeringInstructions {
			triggerType, err := sbeCharEnumValue(
				"ListTriggerType",
				string(instruction.ListTriggerType),
				true,
			)
			if err != nil {
				return nil, fmt.Errorf("encode Orders[%d].ListTriggeringInstructions[%d]: %w", i, j, err)
			}
			triggerAction, err := sbeCharEnumValue(
				"ListTriggerAction",
				string(instruction.ListTriggerAction),
				true,
			)
			if err != nil {
				return nil, fmt.Errorf("encode Orders[%d].ListTriggeringInstructions[%d]: %w", i, j, err)
			}
			triggerIndex, err := sbeRequiredUint8(
				"ListTriggerTriggerIndex",
				instruction.ListTriggerTriggerIndex,
			)
			if err != nil {
				return nil, fmt.Errorf("encode Orders[%d].ListTriggeringInstructions[%d]: %w", i, j, err)
			}
			instructions[j] = fixsbe.NewOrderListOrdersListTriggeringInstructions{
				ListTriggerType:         fixsbe.ListTriggerTypeEnum(triggerType),
				ListTriggerTriggerIndex: triggerIndex,
				ListTriggerAction:       fixsbe.ListTriggerActionEnum(triggerAction),
			}
		}

		orders[i] = fixsbe.NewOrderListOrders{
			PriceExponent:              order.PriceExponent,
			QtyExponent:                order.QtyExponent,
			OrderQty:                   order.OrderQty,
			OrdType:                    order.OrdType,
			ExecInst:                   order.ExecInst,
			Price:                      order.Price,
			TriggerType:                order.TriggerType,
			TriggerAction:              order.TriggerAction,
			TriggerPrice:               order.TriggerPrice,
			TriggerPriceType:           order.TriggerPriceType,
			TriggerPriceDirection:      order.TriggerPriceDirection,
			TriggerTrailingDeltaBips:   order.TriggerTrailingDeltaBips,
			PegOffsetValue:             order.PegOffsetValue,
			PegPriceType:               order.PegPriceType,
			PegMoveType:                order.PegMoveType,
			PegOffsetType:              order.PegOffsetType,
			Side:                       order.Side,
			TimeInForce:                order.TimeInForce,
			MaxFloor:                   order.MaxFloor,
			CashOrderQty:               order.CashOrderQty,
			TargetStrategy:             order.TargetStrategy,
			StrategyID:                 order.StrategyID,
			SelfTradePreventionMode:    order.SelfTradePreventionMode,
			ListTriggeringInstructions: instructions,
			ClOrdID:                    order.ClOrdID,
			Symbol:                     order.Symbol,
		}
	}

	return &fixsbe.NewOrderList{
		ContingencyType: fixsbe.ContingencyTypeEnum(contingencyType),
		OPO:             encodeSBEOptionalBool(req.OPO),
		Orders:          orders,
		ClListID:        clListID,
	}, nil
}

func adaptSBEOrderCancelRequest(req *message.OrderCancelRequest) (*fixsbe.OrderCancelRequest, error) {
	orderID, err := encodeSBEOptionalInt64("OrderID", req.OrderID)
	if err != nil {
		return nil, err
	}
	listID, err := sbeOptionalInt64String("ListID", req.ListID)
	if err != nil {
		return nil, err
	}
	cancelRestrictions, err := sbeUint8Enum(
		"CancelRestrictions",
		string(req.CancelRestrictions),
		false,
	)
	if err != nil {
		return nil, err
	}
	clOrdID, err := sbeString8("ClOrdID", req.ClOrdID, true)
	if err != nil {
		return nil, err
	}
	origClOrdID, err := sbeString8("OrigClOrdID", req.OrigClOrdID, false)
	if err != nil {
		return nil, err
	}
	origClListID, err := sbeString8("OrigClListID", req.OrigClListID, false)
	if err != nil {
		return nil, err
	}
	symbol, err := sbeString8("Symbol", req.Symbol, true)
	if err != nil {
		return nil, err
	}
	return &fixsbe.OrderCancelRequest{
		OrderID:            orderID,
		ListID:             listID,
		CancelRestrictions: fixsbe.CancelRestrictionsEnum(cancelRestrictions),
		ClOrdID:            clOrdID,
		OrigClOrdID:        origClOrdID,
		OrigClListID:       origClListID,
		Symbol:             symbol,
	}, nil
}

func adaptSBEOrderMassCancelRequest(
	req *message.OrderMassCancelRequest,
) (*fixsbe.OrderMassCancelRequest, error) {
	requestType, err := sbeCharEnumValue(
		"MassCancelRequestType",
		string(req.MassCancelRequestType),
		true,
	)
	if err != nil {
		return nil, err
	}
	symbol, err := sbeString8("Symbol", req.Symbol, true)
	if err != nil {
		return nil, err
	}
	clOrdID, err := sbeString8("ClOrdID", req.ClOrdID, true)
	if err != nil {
		return nil, err
	}
	return &fixsbe.OrderMassCancelRequest{
		MassCancelRequestType: fixsbe.MassCancelRequestTypeEnum(requestType),
		Symbol:                symbol,
		ClOrdID:               clOrdID,
	}, nil
}

func adaptSBEOrderAmendKeepPriorityRequest(
	req *message.OrderAmendKeepPriorityRequest,
) (*fixsbe.OrderAmendKeepPriorityRequest, error) {
	orderID, err := encodeSBEOptionalInt64("OrderID", req.OrderID)
	if err != nil {
		return nil, err
	}
	qty, err := newSBEFixedPoint(req.OrderQty)
	if err != nil {
		return nil, fmt.Errorf("encode OrderQty: %w", err)
	}
	orderQty, err := qty.mantissa(req.OrderQty)
	if err != nil {
		return nil, fmt.Errorf("encode OrderQty: %w", err)
	}
	clOrdID, err := sbeString8("ClOrdID", req.ClOrdID, true)
	if err != nil {
		return nil, err
	}
	origClOrdID, err := sbeString8("OrigClOrdID", req.OrigClOrdID, false)
	if err != nil {
		return nil, err
	}
	symbol, err := sbeString8("Symbol", req.Symbol, true)
	if err != nil {
		return nil, err
	}
	return &fixsbe.OrderAmendKeepPriorityRequest{
		OrderID:     orderID,
		QtyExponent: qty.exponent,
		OrderQty:    orderQty,
		ClOrdID:     clOrdID,
		OrigClOrdID: origClOrdID,
		Symbol:      symbol,
	}, nil
}

func adaptSBEInstrumentListRequest(
	req *message.InstrumentListRequest,
) (*fixsbe.InstrumentListRequest, error) {
	requestType, err := sbeUint8Enum(
		"InstrumentListRequestType",
		string(req.InstrumentListRequestType),
		true,
	)
	if err != nil {
		return nil, err
	}
	instrumentReqID, err := sbeString8("InstrumentReqID", req.InstrumentReqID, true)
	if err != nil {
		return nil, err
	}
	symbol, err := sbeString8("Symbol", req.Symbol, false)
	if err != nil {
		return nil, err
	}
	return &fixsbe.InstrumentListRequest{
		InstrumentListRequestType: fixsbe.InstrumentListRequestTypeEnum(requestType),
		InstrumentReqID:           instrumentReqID,
		Symbol:                    symbol,
	}, nil
}

func adaptSBEMarketDataRequest(req *message.MarketDataRequest) (*fixsbe.MarketDataRequest, error) {
	requestType, err := sbeCharEnumValue(
		"SubscriptionRequestType",
		string(req.SubscriptionRequestType),
		true,
	)
	if err != nil {
		return nil, err
	}
	marketDepth, err := sbeOptionalUint16("MarketDepth", req.MarketDepth)
	if err != nil {
		return nil, err
	}
	mdReqID, err := sbeString8("MDReqID", req.MDReqID, true)
	if err != nil {
		return nil, err
	}
	if len(req.Symbols) > math.MaxUint16 {
		return nil, fmt.Errorf("RelatedSym has %d entries, maximum is %d", len(req.Symbols), math.MaxUint16)
	}
	if len(req.MDEntryTypes) > math.MaxUint8 {
		return nil, fmt.Errorf("MDEntryTypes has %d entries, maximum is %d", len(req.MDEntryTypes), math.MaxUint8)
	}

	symbols := make([]fixsbe.MarketDataRequestRelatedSym, len(req.Symbols))
	for i, value := range req.Symbols {
		symbol, err := sbeString8("Symbol", value, true)
		if err != nil {
			return nil, fmt.Errorf("encode RelatedSym[%d]: %w", i, err)
		}
		symbols[i] = fixsbe.MarketDataRequestRelatedSym{Symbol: symbol}
	}
	entryTypes := make([]fixsbe.MarketDataRequestMDEntryTypes, len(req.MDEntryTypes))
	for i, value := range req.MDEntryTypes {
		entryType, err := sbeCharEnumValue("MDEntryType", string(value), true)
		if err != nil {
			return nil, fmt.Errorf("encode MDEntryTypes[%d]: %w", i, err)
		}
		entryTypes[i] = fixsbe.MarketDataRequestMDEntryTypes{
			MDEntryType: fixsbe.MdEntryTypeEnum(entryType),
		}
	}

	return &fixsbe.MarketDataRequest{
		SubscriptionRequestType: fixsbe.SubscriptionRequestTypeEnum(requestType),
		MarketDepth:             marketDepth,
		AggregatedBook:          encodeSBEOptionalBool(req.AggregatedBook),
		RelatedSym:              symbols,
		MDEntryTypes:            entryTypes,
		MDReqID:                 mdReqID,
	}, nil
}

type sbeFixedPoint struct {
	exponent int8
	scale    int
}

func newSBEFixedPoint(values ...float64) (sbeFixedPoint, error) {
	scale := 0
	for _, value := range values {
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return sbeFixedPoint{}, fmt.Errorf("non-finite value %v", value)
		}
		text := strconv.FormatFloat(math.Abs(value), 'f', -1, 64)
		if decimalPoint := strings.IndexByte(text, '.'); decimalPoint >= 0 {
			valueScale := len(text) - decimalPoint - 1
			if valueScale > scale {
				scale = valueScale
			}
		}
	}
	// -128 is the generated null exponent, so the most precise usable exponent is -127.
	if scale > math.MaxInt8 {
		return sbeFixedPoint{}, fmt.Errorf("decimal scale %d exceeds maximum %d", scale, math.MaxInt8)
	}
	return sbeFixedPoint{exponent: int8(-scale), scale: scale}, nil
}

func (e sbeFixedPoint) optionalMantissa(value float64) (int64, error) {
	if value == 0 {
		return math.MinInt64, nil
	}
	return e.mantissa(value)
}

func (e sbeFixedPoint) mantissa(value float64) (int64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("non-finite value %v", value)
	}
	text := strconv.FormatFloat(value, 'f', e.scale, 64)
	text = strings.Replace(text, ".", "", 1)
	mantissa, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value %v with exponent %d overflows int64 mantissa", value, e.exponent)
	}
	if mantissa == math.MinInt64 {
		return 0, fmt.Errorf("value %v encodes to reserved null mantissa", value)
	}
	return mantissa, nil
}

func sbeString8(field string, value string, required bool) ([]byte, error) {
	return sbeString(field, value, required, math.MaxUint8)
}

func sbeString16(field string, value string, required bool) ([]byte, error) {
	return sbeString(field, value, required, math.MaxUint16)
}

func sbeString(field string, value string, required bool, maxLength int) ([]byte, error) {
	if required && value == "" {
		return nil, fmt.Errorf("%s is required", field)
	}
	bytesValue := []byte(value)
	if len(bytesValue) > maxLength {
		return nil, fmt.Errorf("%s length %d exceeds maximum %d", field, len(bytesValue), maxLength)
	}
	return bytesValue, nil
}

func sbeCharEnumValue(field string, value string, required bool) (uint8, error) {
	if value == "" {
		if required {
			return 0, fmt.Errorf("%s is required", field)
		}
		return 0, nil
	}
	if len(value) != 1 {
		return 0, fmt.Errorf("%s must contain one ASCII character, got %q", field, value)
	}
	if value[0] > 0x7f {
		return 0, fmt.Errorf("%s must contain one ASCII character, got %q", field, value)
	}
	return value[0], nil
}

func sbeUint8Enum(field string, value string, required bool) (uint8, error) {
	if value == "" {
		if required {
			return 0, fmt.Errorf("%s is required", field)
		}
		return math.MaxUint8, nil
	}
	parsed, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("%s must be a uint8 enum, got %q", field, value)
	}
	if parsed == math.MaxUint8 {
		return 0, fmt.Errorf("%s uses reserved null value %d", field, parsed)
	}
	return uint8(parsed), nil
}

func sbeBool(value bool) fixsbe.BoolEnumEnum {
	if value {
		return fixsbe.BoolEnum.True
	}
	return fixsbe.BoolEnum.False
}

func encodeSBEOptionalBool(value *bool) fixsbe.BoolEnumEnum {
	if value == nil {
		return fixsbe.BoolEnum.NullValue
	}
	return sbeBool(*value)
}

func sbeRequiredUint8(field string, value int64) (uint8, error) {
	if value < 0 || value > math.MaxUint8 {
		return 0, fmt.Errorf("%s value %d exceeds uint8 range", field, value)
	}
	return uint8(value), nil
}

func sbeRequiredUint32(field string, value int64) (uint32, error) {
	if value < 0 || uint64(value) > math.MaxUint32 {
		return 0, fmt.Errorf("%s value %d exceeds uint32 range", field, value)
	}
	return uint32(value), nil
}

func sbeOptionalUint16(field string, value int64) (uint16, error) {
	if value == 0 {
		return math.MaxUint16, nil
	}
	if value < 0 || value >= math.MaxUint16 {
		return 0, fmt.Errorf("%s value %d exceeds optional uint16 range", field, value)
	}
	return uint16(value), nil
}

func sbeOptionalUint64(field string, value int64) (uint64, error) {
	if value == 0 {
		return math.MaxUint64, nil
	}
	if value < 0 {
		return 0, fmt.Errorf("%s value %d must not be negative", field, value)
	}
	return uint64(value), nil
}

func sbeOptionalUint8Float(field string, value float64) (uint8, error) {
	if value == 0 {
		return math.MaxUint8, nil
	}
	if math.IsNaN(value) || math.IsInf(value, 0) || value != math.Trunc(value) || value < 0 || value >= math.MaxUint8 {
		return 0, fmt.Errorf("%s value %v must be an integer in [0, %d]", field, value, math.MaxUint8-1)
	}
	return uint8(value), nil
}

func sbeOptionalInt32(field string, value int64) (int32, error) {
	if value == 0 {
		return math.MinInt32, nil
	}
	if value <= math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("%s value %d exceeds optional int32 range", field, value)
	}
	return int32(value), nil
}

func encodeSBEOptionalInt64(field string, value int64) (int64, error) {
	if value == 0 {
		return math.MinInt64, nil
	}
	if value == math.MinInt64 {
		return 0, fmt.Errorf("%s uses reserved null value %d", field, value)
	}
	return value, nil
}

func sbeOptionalInt64String(field string, value string) (int64, error) {
	if value == "" {
		return math.MinInt64, nil
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be an int64, got %q", field, value)
	}
	if parsed == math.MinInt64 {
		return 0, fmt.Errorf("%s uses reserved null value %d", field, parsed)
	}
	return parsed, nil
}
