// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type InstrumentListRequestTypeEnum uint8
type InstrumentListRequestTypeValues struct {
	SingleInstrument InstrumentListRequestTypeEnum
	AllInstruments   InstrumentListRequestTypeEnum
	NonRepresentable InstrumentListRequestTypeEnum
	NullValue        InstrumentListRequestTypeEnum
}

var InstrumentListRequestType = InstrumentListRequestTypeValues{0, 4, 254, 255}

func (i InstrumentListRequestTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(i)); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentListRequestTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(i)); err != nil {
		return err
	}
	return nil
}

func (i InstrumentListRequestTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(InstrumentListRequestType)
	for idx := 0; idx < value.NumField(); idx++ {
		if i == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on InstrumentListRequestType, unknown enumeration value %d", i)
}

func (*InstrumentListRequestTypeEnum) EncodedLength() int64 {
	return 1
}

func (*InstrumentListRequestTypeEnum) SingleInstrumentSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequestTypeEnum) SingleInstrumentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SingleInstrumentSinceVersion()
}

func (*InstrumentListRequestTypeEnum) SingleInstrumentDeprecated() uint16 {
	return 0
}

func (*InstrumentListRequestTypeEnum) AllInstrumentsSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequestTypeEnum) AllInstrumentsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.AllInstrumentsSinceVersion()
}

func (*InstrumentListRequestTypeEnum) AllInstrumentsDeprecated() uint16 {
	return 0
}

func (*InstrumentListRequestTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRequestTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.NonRepresentableSinceVersion()
}

func (*InstrumentListRequestTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
