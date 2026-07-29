// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ListTriggerActionEnum byte
type ListTriggerActionValues struct {
	Release          ListTriggerActionEnum
	Cancel           ListTriggerActionEnum
	NonRepresentable ListTriggerActionEnum
	NullValue        ListTriggerActionEnum
}

var ListTriggerAction = ListTriggerActionValues{49, 50, 126, 0}

func (l ListTriggerActionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(l)); err != nil {
		return err
	}
	return nil
}

func (l *ListTriggerActionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(l)); err != nil {
		return err
	}
	return nil
}

func (l ListTriggerActionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ListTriggerAction)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ListTriggerAction, unknown enumeration value %d", l)
}

func (*ListTriggerActionEnum) EncodedLength() int64 {
	return 1
}

func (*ListTriggerActionEnum) ReleaseSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerActionEnum) ReleaseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ReleaseSinceVersion()
}

func (*ListTriggerActionEnum) ReleaseDeprecated() uint16 {
	return 0
}

func (*ListTriggerActionEnum) CancelSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerActionEnum) CancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.CancelSinceVersion()
}

func (*ListTriggerActionEnum) CancelDeprecated() uint16 {
	return 0
}

func (*ListTriggerActionEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (l *ListTriggerActionEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.NonRepresentableSinceVersion()
}

func (*ListTriggerActionEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
