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

type Reject struct {
	RefSeqNum           uint32
	RefTagID            uint32
	SessionRejectReason SessionRejectReasonEnum
	ErrorCode           int32
	RefMsgType          []uint8
	Text                []uint8
}

func (r *Reject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := r.RangeCheck(r.SbeSchemaVersion(), r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint32(_w, r.RefSeqNum); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, r.RefTagID); err != nil {
		return err
	}
	if err := r.SessionRejectReason.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, r.ErrorCode); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(r.RefMsgType))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.RefMsgType); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(r.Text))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, r.Text); err != nil {
		return err
	}
	return nil
}

func (r *Reject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !r.RefSeqNumInActingVersion(actingVersion) {
		r.RefSeqNum = r.RefSeqNumNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.RefSeqNum); err != nil {
			return err
		}
	}
	if !r.RefTagIDInActingVersion(actingVersion) {
		r.RefTagID = r.RefTagIDNullValue()
	} else {
		if err := _m.ReadUint32(_r, &r.RefTagID); err != nil {
			return err
		}
	}
	if r.SessionRejectReasonInActingVersion(actingVersion) {
		if err := r.SessionRejectReason.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !r.ErrorCodeInActingVersion(actingVersion) {
		r.ErrorCode = r.ErrorCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &r.ErrorCode); err != nil {
			return err
		}
	}
	if actingVersion > r.SbeSchemaVersion() && blockLength > r.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-r.SbeBlockLength()))
	}

	if r.RefMsgTypeInActingVersion(actingVersion) {
		var RefMsgTypeLength uint8
		if err := _m.ReadUint8(_r, &RefMsgTypeLength); err != nil {
			return err
		}
		if cap(r.RefMsgType) < int(RefMsgTypeLength) {
			r.RefMsgType = make([]uint8, RefMsgTypeLength)
		}
		r.RefMsgType = r.RefMsgType[:RefMsgTypeLength]
		if err := _m.ReadBytes(_r, r.RefMsgType); err != nil {
			return err
		}
	}

	if r.TextInActingVersion(actingVersion) {
		var TextLength uint16
		if err := _m.ReadUint16(_r, &TextLength); err != nil {
			return err
		}
		if cap(r.Text) < int(TextLength) {
			r.Text = make([]uint8, TextLength)
		}
		r.Text = r.Text[:TextLength]
		if err := _m.ReadBytes(_r, r.Text); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := r.RangeCheck(actingVersion, r.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if r.RefSeqNumInActingVersion(actingVersion) {
		if r.RefSeqNum != r.RefSeqNumNullValue() && (r.RefSeqNum < r.RefSeqNumMinValue() || r.RefSeqNum > r.RefSeqNumMaxValue()) {
			return fmt.Errorf("Range check failed on r.RefSeqNum (%v < %v > %v)", r.RefSeqNumMinValue(), r.RefSeqNum, r.RefSeqNumMaxValue())
		}
	}
	if r.RefTagIDInActingVersion(actingVersion) {
		if r.RefTagID != r.RefTagIDNullValue() && (r.RefTagID < r.RefTagIDMinValue() || r.RefTagID > r.RefTagIDMaxValue()) {
			return fmt.Errorf("Range check failed on r.RefTagID (%v < %v > %v)", r.RefTagIDMinValue(), r.RefTagID, r.RefTagIDMaxValue())
		}
	}
	if err := r.SessionRejectReason.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if r.ErrorCodeInActingVersion(actingVersion) {
		if r.ErrorCode != r.ErrorCodeNullValue() && (r.ErrorCode < r.ErrorCodeMinValue() || r.ErrorCode > r.ErrorCodeMaxValue()) {
			return fmt.Errorf("Range check failed on r.ErrorCode (%v < %v > %v)", r.ErrorCodeMinValue(), r.ErrorCode, r.ErrorCodeMaxValue())
		}
	}
	if !utf8.Valid(r.RefMsgType[:]) {
		return errors.New("r.RefMsgType failed UTF-8 validation")
	}
	if !utf8.Valid(r.Text[:]) {
		return errors.New("r.Text failed UTF-8 validation")
	}
	return nil
}

func RejectInit(r *Reject) {
	r.RefSeqNum = math.MaxUint32
	r.RefTagID = math.MaxUint32
	r.ErrorCode = math.MinInt32
	return
}

func (*Reject) SbeBlockLength() (blockLength uint16) {
	return 13
}

func (*Reject) SbeTemplateId() (templateId uint16) {
	return 20003
}

func (*Reject) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Reject) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Reject) SbeSemanticType() (semanticType []byte) {
	return []byte("3")
}

func (*Reject) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Reject) RefSeqNumId() uint16 {
	return 45
}

func (*Reject) RefSeqNumSinceVersion() uint16 {
	return 0
}

func (r *Reject) RefSeqNumInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RefSeqNumSinceVersion()
}

func (*Reject) RefSeqNumDeprecated() uint16 {
	return 0
}

func (*Reject) RefSeqNumMetaAttribute(meta int) string {
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

func (*Reject) RefSeqNumMinValue() uint32 {
	return 0
}

func (*Reject) RefSeqNumMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Reject) RefSeqNumNullValue() uint32 {
	return math.MaxUint32
}

func (*Reject) RefTagIDId() uint16 {
	return 371
}

func (*Reject) RefTagIDSinceVersion() uint16 {
	return 0
}

func (r *Reject) RefTagIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RefTagIDSinceVersion()
}

func (*Reject) RefTagIDDeprecated() uint16 {
	return 0
}

func (*Reject) RefTagIDMetaAttribute(meta int) string {
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

func (*Reject) RefTagIDMinValue() uint32 {
	return 0
}

func (*Reject) RefTagIDMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Reject) RefTagIDNullValue() uint32 {
	return math.MaxUint32
}

func (*Reject) SessionRejectReasonId() uint16 {
	return 373
}

func (*Reject) SessionRejectReasonSinceVersion() uint16 {
	return 0
}

func (r *Reject) SessionRejectReasonInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.SessionRejectReasonSinceVersion()
}

func (*Reject) SessionRejectReasonDeprecated() uint16 {
	return 0
}

func (*Reject) SessionRejectReasonMetaAttribute(meta int) string {
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

func (*Reject) ErrorCodeId() uint16 {
	return 25016
}

func (*Reject) ErrorCodeSinceVersion() uint16 {
	return 0
}

func (r *Reject) ErrorCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.ErrorCodeSinceVersion()
}

func (*Reject) ErrorCodeDeprecated() uint16 {
	return 0
}

func (*Reject) ErrorCodeMetaAttribute(meta int) string {
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

func (*Reject) ErrorCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Reject) ErrorCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*Reject) ErrorCodeNullValue() int32 {
	return math.MinInt32
}

func (*Reject) RefMsgTypeMetaAttribute(meta int) string {
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

func (*Reject) RefMsgTypeSinceVersion() uint16 {
	return 0
}

func (r *Reject) RefMsgTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.RefMsgTypeSinceVersion()
}

func (*Reject) RefMsgTypeDeprecated() uint16 {
	return 0
}

func (Reject) RefMsgTypeCharacterEncoding() string {
	return "UTF-8"
}

func (Reject) RefMsgTypeHeaderLength() uint64 {
	return 1
}

func (*Reject) TextMetaAttribute(meta int) string {
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

func (*Reject) TextSinceVersion() uint16 {
	return 0
}

func (r *Reject) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.TextSinceVersion()
}

func (*Reject) TextDeprecated() uint16 {
	return 0
}

func (Reject) TextCharacterEncoding() string {
	return "UTF-8"
}

func (Reject) TextHeaderLength() uint64 {
	return 2
}
