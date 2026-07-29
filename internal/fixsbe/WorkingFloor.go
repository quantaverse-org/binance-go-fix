// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type WorkingFloorEnum uint8
type WorkingFloorValues struct {
	Exchange         WorkingFloorEnum
	Broker           WorkingFloorEnum
	Sor              WorkingFloorEnum
	NonRepresentable WorkingFloorEnum
	NullValue        WorkingFloorEnum
}

var WorkingFloor = WorkingFloorValues{1, 2, 3, 254, 255}

func (w WorkingFloorEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(w)); err != nil {
		return err
	}
	return nil
}

func (w *WorkingFloorEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(w)); err != nil {
		return err
	}
	return nil
}

func (w WorkingFloorEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(WorkingFloor)
	for idx := 0; idx < value.NumField(); idx++ {
		if w == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on WorkingFloor, unknown enumeration value %d", w)
}

func (*WorkingFloorEnum) EncodedLength() int64 {
	return 1
}

func (*WorkingFloorEnum) ExchangeSinceVersion() uint16 {
	return 0
}

func (w *WorkingFloorEnum) ExchangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= w.ExchangeSinceVersion()
}

func (*WorkingFloorEnum) ExchangeDeprecated() uint16 {
	return 0
}

func (*WorkingFloorEnum) BrokerSinceVersion() uint16 {
	return 0
}

func (w *WorkingFloorEnum) BrokerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= w.BrokerSinceVersion()
}

func (*WorkingFloorEnum) BrokerDeprecated() uint16 {
	return 0
}

func (*WorkingFloorEnum) SorSinceVersion() uint16 {
	return 0
}

func (w *WorkingFloorEnum) SorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= w.SorSinceVersion()
}

func (*WorkingFloorEnum) SorDeprecated() uint16 {
	return 0
}

func (*WorkingFloorEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (w *WorkingFloorEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= w.NonRepresentableSinceVersion()
}

func (*WorkingFloorEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
