// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type TriggerPriceTypeEnum byte
type TriggerPriceTypeValues struct {
	LastTrade        TriggerPriceTypeEnum
	NonRepresentable TriggerPriceTypeEnum
	NullValue        TriggerPriceTypeEnum
}

var TriggerPriceType = TriggerPriceTypeValues{50, 126, 0}

func (t TriggerPriceTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TriggerPriceTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TriggerPriceTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TriggerPriceType)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TriggerPriceType, unknown enumeration value %d", t)
}

func (*TriggerPriceTypeEnum) EncodedLength() int64 {
	return 1
}

func (*TriggerPriceTypeEnum) LastTradeSinceVersion() uint16 {
	return 0
}

func (t *TriggerPriceTypeEnum) LastTradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LastTradeSinceVersion()
}

func (*TriggerPriceTypeEnum) LastTradeDeprecated() uint16 {
	return 0
}

func (*TriggerPriceTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (t *TriggerPriceTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NonRepresentableSinceVersion()
}

func (*TriggerPriceTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
