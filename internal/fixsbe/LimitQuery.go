// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type LimitQuery struct {
	ReqID []uint8
}

func (l *LimitQuery) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(l.ReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ReqID); err != nil {
		return err
	}
	return nil
}

func (l *LimitQuery) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.ReqIDInActingVersion(actingVersion) {
		var ReqIDLength uint8
		if err := _m.ReadUint8(_r, &ReqIDLength); err != nil {
			return err
		}
		if cap(l.ReqID) < int(ReqIDLength) {
			l.ReqID = make([]uint8, ReqIDLength)
		}
		l.ReqID = l.ReqID[:ReqIDLength]
		if err := _m.ReadBytes(_r, l.ReqID); err != nil {
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

func (l *LimitQuery) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(l.ReqID[:]) {
		return errors.New("l.ReqID failed UTF-8 validation")
	}
	return nil
}

func LimitQueryInit(l *LimitQuery) {
	return
}

func (*LimitQuery) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*LimitQuery) SbeTemplateId() (templateId uint16) {
	return 120
}

func (*LimitQuery) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LimitQuery) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*LimitQuery) SbeSemanticType() (semanticType []byte) {
	return []byte("XLQ")
}

func (*LimitQuery) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LimitQuery) ReqIDMetaAttribute(meta int) string {
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

func (*LimitQuery) ReqIDSinceVersion() uint16 {
	return 0
}

func (l *LimitQuery) ReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ReqIDSinceVersion()
}

func (*LimitQuery) ReqIDDeprecated() uint16 {
	return 0
}

func (LimitQuery) ReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (LimitQuery) ReqIDHeaderLength() uint64 {
	return 1
}
