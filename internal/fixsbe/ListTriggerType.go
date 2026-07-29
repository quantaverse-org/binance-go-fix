// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ListTriggerTypeEnum byte
type ListTriggerTypeValues struct {
	Activated        ListTriggerTypeEnum
	PartiallyFilled  ListTriggerTypeEnum
	Filled           ListTriggerTypeEnum
	NonRepresentable ListTriggerTypeEnum
	NullValue        ListTriggerTypeEnum
}

var ListTriggerType = ListTriggerTypeValues{49, 50, 51, 126, 0}

func (l ListTriggerTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(l)); err != nil {
		return err
	}
	return nil
}

func (l *ListTriggerTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(l)); err != nil {
		return err
	}
	return nil
}

func (l ListTriggerTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ListTriggerType)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ListTriggerType, unknown enumeration value %d", l)
}

func (*ListTriggerTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ListTriggerTypeEnum) ActivatedSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerTypeEnum) ActivatedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ActivatedSinceVersion()
}

func (*ListTriggerTypeEnum) ActivatedDeprecated() uint16 {
	return 0
}

func (*ListTriggerTypeEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerTypeEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.PartiallyFilledSinceVersion()
}

func (*ListTriggerTypeEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*ListTriggerTypeEnum) FilledSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerTypeEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.FilledSinceVersion()
}

func (*ListTriggerTypeEnum) FilledDeprecated() uint16 {
	return 0
}

func (*ListTriggerTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.NonRepresentableSinceVersion()
}

func (*ListTriggerTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
