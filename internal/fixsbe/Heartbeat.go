// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type Heartbeat struct {
	TestReqID []uint8
}

func (h *Heartbeat) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := h.RangeCheck(h.SbeSchemaVersion(), h.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(h.TestReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, h.TestReqID); err != nil {
		return err
	}
	return nil
}

func (h *Heartbeat) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > h.SbeSchemaVersion() && blockLength > h.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-h.SbeBlockLength()))
	}

	if h.TestReqIDInActingVersion(actingVersion) {
		var TestReqIDLength uint8
		if err := _m.ReadUint8(_r, &TestReqIDLength); err != nil {
			return err
		}
		if cap(h.TestReqID) < int(TestReqIDLength) {
			h.TestReqID = make([]uint8, TestReqIDLength)
		}
		h.TestReqID = h.TestReqID[:TestReqIDLength]
		if err := _m.ReadBytes(_r, h.TestReqID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := h.RangeCheck(actingVersion, h.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (h *Heartbeat) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(h.TestReqID[:]) {
		return errors.New("h.TestReqID failed UTF-8 validation")
	}
	return nil
}

func HeartbeatInit(h *Heartbeat) {
	return
}

func (*Heartbeat) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*Heartbeat) SbeTemplateId() (templateId uint16) {
	return 20001
}

func (*Heartbeat) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Heartbeat) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Heartbeat) SbeSemanticType() (semanticType []byte) {
	return []byte("0")
}

func (*Heartbeat) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Heartbeat) TestReqIDMetaAttribute(meta int) string {
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

func (*Heartbeat) TestReqIDSinceVersion() uint16 {
	return 0
}

func (h *Heartbeat) TestReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.TestReqIDSinceVersion()
}

func (*Heartbeat) TestReqIDDeprecated() uint16 {
	return 0
}

func (Heartbeat) TestReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (Heartbeat) TestReqIDHeaderLength() uint64 {
	return 1
}
