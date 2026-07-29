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

type ListStatus struct {
	ListID           int64
	ContingencyType  ContingencyTypeEnum
	ListStatusType   ListStatusTypeEnum
	ListOrderStatus  ListOrderStatusEnum
	ListRejectReason ListRejectReasonEnum
	TransactTime     int64
	Orders           []ListStatusOrders
	ClListID         []uint8
	OrigClListID     []uint8
}
type ListStatusOrders struct {
	OrderID                    int64
	OrdRejReason               OrdRejReasonEnum
	ErrorCode                  int32
	ListTriggeringInstructions []ListStatusOrdersListTriggeringInstructions
	ClOrdID                    []uint8
	Symbol                     []uint8
	ErrorText                  []uint8
}
type ListStatusOrdersListTriggeringInstructions struct {
	ListTriggerType         ListTriggerTypeEnum
	ListTriggerTriggerIndex uint8
	ListTriggerAction       ListTriggerActionEnum
}

func (l *ListStatus) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, l.ListID); err != nil {
		return err
	}
	if err := l.ContingencyType.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.ListStatusType.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.ListOrderStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.ListRejectReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, l.TransactTime); err != nil {
		return err
	}
	var OrdersBlockLength uint8 = 13
	if err := _m.WriteUint8(_w, OrdersBlockLength); err != nil {
		return err
	}
	var OrdersNumInGroup uint8 = uint8(len(l.Orders))
	if err := _m.WriteUint8(_w, OrdersNumInGroup); err != nil {
		return err
	}
	for i := range l.Orders {
		if err := l.Orders[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(l.ClListID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ClListID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(l.OrigClListID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.OrigClListID); err != nil {
		return err
	}
	return nil
}

func (l *ListStatus) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.ListIDInActingVersion(actingVersion) {
		l.ListID = l.ListIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.ListID); err != nil {
			return err
		}
	}
	if l.ContingencyTypeInActingVersion(actingVersion) {
		if err := l.ContingencyType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.ListStatusTypeInActingVersion(actingVersion) {
		if err := l.ListStatusType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.ListOrderStatusInActingVersion(actingVersion) {
		if err := l.ListOrderStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.ListRejectReasonInActingVersion(actingVersion) {
		if err := l.ListRejectReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.TransactTimeInActingVersion(actingVersion) {
		l.TransactTime = l.TransactTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.TransactTime); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.OrdersInActingVersion(actingVersion) {
		var OrdersBlockLength uint8
		if err := _m.ReadUint8(_r, &OrdersBlockLength); err != nil {
			return err
		}
		var OrdersNumInGroup uint8
		if err := _m.ReadUint8(_r, &OrdersNumInGroup); err != nil {
			return err
		}
		if cap(l.Orders) < int(OrdersNumInGroup) {
			l.Orders = make([]ListStatusOrders, OrdersNumInGroup)
		}
		l.Orders = l.Orders[:OrdersNumInGroup]
		for i := range l.Orders {
			if err := l.Orders[i].Decode(_m, _r, actingVersion, uint(OrdersBlockLength)); err != nil {
				return err
			}
		}
	}

	if l.ClListIDInActingVersion(actingVersion) {
		var ClListIDLength uint8
		if err := _m.ReadUint8(_r, &ClListIDLength); err != nil {
			return err
		}
		if cap(l.ClListID) < int(ClListIDLength) {
			l.ClListID = make([]uint8, ClListIDLength)
		}
		l.ClListID = l.ClListID[:ClListIDLength]
		if err := _m.ReadBytes(_r, l.ClListID); err != nil {
			return err
		}
	}

	if l.OrigClListIDInActingVersion(actingVersion) {
		var OrigClListIDLength uint8
		if err := _m.ReadUint8(_r, &OrigClListIDLength); err != nil {
			return err
		}
		if cap(l.OrigClListID) < int(OrigClListIDLength) {
			l.OrigClListID = make([]uint8, OrigClListIDLength)
		}
		l.OrigClListID = l.OrigClListID[:OrigClListIDLength]
		if err := _m.ReadBytes(_r, l.OrigClListID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListStatus) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.ListIDInActingVersion(actingVersion) {
		if l.ListID != l.ListIDNullValue() && (l.ListID < l.ListIDMinValue() || l.ListID > l.ListIDMaxValue()) {
			return fmt.Errorf("Range check failed on l.ListID (%v < %v > %v)", l.ListIDMinValue(), l.ListID, l.ListIDMaxValue())
		}
	}
	if err := l.ContingencyType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.ListStatusType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.ListOrderStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.ListRejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if l.TransactTimeInActingVersion(actingVersion) {
		if l.TransactTime != l.TransactTimeNullValue() && (l.TransactTime < l.TransactTimeMinValue() || l.TransactTime > l.TransactTimeMaxValue()) {
			return fmt.Errorf("Range check failed on l.TransactTime (%v < %v > %v)", l.TransactTimeMinValue(), l.TransactTime, l.TransactTimeMaxValue())
		}
	}
	for i := range l.Orders {
		if err := l.Orders[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(l.ClListID[:]) {
		return errors.New("l.ClListID failed UTF-8 validation")
	}
	if !utf8.Valid(l.OrigClListID[:]) {
		return errors.New("l.OrigClListID failed UTF-8 validation")
	}
	return nil
}

func ListStatusInit(l *ListStatus) {
	l.ListID = math.MinInt64
	l.TransactTime = math.MinInt64
	return
}

func (l *ListStatusOrders) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, l.OrderID); err != nil {
		return err
	}
	if err := l.OrdRejReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.ErrorCode); err != nil {
		return err
	}
	var ListTriggeringInstructionsBlockLength uint8 = 3
	if err := _m.WriteUint8(_w, ListTriggeringInstructionsBlockLength); err != nil {
		return err
	}
	var ListTriggeringInstructionsNumInGroup uint8 = uint8(len(l.ListTriggeringInstructions))
	if err := _m.WriteUint8(_w, ListTriggeringInstructionsNumInGroup); err != nil {
		return err
	}
	for i := range l.ListTriggeringInstructions {
		if err := l.ListTriggeringInstructions[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(l.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(l.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.ErrorText))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ErrorText); err != nil {
		return err
	}
	return nil
}

func (l *ListStatusOrders) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !l.OrderIDInActingVersion(actingVersion) {
		l.OrderID = l.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &l.OrderID); err != nil {
			return err
		}
	}
	if l.OrdRejReasonInActingVersion(actingVersion) {
		if err := l.OrdRejReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.ErrorCodeInActingVersion(actingVersion) {
		l.ErrorCode = l.ErrorCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.ErrorCode); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.ListTriggeringInstructionsInActingVersion(actingVersion) {
		var ListTriggeringInstructionsBlockLength uint8
		if err := _m.ReadUint8(_r, &ListTriggeringInstructionsBlockLength); err != nil {
			return err
		}
		var ListTriggeringInstructionsNumInGroup uint8
		if err := _m.ReadUint8(_r, &ListTriggeringInstructionsNumInGroup); err != nil {
			return err
		}
		if cap(l.ListTriggeringInstructions) < int(ListTriggeringInstructionsNumInGroup) {
			l.ListTriggeringInstructions = make([]ListStatusOrdersListTriggeringInstructions, ListTriggeringInstructionsNumInGroup)
		}
		l.ListTriggeringInstructions = l.ListTriggeringInstructions[:ListTriggeringInstructionsNumInGroup]
		for i := range l.ListTriggeringInstructions {
			if err := l.ListTriggeringInstructions[i].Decode(_m, _r, actingVersion, uint(ListTriggeringInstructionsBlockLength)); err != nil {
				return err
			}
		}
	}

	if l.ClOrdIDInActingVersion(actingVersion) {
		var ClOrdIDLength uint8
		if err := _m.ReadUint8(_r, &ClOrdIDLength); err != nil {
			return err
		}
		if cap(l.ClOrdID) < int(ClOrdIDLength) {
			l.ClOrdID = make([]uint8, ClOrdIDLength)
		}
		l.ClOrdID = l.ClOrdID[:ClOrdIDLength]
		if err := _m.ReadBytes(_r, l.ClOrdID); err != nil {
			return err
		}
	}

	if l.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(l.Symbol) < int(SymbolLength) {
			l.Symbol = make([]uint8, SymbolLength)
		}
		l.Symbol = l.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, l.Symbol); err != nil {
			return err
		}
	}

	if l.ErrorTextInActingVersion(actingVersion) {
		var ErrorTextLength uint16
		if err := _m.ReadUint16(_r, &ErrorTextLength); err != nil {
			return err
		}
		if cap(l.ErrorText) < int(ErrorTextLength) {
			l.ErrorText = make([]uint8, ErrorTextLength)
		}
		l.ErrorText = l.ErrorText[:ErrorTextLength]
		if err := _m.ReadBytes(_r, l.ErrorText); err != nil {
			return err
		}
	}
	return nil
}

func (l *ListStatusOrders) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.OrderIDInActingVersion(actingVersion) {
		if l.OrderID != l.OrderIDNullValue() && (l.OrderID < l.OrderIDMinValue() || l.OrderID > l.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on l.OrderID (%v < %v > %v)", l.OrderIDMinValue(), l.OrderID, l.OrderIDMaxValue())
		}
	}
	if err := l.OrdRejReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if l.ErrorCodeInActingVersion(actingVersion) {
		if l.ErrorCode != l.ErrorCodeNullValue() && (l.ErrorCode < l.ErrorCodeMinValue() || l.ErrorCode > l.ErrorCodeMaxValue()) {
			return fmt.Errorf("Range check failed on l.ErrorCode (%v < %v > %v)", l.ErrorCodeMinValue(), l.ErrorCode, l.ErrorCodeMaxValue())
		}
	}
	for i := range l.ListTriggeringInstructions {
		if err := l.ListTriggeringInstructions[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(l.ClOrdID[:]) {
		return errors.New("l.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(l.Symbol[:]) {
		return errors.New("l.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(l.ErrorText[:]) {
		return errors.New("l.ErrorText failed UTF-8 validation")
	}
	return nil
}

func ListStatusOrdersInit(l *ListStatusOrders) {
	l.OrderID = math.MinInt64
	l.ErrorCode = math.MinInt32
	return
}

func (l *ListStatusOrdersListTriggeringInstructions) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := l.ListTriggerType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, l.ListTriggerTriggerIndex); err != nil {
		return err
	}
	if err := l.ListTriggerAction.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (l *ListStatusOrdersListTriggeringInstructions) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if l.ListTriggerTypeInActingVersion(actingVersion) {
		if err := l.ListTriggerType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.ListTriggerTriggerIndexInActingVersion(actingVersion) {
		l.ListTriggerTriggerIndex = l.ListTriggerTriggerIndexNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.ListTriggerTriggerIndex); err != nil {
			return err
		}
	}
	if l.ListTriggerActionInActingVersion(actingVersion) {
		if err := l.ListTriggerAction.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}
	return nil
}

func (l *ListStatusOrdersListTriggeringInstructions) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := l.ListTriggerType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if l.ListTriggerTriggerIndexInActingVersion(actingVersion) {
		if l.ListTriggerTriggerIndex < l.ListTriggerTriggerIndexMinValue() || l.ListTriggerTriggerIndex > l.ListTriggerTriggerIndexMaxValue() {
			return fmt.Errorf("Range check failed on l.ListTriggerTriggerIndex (%v < %v > %v)", l.ListTriggerTriggerIndexMinValue(), l.ListTriggerTriggerIndex, l.ListTriggerTriggerIndexMaxValue())
		}
	}
	if err := l.ListTriggerAction.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func ListStatusOrdersListTriggeringInstructionsInit(l *ListStatusOrdersListTriggeringInstructions) {
	return
}

func (*ListStatus) SbeBlockLength() (blockLength uint16) {
	return 20
}

func (*ListStatus) SbeTemplateId() (templateId uint16) {
	return 102
}

func (*ListStatus) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ListStatus) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ListStatus) SbeSemanticType() (semanticType []byte) {
	return []byte("N")
}

func (*ListStatus) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*ListStatus) ListIDId() uint16 {
	return 66
}

func (*ListStatus) ListIDSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListIDSinceVersion()
}

func (*ListStatus) ListIDDeprecated() uint16 {
	return 0
}

func (*ListStatus) ListIDMetaAttribute(meta int) string {
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

func (*ListStatus) ListIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ListStatus) ListIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ListStatus) ListIDNullValue() int64 {
	return math.MinInt64
}

func (*ListStatus) ContingencyTypeId() uint16 {
	return 1385
}

func (*ListStatus) ContingencyTypeSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ContingencyTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ContingencyTypeSinceVersion()
}

func (*ListStatus) ContingencyTypeDeprecated() uint16 {
	return 0
}

func (*ListStatus) ContingencyTypeMetaAttribute(meta int) string {
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

func (*ListStatus) ListStatusTypeId() uint16 {
	return 429
}

func (*ListStatus) ListStatusTypeSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ListStatusTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListStatusTypeSinceVersion()
}

func (*ListStatus) ListStatusTypeDeprecated() uint16 {
	return 0
}

func (*ListStatus) ListStatusTypeMetaAttribute(meta int) string {
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

func (*ListStatus) ListOrderStatusId() uint16 {
	return 431
}

func (*ListStatus) ListOrderStatusSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ListOrderStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListOrderStatusSinceVersion()
}

func (*ListStatus) ListOrderStatusDeprecated() uint16 {
	return 0
}

func (*ListStatus) ListOrderStatusMetaAttribute(meta int) string {
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

func (*ListStatus) ListRejectReasonId() uint16 {
	return 1386
}

func (*ListStatus) ListRejectReasonSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ListRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListRejectReasonSinceVersion()
}

func (*ListStatus) ListRejectReasonDeprecated() uint16 {
	return 0
}

func (*ListStatus) ListRejectReasonMetaAttribute(meta int) string {
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

func (*ListStatus) TransactTimeId() uint16 {
	return 60
}

func (*ListStatus) TransactTimeSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TransactTimeSinceVersion()
}

func (*ListStatus) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ListStatus) TransactTimeMetaAttribute(meta int) string {
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

func (*ListStatus) TransactTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ListStatus) TransactTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ListStatus) TransactTimeNullValue() int64 {
	return math.MinInt64
}

func (*ListStatusOrders) OrderIDId() uint16 {
	return 37
}

func (*ListStatusOrders) OrderIDSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderIDSinceVersion()
}

func (*ListStatusOrders) OrderIDDeprecated() uint16 {
	return 0
}

func (*ListStatusOrders) OrderIDMetaAttribute(meta int) string {
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

func (*ListStatusOrders) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ListStatusOrders) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ListStatusOrders) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*ListStatusOrders) OrdRejReasonId() uint16 {
	return 103
}

func (*ListStatusOrders) OrdRejReasonSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) OrdRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrdRejReasonSinceVersion()
}

func (*ListStatusOrders) OrdRejReasonDeprecated() uint16 {
	return 0
}

func (*ListStatusOrders) OrdRejReasonMetaAttribute(meta int) string {
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

func (*ListStatusOrders) ErrorCodeId() uint16 {
	return 25016
}

func (*ListStatusOrders) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ErrorCodeSinceVersion()
}

func (*ListStatusOrders) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*ListStatusOrders) ErrorCodeMetaAttribute(meta int) string {
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

func (*ListStatusOrders) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ListStatusOrders) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*ListStatusOrders) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTypeId() uint16 {
	return 25011
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTypeSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrdersListTriggeringInstructions) ListTriggerTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListTriggerTypeSinceVersion()
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTypeDeprecated() uint16 {
	return 0
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTypeMetaAttribute(meta int) string {
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

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexId() uint16 {
	return 25012
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListTriggerTriggerIndexSinceVersion()
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexDeprecated() uint16 {
	return 0
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexMetaAttribute(meta int) string {
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

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexMinValue() uint8 {
	return 0
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerTriggerIndexNullValue() uint8 {
	return math.MaxUint8
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerActionId() uint16 {
	return 25013
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerActionSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrdersListTriggeringInstructions) ListTriggerActionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListTriggerActionSinceVersion()
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerActionDeprecated() uint16 {
	return 0
}

func (*ListStatusOrdersListTriggeringInstructions) ListTriggerActionMetaAttribute(meta int) string {
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

func (*ListStatusOrders) ClOrdIDMetaAttribute(meta int) string {
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

func (*ListStatusOrders) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ClOrdIDSinceVersion()
}

func (*ListStatusOrders) ClOrdIDDeprecated() uint16 {
	return 0
}

func (ListStatusOrders) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (ListStatusOrders) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*ListStatusOrders) SymbolMetaAttribute(meta int) string {
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

func (*ListStatusOrders) SymbolSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SymbolSinceVersion()
}

func (*ListStatusOrders) SymbolDeprecated() uint16 {
	return 0
}

func (ListStatusOrders) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (ListStatusOrders) SymbolHeaderLength() uint64 {
	return 1
}

func (*ListStatusOrders) ErrorTextMetaAttribute(meta int) string {
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

func (*ListStatusOrders) ErrorTextSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ErrorTextSinceVersion()
}

func (*ListStatusOrders) ErrorTextDeprecated() uint16 {
	return 0
}

func (ListStatusOrders) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (ListStatusOrders) ErrorTextHeaderLength() uint64 {
	return 2
}

func (*ListStatus) OrdersId() uint16 {
	return 73
}

func (*ListStatus) OrdersSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) OrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrdersSinceVersion()
}

func (*ListStatus) OrdersDeprecated() uint16 {
	return 0
}

func (*ListStatusOrders) SbeBlockLength() (blockLength uint) {
	return 13
}

func (*ListStatusOrders) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ListStatusOrders) ListTriggeringInstructionsId() uint16 {
	return 25010
}

func (*ListStatusOrders) ListTriggeringInstructionsSinceVersion() uint16 {
	return 0
}

func (l *ListStatusOrders) ListTriggeringInstructionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ListTriggeringInstructionsSinceVersion()
}

func (*ListStatusOrders) ListTriggeringInstructionsDeprecated() uint16 {
	return 0
}

func (*ListStatusOrdersListTriggeringInstructions) SbeBlockLength() (blockLength uint) {
	return 3
}

func (*ListStatusOrdersListTriggeringInstructions) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ListStatus) ClListIDMetaAttribute(meta int) string {
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

func (*ListStatus) ClListIDSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) ClListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ClListIDSinceVersion()
}

func (*ListStatus) ClListIDDeprecated() uint16 {
	return 0
}

func (ListStatus) ClListIDCharacterEncoding() string {
	return "UTF-8"
}

func (ListStatus) ClListIDHeaderLength() uint64 {
	return 1
}

func (*ListStatus) OrigClListIDMetaAttribute(meta int) string {
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

func (*ListStatus) OrigClListIDSinceVersion() uint16 {
	return 0
}

func (l *ListStatus) OrigClListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrigClListIDSinceVersion()
}

func (*ListStatus) OrigClListIDDeprecated() uint16 {
	return 0
}

func (ListStatus) OrigClListIDCharacterEncoding() string {
	return "UTF-8"
}

func (ListStatus) OrigClListIDHeaderLength() uint64 {
	return 1
}
