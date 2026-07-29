// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type PegMoveTypeEnum uint8
type PegMoveTypeValues struct {
	Fixed            PegMoveTypeEnum
	NonRepresentable PegMoveTypeEnum
	NullValue        PegMoveTypeEnum
}

var PegMoveType = PegMoveTypeValues{1, 254, 255}

func (p PegMoveTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(p)); err != nil {
		return err
	}
	return nil
}

func (p *PegMoveTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(p)); err != nil {
		return err
	}
	return nil
}

func (p PegMoveTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PegMoveType)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PegMoveType, unknown enumeration value %d", p)
}

func (*PegMoveTypeEnum) EncodedLength() int64 {
	return 1
}

func (*PegMoveTypeEnum) FixedSinceVersion() uint16 {
	return 0
}

func (p *PegMoveTypeEnum) FixedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.FixedSinceVersion()
}

func (*PegMoveTypeEnum) FixedDeprecated() uint16 {
	return 0
}

func (*PegMoveTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (p *PegMoveTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NonRepresentableSinceVersion()
}

func (*PegMoveTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
