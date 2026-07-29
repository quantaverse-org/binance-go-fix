// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MassCancelRejectReasonEnum uint8
type MassCancelRejectReasonValues struct {
	Other            MassCancelRejectReasonEnum
	NonRepresentable MassCancelRejectReasonEnum
	NullValue        MassCancelRejectReasonEnum
}

var MassCancelRejectReason = MassCancelRejectReasonValues{99, 254, 255}

func (m MassCancelRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassCancelRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassCancelRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassCancelRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassCancelRejectReason, unknown enumeration value %d", m)
}

func (*MassCancelRejectReasonEnum) EncodedLength() int64 {
	return 1
}

func (*MassCancelRejectReasonEnum) OtherSinceVersion() uint16 {
	return 0
}

func (m *MassCancelRejectReasonEnum) OtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OtherSinceVersion()
}

func (*MassCancelRejectReasonEnum) OtherDeprecated() uint16 {
	return 0
}

func (*MassCancelRejectReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MassCancelRejectReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MassCancelRejectReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
