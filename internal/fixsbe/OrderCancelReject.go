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

type OrderCancelReject struct {
	OrderID            int64
	ListID             int64
	CancelRestrictions CancelRestrictionsEnum
	CxlRejResponseTo   CxlRejResponseToEnum
	ErrorCode          int32
	ClOrdID            []uint8
	OrigClOrdID        []uint8
	OrigClListID       []uint8
	Symbol             []uint8
	ErrorText          []uint8
}

func (o *OrderCancelReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.OrderID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.ListID); err != nil {
		return err
	}
	if err := o.CancelRestrictions.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.CxlRejResponseTo.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.ErrorCode); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.OrigClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.OrigClListID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigClListID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.ErrorText))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ErrorText); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderID); err != nil {
			return err
		}
	}
	if !o.ListIDInActingVersion(actingVersion) {
		o.ListID = o.ListIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.ListID); err != nil {
			return err
		}
	}
	if o.CancelRestrictionsInActingVersion(actingVersion) {
		if err := o.CancelRestrictions.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.CxlRejResponseToInActingVersion(actingVersion) {
		if err := o.CxlRejResponseTo.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.ErrorCodeInActingVersion(actingVersion) {
		o.ErrorCode = o.ErrorCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.ErrorCode); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
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

	if o.OrigClListIDInActingVersion(actingVersion) {
		var OrigClListIDLength uint8
		if err := _m.ReadUint8(_r, &OrigClListIDLength); err != nil {
			return err
		}
		if cap(o.OrigClListID) < int(OrigClListIDLength) {
			o.OrigClListID = make([]uint8, OrigClListIDLength)
		}
		o.OrigClListID = o.OrigClListID[:OrigClListIDLength]
		if err := _m.ReadBytes(_r, o.OrigClListID); err != nil {
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

	if o.ErrorTextInActingVersion(actingVersion) {
		var ErrorTextLength uint16
		if err := _m.ReadUint16(_r, &ErrorTextLength); err != nil {
			return err
		}
		if cap(o.ErrorText) < int(ErrorTextLength) {
			o.ErrorText = make([]uint8, ErrorTextLength)
		}
		o.ErrorText = o.ErrorText[:ErrorTextLength]
		if err := _m.ReadBytes(_r, o.ErrorText); err != nil {
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

func (o *OrderCancelReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID != o.OrderIDNullValue() && (o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.ListIDInActingVersion(actingVersion) {
		if o.ListID != o.ListIDNullValue() && (o.ListID < o.ListIDMinValue() || o.ListID > o.ListIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.ListID (%v < %v > %v)", o.ListIDMinValue(), o.ListID, o.ListIDMaxValue())
		}
	}
	if err := o.CancelRestrictions.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.CxlRejResponseTo.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.ErrorCodeInActingVersion(actingVersion) {
		if o.ErrorCode < o.ErrorCodeMinValue() || o.ErrorCode > o.ErrorCodeMaxValue() {
			return fmt.Errorf("Range check failed on o.ErrorCode (%v < %v > %v)", o.ErrorCodeMinValue(), o.ErrorCode, o.ErrorCodeMaxValue())
		}
	}
	if !utf8.Valid(o.ClOrdID[:]) {
		return errors.New("o.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigClOrdID[:]) {
		return errors.New("o.OrigClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigClListID[:]) {
		return errors.New("o.OrigClListID failed UTF-8 validation")
	}
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(o.ErrorText[:]) {
		return errors.New("o.ErrorText failed UTF-8 validation")
	}
	return nil
}

func OrderCancelRejectInit(o *OrderCancelReject) {
	o.OrderID = math.MinInt64
	o.ListID = math.MinInt64
	return
}

func (*OrderCancelReject) SbeBlockLength() (blockLength uint16) {
	return 22
}

func (*OrderCancelReject) SbeTemplateId() (templateId uint16) {
	return 96
}

func (*OrderCancelReject) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancelReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderCancelReject) SbeSemanticType() (semanticType []byte) {
	return []byte("9")
}

func (*OrderCancelReject) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderCancelReject) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelReject) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelReject) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelReject) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelReject) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelReject) ListIDId() uint16 {
	return 66
}

func (*OrderCancelReject) ListIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) ListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ListIDSinceVersion()
}

func (*OrderCancelReject) ListIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject) ListIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject) ListIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelReject) ListIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelReject) ListIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelReject) CancelRestrictionsId() uint16 {
	return 25002
}

func (*OrderCancelReject) CancelRestrictionsSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) CancelRestrictionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelRestrictionsSinceVersion()
}

func (*OrderCancelReject) CancelRestrictionsDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject) CancelRestrictionsMetaAttribute(meta int) string {
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

func (*OrderCancelReject) CxlRejResponseToId() uint16 {
	return 434
}

func (*OrderCancelReject) CxlRejResponseToSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) CxlRejResponseToInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CxlRejResponseToSinceVersion()
}

func (*OrderCancelReject) CxlRejResponseToDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject) CxlRejResponseToMetaAttribute(meta int) string {
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

func (*OrderCancelReject) ErrorCodeId() uint16 {
	return 25016
}

func (*OrderCancelReject) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorCodeSinceVersion()
}

func (*OrderCancelReject) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*OrderCancelReject) ErrorCodeMetaAttribute(meta int) string {
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

func (*OrderCancelReject) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancelReject) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancelReject) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancelReject) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelReject) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelReject) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelReject) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelReject) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderCancelReject) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelReject) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelReject) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelReject) OrigClListIDMetaAttribute(meta int) string {
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

func (*OrderCancelReject) OrigClListIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) OrigClListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClListIDSinceVersion()
}

func (*OrderCancelReject) OrigClListIDDeprecated() uint16 {
	return 0
}

func (OrderCancelReject) OrigClListIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelReject) OrigClListIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelReject) SymbolMetaAttribute(meta int) string {
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

func (*OrderCancelReject) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderCancelReject) SymbolDeprecated() uint16 {
	return 0
}

func (OrderCancelReject) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelReject) SymbolHeaderLength() uint64 {
	return 1
}

func (*OrderCancelReject) ErrorTextMetaAttribute(meta int) string {
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

func (*OrderCancelReject) ErrorTextSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelReject) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorTextSinceVersion()
}

func (*OrderCancelReject) ErrorTextDeprecated() uint16 {
	return 0
}

func (OrderCancelReject) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelReject) ErrorTextHeaderLength() uint64 {
	return 2
}
