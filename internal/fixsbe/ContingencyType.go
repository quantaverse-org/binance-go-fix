// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ContingencyTypeEnum uint8
type ContingencyTypeValues struct {
	OneCancelsTheOther  ContingencyTypeEnum
	OneTriggersTheOther ContingencyTypeEnum
	NonRepresentable    ContingencyTypeEnum
	NullValue           ContingencyTypeEnum
}

var ContingencyType = ContingencyTypeValues{1, 2, 254, 255}

func (c ContingencyTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(c)); err != nil {
		return err
	}
	return nil
}

func (c *ContingencyTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(c)); err != nil {
		return err
	}
	return nil
}

func (c ContingencyTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ContingencyType)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ContingencyType, unknown enumeration value %d", c)
}

func (*ContingencyTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ContingencyTypeEnum) OneCancelsTheOtherSinceVersion() uint16 {
	return 0
}

func (c *ContingencyTypeEnum) OneCancelsTheOtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OneCancelsTheOtherSinceVersion()
}

func (*ContingencyTypeEnum) OneCancelsTheOtherDeprecated() uint16 {
	return 0
}

func (*ContingencyTypeEnum) OneTriggersTheOtherSinceVersion() uint16 {
	return 0
}

func (c *ContingencyTypeEnum) OneTriggersTheOtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OneTriggersTheOtherSinceVersion()
}

func (*ContingencyTypeEnum) OneTriggersTheOtherDeprecated() uint16 {
	return 0
}

func (*ContingencyTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (c *ContingencyTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NonRepresentableSinceVersion()
}

func (*ContingencyTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
