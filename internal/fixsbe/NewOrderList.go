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

type NewOrderList struct {
	ContingencyType ContingencyTypeEnum
	OPO             BoolEnumEnum
	Orders          []NewOrderListOrders
	ClListID        []uint8
}
type NewOrderListOrders struct {
	PriceExponent              int8
	QtyExponent                int8
	OrderQty                   int64
	OrdType                    OrdTypeEnum
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
	Side                       SideEnum
	TimeInForce                TimeInForceEnum
	MaxFloor                   int64
	CashOrderQty               int64
	TargetStrategy             int32
	StrategyID                 int64
	SelfTradePreventionMode    SelfTradePreventionModeEnum
	ListTriggeringInstructions []NewOrderListOrdersListTriggeringInstructions
	ClOrdID                    []uint8
	Symbol                     []uint8
}
type NewOrderListOrdersListTriggeringInstructions struct {
	ListTriggerType         ListTriggerTypeEnum
	ListTriggerTriggerIndex uint8
	ListTriggerAction       ListTriggerActionEnum
}

func (n *NewOrderList) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := n.ContingencyType.Encode(_m, _w); err != nil {
		return err
	}
	if err := n.OPO.Encode(_m, _w); err != nil {
		return err
	}
	var OrdersBlockLength uint16 = 75
	if err := _m.WriteUint16(_w, OrdersBlockLength); err != nil {
		return err
	}
	var OrdersNumInGroup uint8 = uint8(len(n.Orders))
	if err := _m.WriteUint8(_w, OrdersNumInGroup); err != nil {
		return err
	}
	for i := range n.Orders {
		if err := n.Orders[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(n.ClListID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.ClListID); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if n.ContingencyTypeInActingVersion(actingVersion) {
		if err := n.ContingencyType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if n.OPOInActingVersion(actingVersion) {
		if err := n.OPO.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.OrdersInActingVersion(actingVersion) {
		var OrdersBlockLength uint16
		if err := _m.ReadUint16(_r, &OrdersBlockLength); err != nil {
			return err
		}
		var OrdersNumInGroup uint8
		if err := _m.ReadUint8(_r, &OrdersNumInGroup); err != nil {
			return err
		}
		if cap(n.Orders) < int(OrdersNumInGroup) {
			n.Orders = make([]NewOrderListOrders, OrdersNumInGroup)
		}
		n.Orders = n.Orders[:OrdersNumInGroup]
		for i := range n.Orders {
			if err := n.Orders[i].Decode(_m, _r, actingVersion, uint(OrdersBlockLength)); err != nil {
				return err
			}
		}
	}

	if n.ClListIDInActingVersion(actingVersion) {
		var ClListIDLength uint8
		if err := _m.ReadUint8(_r, &ClListIDLength); err != nil {
			return err
		}
		if cap(n.ClListID) < int(ClListIDLength) {
			n.ClListID = make([]uint8, ClListIDLength)
		}
		n.ClListID = n.ClListID[:ClListIDLength]
		if err := _m.ReadBytes(_r, n.ClListID); err != nil {
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

func (n *NewOrderList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := n.ContingencyType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := n.OPO.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for i := range n.Orders {
		if err := n.Orders[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(n.ClListID[:]) {
		return errors.New("n.ClListID failed UTF-8 validation")
	}
	return nil
}

func NewOrderListInit(n *NewOrderList) {
	return
}

func (n *NewOrderListOrders) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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
	var ListTriggeringInstructionsBlockLength uint8 = 3
	if err := _m.WriteUint8(_w, ListTriggeringInstructionsBlockLength); err != nil {
		return err
	}
	var ListTriggeringInstructionsNumInGroup uint8 = uint8(len(n.ListTriggeringInstructions))
	if err := _m.WriteUint8(_w, ListTriggeringInstructionsNumInGroup); err != nil {
		return err
	}
	for i := range n.ListTriggeringInstructions {
		if err := n.ListTriggeringInstructions[i].Encode(_m, _w); err != nil {
			return err
		}
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

func (n *NewOrderListOrders) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.ListTriggeringInstructionsInActingVersion(actingVersion) {
		var ListTriggeringInstructionsBlockLength uint8
		if err := _m.ReadUint8(_r, &ListTriggeringInstructionsBlockLength); err != nil {
			return err
		}
		var ListTriggeringInstructionsNumInGroup uint8
		if err := _m.ReadUint8(_r, &ListTriggeringInstructionsNumInGroup); err != nil {
			return err
		}
		if cap(n.ListTriggeringInstructions) < int(ListTriggeringInstructionsNumInGroup) {
			n.ListTriggeringInstructions = make([]NewOrderListOrdersListTriggeringInstructions, ListTriggeringInstructionsNumInGroup)
		}
		n.ListTriggeringInstructions = n.ListTriggeringInstructions[:ListTriggeringInstructionsNumInGroup]
		for i := range n.ListTriggeringInstructions {
			if err := n.ListTriggeringInstructions[i].Decode(_m, _r, actingVersion, uint(ListTriggeringInstructionsBlockLength)); err != nil {
				return err
			}
		}
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
	return nil
}

func (n *NewOrderListOrders) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	for i := range n.ListTriggeringInstructions {
		if err := n.ListTriggeringInstructions[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(n.ClOrdID[:]) {
		return errors.New("n.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(n.Symbol[:]) {
		return errors.New("n.Symbol failed UTF-8 validation")
	}
	return nil
}

func NewOrderListOrdersInit(n *NewOrderListOrders) {
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

func (n *NewOrderListOrdersListTriggeringInstructions) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := n.ListTriggerType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, n.ListTriggerTriggerIndex); err != nil {
		return err
	}
	if err := n.ListTriggerAction.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderListOrdersListTriggeringInstructions) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if n.ListTriggerTypeInActingVersion(actingVersion) {
		if err := n.ListTriggerType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.ListTriggerTriggerIndexInActingVersion(actingVersion) {
		n.ListTriggerTriggerIndex = n.ListTriggerTriggerIndexNullValue()
	} else {
		if err := _m.ReadUint8(_r, &n.ListTriggerTriggerIndex); err != nil {
			return err
		}
	}
	if n.ListTriggerActionInActingVersion(actingVersion) {
		if err := n.ListTriggerAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	return nil
}

func (n *NewOrderListOrdersListTriggeringInstructions) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := n.ListTriggerType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.ListTriggerTriggerIndexInActingVersion(actingVersion) {
		if n.ListTriggerTriggerIndex < n.ListTriggerTriggerIndexMinValue() || n.ListTriggerTriggerIndex > n.ListTriggerTriggerIndexMaxValue() {
			return fmt.Errorf("Range check failed on n.ListTriggerTriggerIndex (%v < %v > %v)", n.ListTriggerTriggerIndexMinValue(), n.ListTriggerTriggerIndex, n.ListTriggerTriggerIndexMaxValue())
		}
	}
	if err := n.ListTriggerAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func NewOrderListOrdersListTriggeringInstructionsInit(n *NewOrderListOrdersListTriggeringInstructions) {
	return
}

func (*NewOrderList) SbeBlockLength() (blockLength uint16) {
	return 2
}

func (*NewOrderList) SbeTemplateId() (templateId uint16) {
	return 100
}

func (*NewOrderList) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*NewOrderList) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*NewOrderList) SbeSemanticType() (semanticType []byte) {
	return []byte("E")
}

func (*NewOrderList) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*NewOrderList) ContingencyTypeId() uint16 {
	return 1385
}

func (*NewOrderList) ContingencyTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderList) ContingencyTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ContingencyTypeSinceVersion()
}

func (*NewOrderList) ContingencyTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderList) ContingencyTypeMetaAttribute(meta int) string {
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

func (*NewOrderList) OPOId() uint16 {
	return 25046
}

func (*NewOrderList) OPOSinceVersion() uint16 {
	return 0
}

func (n *NewOrderList) OPOInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OPOSinceVersion()
}

func (*NewOrderList) OPODeprecated() uint16 {
	return 0
}

func (*NewOrderList) OPOMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PriceExponentId() uint16 {
	return 25054
}

func (*NewOrderListOrders) PriceExponentSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceExponentSinceVersion()
}

func (*NewOrderListOrders) PriceExponentDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PriceExponentMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*NewOrderListOrders) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*NewOrderListOrders) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*NewOrderListOrders) QtyExponentId() uint16 {
	return 25055
}

func (*NewOrderListOrders) QtyExponentSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.QtyExponentSinceVersion()
}

func (*NewOrderListOrders) QtyExponentDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) QtyExponentMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*NewOrderListOrders) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*NewOrderListOrders) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*NewOrderListOrders) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderListOrders) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderListOrders) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderListOrders) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderListOrders) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) ExecInstId() uint16 {
	return 18
}

func (*NewOrderListOrders) ExecInstSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) ExecInstInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ExecInstSinceVersion()
}

func (*NewOrderListOrders) ExecInstDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) ExecInstMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PriceId() uint16 {
	return 44
}

func (*NewOrderListOrders) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderListOrders) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PriceMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) PriceNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) TriggerTypeId() uint16 {
	return 1100
}

func (*NewOrderListOrders) TriggerTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerTypeSinceVersion()
}

func (*NewOrderListOrders) TriggerTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerActionId() uint16 {
	return 1101
}

func (*NewOrderListOrders) TriggerActionSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerActionSinceVersion()
}

func (*NewOrderListOrders) TriggerActionDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerActionMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerPriceId() uint16 {
	return 1102
}

func (*NewOrderListOrders) TriggerPriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceSinceVersion()
}

func (*NewOrderListOrders) TriggerPriceDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerPriceMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerPriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) TriggerPriceMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) TriggerPriceNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) TriggerPriceTypeId() uint16 {
	return 1107
}

func (*NewOrderListOrders) TriggerPriceTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceTypeSinceVersion()
}

func (*NewOrderListOrders) TriggerPriceTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerPriceTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerPriceDirectionId() uint16 {
	return 1109
}

func (*NewOrderListOrders) TriggerPriceDirectionSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerPriceDirectionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerPriceDirectionSinceVersion()
}

func (*NewOrderListOrders) TriggerPriceDirectionDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerPriceDirectionMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerTrailingDeltaBipsId() uint16 {
	return 25009
}

func (*NewOrderListOrders) TriggerTrailingDeltaBipsSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TriggerTrailingDeltaBipsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TriggerTrailingDeltaBipsSinceVersion()
}

func (*NewOrderListOrders) TriggerTrailingDeltaBipsDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TriggerTrailingDeltaBipsMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TriggerTrailingDeltaBipsMinValue() uint64 {
	return 0
}

func (*NewOrderListOrders) TriggerTrailingDeltaBipsMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderListOrders) TriggerTrailingDeltaBipsNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderListOrders) PegOffsetValueId() uint16 {
	return 211
}

func (*NewOrderListOrders) PegOffsetValueSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PegOffsetValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegOffsetValueSinceVersion()
}

func (*NewOrderListOrders) PegOffsetValueDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PegOffsetValueMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PegOffsetValueMinValue() uint8 {
	return 0
}

func (*NewOrderListOrders) PegOffsetValueMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*NewOrderListOrders) PegOffsetValueNullValue() uint8 {
	return math.MaxUint8
}

func (*NewOrderListOrders) PegPriceTypeId() uint16 {
	return 1094
}

func (*NewOrderListOrders) PegPriceTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PegPriceTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegPriceTypeSinceVersion()
}

func (*NewOrderListOrders) PegPriceTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PegPriceTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PegMoveTypeId() uint16 {
	return 835
}

func (*NewOrderListOrders) PegMoveTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PegMoveTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegMoveTypeSinceVersion()
}

func (*NewOrderListOrders) PegMoveTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PegMoveTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) PegOffsetTypeId() uint16 {
	return 836
}

func (*NewOrderListOrders) PegOffsetTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) PegOffsetTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PegOffsetTypeSinceVersion()
}

func (*NewOrderListOrders) PegOffsetTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) PegOffsetTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) SideId() uint16 {
	return 54
}

func (*NewOrderListOrders) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderListOrders) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) SideMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TimeInForceId() uint16 {
	return 59
}

func (*NewOrderListOrders) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrderListOrders) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TimeInForceMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) MaxFloorId() uint16 {
	return 111
}

func (*NewOrderListOrders) MaxFloorSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) MaxFloorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.MaxFloorSinceVersion()
}

func (*NewOrderListOrders) MaxFloorDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) MaxFloorMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) MaxFloorMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) MaxFloorMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) MaxFloorNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) CashOrderQtyId() uint16 {
	return 152
}

func (*NewOrderListOrders) CashOrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) CashOrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.CashOrderQtySinceVersion()
}

func (*NewOrderListOrders) CashOrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) CashOrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) CashOrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) CashOrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) CashOrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) TargetStrategyId() uint16 {
	return 847
}

func (*NewOrderListOrders) TargetStrategySinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) TargetStrategyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TargetStrategySinceVersion()
}

func (*NewOrderListOrders) TargetStrategyDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) TargetStrategyMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) TargetStrategyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*NewOrderListOrders) TargetStrategyMaxValue() int32 {
	return math.MaxInt32
}

func (*NewOrderListOrders) TargetStrategyNullValue() int32 {
	return math.MinInt32
}

func (*NewOrderListOrders) StrategyIDId() uint16 {
	return 7940
}

func (*NewOrderListOrders) StrategyIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) StrategyIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.StrategyIDSinceVersion()
}

func (*NewOrderListOrders) StrategyIDDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) StrategyIDMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) StrategyIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderListOrders) StrategyIDMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderListOrders) StrategyIDNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderListOrders) SelfTradePreventionModeId() uint16 {
	return 25001
}

func (*NewOrderListOrders) SelfTradePreventionModeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) SelfTradePreventionModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SelfTradePreventionModeSinceVersion()
}

func (*NewOrderListOrders) SelfTradePreventionModeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) SelfTradePreventionModeMetaAttribute(meta int) string {
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

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTypeId() uint16 {
	return 25011
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrdersListTriggeringInstructions) ListTriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ListTriggerTypeSinceVersion()
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTypeMetaAttribute(meta int) string {
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

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexId() uint16 {
	return 25012
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ListTriggerTriggerIndexSinceVersion()
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexMetaAttribute(meta int) string {
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

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexMinValue() uint8 {
	return 0
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerTriggerIndexNullValue() uint8 {
	return math.MaxUint8
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerActionId() uint16 {
	return 25013
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerActionSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrdersListTriggeringInstructions) ListTriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ListTriggerActionSinceVersion()
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerActionDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrdersListTriggeringInstructions) ListTriggerActionMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderListOrders) ClOrdIDDeprecated() uint16 {
	return 0
}

func (NewOrderListOrders) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (NewOrderListOrders) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*NewOrderListOrders) SymbolMetaAttribute(meta int) string {
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

func (*NewOrderListOrders) SymbolSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SymbolSinceVersion()
}

func (*NewOrderListOrders) SymbolDeprecated() uint16 {
	return 0
}

func (NewOrderListOrders) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (NewOrderListOrders) SymbolHeaderLength() uint64 {
	return 1
}

func (*NewOrderList) OrdersId() uint16 {
	return 73
}

func (*NewOrderList) OrdersSinceVersion() uint16 {
	return 0
}

func (n *NewOrderList) OrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdersSinceVersion()
}

func (*NewOrderList) OrdersDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrders) SbeBlockLength() (blockLength uint) {
	return 75
}

func (*NewOrderListOrders) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*NewOrderListOrders) ListTriggeringInstructionsId() uint16 {
	return 25010
}

func (*NewOrderListOrders) ListTriggeringInstructionsSinceVersion() uint16 {
	return 0
}

func (n *NewOrderListOrders) ListTriggeringInstructionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ListTriggeringInstructionsSinceVersion()
}

func (*NewOrderListOrders) ListTriggeringInstructionsDeprecated() uint16 {
	return 0
}

func (*NewOrderListOrdersListTriggeringInstructions) SbeBlockLength() (blockLength uint) {
	return 3
}

func (*NewOrderListOrdersListTriggeringInstructions) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*NewOrderList) ClListIDMetaAttribute(meta int) string {
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

func (*NewOrderList) ClListIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderList) ClListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClListIDSinceVersion()
}

func (*NewOrderList) ClListIDDeprecated() uint16 {
	return 0
}

func (NewOrderList) ClListIDCharacterEncoding() string {
	return "UTF-8"
}

func (NewOrderList) ClListIDHeaderLength() uint64 {
	return 1
}
