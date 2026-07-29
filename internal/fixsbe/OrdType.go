// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrdTypeEnum byte
type OrdTypeValues struct {
	Market           OrdTypeEnum
	Limit            OrdTypeEnum
	Stop             OrdTypeEnum
	StopLimit        OrdTypeEnum
	Pegged           OrdTypeEnum
	NonRepresentable OrdTypeEnum
	NullValue        OrdTypeEnum
}

var OrdType = OrdTypeValues{49, 50, 51, 52, 80, 126, 0}

func (o OrdTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdType, unknown enumeration value %d", o)
}

func (*OrdTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrdTypeEnum) MarketSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) MarketInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MarketSinceVersion()
}

func (*OrdTypeEnum) MarketDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) LimitSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) LimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LimitSinceVersion()
}

func (*OrdTypeEnum) LimitDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) StopSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) StopInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopSinceVersion()
}

func (*OrdTypeEnum) StopDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) StopLimitSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) StopLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopLimitSinceVersion()
}

func (*OrdTypeEnum) StopLimitDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) PeggedSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) PeggedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PeggedSinceVersion()
}

func (*OrdTypeEnum) PeggedDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NonRepresentableSinceVersion()
}

func (*OrdTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
