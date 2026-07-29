// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type TriggerTypeEnum byte
type TriggerTypeValues struct {
	PriceMovement    TriggerTypeEnum
	NonRepresentable TriggerTypeEnum
	NullValue        TriggerTypeEnum
}

var TriggerType = TriggerTypeValues{52, 126, 0}

func (t TriggerTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TriggerTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TriggerTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TriggerType)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TriggerType, unknown enumeration value %d", t)
}

func (*TriggerTypeEnum) EncodedLength() int64 {
	return 1
}

func (*TriggerTypeEnum) PriceMovementSinceVersion() uint16 {
	return 0
}

func (t *TriggerTypeEnum) PriceMovementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PriceMovementSinceVersion()
}

func (*TriggerTypeEnum) PriceMovementDeprecated() uint16 {
	return 0
}

func (*TriggerTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (t *TriggerTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NonRepresentableSinceVersion()
}

func (*TriggerTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
