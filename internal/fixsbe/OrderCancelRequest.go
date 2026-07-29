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

type OrderCancelRequest struct {
	OrderID            int64
	ListID             int64
	CancelRestrictions CancelRestrictionsEnum
	ClOrdID            []uint8
	OrigClOrdID        []uint8
	OrigClListID       []uint8
	Symbol             []uint8
}

func (o *OrderCancelRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	return nil
}

func (o *OrderCancelRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderCancelRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func OrderCancelRequestInit(o *OrderCancelRequest) {
	o.OrderID = math.MinInt64
	o.ListID = math.MinInt64
	return
}

func (*OrderCancelRequest) SbeBlockLength() (blockLength uint16) {
	return 17
}

func (*OrderCancelRequest) SbeTemplateId() (templateId uint16) {
	return 101
}

func (*OrderCancelRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancelRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderCancelRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("F")
}

func (*OrderCancelRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderCancelRequest) OrderIDId() uint16 {
	return 37
}

func (*OrderCancelRequest) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderCancelRequest) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) OrderIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequest) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequest) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequest) ListIDId() uint16 {
	return 66
}

func (*OrderCancelRequest) ListIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) ListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ListIDSinceVersion()
}

func (*OrderCancelRequest) ListIDDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) ListIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) ListIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderCancelRequest) ListIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderCancelRequest) ListIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderCancelRequest) CancelRestrictionsId() uint16 {
	return 25002
}

func (*OrderCancelRequest) CancelRestrictionsSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) CancelRestrictionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelRestrictionsSinceVersion()
}

func (*OrderCancelRequest) CancelRestrictionsDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequest) CancelRestrictionsMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderCancelRequest) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequest) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequest) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequest) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderCancelRequest) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequest) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequest) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequest) OrigClListIDMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) OrigClListIDSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) OrigClListIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClListIDSinceVersion()
}

func (*OrderCancelRequest) OrigClListIDDeprecated() uint16 {
	return 0
}

func (OrderCancelRequest) OrigClListIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequest) OrigClListIDHeaderLength() uint64 {
	return 1
}

func (*OrderCancelRequest) SymbolMetaAttribute(meta int) string {
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

func (*OrderCancelRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderCancelRequest) SymbolDeprecated() uint16 {
	return 0
}

func (OrderCancelRequest) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancelRequest) SymbolHeaderLength() uint64 {
	return 1
}
