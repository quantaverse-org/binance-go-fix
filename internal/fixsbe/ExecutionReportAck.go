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

type ExecutionReportAck struct {
	OrderID      int64
	ListID       int64
	TransactTime int64
	ExecType     ExecTypeEnum
	OrdStatus    OrdStatusEnum
	OrdRejReason OrdRejReasonEnum
	ErrorCode    int32
	ClOrdID      []uint8
	Symbol       []uint8
	ErrorText    []uint8
}

func (e *ExecutionReportAck) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := e.RangeCheck(e.SbeSchemaVersion(), e.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, e.OrderID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.ListID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, e.TransactTime); err != nil {
		return err
	}
	if err := e.ExecType.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.OrdStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := e.OrdRejReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, e.ErrorCode); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.ClOrdID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.ClOrdID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(e.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, e.Symbol); err != nil {
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

func (e *ExecutionReportAck) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !e.OrderIDInActingVersion(actingVersion) {
		e.OrderID = e.OrderIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.OrderID); err != nil {
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
	if !e.TransactTimeInActingVersion(actingVersion) {
		e.TransactTime = e.TransactTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &e.TransactTime); err != nil {
			return err
		}
	}
	if e.ExecTypeInActingVersion(actingVersion) {
		if err := e.ExecType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if e.OrdStatusInActingVersion(actingVersion) {
		if err := e.OrdStatus.Decode(_m, _r, actingVersion); err != nil {
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
	if actingVersion > e.SbeSchemaVersion() && blockLength > e.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-e.SbeBlockLength()))
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

func (e *ExecutionReportAck) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if e.OrderIDInActingVersion(actingVersion) {
		if e.OrderID != e.OrderIDNullValue() && (e.OrderID < e.OrderIDMinValue() || e.OrderID > e.OrderIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.OrderID (%v < %v > %v)", e.OrderIDMinValue(), e.OrderID, e.OrderIDMaxValue())
		}
	}
	if e.ListIDInActingVersion(actingVersion) {
		if e.ListID != e.ListIDNullValue() && (e.ListID < e.ListIDMinValue() || e.ListID > e.ListIDMaxValue()) {
			return fmt.Errorf("Range check failed on e.ListID (%v < %v > %v)", e.ListIDMinValue(), e.ListID, e.ListIDMaxValue())
		}
	}
	if e.TransactTimeInActingVersion(actingVersion) {
		if e.TransactTime != e.TransactTimeNullValue() && (e.TransactTime < e.TransactTimeMinValue() || e.TransactTime > e.TransactTimeMaxValue()) {
			return fmt.Errorf("Range check failed on e.TransactTime (%v < %v > %v)", e.TransactTimeMinValue(), e.TransactTime, e.TransactTimeMaxValue())
		}
	}
	if err := e.ExecType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := e.OrdStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
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
	if !utf8.Valid(e.ClOrdID[:]) {
		return errors.New("e.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(e.Symbol[:]) {
		return errors.New("e.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(e.ErrorText[:]) {
		return errors.New("e.ErrorText failed UTF-8 validation")
	}
	return nil
}

func ExecutionReportAckInit(e *ExecutionReportAck) {
	e.OrderID = math.MinInt64
	e.ListID = math.MinInt64
	e.TransactTime = math.MinInt64
	e.ErrorCode = math.MinInt32
	return
}

func (*ExecutionReportAck) SbeBlockLength() (blockLength uint16) {
	return 31
}

func (*ExecutionReportAck) SbeTemplateId() (templateId uint16) {
	return 198
}

func (*ExecutionReportAck) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ExecutionReportAck) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*ExecutionReportAck) SbeSemanticType() (semanticType []byte) {
	return []byte("8")
}

func (*ExecutionReportAck) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*ExecutionReportAck) OrderIDId() uint16 {
	return 37
}

func (*ExecutionReportAck) OrderIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrderIDSinceVersion()
}

func (*ExecutionReportAck) OrderIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) OrderIDMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReportAck) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReportAck) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReportAck) ListIDId() uint16 {
	return 66
}

func (*ExecutionReportAck) ListIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) ListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ListIDSinceVersion()
}

func (*ExecutionReportAck) ListIDDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) ListIDMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) ListIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReportAck) ListIDMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReportAck) ListIDNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReportAck) TransactTimeId() uint16 {
	return 60
}

func (*ExecutionReportAck) TransactTimeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TransactTimeSinceVersion()
}

func (*ExecutionReportAck) TransactTimeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) TransactTimeMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) TransactTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ExecutionReportAck) TransactTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ExecutionReportAck) TransactTimeNullValue() int64 {
	return math.MinInt64
}

func (*ExecutionReportAck) ExecTypeId() uint16 {
	return 150
}

func (*ExecutionReportAck) ExecTypeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) ExecTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecTypeSinceVersion()
}

func (*ExecutionReportAck) ExecTypeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) ExecTypeMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) OrdStatusId() uint16 {
	return 39
}

func (*ExecutionReportAck) OrdStatusSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdStatusSinceVersion()
}

func (*ExecutionReportAck) OrdStatusDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) OrdStatusMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) OrdRejReasonId() uint16 {
	return 103
}

func (*ExecutionReportAck) OrdRejReasonSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) OrdRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OrdRejReasonSinceVersion()
}

func (*ExecutionReportAck) OrdRejReasonDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) OrdRejReasonMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) ErrorCodeId() uint16 {
	return 25016
}

func (*ExecutionReportAck) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorCodeSinceVersion()
}

func (*ExecutionReportAck) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*ExecutionReportAck) ErrorCodeMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ExecutionReportAck) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*ExecutionReportAck) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*ExecutionReportAck) ClOrdIDMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ClOrdIDSinceVersion()
}

func (*ExecutionReportAck) ClOrdIDDeprecated() uint16 {
	return 0
}

func (ExecutionReportAck) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReportAck) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*ExecutionReportAck) SymbolMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) SymbolSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.SymbolSinceVersion()
}

func (*ExecutionReportAck) SymbolDeprecated() uint16 {
	return 0
}

func (ExecutionReportAck) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReportAck) SymbolHeaderLength() uint64 {
	return 1
}

func (*ExecutionReportAck) ErrorTextMetaAttribute(meta int) string {
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

func (*ExecutionReportAck) ErrorTextSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportAck) ErrorTextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ErrorTextSinceVersion()
}

func (*ExecutionReportAck) ErrorTextDeprecated() uint16 {
	return 0
}

func (ExecutionReportAck) ErrorTextCharacterEncoding() string {
	return "UTF-8"
}

func (ExecutionReportAck) ErrorTextHeaderLength() uint64 {
	return 2
}
