// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type OrderMassCancelRequest struct {
	MassCancelRequestType MassCancelRequestTypeEnum
	Symbol                []uint8
	ClOrdID               []uint8
}

func (o *OrderMassCancelRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := o.MassCancelRequestType.Encode(_m, _w); err != nil {
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
	return nil
}

func (o *OrderMassCancelRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if o.MassCancelRequestTypeInActingVersion(actingVersion) {
		if err := o.MassCancelRequestType.Decode(_m, _r, actingVersion); err != nil {
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
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderMassCancelRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := o.MassCancelRequestType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(o.Symbol[:]) {
		return errors.New("o.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(o.ClOrdID[:]) {
		return errors.New("o.ClOrdID failed UTF-8 validation")
	}
	return nil
}

func OrderMassCancelRequestInit(o *OrderMassCancelRequest) {
	return
}

func (*OrderMassCancelRequest) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*OrderMassCancelRequest) SbeTemplateId() (templateId uint16) {
	return 103
}

func (*OrderMassCancelRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderMassCancelRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*OrderMassCancelRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("q")
}

func (*OrderMassCancelRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderMassCancelRequest) MassCancelRequestTypeId() uint16 {
	return 530
}

func (*OrderMassCancelRequest) MassCancelRequestTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelRequest) MassCancelRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MassCancelRequestTypeSinceVersion()
}

func (*OrderMassCancelRequest) MassCancelRequestTypeDeprecated() uint16 {
	return 0
}

func (*OrderMassCancelRequest) MassCancelRequestTypeMetaAttribute(meta int) string {
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

func (*OrderMassCancelRequest) SymbolMetaAttribute(meta int) string {
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

func (*OrderMassCancelRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderMassCancelRequest) SymbolDeprecated() uint16 {
	return 0
}

func (OrderMassCancelRequest) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (OrderMassCancelRequest) SymbolHeaderLength() uint64 {
	return 1
}

func (*OrderMassCancelRequest) ClOrdIDMetaAttribute(meta int) string {
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

func (*OrderMassCancelRequest) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderMassCancelRequest) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderMassCancelRequest) ClOrdIDDeprecated() uint16 {
	return 0
}

func (OrderMassCancelRequest) ClOrdIDCharacterEncoding() string {
	return "UTF-8"
}

func (OrderMassCancelRequest) ClOrdIDHeaderLength() uint64 {
	return 1
}
