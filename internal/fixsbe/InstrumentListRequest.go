// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type InstrumentListRequest struct {
	InstrumentListRequestType InstrumentListRequestTypeEnum
	InstrumentReqID           []uint8
	Symbol                    []uint8
}

func (i *InstrumentListRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := i.InstrumentListRequestType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(i.InstrumentReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.InstrumentReqID); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(i.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.Symbol); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentListRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if i.InstrumentListRequestTypeInActingVersion(actingVersion) {
		if err := i.InstrumentListRequestType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}

	if i.InstrumentReqIDInActingVersion(actingVersion) {
		var InstrumentReqIDLength uint8
		if err := _m.ReadUint8(_r, &InstrumentReqIDLength); err != nil {
			return err
		}
		if cap(i.InstrumentReqID) < int(InstrumentReqIDLength) {
			i.InstrumentReqID = make([]uint8, InstrumentReqIDLength)
		}
		i.InstrumentReqID = i.InstrumentReqID[:InstrumentReqIDLength]
		if err := _m.ReadBytes(_r, i.InstrumentReqID); err != nil {
			return err
		}
	}

	if i.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(i.Symbol) < int(SymbolLength) {
			i.Symbol = make([]uint8, SymbolLength)
		}
		i.Symbol = i.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, i.Symbol); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *InstrumentListRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := i.InstrumentListRequestType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(i.InstrumentReqID[:]) {
		return errors.New("i.InstrumentReqID failed UTF-8 validation")
	}
	if !utf8.Valid(i.Symbol[:]) {
		return errors.New("i.Symbol failed UTF-8 validation")
	}
	return nil
}

func InstrumentListRequestInit(i *InstrumentListRequest) {
	return
}

func (*InstrumentListRequest) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*InstrumentListRequest) SbeTemplateId() (templateId uint16) {
	return 200
}

func (*InstrumentListRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*InstrumentListRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*InstrumentListRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("x")
}

func (*InstrumentListRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*InstrumentListRequest) InstrumentListRequestTypeId() uint16 {
	return 559
}

func (*InstrumentListRequest) InstrumentListRequestTypeSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequest) InstrumentListRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentListRequestTypeSinceVersion()
}

func (*InstrumentListRequest) InstrumentListRequestTypeDeprecated() uint16 {
	return 0
}

func (*InstrumentListRequest) InstrumentListRequestTypeMetaAttribute(meta int) string {
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

func (*InstrumentListRequest) InstrumentReqIDMetaAttribute(meta int) string {
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

func (*InstrumentListRequest) InstrumentReqIDSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequest) InstrumentReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentReqIDSinceVersion()
}

func (*InstrumentListRequest) InstrumentReqIDDeprecated() uint16 {
	return 0
}

func (InstrumentListRequest) InstrumentReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (InstrumentListRequest) InstrumentReqIDHeaderLength() uint64 {
	return 1
}

func (*InstrumentListRequest) SymbolMetaAttribute(meta int) string {
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

func (*InstrumentListRequest) SymbolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequest) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SymbolSinceVersion()
}

func (*InstrumentListRequest) SymbolDeprecated() uint16 {
	return 0
}

func (InstrumentListRequest) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (InstrumentListRequest) SymbolHeaderLength() uint64 {
	return 1
}
