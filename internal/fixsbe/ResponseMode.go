// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ResponseModeEnum uint8
type ResponseModeValues struct {
	Everything       ResponseModeEnum
	OnlyAcks         ResponseModeEnum
	NonRepresentable ResponseModeEnum
	NullValue        ResponseModeEnum
}

var ResponseMode = ResponseModeValues{1, 2, 254, 255}

func (r ResponseModeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(r)); err != nil {
		return err
	}
	return nil
}

func (r *ResponseModeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(r)); err != nil {
		return err
	}
	return nil
}

func (r ResponseModeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ResponseMode)
	for idx := 0; idx < value.NumField(); idx++ {
		if r == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ResponseMode, unknown enumeration value %d", r)
}

func (*ResponseModeEnum) EncodedLength() int64 {
	return 1
}

func (*ResponseModeEnum) EverythingSinceVersion() uint16 {
	return 0
}

func (r *ResponseModeEnum) EverythingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.EverythingSinceVersion()
}

func (*ResponseModeEnum) EverythingDeprecated() uint16 {
	return 0
}

func (*ResponseModeEnum) OnlyAcksSinceVersion() uint16 {
	return 0
}

func (r *ResponseModeEnum) OnlyAcksInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.OnlyAcksSinceVersion()
}

func (*ResponseModeEnum) OnlyAcksDeprecated() uint16 {
	return 0
}

func (*ResponseModeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (r *ResponseModeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= r.NonRepresentableSinceVersion()
}

func (*ResponseModeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
