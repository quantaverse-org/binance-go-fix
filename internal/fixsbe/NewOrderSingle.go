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

type NewOrderSingle struct {
	PriceExponent            int8
	QtyExponent              int8
	OrderQty                 int64
	OrdType                  OrdTypeEnum
	ExecInst                 ExecInstEnum
	Price                    int64
	TriggerType              TriggerTypeEnum
	TriggerAction            TriggerActionEnum
	TriggerPrice             int64
	TriggerPriceType         TriggerPriceTypeEnum
	TriggerPriceDirection    TriggerPriceDirectionEnum
	TriggerTrailingDeltaBips uint64
	PegOffsetValue           uint8
	PegPriceType             PegPriceTypeEnum
	PegMoveType              PegMoveTypeEnum
	PegOffsetType            PegOffsetTypeEnum
	Side                     SideEnum
	TimeInForce              TimeInForceEnum
	MaxFloor                 int64
	CashOrderQty             int64
	TargetStrategy           int32
	StrategyID               int64
	SelfTradePreventionMode  SelfTradePreventionModeEnum
	SOR                      BoolEnumEnum
	ClOrdID                  []uint8
	Symbol                   []uint8
}

func (n *NewOrderSingle) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt8(_w, n.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, n.QtyExponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.OrderQty); err != nil {
		return err
	}
	if err := n.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.ExecInst.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.Price); err != nil {
		return err
	}
	if err := n.TriggerType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.TriggerAction.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.TriggerPrice); err != nil {
		return err
	}
	if err := n.TriggerPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.TriggerPriceDirection.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.TriggerTrailingDeltaBips); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, n.PegOffsetValue); err != nil {
		return err
	}
	if err := n.PegPriceType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.PegMoveType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.PegOffsetType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.MaxFloor); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.CashOrderQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, n.TargetStrategy); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.StrategyID); err != nil {
		return err
	}
	if err := n.SelfTradePreventionMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.SOR.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(n.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(n.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Symbol); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderSingle) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.PriceExponentInActingVersion(actingVersion) {
		n.PriceExponent = n.PriceExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &n.PriceExponent); err != nil {
			return err
		}
	}
	if !n.QtyExponentInActingVersion(actingVersion) {
		n.QtyExponent = n.QtyExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &n.QtyExponent); err != nil {
			return err
		}
	}
	if !n.OrderQtyInActingVersion(actingVersion) {
		n.OrderQty = n.OrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.OrderQty); err != nil {
			return err
		}
	}
	if n.OrdTypeInActingVersion(actingVersion) {
		if err := n.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.ExecInstInActingVersion(actingVersion) {
		if err := n.ExecInst.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.PriceInActingVersion(actingVersion) {
		n.Price = n.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.Price); err != nil {
			return err
		}
	}
	if n.TriggerTypeInActingVersion(actingVersion) {
		if err := n.TriggerType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.TriggerActionInActingVersion(actingVersion) {
		if err := n.TriggerAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.TriggerPriceInActingVersion(actingVersion) {
		n.TriggerPrice = n.TriggerPriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.TriggerPrice); err != nil {
			return err
		}
	}
	if n.TriggerPriceTypeInActingVersion(actingVersion) {
		if err := n.TriggerPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.TriggerPriceDirectionInActingVersion(actingVersion) {
		if err := n.TriggerPriceDirection.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		n.TriggerTrailingDeltaBips = n.TriggerTrailingDeltaBipsNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.TriggerTrailingDeltaBips); err != nil {
			return err
		}
	}
	if !n.PegOffsetValueInActingVersion(actingVersion) {
		n.PegOffsetValue = n.PegOffsetValueNullValue()
	} else {
		if err := _m.ReadUint8(_r, &n.PegOffsetValue); err != nil {
			return err
		}
	}
	if n.PegPriceTypeInActingVersion(actingVersion) {
		if err := n.PegPriceType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.PegMoveTypeInActingVersion(actingVersion) {
		if err := n.PegMoveType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.PegOffsetTypeInActingVersion(actingVersion) {
		if err := n.PegOffsetType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.TimeInForceInActingVersion(actingVersion) {
		if err := n.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.MaxFloorInActingVersion(actingVersion) {
		n.MaxFloor = n.MaxFloorNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.MaxFloor); err != nil {
			return err
		}
	}
	if !n.CashOrderQtyInActingVersion(actingVersion) {
		n.CashOrderQty = n.CashOrderQtyNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.CashOrderQty); err != nil {
			return err
		}
	}
	if !n.TargetStrategyInActingVersion(actingVersion) {
		n.TargetStrategy = n.TargetStrategyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &n.TargetStrategy); err != nil {
			return err
		}
	}
	if !n.StrategyIDInActingVersion(actingVersion) {
		n.StrategyID = n.StrategyIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.StrategyID); err != nil {
			return err
		}
	}
	if n.SelfTradePreventionModeInActingVersion(actingVersion) {
		if err := n.SelfTradePreventionMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.SORInActingVersion(actingVersion) {
		if err := n.SOR.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.ClOrdIDInActingVersion(actingVersion) {
		var ClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &ClOrdIDLength); err != nil {
			return err
		}
		if cap(n.ClOrdID) < int(ClOrdIDLength) {
			n.ClOrdID = make([]uint8, ClOrdIDLength)
		}
		n.ClOrdID = n.ClOrdID[:ClOrdIDLength]
		if err := _m.ReadBytes(_r, n.ClOrdID); err != nil {
			return err
		}
	}

	if n.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(n.Symbol) < int(SymbolLength) {
			n.Symbol = make([]uint8, SymbolLength)
		}
		n.Symbol = n.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, n.Symbol); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderSingle) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.PriceExponentInActingVersion(actingVersion) {
		if n.PriceExponent < n.PriceExponentMinValue() || n.PriceExponent > n.PriceExponentMaxValue() {
			return fmt.Errorf("Range check failed on n.PriceExponent (%v < %v > %v)", n.PriceExponentMinValue(), n.PriceExponent, n.PriceExponentMaxValue())
		}
	}
	if n.QtyExponentInActingVersion(actingVersion) {
		if n.QtyExponent < n.QtyExponentMinValue() || n.QtyExponent > n.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on n.QtyExponent (%v < %v > %v)", n.QtyExponentMinValue(), n.QtyExponent, n.QtyExponentMaxValue())
		}
	}
	if n.OrderQtyInActingVersion(actingVersion) {
		if n.OrderQty != n.OrderQtyNullValue() && (n.OrderQty < n.OrderQtyMinValue() || n.OrderQty > n.OrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on n.OrderQty (%v < %v > %v)", n.OrderQtyMinValue(), n.OrderQty, n.OrderQtyMaxValue())
		}
	}
	if err := n.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.ExecInst.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.PriceInActingVersion(actingVersion) {
		if n.Price != n.PriceNullValue() && (n.Price < n.PriceMinValue() || n.Price > n.PriceMaxValue()) {
			return fmt.Errorf("Range check failed on n.Price (%v < %v > %v)", n.PriceMinValue(), n.Price, n.PriceMaxValue())
		}
	}
	if err := n.TriggerType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.TriggerAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.TriggerPriceInActingVersion(actingVersion) {
		if n.TriggerPrice != n.TriggerPriceNullValue() && (n.TriggerPrice < n.TriggerPriceMinValue() || n.TriggerPrice > n.TriggerPriceMaxValue()) {
			return fmt.Errorf("Range check failed on n.TriggerPrice (%v < %v > %v)", n.TriggerPriceMinValue(), n.TriggerPrice, n.TriggerPriceMaxValue())
		}
	}
	if err := n.TriggerPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.TriggerPriceDirection.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.TriggerTrailingDeltaBipsInActingVersion(actingVersion) {
		if n.TriggerTrailingDeltaBips != n.TriggerTrailingDeltaBipsNullValue() && (n.TriggerTrailingDeltaBips < n.TriggerTrailingDeltaBipsMinValue() || n.TriggerTrailingDeltaBips > n.TriggerTrailingDeltaBipsMaxValue()) {
			return fmt.Errorf("Range check failed on n.TriggerTrailingDeltaBips (%v < %v > %v)", n.TriggerTrailingDeltaBipsMinValue(), n.TriggerTrailingDeltaBips, n.TriggerTrailingDeltaBipsMaxValue())
		}
	}
	if n.PegOffsetValueInActingVersion(actingVersion) {
		if n.PegOffsetValue != n.PegOffsetValueNullValue() && (n.PegOffsetValue < n.PegOffsetValueMinValue() || n.PegOffsetValue > n.PegOffsetValueMaxValue()) {
			return fmt.Errorf("Range check failed on n.PegOffsetValue (%v < %v > %v)", n.PegOffsetValueMinValue(), n.PegOffsetValue, n.PegOffsetValueMaxValue())
		}
	}
	if err := n.PegPriceType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.PegMoveType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.PegOffsetType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.MaxFloorInActingVersion(actingVersion) {
		if n.MaxFloor != n.MaxFloorNullValue() && (n.MaxFloor < n.MaxFloorMinValue() || n.MaxFloor > n.MaxFloorMaxValue()) {
			return fmt.Errorf("Range check failed on n.MaxFloor (%v < %v > %v)", n.MaxFloorMinValue(), n.MaxFloor, n.MaxFloorMaxValue())
		}
	}
	if n.CashOrderQtyInActingVersion(actingVersion) {
		if n.CashOrderQty != n.CashOrderQtyNullValue() && (n.CashOrderQty < n.CashOrderQtyMinValue() || n.CashOrderQty > n.CashOrderQtyMaxValue()) {
			return fmt.Errorf("Range check failed on n.CashOrderQty (%v < %v > %v)", n.CashOrderQtyMinValue(), n.CashOrderQty, n.CashOrderQtyMaxValue())
		}
	}
	if n.TargetStrategyInActingVersion(actingVersion) {
		if n.TargetStrategy != n.TargetStrategyNullValue() && (n.TargetStrategy < n.TargetStrategyMinValue() || n.TargetStrategy > n.TargetStrategyMaxValue()) {
			return fmt.Errorf("Range check failed on n.TargetStrategy (%v < %v > %v)", n.TargetStrategyMinValue(), n.TargetStrategy, n.TargetStrategyMaxValue())
		}
	}
	if n.StrategyIDInActingVersion(actingVersion) {
		if n.StrategyID != n.StrategyIDNullValue() && (n.StrategyID < n.StrategyIDMinValue() || n.StrategyID > n.StrategyIDMaxValue()) {
			return fmt.Errorf("Range check failed on n.StrategyID (%v < %v > %v)", n.StrategyIDMinValue(), n.StrategyID, n.StrategyIDMaxValue())
		}
	}
	if err := n.SelfTradePreventionMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.SOR.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(n.ClOrdID[:]) {
		return errors.New("n.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(n.Symbol[:]) {
		return errors.New("n.Symbol failed UTF-8 validation")
	}
	return nil
}

func NewOrderSingleInit(n *NewOrderSingle) {
	n.OrderQty = math.MinInt64
	n.Price = math.MinInt64
	n.TriggerPrice = math.MinInt64
	n.TriggerTrailingDeltaBips = math.MaxUint64
	n.PegOffsetValue = math.MaxUint8
	n.MaxFloor = math.MinInt64
	n.CashOrderQty = math.MinInt64
	n.TargetStrategy = math.MinInt32
	n.StrategyID = math.MinInt64
	return
}

func (*NewOrderSingle) SbeBlockLength() (blockLength uint16) {
	return 76
}

func (*NewOrderSingle) SbeTemplateId() (templateId uint16) {
	return 99
}

func (*NewOrderSingle) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*NewOrderSingle) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*NewOrderSingle) SbeSemanticType() (semanticType []byte) {
	return []byte("D")
}

func (*NewOrderSingle) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*NewOrderSingle) PriceExponentId() uint16 {
	return 25054
}

func (*NewOrderSingle) PriceExponentSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceExponentSinceVersion()
}

func (*NewOrderSingle) PriceExponentDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PriceExponentMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*NewOrderSingle) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*NewOrderSingle) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*NewOrderSingle) QtyExponentId() uint16 {
	return 25055
}

func (*NewOrderSingle) QtyExponentSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.QtyExponentSinceVersion()
}

func (*NewOrderSingle) QtyExponentDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) QtyExponentMetaAttribute(meta int) string {
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

func (*NewOrderSingle) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*NewOrderSingle) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*NewOrderSingle) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*NewOrderSingle) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderSingle) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderSingle) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderSingle) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderSingle) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) ExecInstId() uint16 {
	return 18
}

func (*NewOrderSingle) ExecInstSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExecInstSinceVersion()
}

func (*NewOrderSingle) ExecInstDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) ExecInstMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PriceId() uint16 {
	return 44
}

func (*NewOrderSingle) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderSingle) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PriceMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) PriceNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) TriggerTypeId() uint16 {
	return 1100
}

func (*NewOrderSingle) TriggerTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerTypeSinceVersion()
}

func (*NewOrderSingle) TriggerTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerActionId() uint16 {
	return 1101
}

func (*NewOrderSingle) TriggerActionSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerActionSinceVersion()
}

func (*NewOrderSingle) TriggerActionDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerActionMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerPriceId() uint16 {
	return 1102
}

func (*NewOrderSingle) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceSinceVersion()
}

func (*NewOrderSingle) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerPriceMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) TriggerPriceTypeId() uint16 {
	return 1107
}

func (*NewOrderSingle) TriggerPriceTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceTypeSinceVersion()
}

func (*NewOrderSingle) TriggerPriceTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerPriceTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerPriceDirectionId() uint16 {
	return 1109
}

func (*NewOrderSingle) TriggerPriceDirectionSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerPriceDirectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceDirectionSinceVersion()
}

func (*NewOrderSingle) TriggerPriceDirectionDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerPriceDirectionMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerTrailingDeltaBipsId() uint16 {
	return 25009
}

func (*NewOrderSingle) TriggerTrailingDeltaBipsSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TriggerTrailingDeltaBipsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerTrailingDeltaBipsSinceVersion()
}

func (*NewOrderSingle) TriggerTrailingDeltaBipsDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TriggerTrailingDeltaBipsMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TriggerTrailingDeltaBipsMinValue() uint64 {
	return 0
}

func (*NewOrderSingle) TriggerTrailingDeltaBipsMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle) TriggerTrailingDeltaBipsNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle) PegOffsetValueId() uint16 {
	return 211
}

func (*NewOrderSingle) PegOffsetValueSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PegOffsetValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegOffsetValueSinceVersion()
}

func (*NewOrderSingle) PegOffsetValueDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PegOffsetValueMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PegOffsetValueMinValue() uint8 {
	return 0
}

func (*NewOrderSingle) PegOffsetValueMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*NewOrderSingle) PegOffsetValueNullValue() uint8 {
	return math.MaxUint8
}

func (*NewOrderSingle) PegPriceTypeId() uint16 {
	return 1094
}

func (*NewOrderSingle) PegPriceTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PegPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegPriceTypeSinceVersion()
}

func (*NewOrderSingle) PegPriceTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PegPriceTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PegMoveTypeId() uint16 {
	return 835
}

func (*NewOrderSingle) PegMoveTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PegMoveTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegMoveTypeSinceVersion()
}

func (*NewOrderSingle) PegMoveTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PegMoveTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PegOffsetTypeId() uint16 {
	return 836
}

func (*NewOrderSingle) PegOffsetTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PegOffsetTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegOffsetTypeSinceVersion()
}

func (*NewOrderSingle) PegOffsetTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PegOffsetTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) SideId() uint16 {
	return 54
}

func (*NewOrderSingle) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderSingle) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SideMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TimeInForceId() uint16 {
	return 59
}

func (*NewOrderSingle) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrderSingle) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TimeInForceMetaAttribute(meta int) string {
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

func (*NewOrderSingle) MaxFloorId() uint16 {
	return 111
}

func (*NewOrderSingle) MaxFloorSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) MaxFloorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MaxFloorSinceVersion()
}

func (*NewOrderSingle) MaxFloorDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) MaxFloorMetaAttribute(meta int) string {
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

func (*NewOrderSingle) MaxFloorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) MaxFloorMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) MaxFloorNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) CashOrderQtyId() uint16 {
	return 152
}

func (*NewOrderSingle) CashOrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) CashOrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CashOrderQtySinceVersion()
}

func (*NewOrderSingle) CashOrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) CashOrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle) CashOrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) CashOrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) CashOrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) TargetStrategyId() uint16 {
	return 847
}

func (*NewOrderSingle) TargetStrategySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TargetStrategyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TargetStrategySinceVersion()
}

func (*NewOrderSingle) TargetStrategyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TargetStrategyMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TargetStrategyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderSingle) TargetStrategyMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderSingle) TargetStrategyNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderSingle) StrategyIDId() uint16 {
	return 7940
}

func (*NewOrderSingle) StrategyIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) StrategyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.StrategyIDSinceVersion()
}

func (*NewOrderSingle) StrategyIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) StrategyIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle) StrategyIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) StrategyIDMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) StrategyIDNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) SelfTradePreventionModeId() uint16 {
	return 25001
}

func (*NewOrderSingle) SelfTradePreventionModeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SelfTradePreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SelfTradePreventionModeSinceVersion()
}

func (*NewOrderSingle) SelfTradePreventionModeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SelfTradePreventionModeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) SORId() uint16 {
	return 25032
}

func (*NewOrderSingle) SORSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SORInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SORSinceVersion()
}

func (*NewOrderSingle) SORDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SORMetaAttribute(meta int) string {
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

func (*NewOrderSingle) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderSingle) ClOrdIDDeprecated() uint16 {
	return 0
}

func (NewOrderSingle) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (NewOrderSingle) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*NewOrderSingle) SymbolMetaAttribute(meta int) string {
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

func (*NewOrderSingle) SymbolSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SymbolSinceVersion()
}

func (*NewOrderSingle) SymbolDeprecated() uint16 {
	return 0
}

func (NewOrderSingle) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (NewOrderSingle) SymbolHeaderLength() uint64 {
	return 1
}
