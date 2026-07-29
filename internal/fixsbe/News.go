// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type News struct {
	Headline []uint8
}

func (n *News) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, uint16(len(n.Headline))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Headline); err != nil {
		return err
	}
	return nil
}

func (n *News) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}

	if n.HeadlineInActingVersion(actingVersion) {
		var HeadlineLength uint16
		if err := _m.ReadUint16(_r, &HeadlineLength); err != nil {
			return err
		}
		if cap(n.Headline) < int(HeadlineLength) {
			n.Headline = make([]uint8, HeadlineLength)
		}
		n.Headline = n.Headline[:HeadlineLength]
		if err := _m.ReadBytes(_r, n.Headline); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *News) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(n.Headline[:]) {
		return errors.New("n.Headline failed UTF-8 validation")
	}
	return nil
}

func NewsInit(n *News) {
	return
}

func (*News) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*News) SbeTemplateId() (templateId uint16) {
	return 20100
}

func (*News) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*News) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*News) SbeSemanticType() (semanticType []byte) {
	return []byte("B")
}

func (*News) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*News) HeadlineMetaAttribute(meta int) string {
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

func (*News) HeadlineSinceVersion() uint16 {
	return 0
}

func (n *News) HeadlineInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.HeadlineSinceVersion()
}

func (*News) HeadlineDeprecated() uint16 {
	return 0
}

func (News) HeadlineCharacterEncoding() string {
	return "UTF-8"
}

func (News) HeadlineHeaderLength() uint64 {
	return 2
}
