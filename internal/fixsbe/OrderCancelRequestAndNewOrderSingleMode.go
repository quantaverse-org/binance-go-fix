// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrderCancelRequestAndNewOrderSingleModeEnum uint8
type OrderCancelRequestAndNewOrderSingleModeValues struct {
	StopOnFailure    OrderCancelRequestAndNewOrderSingleModeEnum
	AllowFailure     OrderCancelRequestAndNewOrderSingleModeEnum
	NonRepresentable OrderCancelRequestAndNewOrderSingleModeEnum
	NullValue        OrderCancelRequestAndNewOrderSingleModeEnum
}

var OrderCancelRequestAndNewOrderSingleMode = OrderCancelRequestAndNewOrderSingleModeValues{1, 2, 254, 255}

func (o OrderCancelRequestAndNewOrderSingleModeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancelRequestAndNewOrderSingleModeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrderCancelRequestAndNewOrderSingleModeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrderCancelRequestAndNewOrderSingleMode)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrderCancelRequestAndNewOrderSingleMode, unknown enumeration value %d", o)
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) EncodedLength() int64 {
	return 1
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) StopOnFailureSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingleModeEnum) StopOnFailureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopOnFailureSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) StopOnFailureDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) AllowFailureSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingleModeEnum) AllowFailureInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AllowFailureSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) AllowFailureDeprecated() uint16 {
	return 0
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (o *OrderCancelRequestAndNewOrderSingleModeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NonRepresentableSinceVersion()
}

func (*OrderCancelRequestAndNewOrderSingleModeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
