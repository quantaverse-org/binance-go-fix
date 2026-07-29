// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type PegOffsetTypeEnum byte
type PegOffsetTypeValues struct {
	PriceTier        PegOffsetTypeEnum
	NonRepresentable PegOffsetTypeEnum
	NullValue        PegOffsetTypeEnum
}

var PegOffsetType = PegOffsetTypeValues{51, 126, 0}

func (p PegOffsetTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(p)); err != nil {
		return err
	}
	return nil
}

func (p *PegOffsetTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(p)); err != nil {
		return err
	}
	return nil
}

func (p PegOffsetTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(PegOffsetType)
	for idx := 0; idx < value.NumField(); idx++ {
		if p == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on PegOffsetType, unknown enumeration value %d", p)
}

func (*PegOffsetTypeEnum) EncodedLength() int64 {
	return 1
}

func (*PegOffsetTypeEnum) PriceTierSinceVersion() uint16 {
	return 0
}

func (p *PegOffsetTypeEnum) PriceTierInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PriceTierSinceVersion()
}

func (*PegOffsetTypeEnum) PriceTierDeprecated() uint16 {
	return 0
}

func (*PegOffsetTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (p *PegOffsetTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.NonRepresentableSinceVersion()
}

func (*PegOffsetTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
