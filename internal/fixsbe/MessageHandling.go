// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MessageHandlingEnum uint8
type MessageHandlingValues struct {
	Unordered        MessageHandlingEnum
	Sequential       MessageHandlingEnum
	NonRepresentable MessageHandlingEnum
	NullValue        MessageHandlingEnum
}

var MessageHandling = MessageHandlingValues{1, 2, 254, 255}

func (m MessageHandlingEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MessageHandlingEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MessageHandlingEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MessageHandling)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MessageHandling, unknown enumeration value %d", m)
}

func (*MessageHandlingEnum) EncodedLength() int64 {
	return 1
}

func (*MessageHandlingEnum) UnorderedSinceVersion() uint16 {
	return 0
}

func (m *MessageHandlingEnum) UnorderedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.UnorderedSinceVersion()
}

func (*MessageHandlingEnum) UnorderedDeprecated() uint16 {
	return 0
}

func (*MessageHandlingEnum) SequentialSinceVersion() uint16 {
	return 0
}

func (m *MessageHandlingEnum) SequentialInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SequentialSinceVersion()
}

func (*MessageHandlingEnum) SequentialDeprecated() uint16 {
	return 0
}

func (*MessageHandlingEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MessageHandlingEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MessageHandlingEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
