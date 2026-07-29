// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type TriggerPriceDirectionEnum byte
type TriggerPriceDirectionValues struct {
	Up               TriggerPriceDirectionEnum
	Down             TriggerPriceDirectionEnum
	NonRepresentable TriggerPriceDirectionEnum
	NullValue        TriggerPriceDirectionEnum
}

var TriggerPriceDirection = TriggerPriceDirectionValues{85, 68, 126, 0}

func (t TriggerPriceDirectionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TriggerPriceDirectionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TriggerPriceDirectionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TriggerPriceDirection)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TriggerPriceDirection, unknown enumeration value %d", t)
}

func (*TriggerPriceDirectionEnum) EncodedLength() int64 {
	return 1
}

func (*TriggerPriceDirectionEnum) UpSinceVersion() uint16 {
	return 0
}

func (t *TriggerPriceDirectionEnum) UpInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.UpSinceVersion()
}

func (*TriggerPriceDirectionEnum) UpDeprecated() uint16 {
	return 0
}

func (*TriggerPriceDirectionEnum) DownSinceVersion() uint16 {
	return 0
}

func (t *TriggerPriceDirectionEnum) DownInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.DownSinceVersion()
}

func (*TriggerPriceDirectionEnum) DownDeprecated() uint16 {
	return 0
}

func (*TriggerPriceDirectionEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (t *TriggerPriceDirectionEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NonRepresentableSinceVersion()
}

func (*TriggerPriceDirectionEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
