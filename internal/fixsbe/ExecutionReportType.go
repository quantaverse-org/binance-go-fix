// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ExecutionReportTypeEnum uint8
type ExecutionReportTypeValues struct {
	Full             ExecutionReportTypeEnum
	Mini             ExecutionReportTypeEnum
	NonRepresentable ExecutionReportTypeEnum
	NullValue        ExecutionReportTypeEnum
}

var ExecutionReportType = ExecutionReportTypeValues{1, 2, 254, 255}

func (e ExecutionReportTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecutionReportTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecutionReportTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecutionReportType)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecutionReportType, unknown enumeration value %d", e)
}

func (*ExecutionReportTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ExecutionReportTypeEnum) FullSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTypeEnum) FullInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.FullSinceVersion()
}

func (*ExecutionReportTypeEnum) FullDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTypeEnum) MiniSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTypeEnum) MiniInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.MiniSinceVersion()
}

func (*ExecutionReportTypeEnum) MiniDeprecated() uint16 {
	return 0
}

func (*ExecutionReportTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (e *ExecutionReportTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NonRepresentableSinceVersion()
}

func (*ExecutionReportTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
