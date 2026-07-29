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

type OrderAmendKeepPriorityRequest struct {
	OrderID     int64
	QtyExponent int8
	OrderQty    int64
	ClOrdID     []uint8
	OrigClOrdID []uint8
	Symbol      []uint8
}

func (o *OrderAmendKeepPriorityRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	return nil
}

func (o *OrderAmendKeepPriorityRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderAmendKeepPriorityRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if !utf8.Valid(o.ClOrdID[:]) {
		return errors.New("o.ClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigClOrdID[:]) {
		return errors.New("o.OrigClOrdID failed UTF-8 validation")
	}
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	return nil
}

func OrderAmendKeepPriorityRequestInit(o *OrderAmendKeepPriorityRequest) {
	o.OrderID = math.MinInt64
	return
}

func (*OrderAmendKeepPriorityRequest) SbeBlockLength() (blockLength uint16) {
	return 17
}

func (*OrderAmendKeepPriorityRequest) SbeTemplateId() (templateId uint16) {
	return 105
}

func (*OrderAmendKeepPriorityRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderAmendKeepPriorityRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderAmendKeepPriorityRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("XAK")
}

func (*OrderAmendKeepPriorityRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderAmendKeepPriorityRequest) OrderIDId() uint16 {
	return 37
}

func (*OrderAmendKeepPriorityRequest) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderAmendKeepPriorityRequest) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderAmendKeepPriorityRequest) OrderIDMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) OrderIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderAmendKeepPriorityRequest) OrderIDMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderAmendKeepPriorityRequest) OrderIDNullValue() int64 {
	return math.MinInt64
}

func (*OrderAmendKeepPriorityRequest) QtyExponentId() uint16 {
	return 25055
}

func (*OrderAmendKeepPriorityRequest) QtyExponentSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.QtyExponentSinceVersion()
}

func (*OrderAmendKeepPriorityRequest) QtyExponentDeprecated() uint16 {
	return 0
}

func (*OrderAmendKeepPriorityRequest) QtyExponentMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*OrderAmendKeepPriorityRequest) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*OrderAmendKeepPriorityRequest) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*OrderAmendKeepPriorityRequest) OrderQtyId() uint16 {
	return 38
}

func (*OrderAmendKeepPriorityRequest) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderAmendKeepPriorityRequest) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderAmendKeepPriorityRequest) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) OrderQtyMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderAmendKeepPriorityRequest) OrderQtyMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderAmendKeepPriorityRequest) OrderQtyNullValue() int64 {
	return math.MinInt64
}

func (*OrderAmendKeepPriorityRequest) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderAmendKeepPriorityRequest) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderAmendKeepPriorityRequest) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendKeepPriorityRequest) ClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderAmendKeepPriorityRequest) OrigClOrdIDMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) OrigClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) OrigClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigClOrdIDSinceVersion()
}

func (*OrderAmendKeepPriorityRequest) OrigClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderAmendKeepPriorityRequest) OrigClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendKeepPriorityRequest) OrigClOrdIDHeaderLength() uint64 {
	return 1
}

func (*OrderAmendKeepPriorityRequest) SymbolMetaAttribute(meta int) string {
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

func (*OrderAmendKeepPriorityRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderAmendKeepPriorityRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderAmendKeepPriorityRequest) SymbolDeprecated() uint16 {
	return 0
}

func (OrderAmendKeepPriorityRequest) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderAmendKeepPriorityRequest) SymbolHeaderLength() uint64 {
	return 1
}
