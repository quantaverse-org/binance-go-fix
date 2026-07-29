// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type ExecutionReport struct {
	PriceExponent              int8
	QtyExponent                int8
	ExecID                     int64
	OrderID                    int64
	OrderQty                   int64
	OrdType                    OrdTypeEnum
	Side                       SideEnum
	ExecInst                   ExecInstEnum
	Price                      int64
	TriggerType                TriggerTypeEnum
	TriggerAction              TriggerActionEnum
	TriggerPrice               int64
	TriggerPriceType           TriggerPriceTypeEnum
	TriggerPriceDirection      TriggerPriceDirectionEnum
	TriggerTrailingDeltaBips   uint64
	PegOffsetValue             uint8
	PegPriceType               PegPriceTypeEnum
	PegMoveType                PegMoveTypeEnum
	PegOffsetType              PegOffsetTypeEnum
	PeggedPrice                int64
	TimeInForce                TimeInForceEnum
	TransactTime               int64
	OrderCreationTime          int64
	MaxFloor                   int64
	ListID                     int64
	CashOrderQty               int64
	TargetStrategy             int32
	StrategyID                 int64
	OrderCapacity              OrderCapacityEnum
	SelfTradePreventionMode    SelfTradePreventionModeEnum
	ExecType                   ExecTypeEnum
	CumQty                     int64
	LeavesQty                  int64
	CumQuoteQty                int64
	AggressorIndicator         BoolEnumEnum
	TradeID                    int64
	LastPx                     int64
	LastQty                    int64
	OrdStatus                  OrdStatusEnum
	SecondaryOrderID           int64
	SecondaryExternalAccountID int64
	AllocID                    int64
	MatchType                  MatchTypeEnum
	WorkingFloor               WorkingFloorEnum
	WorkingIndicator           BoolEnumEnum
	WorkingTime                int64
	TrailingTime               int64
	PreventedMatchID           int64
	PreventedExecutionPrice    int64
	PreventedExecutionQty      int64
	TradeGroupID               int64
	CounterOrderID             int64
	PreventedQty               int64
	LastPreventedQty           int64
	SOR                        BoolEnumEnum
	OrdRejReason               OrdRejReasonEnum
	ErrorCode                  int32
	ExpiryReason               ExpiryReasonEnum
	MiscFees                   []ExecutionReportMiscFees
	ClOrdID                    []uint8
	OrigClOrdID                []uint8
	Symbol                     []uint8
	SecondarySymbol            []uint8
	CounterSymbol              []uint8
	ErrorText                  []uint8
}
type ExecutionReportMiscFees struct {
	Exponent    int8
	MiscFeeAmt  int64
	MiscFeeType MiscFeeTypeEnum
	MiscFeeCurr []uint8
}

func (e *ExecutionReport) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt8(_w, e.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, e.QtyExponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ExecID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderQty); err != nil {
		return err
	}
	if err := e.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.Price); err != nil {
		return err
	}
	if err := e.TriggerType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TriggerAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TriggerPrice); err != nil {
		return err
	}
	if err := e.TriggerPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.TriggerPriceDirection.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, e.TriggerTrailingDeltaBips); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, e.PegOffsetValue); err != nil {
		return err
	}
	if err := e.PegPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PegMoveType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.PegOffsetType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PeggedPrice); err != nil {
		return err
	}
	if err := e.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.OrderCreationTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MaxFloor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ListID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.CashOrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.TargetStrategy); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.StrategyID); err != nil {
		return err
	}
	if err := e.OrderCapacity.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.SelfTradePreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.ExecType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.CumQty); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LeavesQty); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.CumQuoteQty); err != nil {
		return err
	}
	if err := e.AggressorIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TradeID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastQty); err != nil {
		return err
	}
	if err := e.OrdStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.SecondaryOrderID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.SecondaryExternalAccountID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.AllocID); err != nil {
		return err
	}
	if err := e.MatchType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.WorkingFloor.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.WorkingIndicator.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.WorkingTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TrailingTime); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PreventedMatchID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PreventedExecutionPrice); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PreventedExecutionQty); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TradeGroupID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.CounterOrderID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.PreventedQty); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.LastPreventedQty); err != nil {
		return err
	}
	if err := e.SOR.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.OrdRejReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.ErrorCode); err != nil {
		return err
	}
	if err := e.ExpiryReason.Encode(_m, _w); err != nil {
		return err
	}
	var MiscFeesBlockLength uint8 = 10
	if err := _m.WriteUint8(_w, MiscFeesBlockLength); err != nil {
		return err
	}
	var MiscFeesNumInGroup uint16 = uint16(len(e.MiscFees))
	if err := _m.WriteUint16(_w, MiscFeesNumInGroup); err != nil {
		return err
	}
	for i := range e.MiscFees {
		if err := e.MiscFees[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(e.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.OrigClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.OrigClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.SecondarySymbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.SecondarySymbol); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.CounterSymbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.CounterSymbol); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(e.ErrorText))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ErrorText); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReport) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.PriceExponentInActingVersion(actingVersion) {
		e.PriceExponent = e.PriceExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &e.PriceExponent); err != nil {
			return err
		}
	}
	if !e.QtyExponentInActingVersion(actingVersion) {
		e.QtyExponent = e.QtyExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &e.QtyExponent); err != nil {
			return err
		}
	}
	if !e.ExecIDInActingVersion(actingVersion) {
		e.ExecID = e.ExecIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.ExecID); err != nil {
			return err
		}
	}
	if !e.OrderIDInActingVersion(actingVersion) {
		e.OrderID = e.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderID); err != nil {
			return err
		}
	}
	if !e.OrderQtyInActingVersion(actingVersion) {
		e.OrderQty = e.OrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderQty); err != nil {
			return err
		}
	}
	if e.OrdTypeInActingVersion(actingVersion) {
		if err := e.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SideInActingVersion(actingVersion) {
		if err := e.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ExecInstInActingVersion(actingVersion) {
		if err := e.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.PriceInActingVersion(actingVersion) {
		e.Price = e.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.Price); err != nil {
			return err
		}
	}
	if e.TriggerTypeInActingVersion(actingVersion) {
		if err := e.TriggerType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.TriggerActionInActingVersion(actingVersion) {
		if err := e.TriggerAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TriggerPriceInActingVersion(actingVersion) {
		e.TriggerPrice = e.TriggerPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TriggerPrice); err != nil {
			return err
		}
	}
	if e.TriggerPriceTypeInActingVersion(actingVersion) {
		if err := e.TriggerPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.TriggerPriceDirectionInActingVersion(actingVersion) {
		if err := e.TriggerPriceDirection.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		e.TriggerTrailingDeltaBips = e.TriggerTrailingDeltaBipsNullValue()
	} else {
		if err := _m.ReadUint64(_r, &e.TriggerTrailingDeltaBips); err != nil {
			return err
		}
	}
	if !e.PegOffsetValueInActingVersion(actingVersion) {
		e.PegOffsetValue = e.PegOffsetValueNullValue()
	} else {
		if err := _m.ReadUint8(_r, &e.PegOffsetValue); err != nil {
			return err
		}
	}
	if e.PegPriceTypeInActingVersion(actingVersion) {
		if err := e.PegPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.PegMoveTypeInActingVersion(actingVersion) {
		if err := e.PegMoveType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.PegOffsetTypeInActingVersion(actingVersion) {
		if err := e.PegOffsetType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.PeggedPriceInActingVersion(actingVersion) {
		e.PeggedPrice = e.PeggedPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PeggedPrice); err != nil {
			return err
		}
	}
	if e.TimeInForceInActingVersion(actingVersion) {
		if err := e.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TransactTimeInActingVersion(actingVersion) {
		e.TransactTime = e.TransactTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TransactTime); err != nil {
			return err
		}
	}
	if !e.OrderCreationTimeInActingVersion(actingVersion) {
		e.OrderCreationTime = e.OrderCreationTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderCreationTime); err != nil {
			return err
		}
	}
	if !e.MaxFloorInActingVersion(actingVersion) {
		e.MaxFloor = e.MaxFloorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MaxFloor); err != nil {
			return err
		}
	}
	if !e.ListIDInActingVersion(actingVersion) {
		e.ListID = e.ListIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.ListID); err != nil {
			return err
		}
	}
	if !e.CashOrderQtyInActingVersion(actingVersion) {
		e.CashOrderQty = e.CashOrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CashOrderQty); err != nil {
			return err
		}
	}
	if !e.TargetStrategyInActingVersion(actingVersion) {
		e.TargetStrategy = e.TargetStrategyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.TargetStrategy); err != nil {
			return err
		}
	}
	if !e.StrategyIDInActingVersion(actingVersion) {
		e.StrategyID = e.StrategyIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.StrategyID); err != nil {
			return err
		}
	}
	if e.OrderCapacityInActingVersion(actingVersion) {
		if err := e.OrderCapacity.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.SelfTradePreventionModeInActingVersion(actingVersion) {
		if err := e.SelfTradePreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.ExecTypeInActingVersion(actingVersion) {
		if err := e.ExecType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.CumQtyInActingVersion(actingVersion) {
		e.CumQty = e.CumQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CumQty); err != nil {
			return err
		}
	}
	if !e.LeavesQtyInActingVersion(actingVersion) {
		e.LeavesQty = e.LeavesQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LeavesQty); err != nil {
			return err
		}
	}
	if !e.CumQuoteQtyInActingVersion(actingVersion) {
		e.CumQuoteQty = e.CumQuoteQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CumQuoteQty); err != nil {
			return err
		}
	}
	if e.AggressorIndicatorInActingVersion(actingVersion) {
		if err := e.AggressorIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.TradeIDInActingVersion(actingVersion) {
		e.TradeID = e.TradeIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TradeID); err != nil {
			return err
		}
	}
	if !e.LastPxInActingVersion(actingVersion) {
		e.LastPx = e.LastPxNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastPx); err != nil {
			return err
		}
	}
	if !e.LastQtyInActingVersion(actingVersion) {
		e.LastQty = e.LastQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastQty); err != nil {
			return err
		}
	}
	if e.OrdStatusInActingVersion(actingVersion) {
		if err := e.OrdStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.SecondaryOrderIDInActingVersion(actingVersion) {
		e.SecondaryOrderID = e.SecondaryOrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.SecondaryOrderID); err != nil {
			return err
		}
	}
	if !e.SecondaryExternalAccountIDInActingVersion(actingVersion) {
		e.SecondaryExternalAccountID = e.SecondaryExternalAccountIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.SecondaryExternalAccountID); err != nil {
			return err
		}
	}
	if !e.AllocIDInActingVersion(actingVersion) {
		e.AllocID = e.AllocIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.AllocID); err != nil {
			return err
		}
	}
	if e.MatchTypeInActingVersion(actingVersion) {
		if err := e.MatchType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.WorkingFloorInActingVersion(actingVersion) {
		if err := e.WorkingFloor.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.WorkingIndicatorInActingVersion(actingVersion) {
		if err := e.WorkingIndicator.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.WorkingTimeInActingVersion(actingVersion) {
		e.WorkingTime = e.WorkingTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.WorkingTime); err != nil {
			return err
		}
	}
	if !e.TrailingTimeInActingVersion(actingVersion) {
		e.TrailingTime = e.TrailingTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TrailingTime); err != nil {
			return err
		}
	}
	if !e.PreventedMatchIDInActingVersion(actingVersion) {
		e.PreventedMatchID = e.PreventedMatchIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PreventedMatchID); err != nil {
			return err
		}
	}
	if !e.PreventedExecutionPriceInActingVersion(actingVersion) {
		e.PreventedExecutionPrice = e.PreventedExecutionPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PreventedExecutionPrice); err != nil {
			return err
		}
	}
	if !e.PreventedExecutionQtyInActingVersion(actingVersion) {
		e.PreventedExecutionQty = e.PreventedExecutionQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PreventedExecutionQty); err != nil {
			return err
		}
	}
	if !e.TradeGroupIDInActingVersion(actingVersion) {
		e.TradeGroupID = e.TradeGroupIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TradeGroupID); err != nil {
			return err
		}
	}
	if !e.CounterOrderIDInActingVersion(actingVersion) {
		e.CounterOrderID = e.CounterOrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.CounterOrderID); err != nil {
			return err
		}
	}
	if !e.PreventedQtyInActingVersion(actingVersion) {
		e.PreventedQty = e.PreventedQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.PreventedQty); err != nil {
			return err
		}
	}
	if !e.LastPreventedQtyInActingVersion(actingVersion) {
		e.LastPreventedQty = e.LastPreventedQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.LastPreventedQty); err != nil {
			return err
		}
	}
	if e.SORInActingVersion(actingVersion) {
		if err := e.SOR.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.OrdRejReasonInActingVersion(actingVersion) {
		if err := e.OrdRejReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !e.ErrorCodeInActingVersion(actingVersion) {
		e.ErrorCode = e.ErrorCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &e.ErrorCode); err != nil {
			return err
		}
	}
	if e.ExpiryReasonInActingVersion(actingVersion) {
		if err := e.ExpiryReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.MiscFeesInActingVersion(actingVersion) {
		var MiscFeesBlockLength uint8
		if err := _m.ReadUint8(_r, &MiscFeesBlockLength); err != nil {
			return err
		}
		var MiscFeesNumInGroup uint16
		if err := _m.ReadUint16(_r, &MiscFeesNumInGroup); err != nil {
			return err
		}
		if cap(e.MiscFees) < int(MiscFeesNumInGroup) {
			e.MiscFees = make([]ExecutionReportMiscFees, MiscFeesNumInGroup)
		}
		e.MiscFees = e.MiscFees[:MiscFeesNumInGroup]
		for i := range e.MiscFees {
			if err := e.MiscFees[i].Decode(_m, _r, actingVersion, uint(MiscFeesBlockLength)); err != nil {
				return err
			}
		}
	}

	if e.ClOrdIDInActingVersion(actingVersion) {
		var ClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &ClOrdIDLength); err != nil {
			return err
		}
		if cap(e.ClOrdID) < int(ClOrdIDLength) {
			e.ClOrdID = make([]uint8, ClOrdIDLength)
		}
		e.ClOrdID = e.ClOrdID[:ClOrdIDLength]
		if err := _m.ReadBytes(_r, e.ClOrdID); err != nil {
			return err
		}
	}

	if e.OrigClOrdIDInActingVersion(actingVersion) {
		var OrigClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &OrigClOrdIDLength); err != nil {
			return err
		}
		if cap(e.OrigClOrdID) < int(OrigClOrdIDLength) {
			e.OrigClOrdID = make([]uint8, OrigClOrdIDLength)
		}
		e.OrigClOrdID = e.OrigClOrdID[:OrigClOrdIDLength]
		if err := _m.ReadBytes(_r, e.OrigClOrdID); err != nil {
			return err
		}
	}

	if e.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(e.Symbol) < int(SymbolLength) {
			e.Symbol = make([]uint8, SymbolLength)
		}
		e.Symbol = e.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, e.Symbol); err != nil {
			return err
		}
	}

	if e.SecondarySymbolInActingVersion(actingVersion) {
		var SecondarySymbolLength uint8
		if err := _m.ReadUint8(_r, &SecondarySymbolLength); err != nil {
			return err
		}
		if cap(e.SecondarySymbol) < int(SecondarySymbolLength) {
			e.SecondarySymbol = make([]uint8, SecondarySymbolLength)
		}
		e.SecondarySymbol = e.SecondarySymbol[:SecondarySymbolLength]
		if err := _m.ReadBytes(_r, e.SecondarySymbol); err != nil {
			return err
		}
	}

	if e.CounterSymbolInActingVersion(actingVersion) {
		var CounterSymbolLength uint8
		if err := _m.ReadUint8(_r, &CounterSymbolLength); err != nil {
			return err
		}
		if cap(e.CounterSymbol) < int(CounterSymbolLength) {
			e.CounterSymbol = make([]uint8, CounterSymbolLength)
		}
		e.CounterSymbol = e.CounterSymbol[:CounterSymbolLength]
		if err := _m.ReadBytes(_r, e.CounterSymbol); err != nil {
			return err
		}
	}

	if e.ErrorTextInActingVersion(actingVersion) {
		var ErrorTextLength uint16
		if err := _m.ReadUint16(_r, &ErrorTextLength); err != nil {
			return err
		}
		if cap(e.ErrorText) < int(ErrorTextLength) {
			e.ErrorText = make([]uint8, ErrorTextLength)
		}
		e.ErrorText = e.ErrorText[:ErrorTextLength]
		if err := _m.ReadBytes(_r, e.ErrorText); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := e.RangeCheck(actingVersion, e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionReport) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.PriceExponentInActingVersion(actingVersion) {
		if e.PriceExponent < e.PriceExponentMinValue() || e.PriceExponent > e.PriceExponentMaxValue() {
			return fmt.Errorf("Range check failed on e.PriceExponent (%v < %v > %v)", e.PriceExponentMinValue(), e.PriceExponent, e.PriceExponentMaxValue())
		}
	}
	if e.QtyExponentInActingVersion(actingVersion) {
		if e.QtyExponent < e.QtyExponentMinValue() || e.QtyExponent > e.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on e.QtyExponent (%v < %v > %v)", e.QtyExponentMinValue(), e.QtyExponent, e.QtyExponentMaxValue())
		}
	}
	if e.ExecIDInActingVersion(actingVersion) {
		if e.ExecID != e.ExecIDNullValue() && (e.ExecID < e.ExecIDMinValue() || e.ExecID > e.ExecIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.ExecID (%v < %v > %v)", e.ExecIDMinValue(), e.ExecID, e.ExecIDMaxValue())
		}
	}
	if e.OrderIDInActingVersion(actingVersion) {
		if e.OrderID != e.OrderIDNullValue() && (e.OrderID < e.OrderIDMinValue() || e.OrderID > e.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrderID (%v < %v > %v)", e.OrderIDMinValue(), e.OrderID, e.OrderIDMaxValue())
		}
	}
	if e.OrderQtyInActingVersion(actingVersion) {
		if e.OrderQty != e.OrderQtyNullValue() && (e.OrderQty < e.OrderQtyMinValue() || e.OrderQty > e.OrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrderQty (%v < %v > %v)", e.OrderQtyMinValue(), e.OrderQty, e.OrderQtyMaxValue())
		}
	}
	if err := e.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ExecInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.PriceInActingVersion(actingVersion) {
		if e.Price != e.PriceNullValue() && (e.Price < e.PriceMinValue() || e.Price > e.PriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.Price (%v < %v > %v)", e.PriceMinValue(), e.Price, e.PriceMaxValue())
		}
	}
	if err := e.TriggerType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.TriggerAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.TriggerPriceInActingVersion(actingVersion) {
		if e.TriggerPrice != e.TriggerPriceNullValue() && (e.TriggerPrice < e.TriggerPriceMinValue() || e.TriggerPrice > e.TriggerPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.TriggerPrice (%v < %v > %v)", e.TriggerPriceMinValue(), e.TriggerPrice, e.TriggerPriceMaxValue())
		}
	}
	if err := e.TriggerPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.TriggerPriceDirection.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		if e.TriggerTrailingDeltaBips != e.TriggerTrailingDeltaBipsNullValue() && (e.TriggerTrailingDeltaBips < e.TriggerTrailingDeltaBipsMinValue() || e.TriggerTrailingDeltaBips > e.TriggerTrailingDeltaBipsMaxValue()) {
			return fmt.Errorf("Range check failed on e.TriggerTrailingDeltaBips (%v < %v > %v)", e.TriggerTrailingDeltaBipsMinValue(), e.TriggerTrailingDeltaBips, e.TriggerTrailingDeltaBipsMaxValue())
		}
	}
	if e.PegOffsetValueInActingVersion(actingVersion) {
		if e.PegOffsetValue != e.PegOffsetValueNullValue() && (e.PegOffsetValue < e.PegOffsetValueMinValue() || e.PegOffsetValue > e.PegOffsetValueMaxValue()) {
			return fmt.Errorf("Range check failed on e.PegOffsetValue (%v < %v > %v)", e.PegOffsetValueMinValue(), e.PegOffsetValue, e.PegOffsetValueMaxValue())
		}
	}
	if err := e.PegPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PegMoveType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.PegOffsetType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.PeggedPriceInActingVersion(actingVersion) {
		if e.PeggedPrice != e.PeggedPriceNullValue() && (e.PeggedPrice < e.PeggedPriceMinValue() || e.PeggedPrice > e.PeggedPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.PeggedPrice (%v < %v > %v)", e.PeggedPriceMinValue(), e.PeggedPrice, e.PeggedPriceMaxValue())
		}
	}
	if err := e.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.TransactTimeInActingVersion(actingVersion) {
		if e.TransactTime != e.TransactTimeNullValue() && (e.TransactTime < e.TransactTimeMinValue() || e.TransactTime > e.TransactTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.TransactTime (%v < %v > %v)", e.TransactTimeMinValue(), e.TransactTime, e.TransactTimeMaxValue())
		}
	}
	if e.OrderCreationTimeInActingVersion(actingVersion) {
		if e.OrderCreationTime != e.OrderCreationTimeNullValue() && (e.OrderCreationTime < e.OrderCreationTimeMinValue() || e.OrderCreationTime > e.OrderCreationTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrderCreationTime (%v < %v > %v)", e.OrderCreationTimeMinValue(), e.OrderCreationTime, e.OrderCreationTimeMaxValue())
		}
	}
	if e.MaxFloorInActingVersion(actingVersion) {
		if e.MaxFloor != e.MaxFloorNullValue() && (e.MaxFloor < e.MaxFloorMinValue() || e.MaxFloor > e.MaxFloorMaxValue()) {
			return fmt.Errorf("Range check failed on e.MaxFloor (%v < %v > %v)", e.MaxFloorMinValue(), e.MaxFloor, e.MaxFloorMaxValue())
		}
	}
	if e.ListIDInActingVersion(actingVersion) {
		if e.ListID != e.ListIDNullValue() && (e.ListID < e.ListIDMinValue() || e.ListID > e.ListIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.ListID (%v < %v > %v)", e.ListIDMinValue(), e.ListID, e.ListIDMaxValue())
		}
	}
	if e.CashOrderQtyInActingVersion(actingVersion) {
		if e.CashOrderQty != e.CashOrderQtyNullValue() && (e.CashOrderQty < e.CashOrderQtyMinValue() || e.CashOrderQty > e.CashOrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.CashOrderQty (%v < %v > %v)", e.CashOrderQtyMinValue(), e.CashOrderQty, e.CashOrderQtyMaxValue())
		}
	}
	if e.TargetStrategyInActingVersion(actingVersion) {
		if e.TargetStrategy != e.TargetStrategyNullValue() && (e.TargetStrategy < e.TargetStrategyMinValue() || e.TargetStrategy > e.TargetStrategyMaxValue()) {
			return fmt.Errorf("Range check failed on e.TargetStrategy (%v < %v > %v)", e.TargetStrategyMinValue(), e.TargetStrategy, e.TargetStrategyMaxValue())
		}
	}
	if e.StrategyIDInActingVersion(actingVersion) {
		if e.StrategyID != e.StrategyIDNullValue() && (e.StrategyID < e.StrategyIDMinValue() || e.StrategyID > e.StrategyIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.StrategyID (%v < %v > %v)", e.StrategyIDMinValue(), e.StrategyID, e.StrategyIDMaxValue())
		}
	}
	if err := e.OrderCapacity.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.SelfTradePreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.ExecType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.CumQtyInActingVersion(actingVersion) {
		if e.CumQty < e.CumQtyMinValue() || e.CumQty > e.CumQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.CumQty (%v < %v > %v)", e.CumQtyMinValue(), e.CumQty, e.CumQtyMaxValue())
		}
	}
	if e.LeavesQtyInActingVersion(actingVersion) {
		if e.LeavesQty != e.LeavesQtyNullValue() && (e.LeavesQty < e.LeavesQtyMinValue() || e.LeavesQty > e.LeavesQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.LeavesQty (%v < %v > %v)", e.LeavesQtyMinValue(), e.LeavesQty, e.LeavesQtyMaxValue())
		}
	}
	if e.CumQuoteQtyInActingVersion(actingVersion) {
		if e.CumQuoteQty != e.CumQuoteQtyNullValue() && (e.CumQuoteQty < e.CumQuoteQtyMinValue() || e.CumQuoteQty > e.CumQuoteQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.CumQuoteQty (%v < %v > %v)", e.CumQuoteQtyMinValue(), e.CumQuoteQty, e.CumQuoteQtyMaxValue())
		}
	}
	if err := e.AggressorIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.TradeIDInActingVersion(actingVersion) {
		if e.TradeID != e.TradeIDNullValue() && (e.TradeID < e.TradeIDMinValue() || e.TradeID > e.TradeIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.TradeID (%v < %v > %v)", e.TradeIDMinValue(), e.TradeID, e.TradeIDMaxValue())
		}
	}
	if e.LastPxInActingVersion(actingVersion) {
		if e.LastPx != e.LastPxNullValue() && (e.LastPx < e.LastPxMinValue() || e.LastPx > e.LastPxMaxValue()) {
			return fmt.Errorf("Range check failed on e.LastPx (%v < %v > %v)", e.LastPxMinValue(), e.LastPx, e.LastPxMaxValue())
		}
	}
	if e.LastQtyInActingVersion(actingVersion) {
		if e.LastQty < e.LastQtyMinValue() || e.LastQty > e.LastQtyMaxValue() {
			return fmt.Errorf("Range check failed on e.LastQty (%v < %v > %v)", e.LastQtyMinValue(), e.LastQty, e.LastQtyMaxValue())
		}
	}
	if err := e.OrdStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.SecondaryOrderIDInActingVersion(actingVersion) {
		if e.SecondaryOrderID != e.SecondaryOrderIDNullValue() && (e.SecondaryOrderID < e.SecondaryOrderIDMinValue() || e.SecondaryOrderID > e.SecondaryOrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.SecondaryOrderID (%v < %v > %v)", e.SecondaryOrderIDMinValue(), e.SecondaryOrderID, e.SecondaryOrderIDMaxValue())
		}
	}
	if e.SecondaryExternalAccountIDInActingVersion(actingVersion) {
		if e.SecondaryExternalAccountID != e.SecondaryExternalAccountIDNullValue() && (e.SecondaryExternalAccountID < e.SecondaryExternalAccountIDMinValue() || e.SecondaryExternalAccountID > e.SecondaryExternalAccountIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.SecondaryExternalAccountID (%v < %v > %v)", e.SecondaryExternalAccountIDMinValue(), e.SecondaryExternalAccountID, e.SecondaryExternalAccountIDMaxValue())
		}
	}
	if e.AllocIDInActingVersion(actingVersion) {
		if e.AllocID != e.AllocIDNullValue() && (e.AllocID < e.AllocIDMinValue() || e.AllocID > e.AllocIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.AllocID (%v < %v > %v)", e.AllocIDMinValue(), e.AllocID, e.AllocIDMaxValue())
		}
	}
	if err := e.MatchType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.WorkingFloor.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.WorkingIndicator.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.WorkingTimeInActingVersion(actingVersion) {
		if e.WorkingTime != e.WorkingTimeNullValue() && (e.WorkingTime < e.WorkingTimeMinValue() || e.WorkingTime > e.WorkingTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.WorkingTime (%v < %v > %v)", e.WorkingTimeMinValue(), e.WorkingTime, e.WorkingTimeMaxValue())
		}
	}
	if e.TrailingTimeInActingVersion(actingVersion) {
		if e.TrailingTime != e.TrailingTimeNullValue() && (e.TrailingTime < e.TrailingTimeMinValue() || e.TrailingTime > e.TrailingTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.TrailingTime (%v < %v > %v)", e.TrailingTimeMinValue(), e.TrailingTime, e.TrailingTimeMaxValue())
		}
	}
	if e.PreventedMatchIDInActingVersion(actingVersion) {
		if e.PreventedMatchID != e.PreventedMatchIDNullValue() && (e.PreventedMatchID < e.PreventedMatchIDMinValue() || e.PreventedMatchID > e.PreventedMatchIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.PreventedMatchID (%v < %v > %v)", e.PreventedMatchIDMinValue(), e.PreventedMatchID, e.PreventedMatchIDMaxValue())
		}
	}
	if e.PreventedExecutionPriceInActingVersion(actingVersion) {
		if e.PreventedExecutionPrice != e.PreventedExecutionPriceNullValue() && (e.PreventedExecutionPrice < e.PreventedExecutionPriceMinValue() || e.PreventedExecutionPrice > e.PreventedExecutionPriceMaxValue()) {
			return fmt.Errorf("Range check failed on e.PreventedExecutionPrice (%v < %v > %v)", e.PreventedExecutionPriceMinValue(), e.PreventedExecutionPrice, e.PreventedExecutionPriceMaxValue())
		}
	}
	if e.PreventedExecutionQtyInActingVersion(actingVersion) {
		if e.PreventedExecutionQty != e.PreventedExecutionQtyNullValue() && (e.PreventedExecutionQty < e.PreventedExecutionQtyMinValue() || e.PreventedExecutionQty > e.PreventedExecutionQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.PreventedExecutionQty (%v < %v > %v)", e.PreventedExecutionQtyMinValue(), e.PreventedExecutionQty, e.PreventedExecutionQtyMaxValue())
		}
	}
	if e.TradeGroupIDInActingVersion(actingVersion) {
		if e.TradeGroupID != e.TradeGroupIDNullValue() && (e.TradeGroupID < e.TradeGroupIDMinValue() || e.TradeGroupID > e.TradeGroupIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.TradeGroupID (%v < %v > %v)", e.TradeGroupIDMinValue(), e.TradeGroupID, e.TradeGroupIDMaxValue())
		}
	}
	if e.CounterOrderIDInActingVersion(actingVersion) {
		if e.CounterOrderID != e.CounterOrderIDNullValue() && (e.CounterOrderID < e.CounterOrderIDMinValue() || e.CounterOrderID > e.CounterOrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.CounterOrderID (%v < %v > %v)", e.CounterOrderIDMinValue(), e.CounterOrderID, e.CounterOrderIDMaxValue())
		}
	}
	if e.PreventedQtyInActingVersion(actingVersion) {
		if e.PreventedQty != e.PreventedQtyNullValue() && (e.PreventedQty < e.PreventedQtyMinValue() || e.PreventedQty > e.PreventedQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.PreventedQty (%v < %v > %v)", e.PreventedQtyMinValue(), e.PreventedQty, e.PreventedQtyMaxValue())
		}
	}
	if e.LastPreventedQtyInActingVersion(actingVersion) {
		if e.LastPreventedQty != e.LastPreventedQtyNullValue() && (e.LastPreventedQty < e.LastPreventedQtyMinValue() || e.LastPreventedQty > e.LastPreventedQtyMaxValue()) {
			return fmt.Errorf("Range check failed on e.LastPreventedQty (%v < %v > %v)", e.LastPreventedQtyMinValue(), e.LastPreventedQty, e.LastPreventedQtyMaxValue())
		}
	}
	if err := e.SOR.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.OrdRejReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if e.ErrorCodeInActingVersion(actingVersion) {
		if e.ErrorCode != e.ErrorCodeNullValue() && (e.ErrorCode < e.ErrorCodeMinValue() || e.ErrorCode > e.ErrorCodeMaxValue()) {
			return fmt.Errorf("Range check failed on e.ErrorCode (%v < %v > %v)", e.ErrorCodeMinValue(), e.ErrorCode, e.ErrorCodeMaxValue())
		}
	}
	if err := e.ExpiryReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for i := range e.MiscFees {
		if err := e.MiscFees[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(e.ClOrdID[:]) {
		return errors.New("e.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(e.OrigClOrdID[:]) {
		return errors.New("e.OrigClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(e.Symbol[:]) {
		return errors.New("e.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(e.SecondarySymbol[:]) {
		return errors.New("e.SecondarySymbol failed UTF-8 validation")
	}
	if !utf8.Valid(e.CounterSymbol[:]) {
		return errors.New("e.CounterSymbol failed UTF-8 validation")
	}
	if !utf8.Valid(e.ErrorText[:]) {
		return errors.New("e.ErrorText failed UTF-8 validation")
	}
	return nil
}

func ExecutionReportInit(e *ExecutionReport) {
	e.ExecID = math.MinInt64
	e.OrderID = math.MinInt64
	e.OrderQty = math.MinInt64
	e.Price = math.MinInt64
	e.TriggerPrice = math.MinInt64
	e.TriggerTrailingDeltaBips = math.MaxUint64
	e.PegOffsetValue = math.MaxUint8
	e.PeggedPrice = math.MinInt64
	e.TransactTime = math.MinInt64
	e.OrderCreationTime = math.MinInt64
	e.MaxFloor = math.MinInt64
	e.ListID = math.MinInt64
	e.CashOrderQty = math.MinInt64
	e.TargetStrategy = math.MinInt32
	e.StrategyID = math.MinInt64
	e.LeavesQty = math.MinInt64
	e.CumQuoteQty = math.MinInt64
	e.TradeID = math.MinInt64
	e.LastPx = math.MinInt64
	e.SecondaryOrderID = math.MinInt64
	e.SecondaryExternalAccountID = math.MinInt64
	e.AllocID = math.MinInt64
	e.WorkingTime = math.MinInt64
	e.TrailingTime = math.MinInt64
	e.PreventedMatchID = math.MinInt64
	e.PreventedExecutionPrice = math.MinInt64
	e.PreventedExecutionQty = math.MinInt64
	e.TradeGroupID = math.MinInt64
	e.CounterOrderID = math.MinInt64
	e.PreventedQty = math.MinInt64
	e.LastPreventedQty = math.MinInt64
	e.ErrorCode = math.MinInt32
	return
}

func (e *ExecutionReportMiscFees) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, e.Exponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.MiscFeeAmt); err != nil {
		return err
	}
	if err := e.MiscFeeType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.MiscFeeCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.MiscFeeCurr); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportMiscFees) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !e.ExponentInActingVersion(actingVersion) {
		e.Exponent = e.ExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &e.Exponent); err != nil {
			return err
		}
	}
	if !e.MiscFeeAmtInActingVersion(actingVersion) {
		e.MiscFeeAmt = e.MiscFeeAmtNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.MiscFeeAmt); err != nil {
			return err
		}
	}
	if e.MiscFeeTypeInActingVersion(actingVersion) {
		if err := e.MiscFeeType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
	}

	if e.MiscFeeCurrInActingVersion(actingVersion) {
		var MiscFeeCurrLength uint8
		if err := _m.ReadUint8(_r, &MiscFeeCurrLength); err != nil {
			return err
		}
		if cap(e.MiscFeeCurr) < int(MiscFeeCurrLength) {
			e.MiscFeeCurr = make([]uint8, MiscFeeCurrLength)
		}
		e.MiscFeeCurr = e.MiscFeeCurr[:MiscFeeCurrLength]
		if err := _m.ReadBytes(_r, e.MiscFeeCurr); err != nil {
			return err
		}
	}
	return nil
}

func (e *ExecutionReportMiscFees) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.ExponentInActingVersion(actingVersion) {
		if e.Exponent < e.ExponentMinValue() || e.Exponent > e.ExponentMaxValue() {
			return fmt.Errorf("Range check failed on e.Exponent (%v < %v > %v)", e.ExponentMinValue(), e.Exponent, e.ExponentMaxValue())
		}
	}
	if e.MiscFeeAmtInActingVersion(actingVersion) {
		if e.MiscFeeAmt < e.MiscFeeAmtMinValue() || e.MiscFeeAmt > e.MiscFeeAmtMaxValue() {
			return fmt.Errorf("Range check failed on e.MiscFeeAmt (%v < %v > %v)", e.MiscFeeAmtMinValue(), e.MiscFeeAmt, e.MiscFeeAmtMaxValue())
		}
	}
	if err := e.MiscFeeType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(e.MiscFeeCurr[:]) {
		return errors.New("e.MiscFeeCurr failed UTF-8 validation")
	}
	return nil
}

func ExecutionReportMiscFeesInit(e *ExecutionReportMiscFees) {
	return
}

func (*ExecutionReport) SbeBlockLength() (blockLength uint16) {
	return 281
}

func (*ExecutionReport) SbeTemplateId() (templateId uint16) {
	return 98
}

func (*ExecutionReport) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ExecutionReport) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ExecutionReport) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReport) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*ExecutionReport) PriceExponentId() uint16 {
	return 25054
}

func (*ExecutionReport) PriceExponentSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceExponentSinceVersion()
}

func (*ExecutionReport) PriceExponentDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PriceExponentMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*ExecutionReport) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*ExecutionReport) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*ExecutionReport) QtyExponentId() uint16 {
	return 25055
}

func (*ExecutionReport) QtyExponentSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.QtyExponentSinceVersion()
}

func (*ExecutionReport) QtyExponentDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) QtyExponentMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*ExecutionReport) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*ExecutionReport) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*ExecutionReport) ExecIDId() uint16 {
	return 17
}

func (*ExecutionReport) ExecIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ExecIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecIDSinceVersion()
}

func (*ExecutionReport) ExecIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ExecIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) ExecIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) ExecIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) ExecIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReport) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReport) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrderQtyId() uint16 {
	return 38
}

func (*ExecutionReport) OrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderQtySinceVersion()
}

func (*ExecutionReport) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrdTypeId() uint16 {
	return 40
}

func (*ExecutionReport) OrdTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdTypeSinceVersion()
}

func (*ExecutionReport) OrdTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrdTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) SideId() uint16 {
	return 54
}

func (*ExecutionReport) SideSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SideSinceVersion()
}

func (*ExecutionReport) SideDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) ExecInstId() uint16 {
	return 18
}

func (*ExecutionReport) ExecInstSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecInstSinceVersion()
}

func (*ExecutionReport) ExecInstDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ExecInstMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PriceId() uint16 {
	return 44
}

func (*ExecutionReport) PriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PriceSinceVersion()
}

func (*ExecutionReport) PriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PriceNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TriggerTypeId() uint16 {
	return 1100
}

func (*ExecutionReport) TriggerTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerTypeSinceVersion()
}

func (*ExecutionReport) TriggerTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerActionId() uint16 {
	return 1101
}

func (*ExecutionReport) TriggerActionSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerActionSinceVersion()
}

func (*ExecutionReport) TriggerActionDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerActionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerPriceId() uint16 {
	return 1102
}

func (*ExecutionReport) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerPriceSinceVersion()
}

func (*ExecutionReport) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TriggerPriceTypeId() uint16 {
	return 1107
}

func (*ExecutionReport) TriggerPriceTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerPriceTypeSinceVersion()
}

func (*ExecutionReport) TriggerPriceTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerPriceTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerPriceDirectionId() uint16 {
	return 1109
}

func (*ExecutionReport) TriggerPriceDirectionSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerPriceDirectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerPriceDirectionSinceVersion()
}

func (*ExecutionReport) TriggerPriceDirectionDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerPriceDirectionMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerTrailingDeltaBipsId() uint16 {
	return 25009
}

func (*ExecutionReport) TriggerTrailingDeltaBipsSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TriggerTrailingDeltaBipsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TriggerTrailingDeltaBipsSinceVersion()
}

func (*ExecutionReport) TriggerTrailingDeltaBipsDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TriggerTrailingDeltaBipsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TriggerTrailingDeltaBipsMinValue() uint64 {
	return 0
}

func (*ExecutionReport) TriggerTrailingDeltaBipsMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ExecutionReport) TriggerTrailingDeltaBipsNullValue() uint64 {
	return math.MaxUint64
}

func (*ExecutionReport) PegOffsetValueId() uint16 {
	return 211
}

func (*ExecutionReport) PegOffsetValueSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PegOffsetValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PegOffsetValueSinceVersion()
}

func (*ExecutionReport) PegOffsetValueDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PegOffsetValueMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PegOffsetValueMinValue() uint8 {
	return 0
}

func (*ExecutionReport) PegOffsetValueMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ExecutionReport) PegOffsetValueNullValue() uint8 {
	return math.MaxUint8
}

func (*ExecutionReport) PegPriceTypeId() uint16 {
	return 1094
}

func (*ExecutionReport) PegPriceTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PegPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PegPriceTypeSinceVersion()
}

func (*ExecutionReport) PegPriceTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PegPriceTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PegMoveTypeId() uint16 {
	return 835
}

func (*ExecutionReport) PegMoveTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PegMoveTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PegMoveTypeSinceVersion()
}

func (*ExecutionReport) PegMoveTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PegMoveTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PegOffsetTypeId() uint16 {
	return 836
}

func (*ExecutionReport) PegOffsetTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PegOffsetTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PegOffsetTypeSinceVersion()
}

func (*ExecutionReport) PegOffsetTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PegOffsetTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PeggedPriceId() uint16 {
	return 839
}

func (*ExecutionReport) PeggedPriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PeggedPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PeggedPriceSinceVersion()
}

func (*ExecutionReport) PeggedPriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PeggedPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PeggedPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PeggedPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PeggedPriceNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TimeInForceId() uint16 {
	return 59
}

func (*ExecutionReport) TimeInForceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TimeInForceSinceVersion()
}

func (*ExecutionReport) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TimeInForceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReport) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReport) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TransactTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) TransactTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) TransactTimeNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrderCreationTimeId() uint16 {
	return 25018
}

func (*ExecutionReport) OrderCreationTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrderCreationTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderCreationTimeSinceVersion()
}

func (*ExecutionReport) OrderCreationTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrderCreationTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) OrderCreationTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) OrderCreationTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) OrderCreationTimeNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) MaxFloorId() uint16 {
	return 111
}

func (*ExecutionReport) MaxFloorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) MaxFloorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MaxFloorSinceVersion()
}

func (*ExecutionReport) MaxFloorDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) MaxFloorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) MaxFloorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) MaxFloorMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) MaxFloorNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) ListIDId() uint16 {
	return 66
}

func (*ExecutionReport) ListIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ListIDSinceVersion()
}

func (*ExecutionReport) ListIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ListIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) ListIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) ListIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) ListIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) CashOrderQtyId() uint16 {
	return 152
}

func (*ExecutionReport) CashOrderQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) CashOrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CashOrderQtySinceVersion()
}

func (*ExecutionReport) CashOrderQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) CashOrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) CashOrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) CashOrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) CashOrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TargetStrategyId() uint16 {
	return 847
}

func (*ExecutionReport) TargetStrategySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TargetStrategyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TargetStrategySinceVersion()
}

func (*ExecutionReport) TargetStrategyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TargetStrategyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TargetStrategyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReport) TargetStrategyMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReport) TargetStrategyNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReport) StrategyIDId() uint16 {
	return 7940
}

func (*ExecutionReport) StrategyIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) StrategyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.StrategyIDSinceVersion()
}

func (*ExecutionReport) StrategyIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) StrategyIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) StrategyIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) StrategyIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) StrategyIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrderCapacityId() uint16 {
	return 528
}

func (*ExecutionReport) OrderCapacitySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrderCapacityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderCapacitySinceVersion()
}

func (*ExecutionReport) OrderCapacityDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrderCapacityMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) SelfTradePreventionModeId() uint16 {
	return 25001
}

func (*ExecutionReport) SelfTradePreventionModeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SelfTradePreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SelfTradePreventionModeSinceVersion()
}

func (*ExecutionReport) SelfTradePreventionModeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) SelfTradePreventionModeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReport) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReport) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ExecTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) CumQtyId() uint16 {
	return 14
}

func (*ExecutionReport) CumQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) CumQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQtySinceVersion()
}

func (*ExecutionReport) CumQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) CumQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) CumQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) CumQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) CumQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) LeavesQtyId() uint16 {
	return 151
}

func (*ExecutionReport) LeavesQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) LeavesQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LeavesQtySinceVersion()
}

func (*ExecutionReport) LeavesQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) LeavesQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) LeavesQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) LeavesQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) LeavesQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) CumQuoteQtyId() uint16 {
	return 25017
}

func (*ExecutionReport) CumQuoteQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) CumQuoteQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CumQuoteQtySinceVersion()
}

func (*ExecutionReport) CumQuoteQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) CumQuoteQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) CumQuoteQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) CumQuoteQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) CumQuoteQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) AggressorIndicatorId() uint16 {
	return 1057
}

func (*ExecutionReport) AggressorIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) AggressorIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AggressorIndicatorSinceVersion()
}

func (*ExecutionReport) AggressorIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) AggressorIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TradeIDId() uint16 {
	return 1003
}

func (*ExecutionReport) TradeIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeIDSinceVersion()
}

func (*ExecutionReport) TradeIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TradeIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TradeIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) TradeIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) TradeIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) LastPxId() uint16 {
	return 31
}

func (*ExecutionReport) LastPxSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) LastPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPxSinceVersion()
}

func (*ExecutionReport) LastPxDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) LastPxMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) LastPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) LastPxMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) LastPxNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) LastQtyId() uint16 {
	return 32
}

func (*ExecutionReport) LastQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) LastQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastQtySinceVersion()
}

func (*ExecutionReport) LastQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) LastQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) LastQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) LastQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) LastQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReport) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReport) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrdStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) SecondaryOrderIDId() uint16 {
	return 198
}

func (*ExecutionReport) SecondaryOrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SecondaryOrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecondaryOrderIDSinceVersion()
}

func (*ExecutionReport) SecondaryOrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) SecondaryOrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) SecondaryOrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) SecondaryOrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) SecondaryOrderIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) SecondaryExternalAccountIDId() uint16 {
	return 25020
}

func (*ExecutionReport) SecondaryExternalAccountIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SecondaryExternalAccountIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecondaryExternalAccountIDSinceVersion()
}

func (*ExecutionReport) SecondaryExternalAccountIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) SecondaryExternalAccountIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) SecondaryExternalAccountIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) SecondaryExternalAccountIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) SecondaryExternalAccountIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) AllocIDId() uint16 {
	return 70
}

func (*ExecutionReport) AllocIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) AllocIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.AllocIDSinceVersion()
}

func (*ExecutionReport) AllocIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) AllocIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) AllocIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) AllocIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) AllocIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) MatchTypeId() uint16 {
	return 574
}

func (*ExecutionReport) MatchTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) MatchTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MatchTypeSinceVersion()
}

func (*ExecutionReport) MatchTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) MatchTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) WorkingFloorId() uint16 {
	return 25021
}

func (*ExecutionReport) WorkingFloorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) WorkingFloorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.WorkingFloorSinceVersion()
}

func (*ExecutionReport) WorkingFloorDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) WorkingFloorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) WorkingIndicatorId() uint16 {
	return 636
}

func (*ExecutionReport) WorkingIndicatorSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) WorkingIndicatorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.WorkingIndicatorSinceVersion()
}

func (*ExecutionReport) WorkingIndicatorDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) WorkingIndicatorMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) WorkingTimeId() uint16 {
	return 25023
}

func (*ExecutionReport) WorkingTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) WorkingTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.WorkingTimeSinceVersion()
}

func (*ExecutionReport) WorkingTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) WorkingTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) WorkingTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) WorkingTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) WorkingTimeNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TrailingTimeId() uint16 {
	return 25022
}

func (*ExecutionReport) TrailingTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TrailingTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TrailingTimeSinceVersion()
}

func (*ExecutionReport) TrailingTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TrailingTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TrailingTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) TrailingTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) TrailingTimeNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) PreventedMatchIDId() uint16 {
	return 25024
}

func (*ExecutionReport) PreventedMatchIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PreventedMatchIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreventedMatchIDSinceVersion()
}

func (*ExecutionReport) PreventedMatchIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PreventedMatchIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PreventedMatchIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PreventedMatchIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PreventedMatchIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) PreventedExecutionPriceId() uint16 {
	return 25025
}

func (*ExecutionReport) PreventedExecutionPriceSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PreventedExecutionPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreventedExecutionPriceSinceVersion()
}

func (*ExecutionReport) PreventedExecutionPriceDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PreventedExecutionPriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PreventedExecutionPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PreventedExecutionPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PreventedExecutionPriceNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) PreventedExecutionQtyId() uint16 {
	return 25026
}

func (*ExecutionReport) PreventedExecutionQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PreventedExecutionQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreventedExecutionQtySinceVersion()
}

func (*ExecutionReport) PreventedExecutionQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PreventedExecutionQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PreventedExecutionQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PreventedExecutionQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PreventedExecutionQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) TradeGroupIDId() uint16 {
	return 25027
}

func (*ExecutionReport) TradeGroupIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) TradeGroupIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeGroupIDSinceVersion()
}

func (*ExecutionReport) TradeGroupIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) TradeGroupIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) TradeGroupIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) TradeGroupIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) TradeGroupIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) CounterOrderIDId() uint16 {
	return 25029
}

func (*ExecutionReport) CounterOrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) CounterOrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CounterOrderIDSinceVersion()
}

func (*ExecutionReport) CounterOrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) CounterOrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) CounterOrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) CounterOrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) CounterOrderIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) PreventedQtyId() uint16 {
	return 25030
}

func (*ExecutionReport) PreventedQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) PreventedQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.PreventedQtySinceVersion()
}

func (*ExecutionReport) PreventedQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) PreventedQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) PreventedQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) PreventedQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) PreventedQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) LastPreventedQtyId() uint16 {
	return 25031
}

func (*ExecutionReport) LastPreventedQtySinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) LastPreventedQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.LastPreventedQtySinceVersion()
}

func (*ExecutionReport) LastPreventedQtyDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) LastPreventedQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) LastPreventedQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReport) LastPreventedQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReport) LastPreventedQtyNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReport) SORId() uint16 {
	return 25032
}

func (*ExecutionReport) SORSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SORSinceVersion()
}

func (*ExecutionReport) SORDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) SORMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) OrdRejReasonId() uint16 {
	return 103
}

func (*ExecutionReport) OrdRejReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrdRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdRejReasonSinceVersion()
}

func (*ExecutionReport) OrdRejReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) OrdRejReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) ErrorCodeId() uint16 {
	return 25016
}

func (*ExecutionReport) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorCodeSinceVersion()
}

func (*ExecutionReport) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ErrorCodeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReport) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReport) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReport) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReport) ExpiryReasonId() uint16 {
	return 25056
}

func (*ExecutionReport) ExpiryReasonSinceVersion() uint16 {
	return 1
}

func (e *ExecutionReport) ExpiryReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpiryReasonSinceVersion()
}

func (*ExecutionReport) ExpiryReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReport) ExpiryReasonMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*ExecutionReportMiscFees) ExponentId() uint16 {
	return 25053
}

func (*ExecutionReportMiscFees) ExponentSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportMiscFees) ExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExponentSinceVersion()
}

func (*ExecutionReportMiscFees) ExponentDeprecated() uint16 {
	return 0
}

func (*ExecutionReportMiscFees) ExponentMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportMiscFees) ExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*ExecutionReportMiscFees) ExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*ExecutionReportMiscFees) ExponentNullValue() int8 {
	return math.MinInt8
}

func (*ExecutionReportMiscFees) MiscFeeAmtId() uint16 {
	return 137
}

func (*ExecutionReportMiscFees) MiscFeeAmtSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportMiscFees) MiscFeeAmtInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MiscFeeAmtSinceVersion()
}

func (*ExecutionReportMiscFees) MiscFeeAmtDeprecated() uint16 {
	return 0
}

func (*ExecutionReportMiscFees) MiscFeeAmtMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportMiscFees) MiscFeeAmtMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReportMiscFees) MiscFeeAmtMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReportMiscFees) MiscFeeAmtNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReportMiscFees) MiscFeeTypeId() uint16 {
	return 139
}

func (*ExecutionReportMiscFees) MiscFeeTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportMiscFees) MiscFeeTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MiscFeeTypeSinceVersion()
}

func (*ExecutionReportMiscFees) MiscFeeTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportMiscFees) MiscFeeTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportMiscFees) MiscFeeCurrMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReportMiscFees) MiscFeeCurrSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportMiscFees) MiscFeeCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MiscFeeCurrSinceVersion()
}

func (*ExecutionReportMiscFees) MiscFeeCurrDeprecated() uint16 {
	return 0
}

func (ExecutionReportMiscFees) MiscFeeCurrCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReportMiscFees) MiscFeeCurrHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) MiscFeesId() uint16 {
	return 136
}

func (*ExecutionReport) MiscFeesSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) MiscFeesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MiscFeesSinceVersion()
}

func (*ExecutionReport) MiscFeesDeprecated() uint16 {
	return 0
}

func (*ExecutionReportMiscFees) SbeBlockLength() (blockLength uint) {
	return 10
}

func (*ExecutionReportMiscFees) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ExecutionReport) ClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReport) ClOrdIDDeprecated() uint16 {
	return 0
}

func (ExecutionReport) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) OrigClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrigClOrdIDSinceVersion()
}

func (*ExecutionReport) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (ExecutionReport) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) SymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) SymbolSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*ExecutionReport) SymbolDeprecated() uint16 {
	return 0
}

func (ExecutionReport) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) SymbolHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) SecondarySymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) SecondarySymbolSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) SecondarySymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SecondarySymbolSinceVersion()
}

func (*ExecutionReport) SecondarySymbolDeprecated() uint16 {
	return 0
}

func (ExecutionReport) SecondarySymbolCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) SecondarySymbolHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) CounterSymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) CounterSymbolSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) CounterSymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CounterSymbolSinceVersion()
}

func (*ExecutionReport) CounterSymbolDeprecated() uint16 {
	return 0
}

func (ExecutionReport) CounterSymbolCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) CounterSymbolHeaderLength() uint64 {
	return 1
}

func (*ExecutionReport) ErrorTextMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*ExecutionReport) ErrorTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReport) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorTextSinceVersion()
}

func (*ExecutionReport) ErrorTextDeprecated() uint16 {
	return 0
}

func (ExecutionReport) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReport) ErrorTextHeaderLength() uint64 {
	return 2
}
