// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type LimitResetIntervalResolutionEnum byte
type LimitResetIntervalResolutionValues struct {
	Second           LimitResetIntervalResolutionEnum
	Minute           LimitResetIntervalResolutionEnum
	Hour             LimitResetIntervalResolutionEnum
	Day              LimitResetIntervalResolutionEnum
	NonRepresentable LimitResetIntervalResolutionEnum
	NullValue        LimitResetIntervalResolutionEnum
}

var LimitResetIntervalResolution = LimitResetIntervalResolutionValues{115, 109, 104, 100, 126, 0}

func (l LimitResetIntervalResolutionEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(l)); err != nil {
		return err
	}
	return nil
}

func (l *LimitResetIntervalResolutionEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(l)); err != nil {
		return err
	}
	return nil
}

func (l LimitResetIntervalResolutionEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(LimitResetIntervalResolution)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on LimitResetIntervalResolution, unknown enumeration value %d", l)
}

func (*LimitResetIntervalResolutionEnum) EncodedLength() int64 {
	return 1
}

func (*LimitResetIntervalResolutionEnum) SecondSinceVersion() uint16 {
	return 0
}

func (l *LimitResetIntervalResolutionEnum) SecondInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SecondSinceVersion()
}

func (*LimitResetIntervalResolutionEnum) SecondDeprecated() uint16 {
	return 0
}

func (*LimitResetIntervalResolutionEnum) MinuteSinceVersion() uint16 {
	return 0
}

func (l *LimitResetIntervalResolutionEnum) MinuteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.MinuteSinceVersion()
}

func (*LimitResetIntervalResolutionEnum) MinuteDeprecated() uint16 {
	return 0
}

func (*LimitResetIntervalResolutionEnum) HourSinceVersion() uint16 {
	return 0
}

func (l *LimitResetIntervalResolutionEnum) HourInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.HourSinceVersion()
}

func (*LimitResetIntervalResolutionEnum) HourDeprecated() uint16 {
	return 0
}

func (*LimitResetIntervalResolutionEnum) DaySinceVersion() uint16 {
	return 0
}

func (l *LimitResetIntervalResolutionEnum) DayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.DaySinceVersion()
}

func (*LimitResetIntervalResolutionEnum) DayDeprecated() uint16 {
	return 0
}

func (*LimitResetIntervalResolutionEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (l *LimitResetIntervalResolutionEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.NonRepresentableSinceVersion()
}

func (*LimitResetIntervalResolutionEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
