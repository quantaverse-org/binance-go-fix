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

type OrderMassCancelReport struct {
	MassCancelRequestType  MassCancelRequestTypeEnum
	MassCancelResponse     MassCancelResponseEnum
	MassCancelRejectReason MassCancelRejectReasonEnum
	TotalAffectedOrders    int64
	ErrorCode              int32
	Symbol                 []uint8
	ClOrdID                []uint8
	ErrorText              []uint8
}

func (o *OrderMassCancelReport) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.MassCancelRequestType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.MassCancelResponse.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.MassCancelRejectReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.TotalAffectedOrders); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.ErrorCode); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(o.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID); err != nil {
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

func (o *OrderMassCancelReport) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.MassCancelRequestTypeInActingVersion(actingVersion) {
		if err := o.MassCancelRequestType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.MassCancelResponseInActingVersion(actingVersion) {
		if err := o.MassCancelResponse.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.MassCancelRejectReasonInActingVersion(actingVersion) {
		if err := o.MassCancelRejectReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.TotalAffectedOrdersInActingVersion(actingVersion) {
		o.TotalAffectedOrders = o.TotalAffectedOrdersNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.TotalAffectedOrders); err != nil {
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

func (o *OrderMassCancelReport) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := o.MassCancelRequestType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.MassCancelResponse.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.MassCancelRejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TotalAffectedOrdersInActingVersion(actingVersion) {
		if o.TotalAffectedOrders != o.TotalAffectedOrdersNullValue() && (o.TotalAffectedOrders < o.TotalAffectedOrdersMinValue() || o.TotalAffectedOrders > o.TotalAffectedOrdersMaxValue()) {
			return fmt.Errorf("Range check failed on o.TotalAffectedOrders (%v < %v > %v)", o.TotalAffectedOrdersMinValue(), o.TotalAffectedOrders, o.TotalAffectedOrdersMaxValue())
		}
	}
	if o.ErrorCodeInActingVersion(actingVersion) {
		if o.ErrorCode != o.ErrorCodeNullValue() && (o.ErrorCode < o.ErrorCodeMinValue() || o.ErrorCode > o.ErrorCodeMaxValue()) {
			return fmt.Errorf("Range check failed on o.ErrorCode (%v < %v > %v)", o.ErrorCodeMinValue(), o.ErrorCode, o.ErrorCodeMaxValue())
		}
	}
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(o.ClOrdID[:]) {
		return errors.New("o.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.ErrorText[:]) {
		return errors.New("o.ErrorText failed UTF-8 validation")
	}
	return nil
}

func OrderMassCancelReportInit(o *OrderMassCancelReport) {
	o.TotalAffectedOrders = math.MinInt64
	o.ErrorCode = math.MinInt32
	return
}

func (*OrderMassCancelReport) SbeBlockLength() (blockLength uint16) {
	return 15
}

func (*OrderMassCancelReport) SbeTemplateId() (templateId uint16) {
	return 104
}

func (*OrderMassCancelReport) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderMassCancelReport) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderMassCancelReport) SbeSemanticType() (semanticType []byte) {
	return []byte("r")
}

func (*OrderMassCancelReport) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderMassCancelReport) MassCancelRequestTypeId() uint16 {
	return 530
}

func (*OrderMassCancelReport) MassCancelRequestTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) MassCancelRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelRequestTypeSinceVersion()
}

func (*OrderMassCancelReport) MassCancelRequestTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelReport) MassCancelRequestTypeMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) MassCancelResponseId() uint16 {
	return 531
}

func (*OrderMassCancelReport) MassCancelResponseSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) MassCancelResponseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelResponseSinceVersion()
}

func (*OrderMassCancelReport) MassCancelResponseDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelReport) MassCancelResponseMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) MassCancelRejectReasonId() uint16 {
	return 532
}

func (*OrderMassCancelReport) MassCancelRejectReasonSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) MassCancelRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelRejectReasonSinceVersion()
}

func (*OrderMassCancelReport) MassCancelRejectReasonDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelReport) MassCancelRejectReasonMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) TotalAffectedOrdersId() uint16 {
	return 533
}

func (*OrderMassCancelReport) TotalAffectedOrdersSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) TotalAffectedOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TotalAffectedOrdersSinceVersion()
}

func (*OrderMassCancelReport) TotalAffectedOrdersDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelReport) TotalAffectedOrdersMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) TotalAffectedOrdersMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderMassCancelReport) TotalAffectedOrdersMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderMassCancelReport) TotalAffectedOrdersNullValue() int64 {
	return math.MinInt64
}

func (*OrderMassCancelReport) ErrorCodeId() uint16 {
	return 25016
}

func (*OrderMassCancelReport) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorCodeSinceVersion()
}

func (*OrderMassCancelReport) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelReport) ErrorCodeMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderMassCancelReport) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderMassCancelReport) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*OrderMassCancelReport) SymbolMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderMassCancelReport) SymbolDeprecated() uint16 {
	return 0
}

func (OrderMassCancelReport) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderMassCancelReport) SymbolHeaderLength() uint64 {
	return 1
}

func (*OrderMassCancelReport) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderMassCancelReport) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderMassCancelReport) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderMassCancelReport) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderMassCancelReport) ErrorTextMetaAttribute(meta int) string {
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

func (*OrderMassCancelReport) ErrorTextSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelReport) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ErrorTextSinceVersion()
}

func (*OrderMassCancelReport) ErrorTextDeprecated() uint16 {
	return 0
}

func (OrderMassCancelReport) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (OrderMassCancelReport) ErrorTextHeaderLength() uint64 {
	return 2
}
