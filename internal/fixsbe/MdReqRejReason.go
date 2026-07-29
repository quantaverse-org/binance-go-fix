// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MdReqRejReasonEnum byte
type MdReqRejReasonValues struct {
	DuplicateMdReqID     MdReqRejReasonEnum
	TooManySubscriptions MdReqRejReasonEnum
	NonRepresentable     MdReqRejReasonEnum
	NullValue            MdReqRejReasonEnum
}

var MdReqRejReason = MdReqRejReasonValues{49, 50, 126, 0}

func (m MdReqRejReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MdReqRejReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MdReqRejReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MdReqRejReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MdReqRejReason, unknown enumeration value %d", m)
}

func (*MdReqRejReasonEnum) EncodedLength() int64 {
	return 1
}

func (*MdReqRejReasonEnum) DuplicateMdReqIDSinceVersion() uint16 {
	return 0
}

func (m *MdReqRejReasonEnum) DuplicateMdReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DuplicateMdReqIDSinceVersion()
}

func (*MdReqRejReasonEnum) DuplicateMdReqIDDeprecated() uint16 {
	return 0
}

func (*MdReqRejReasonEnum) TooManySubscriptionsSinceVersion() uint16 {
	return 0
}

func (m *MdReqRejReasonEnum) TooManySubscriptionsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TooManySubscriptionsSinceVersion()
}

func (*MdReqRejReasonEnum) TooManySubscriptionsDeprecated() uint16 {
	return 0
}

func (*MdReqRejReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MdReqRejReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MdReqRejReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
