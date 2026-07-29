// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ListRejectReasonEnum uint8
type ListRejectReasonValues struct {
	Other            ListRejectReasonEnum
	NonRepresentable ListRejectReasonEnum
	NullValue        ListRejectReasonEnum
}

var ListRejectReason = ListRejectReasonValues{99, 254, 255}

func (l ListRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(l)); err != nil {
		return err
	}
	return nil
}

func (l *ListRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(l)); err != nil {
		return err
	}
	return nil
}

func (l ListRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ListRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ListRejectReason, unknown enumeration value %d", l)
}

func (*ListRejectReasonEnum) EncodedLength() int64 {
	return 1
}

func (*ListRejectReasonEnum) OtherSinceVersion() uint16 {
	return 0
}

func (l *ListRejectReasonEnum) OtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OtherSinceVersion()
}

func (*ListRejectReasonEnum) OtherDeprecated() uint16 {
	return 0
}

func (*ListRejectReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (l *ListRejectReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.NonRepresentableSinceVersion()
}

func (*ListRejectReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
