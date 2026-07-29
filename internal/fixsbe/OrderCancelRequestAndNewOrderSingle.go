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

type OrderCancelRequestAndNewOrderSingle struct {
	OrderCancelRequestAndNewOrderSingleMode OrderCancelRequestAndNewOrderSingleModeEnum
	OrderRateLimitExceededMode              OrderRateLimitExceededModeEnum
	OrderID                                 int64
	CancelRestrictions                      CancelRestrictionsEnum
	PriceExponent                           int8
	QtyExponent                             int8
	OrderQty                                int64
	OrdType                                 OrdTypeEnum
	ExecInst                                ExecInstEnum
	Price                                   int64
	TriggerType                             TriggerTypeEnum
	TriggerAction                           TriggerActionEnum
	TriggerPrice                            int64
	TriggerPriceType                        TriggerPriceTypeEnum
	TriggerPriceDirection                   TriggerPriceDirectionEnum
	TriggerTrailingDeltaBips                uint64
	PegOffsetValue                          uint8
	PegPriceType                            PegPriceTypeEnum
	PegMoveType                             PegMoveTypeEnum
	PegOffsetType                           PegOffsetTypeEnum
	Side                                    SideEnum
	TimeInForce                             TimeInForceEnum
	MaxFloor                                int64
	CashOrderQty                            int64
	TargetStrategy                          int32
	StrategyID                              int64
	SelfTradePreventionMode                 SelfTradePreventionModeEnum
	CancelClOrdID                           []uint8
	OrigClOrdID                             []uint8
	ClOrdID                                 []uint8
	Symbol                                  []uint8
}

func (o *OrderCancelRequestAndNewOrderSingle) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.OrderCancelRequestAndNewOrderSingleMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrderRateLimitExceededMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OrderID); err != nil {
		return err
	}
	if err := o.CancelRestrictions.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, o.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, o.QtyExponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OrderQty); err != nil {
		return err
	}
	if err := o.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.Price); err != nil {
		return err
	}
	if err := o.TriggerType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TriggerAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.TriggerPrice); err != nil {
		return err
	}
	if err := o.TriggerPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TriggerPriceDirection.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.TriggerTrailingDeltaBips); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.PegOffsetValue); err != nil {
		return err
	}
	if err := o.PegPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.PegMoveType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.PegOffsetType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.MaxFloor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.CashOrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.TargetStrategy); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.StrategyID); err != nil {
		return err
	}
	if err := o.SelfTradePreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.CancelClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CancelClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.OrigClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelRequestAndNewOrderSingle) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.OrderCancelRequestAndNewOrderSingleModeInActingVersion(actingVersion) {
		if err := o.OrderCancelRequestAndNewOrderSingleMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrderRateLimitExceededModeInActingVersion(actingVersion) {
		if err := o.OrderRateLimitExceededMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderID); err != nil {
			return err
		}
	}
	if o.CancelRestrictionsInActingVersion(actingVersion) {
		if err := o.CancelRestrictions.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.PriceExponentInActingVersion(actingVersion) {
		o.PriceExponent = o.PriceExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &o.PriceExponent); err != nil {
			return err
		}
	}
	if !o.QtyExponentInActingVersion(actingVersion) {
		o.QtyExponent = o.QtyExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &o.QtyExponent); err != nil {
			return err
		}
	}
	if !o.OrderQtyInActingVersion(actingVersion) {
		o.OrderQty = o.OrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderQty); err != nil {
			return err
		}
	}
	if o.OrdTypeInActingVersion(actingVersion) {
		if err := o.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.ExecInstInActingVersion(actingVersion) {
		if err := o.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.PriceInActingVersion(actingVersion) {
		o.Price = o.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Price); err != nil {
			return err
		}
	}
	if o.TriggerTypeInActingVersion(actingVersion) {
		if err := o.TriggerType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TriggerActionInActingVersion(actingVersion) {
		if err := o.TriggerAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.TriggerPriceInActingVersion(actingVersion) {
		o.TriggerPrice = o.TriggerPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.TriggerPrice); err != nil {
			return err
		}
	}
	if o.TriggerPriceTypeInActingVersion(actingVersion) {
		if err := o.TriggerPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TriggerPriceDirectionInActingVersion(actingVersion) {
		if err := o.TriggerPriceDirection.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		o.TriggerTrailingDeltaBips = o.TriggerTrailingDeltaBipsNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.TriggerTrailingDeltaBips); err != nil {
			return err
		}
	}
	if !o.PegOffsetValueInActingVersion(actingVersion) {
		o.PegOffsetValue = o.PegOffsetValueNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.PegOffsetValue); err != nil {
			return err
		}
	}
	if o.PegPriceTypeInActingVersion(actingVersion) {
		if err := o.PegPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.PegMoveTypeInActingVersion(actingVersion) {
		if err := o.PegMoveType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.PegOffsetTypeInActingVersion(actingVersion) {
		if err := o.PegOffsetType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.MaxFloorInActingVersion(actingVersion) {
		o.MaxFloor = o.MaxFloorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.MaxFloor); err != nil {
			return err
		}
	}
	if !o.CashOrderQtyInActingVersion(actingVersion) {
		o.CashOrderQty = o.CashOrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.CashOrderQty); err != nil {
			return err
		}
	}
	if !o.TargetStrategyInActingVersion(actingVersion) {
		o.TargetStrategy = o.TargetStrategyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.TargetStrategy); err != nil {
			return err
		}
	}
	if !o.StrategyIDInActingVersion(actingVersion) {
		o.StrategyID = o.StrategyIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.StrategyID); err != nil {
			return err
		}
	}
	if o.SelfTradePreventionModeInActingVersion(actingVersion) {
		if err := o.SelfTradePreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.CancelClOrdIDInActingVersion(actingVersion) {
		var CancelClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &CancelClOrdIDLength); err != nil {
			return err
		}
		if cap(o.CancelClOrdID) < int(CancelClOrdIDLength) {
			o.CancelClOrdID = make([]uint8, CancelClOrdIDLength)
		}
		o.CancelClOrdID = o.CancelClOrdID[:CancelClOrdIDLength]
		if err := _m.ReadBytes(_r, o.CancelClOrdID); err != nil {
			return err
		}
	}

	if o.OrigClOrdIDInActingVersion(actingVersion) {
		var OrigClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &OrigClOrdIDLength); err != nil {
			return err
		}
		if cap(o.OrigClOrdID) < int(OrigClOrdIDLength) {
			o.OrigClOrdID = make([]uint8, OrigClOrdIDLength)
		}
		o.OrigClOrdID = o.OrigClOrdID[:OrigClOrdIDLength]
		if err := _m.ReadBytes(_r, o.OrigClOrdID); err != nil {
			return err
		}
	}

	if o.ClOrdIDInActingVersion(actingVersion) {
		var ClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &ClOrdIDLength); err != nil {
			return err
		}
		if cap(o.ClOrdID) < int(ClOrdIDLength) {
			o.ClOrdID = make([]uint8, ClOrdIDLength)
		}
		o.ClOrdID = o.ClOrdID[:ClOrdIDLength]
		if err := _m.ReadBytes(_r, o.ClOrdID); err != nil {
			return err
		}
	}

	if o.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(o.Symbol) < int(SymbolLength) {
			o.Symbol = make([]uint8, SymbolLength)
		}
		o.Symbol = o.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, o.Symbol); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderCancelRequestAndNewOrderSingle) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := o.OrderCancelRequestAndNewOrderSingleMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OrderRateLimitExceededMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID != o.OrderIDNullValue() && (o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if err := o.CancelRestrictions.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.PriceExponentInActingVersion(actingVersion) {
		if o.PriceExponent < o.PriceExponentMinValue() || o.PriceExponent > o.PriceExponentMaxValue() {
			return fmt.Errorf("Range check failed on o.PriceExponent (%v < %v > %v)", o.PriceExponentMinValue(), o.PriceExponent, o.PriceExponentMaxValue())
		}
	}
	if o.QtyExponentInActingVersion(actingVersion) {
		if o.QtyExponent < o.QtyExponentMinValue() || o.QtyExponent > o.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on o.QtyExponent (%v < %v > %v)", o.QtyExponentMinValue(), o.QtyExponent, o.QtyExponentMaxValue())
		}
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty != o.OrderQtyNullValue() && (o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
	}
	if err := o.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.ExecInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.PriceInActingVersion(actingVersion) {
		if o.Price != o.PriceNullValue() && (o.Price < o.PriceMinValue() || o.Price > o.PriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.Price (%v < %v > %v)", o.PriceMinValue(), o.Price, o.PriceMaxValue())
		}
	}
	if err := o.TriggerType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TriggerAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TriggerPriceInActingVersion(actingVersion) {
		if o.TriggerPrice != o.TriggerPriceNullValue() && (o.TriggerPrice < o.TriggerPriceMinValue() || o.TriggerPrice > o.TriggerPriceMaxValue()) {
			return fmt.Errorf("Range check failed on o.TriggerPrice (%v < %v > %v)", o.TriggerPriceMinValue(), o.TriggerPrice, o.TriggerPriceMaxValue())
		}
	}
	if err := o.TriggerPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TriggerPriceDirection.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		if o.TriggerTrailingDeltaBips != o.TriggerTrailingDeltaBipsNullValue() && (o.TriggerTrailingDeltaBips < o.TriggerTrailingDeltaBipsMinValue() || o.TriggerTrailingDeltaBips > o.TriggerTrailingDeltaBipsMaxValue()) {
			return fmt.Errorf("Range check failed on o.TriggerTrailingDeltaBips (%v < %v > %v)", o.TriggerTrailingDeltaBipsMinValue(), o.TriggerTrailingDeltaBips, o.TriggerTrailingDeltaBipsMaxValue())
		}
	}
	if o.PegOffsetValueInActingVersion(actingVersion) {
		if o.PegOffsetValue != o.PegOffsetValueNullValue() && (o.PegOffsetValue < o.PegOffsetValueMinValue() || o.PegOffsetValue > o.PegOffsetValueMaxValue()) {
			return fmt.Errorf("Range check failed on o.PegOffsetValue (%v < %v > %v)", o.PegOffsetValueMinValue(), o.PegOffsetValue, o.PegOffsetValueMaxValue())
		}
	}
	if err := o.PegPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.PegMoveType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.PegOffsetType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.MaxFloorInActingVersion(actingVersion) {
		if o.MaxFloor != o.MaxFloorNullValue() && (o.MaxFloor < o.MaxFloorMinValue() || o.MaxFloor > o.MaxFloorMaxValue()) {
			return fmt.Errorf("Range check failed on o.MaxFloor (%v < %v > %v)", o.MaxFloorMinValue(), o.MaxFloor, o.MaxFloorMaxValue())
		}
	}
	if o.CashOrderQtyInActingVersion(actingVersion) {
		if o.CashOrderQty != o.CashOrderQtyNullValue() && (o.CashOrderQty < o.CashOrderQtyMinValue() || o.CashOrderQty > o.CashOrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on o.CashOrderQty (%v < %v > %v)", o.CashOrderQtyMinValue(), o.CashOrderQty, o.CashOrderQtyMaxValue())
		}
	}
	if o.TargetStrategyInActingVersion(actingVersion) {
		if o.TargetStrategy != o.TargetStrategyNullValue() && (o.TargetStrategy < o.TargetStrategyMinValue() || o.TargetStrategy > o.TargetStrategyMaxValue()) {
			return fmt.Errorf("Range check failed on o.TargetStrategy (%v < %v > %v)", o.TargetStrategyMinValue(), o.TargetStrategy, o.TargetStrategyMaxValue())
		}
	}
	if o.StrategyIDInActingVersion(actingVersion) {
		if o.StrategyID != o.StrategyIDNullValue() && (o.StrategyID < o.StrategyIDMinValue() || o.StrategyID > o.StrategyIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.StrategyID (%v < %v > %v)", o.StrategyIDMinValue(), o.StrategyID, o.StrategyIDMaxValue())
		}
	}
	if err := o.SelfTradePreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(o.CancelClOrdID[:]) {
		return errors.New("o.CancelClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigClOrdID[:]) {
		return errors.New("o.OrigClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.ClOrdID[:]) {
		return errors.New("o.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	return nil
}

func OrderCancelRequestAndNewOrderSingleInit(o *OrderCancelRequestAndNewOrderSingle) {
	o.OrderID = math.MinInt64
	o.OrderQty = math.MinInt64
	o.Price = math.MinInt64
	o.TriggerPrice = math.MinInt64
	o.TriggerTrailingDeltaBips = math.MaxUint64
	o.PegOffsetValue = math.MaxUint8
	o.MaxFloor = math.MinInt64
	o.CashOrderQty = math.MinInt64
	o.TargetStrategy = math.MinInt32
	o.StrategyID = math.MinInt64
	return
}

func (*OrderCancelRequestAndNewOrderSingle) SbeBlockLength() (blockLength uint16) {
	return 86
}

func (*OrderCancelRequestAndNewOrderSingle) SbeTemplateId() (templateId uint16) {
	return 97
}

func (*OrderCancelRequestAndNewOrderSingle) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingle) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingle) SbeSemanticType() (semanticType []byte) {
	return []byte("XCN")
}

func (*OrderCancelRequestAndNewOrderSingle) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderCancelRequestAndNewOrderSingle) OrderCancelRequestAndNewOrderSingleModeId() uint16 {
	return 25033
}

func (*OrderCancelRequestAndNewOrderSingle) OrderCancelRequestAndNewOrderSingleModeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrderCancelRequestAndNewOrderSingleModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderCancelRequestAndNewOrderSingleModeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrderCancelRequestAndNewOrderSingleModeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) OrderCancelRequestAndNewOrderSingleModeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) OrderRateLimitExceededModeId() uint16 {
	return 25038
}

func (*OrderCancelRequestAndNewOrderSingle) OrderRateLimitExceededModeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrderRateLimitExceededModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderRateLimitExceededModeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrderRateLimitExceededModeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) OrderRateLimitExceededModeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelRequestAndNewOrderSingle) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) CancelRestrictionsId() uint16 {
	return 25002
}

func (*OrderCancelRequestAndNewOrderSingle) CancelRestrictionsSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) CancelRestrictionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelRestrictionsSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) CancelRestrictionsDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) CancelRestrictionsMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentId() uint16 {
	return 25054
}

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceExponentSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*OrderCancelRequestAndNewOrderSingle) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentId() uint16 {
	return 25055
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QtyExponentSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*OrderCancelRequestAndNewOrderSingle) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyId() uint16 {
	return 38
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) OrdTypeId() uint16 {
	return 40
}

func (*OrderCancelRequestAndNewOrderSingle) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) OrdTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) ExecInstId() uint16 {
	return 18
}

func (*OrderCancelRequestAndNewOrderSingle) ExecInstSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExecInstSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) ExecInstDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) ExecInstMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PriceId() uint16 {
	return 44
}

func (*OrderCancelRequestAndNewOrderSingle) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PriceDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PriceMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) PriceNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTypeId() uint16 {
	return 1100
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerActionId() uint16 {
	return 1101
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerActionSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerActionSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerActionDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerActionMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceId() uint16 {
	return 1102
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerPriceSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceTypeId() uint16 {
	return 1107
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerPriceTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceDirectionId() uint16 {
	return 1109
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceDirectionSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerPriceDirectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerPriceDirectionSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceDirectionDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerPriceDirectionMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsId() uint16 {
	return 25009
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TriggerTrailingDeltaBipsSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsMinValue() uint64 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderCancelRequestAndNewOrderSingle) TriggerTrailingDeltaBipsNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueId() uint16 {
	return 211
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PegOffsetValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PegOffsetValueSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueMinValue() uint8 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetValueNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderCancelRequestAndNewOrderSingle) PegPriceTypeId() uint16 {
	return 1094
}

func (*OrderCancelRequestAndNewOrderSingle) PegPriceTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PegPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PegPriceTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PegPriceTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PegPriceTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PegMoveTypeId() uint16 {
	return 835
}

func (*OrderCancelRequestAndNewOrderSingle) PegMoveTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PegMoveTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PegMoveTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PegMoveTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PegMoveTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetTypeId() uint16 {
	return 836
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) PegOffsetTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PegOffsetTypeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) PegOffsetTypeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) SideId() uint16 {
	return 54
}

func (*OrderCancelRequestAndNewOrderSingle) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) SideDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) SideMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TimeInForceId() uint16 {
	return 59
}

func (*OrderCancelRequestAndNewOrderSingle) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorId() uint16 {
	return 111
}

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) MaxFloorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MaxFloorSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) MaxFloorNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyId() uint16 {
	return 152
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) CashOrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CashOrderQtySinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) CashOrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyId() uint16 {
	return 847
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategySinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) TargetStrategyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TargetStrategySinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancelRequestAndNewOrderSingle) TargetStrategyNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDId() uint16 {
	return 7940
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) StrategyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StrategyIDSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequestAndNewOrderSingle) StrategyIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequestAndNewOrderSingle) SelfTradePreventionModeId() uint16 {
	return 25001
}

func (*OrderCancelRequestAndNewOrderSingle) SelfTradePreventionModeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) SelfTradePreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SelfTradePreventionModeSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) SelfTradePreventionModeDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingle) SelfTradePreventionModeMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) CancelClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) CancelClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) CancelClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelClOrdIDSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) CancelClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequestAndNewOrderSingle) CancelClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequestAndNewOrderSingle) CancelClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingle) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequestAndNewOrderSingle) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequestAndNewOrderSingle) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingle) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequestAndNewOrderSingle) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequestAndNewOrderSingle) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingle) SymbolMetaAttribute(meta int) string {
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

func (*OrderCancelRequestAndNewOrderSingle) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingle) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingle) SymbolDeprecated() uint16 {
	return 0
}

func (OrderCancelRequestAndNewOrderSingle) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequestAndNewOrderSingle) SymbolHeaderLength() uint64 {
	return 1
}
