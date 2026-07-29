// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type TestRequest struct {
	TestReqID []uint8
}

func (t *TestRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(t.TestReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.TestReqID); err != nil {
		return err
	}
	return nil
}

func (t *TestRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.TestReqIDInActingVersion(actingVersion) {
		var TestReqIDLength uint8
		if err := _m.ReadUint8(_r, &TestReqIDLength); err != nil {
			return err
		}
		if cap(t.TestReqID) < int(TestReqIDLength) {
			t.TestReqID = make([]uint8, TestReqIDLength)
		}
		t.TestReqID = t.TestReqID[:TestReqIDLength]
		if err := _m.ReadBytes(_r, t.TestReqID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(t.TestReqID[:]) {
		return errors.New("t.TestReqID failed UTF-8 validation")
	}
	return nil
}

func TestRequestInit(t *TestRequest) {
	return
}

func (*TestRequest) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*TestRequest) SbeTemplateId() (templateId uint16) {
	return 20002
}

func (*TestRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TestRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*TestRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("1")
}

func (*TestRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*TestRequest) TestReqIDMetaAttribute(meta int) string {
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

func (*TestRequest) TestReqIDSinceVersion() uint16 {
	return 0
}

func (t *TestRequest) TestReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TestReqIDSinceVersion()
}

func (*TestRequest) TestReqIDDeprecated() uint16 {
	return 0
}

func (TestRequest) TestReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (TestRequest) TestReqIDHeaderLength() uint64 {
	return 1
}
