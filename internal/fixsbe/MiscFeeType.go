// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MiscFeeTypeEnum uint8
type MiscFeeTypeValues struct {
	ExchangeFees     MiscFeeTypeEnum
	NonRepresentable MiscFeeTypeEnum
	NullValue        MiscFeeTypeEnum
}

var MiscFeeType = MiscFeeTypeValues{4, 254, 255}

func (m MiscFeeTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(m)); err != nil {
		return err
	}
	return nil
}

func (m *MiscFeeTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(m)); err != nil {
		return err
	}
	return nil
}

func (m MiscFeeTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MiscFeeType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MiscFeeType, unknown enumeration value %d", m)
}

func (*MiscFeeTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MiscFeeTypeEnum) ExchangeFeesSinceVersion() uint16 {
	return 0
}

func (m *MiscFeeTypeEnum) ExchangeFeesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.ExchangeFeesSinceVersion()
}

func (*MiscFeeTypeEnum) ExchangeFeesDeprecated() uint16 {
	return 0
}

func (*MiscFeeTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MiscFeeTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MiscFeeTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
