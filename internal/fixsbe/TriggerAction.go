// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type TriggerActionEnum byte
type TriggerActionValues struct {
	Activate         TriggerActionEnum
	NonRepresentable TriggerActionEnum
	NullValue        TriggerActionEnum
}

var TriggerAction = TriggerActionValues{49, 126, 0}

func (t TriggerActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TriggerActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TriggerActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TriggerAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TriggerAction, unknown enumeration value %d", t)
}

func (*TriggerActionEnum) EncodedLength() int64 {
	return 1
}

func (*TriggerActionEnum) ActivateSinceVersion() uint16 {
	return 0
}

func (t *TriggerActionEnum) ActivateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ActivateSinceVersion()
}

func (*TriggerActionEnum) ActivateDeprecated() uint16 {
	return 0
}

func (*TriggerActionEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (t *TriggerActionEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NonRepresentableSinceVersion()
}

func (*TriggerActionEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
