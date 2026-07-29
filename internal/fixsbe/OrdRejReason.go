// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrdRejReasonEnum uint8
type OrdRejReasonValues struct {
	Other            OrdRejReasonEnum
	NonRepresentable OrdRejReasonEnum
	NullValue        OrdRejReasonEnum
}

var OrdRejReason = OrdRejReasonValues{99, 254, 255}

func (o OrdRejReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdRejReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdRejReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdRejReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdRejReason, unknown enumeration value %d", o)
}

func (*OrdRejReasonEnum) EncodedLength() int64 {
	return 1
}

func (*OrdRejReasonEnum) OtherSinceVersion() uint16 {
	return 0
}

func (o *OrdRejReasonEnum) OtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OtherSinceVersion()
}

func (*OrdRejReasonEnum) OtherDeprecated() uint16 {
	return 0
}

func (*OrdRejReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (o *OrdRejReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NonRepresentableSinceVersion()
}

func (*OrdRejReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
