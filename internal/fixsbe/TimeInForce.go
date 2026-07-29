// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type TimeInForceEnum byte
type TimeInForceValues struct {
	GoodTillCancel    TimeInForceEnum
	ImmediateOrCancel TimeInForceEnum
	FillOrKill        TimeInForceEnum
	NonRepresentable  TimeInForceEnum
	NullValue         TimeInForceEnum
}

var TimeInForce = TimeInForceValues{49, 51, 52, 126, 0}

func (t TimeInForceEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(t)); err != nil {
		return err
	}
	return nil
}

func (t *TimeInForceEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(t)); err != nil {
		return err
	}
	return nil
}

func (t TimeInForceEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(TimeInForce)
	for idx := 0; idx < value.NumField(); idx++ {
		if t == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on TimeInForce, unknown enumeration value %d", t)
}

func (*TimeInForceEnum) EncodedLength() int64 {
	return 1
}

func (*TimeInForceEnum) GoodTillCancelSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) GoodTillCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.GoodTillCancelSinceVersion()
}

func (*TimeInForceEnum) GoodTillCancelDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) ImmediateOrCancelSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) ImmediateOrCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ImmediateOrCancelSinceVersion()
}

func (*TimeInForceEnum) ImmediateOrCancelDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) FillOrKillSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) FillOrKillInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.FillOrKillSinceVersion()
}

func (*TimeInForceEnum) FillOrKillDeprecated() uint16 {
	return 0
}

func (*TimeInForceEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (t *TimeInForceEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.NonRepresentableSinceVersion()
}

func (*TimeInForceEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
