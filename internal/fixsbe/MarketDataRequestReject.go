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

type MarketDataRequestReject struct {
	MDReqRejReason MdReqRejReasonEnum
	ErrorCode      int32
	MDReqID        []uint8
	Text           []uint8
}

func (m *MarketDataRequestReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := m.MDReqRejReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, m.ErrorCode); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(m.MDReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.MDReqID); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(m.Text))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Text); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataRequestReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if m.MDReqRejReasonInActingVersion(actingVersion) {
		if err := m.MDReqRejReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.ErrorCodeInActingVersion(actingVersion) {
		m.ErrorCode = m.ErrorCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.ErrorCode); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.MDReqIDInActingVersion(actingVersion) {
		var MDReqIDLength uint8
		if err := _m.ReadUint8(_r, &MDReqIDLength); err != nil {
			return err
		}
		if cap(m.MDReqID) < int(MDReqIDLength) {
			m.MDReqID = make([]uint8, MDReqIDLength)
		}
		m.MDReqID = m.MDReqID[:MDReqIDLength]
		if err := _m.ReadBytes(_r, m.MDReqID); err != nil {
			return err
		}
	}

	if m.TextInActingVersion(actingVersion) {
		var TextLength uint16
		if err := _m.ReadUint16(_r, &TextLength); err != nil {
			return err
		}
		if cap(m.Text) < int(TextLength) {
			m.Text = make([]uint8, TextLength)
		}
		m.Text = m.Text[:TextLength]
		if err := _m.ReadBytes(_r, m.Text); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataRequestReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := m.MDReqRejReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.ErrorCodeInActingVersion(actingVersion) {
		if m.ErrorCode != m.ErrorCodeNullValue() && (m.ErrorCode < m.ErrorCodeMinValue() || m.ErrorCode > m.ErrorCodeMaxValue()) {
			return fmt.Errorf("Range check failed on m.ErrorCode (%v < %v > %v)", m.ErrorCodeMinValue(), m.ErrorCode, m.ErrorCodeMaxValue())
		}
	}
	if !utf8.Valid(m.MDReqID[:]) {
		return errors.New("m.MDReqID failed UTF-8 validation")
	}
	if !utf8.Valid(m.Text[:]) {
		return errors.New("m.Text failed UTF-8 validation")
	}
	return nil
}

func MarketDataRequestRejectInit(m *MarketDataRequestReject) {
	m.ErrorCode = math.MinInt32
	return
}

func (*MarketDataRequestReject) SbeBlockLength() (blockLength uint16) {
	return 5
}

func (*MarketDataRequestReject) SbeTemplateId() (templateId uint16) {
	return 203
}

func (*MarketDataRequestReject) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataRequestReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataRequestReject) SbeSemanticType() (semanticType []byte) {
	return []byte("Y")
}

func (*MarketDataRequestReject) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataRequestReject) MDReqRejReasonId() uint16 {
	return 281
}

func (*MarketDataRequestReject) MDReqRejReasonSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestReject) MDReqRejReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDReqRejReasonSinceVersion()
}

func (*MarketDataRequestReject) MDReqRejReasonDeprecated() uint16 {
	return 0
}

func (*MarketDataRequestReject) MDReqRejReasonMetaAttribute(meta int) string {
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

func (*MarketDataRequestReject) ErrorCodeId() uint16 {
	return 25016
}

func (*MarketDataRequestReject) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestReject) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ErrorCodeSinceVersion()
}

func (*MarketDataRequestReject) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*MarketDataRequestReject) ErrorCodeMetaAttribute(meta int) string {
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

func (*MarketDataRequestReject) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MarketDataRequestReject) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*MarketDataRequestReject) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*MarketDataRequestReject) MDReqIDMetaAttribute(meta int) string {
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

func (*MarketDataRequestReject) MDReqIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestReject) MDReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDReqIDSinceVersion()
}

func (*MarketDataRequestReject) MDReqIDDeprecated() uint16 {
	return 0
}

func (MarketDataRequestReject) MDReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataRequestReject) MDReqIDHeaderLength() uint64 {
	return 1
}

func (*MarketDataRequestReject) TextMetaAttribute(meta int) string {
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

func (*MarketDataRequestReject) TextSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestReject) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TextSinceVersion()
}

func (*MarketDataRequestReject) TextDeprecated() uint16 {
	return 0
}

func (MarketDataRequestReject) TextCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataRequestReject) TextHeaderLength() uint64 {
	return 2
}
