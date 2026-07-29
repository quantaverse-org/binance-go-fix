// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type CancelRestrictionsEnum uint8
type CancelRestrictionsValues struct {
	OnlyNew             CancelRestrictionsEnum
	OnlyPartiallyFilled CancelRestrictionsEnum
	NonRepresentable    CancelRestrictionsEnum
	NullValue           CancelRestrictionsEnum
}

var CancelRestrictions = CancelRestrictionsValues{1, 2, 254, 255}

func (c CancelRestrictionsEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(c)); err != nil {
		return err
	}
	return nil
}

func (c *CancelRestrictionsEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(c)); err != nil {
		return err
	}
	return nil
}

func (c CancelRestrictionsEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CancelRestrictions)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CancelRestrictions, unknown enumeration value %d", c)
}

func (*CancelRestrictionsEnum) EncodedLength() int64 {
	return 1
}

func (*CancelRestrictionsEnum) OnlyNewSinceVersion() uint16 {
	return 0
}

func (c *CancelRestrictionsEnum) OnlyNewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OnlyNewSinceVersion()
}

func (*CancelRestrictionsEnum) OnlyNewDeprecated() uint16 {
	return 0
}

func (*CancelRestrictionsEnum) OnlyPartiallyFilledSinceVersion() uint16 {
	return 0
}

func (c *CancelRestrictionsEnum) OnlyPartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OnlyPartiallyFilledSinceVersion()
}

func (*CancelRestrictionsEnum) OnlyPartiallyFilledDeprecated() uint16 {
	return 0
}

func (*CancelRestrictionsEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (c *CancelRestrictionsEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NonRepresentableSinceVersion()
}

func (*CancelRestrictionsEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
