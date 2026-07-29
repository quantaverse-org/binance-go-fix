// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrderRateLimitExceededModeEnum uint8
type OrderRateLimitExceededModeValues struct {
	DoNothing        OrderRateLimitExceededModeEnum
	CancelOnly       OrderRateLimitExceededModeEnum
	NonRepresentable OrderRateLimitExceededModeEnum
	NullValue        OrderRateLimitExceededModeEnum
}

var OrderRateLimitExceededMode = OrderRateLimitExceededModeValues{1, 2, 254, 255}

func (o OrderRateLimitExceededModeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderRateLimitExceededModeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderRateLimitExceededModeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderRateLimitExceededMode)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderRateLimitExceededMode, unknown enumeration value %d", o)
}

func (*OrderRateLimitExceededModeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderRateLimitExceededModeEnum) DoNothingSinceVersion() uint16 {
	return 0
}

func (o *OrderRateLimitExceededModeEnum) DoNothingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DoNothingSinceVersion()
}

func (*OrderRateLimitExceededModeEnum) DoNothingDeprecated() uint16 {
	return 0
}

func (*OrderRateLimitExceededModeEnum) CancelOnlySinceVersion() uint16 {
	return 0
}

func (o *OrderRateLimitExceededModeEnum) CancelOnlyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelOnlySinceVersion()
}

func (*OrderRateLimitExceededModeEnum) CancelOnlyDeprecated() uint16 {
	return 0
}

func (*OrderRateLimitExceededModeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (o *OrderRateLimitExceededModeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NonRepresentableSinceVersion()
}

func (*OrderRateLimitExceededModeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
