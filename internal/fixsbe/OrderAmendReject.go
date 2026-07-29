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

type OrderAmendReject struct {
	OrderID     int64
	QtyExponent int8
	OrderQty    int64
	ErrorCode   int32
	ClOrdID     []uint8
	OrigClOrdID []uint8
	Symbol      []uint8
	ErrorText   []uint8
}

func (o *OrderAmendReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.OrderID); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, o.QtyExponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.OrderQty); err != nil {
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

func (o *OrderAmendReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIDInActingVersion(actingVersion) {
		o.OrderID = o.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderID); err != nil {
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

func (o *OrderAmendReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIDInActingVersion(actingVersion) {
		if o.OrderID != o.OrderIDNullValue() && (o.OrderID < o.OrderIDMinValue() || o.OrderID > o.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on o.OrderID (%v < %v > %v)", o.OrderIDMinValue(), o.OrderID, o.OrderIDMaxValue())
		}
	}
	if o.QtyExponentInActingVersion(actingVersion) {
		if o.QtyExponent < o.QtyExponentMinValue() || o.QtyExponent > o.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on o.QtyExponent (%v < %v > %v)", o.QtyExponentMinValue(), o.QtyExponent, o.QtyExponentMaxValue())
		}
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
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
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(o.ErrorText[:]) {
		return errors.New("o.ErrorText failed UTF-8 validation")
	}
	return nil
}

func OrderAmendRejectInit(o *OrderAmendReject) {
	o.OrderID = math.MinInt64
	return
}

func (*OrderAmendReject) SbeBlockLength() (blockLength uint16) {
	return 21
}

func (*OrderAmendReject) SbeTemplateId() (templateId uint16) {
	return 106
}

func (*OrderAmendReject) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderAmendReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderAmendReject) SbeSemanticType() (semanticType []byte) {
	return []byte("XAR")
}

func (*OrderAmendReject) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderAmendReject) OrderIDId() uint16 {
	return 37
}

func (*OrderAmendReject) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderAmendReject) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderAmendReject) OrderIDMetaAttribute(meta int) string {
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

func (*OrderAmendReject) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderAmendReject) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderAmendReject) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderAmendReject) QtyExponentId() uint16 {
	return 25055
}

func (*OrderAmendReject) QtyExponentSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QtyExponentSinceVersion()
}

func (*OrderAmendReject) QtyExponentDeprecated() uint16 {
	return 0
}

func (*OrderAmendReject) QtyExponentMetaAttribute(meta int) string {
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

func (*OrderAmendReject) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*OrderAmendReject) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*OrderAmendReject) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*OrderAmendReject) OrderQtyId() uint16 {
	return 38
}

func (*OrderAmendReject) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderAmendReject) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderAmendReject) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderAmendReject) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderAmendReject) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderAmendReject) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*OrderAmendReject) ErrorCodeId() uint16 {
	return 25016
}

func (*OrderAmendReject) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorCodeSinceVersion()
}

func (*OrderAmendReject) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*OrderAmendReject) ErrorCodeMetaAttribute(meta int) string {
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

func (*OrderAmendReject) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderAmendReject) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderAmendReject) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*OrderAmendReject) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderAmendReject) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderAmendReject) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderAmendReject) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendReject) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderAmendReject) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderAmendReject) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderAmendReject) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderAmendReject) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendReject) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderAmendReject) SymbolMetaAttribute(meta int) string {
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

func (*OrderAmendReject) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderAmendReject) SymbolDeprecated() uint16 {
	return 0
}

func (OrderAmendReject) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendReject) SymbolHeaderLength() uint64 {
	return 1
}

func (*OrderAmendReject) ErrorTextMetaAttribute(meta int) string {
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

func (*OrderAmendReject) ErrorTextSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendReject) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorTextSinceVersion()
}

func (*OrderAmendReject) ErrorTextDeprecated() uint16 {
	return 0
}

func (OrderAmendReject) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendReject) ErrorTextHeaderLength() uint64 {
	return 2
}
