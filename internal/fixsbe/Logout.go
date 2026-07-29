// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type Logout struct {
	Text []uint8
}

func (l *Logout) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, uint16(len(l.Text))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.Text); err != nil {
		return err
	}
	return nil
}

func (l *Logout) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.TextInActingVersion(actingVersion) {
		var TextLength uint16
		if err := _m.ReadUint16(_r, &TextLength); err != nil {
			return err
		}
		if cap(l.Text) < int(TextLength) {
			l.Text = make([]uint8, TextLength)
		}
		l.Text = l.Text[:TextLength]
		if err := _m.ReadBytes(_r, l.Text); err != nil {
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

func (l *Logout) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(l.Text[:]) {
		return errors.New("l.Text failed UTF-8 validation")
	}
	return nil
}

func LogoutInit(l *Logout) {
	return
}

func (*Logout) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*Logout) SbeTemplateId() (templateId uint16) {
	return 20004
}

func (*Logout) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Logout) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Logout) SbeSemanticType() (semanticType []byte) {
	return []byte("5")
}

func (*Logout) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Logout) TextMetaAttribute(meta int) string {
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

func (*Logout) TextSinceVersion() uint16 {
	return 0
}

func (l *Logout) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TextSinceVersion()
}

func (*Logout) TextDeprecated() uint16 {
	return 0
}

func (Logout) TextCharacterEncoding() string {
	return "UTF-8"
}

func (Logout) TextHeaderLength() uint64 {
	return 2
}
