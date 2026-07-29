// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ExecInstEnum byte
type ExecInstValues struct {
	ParticipateDontInitiate ExecInstEnum
	NonRepresentable        ExecInstEnum
	NullValue               ExecInstEnum
}

var ExecInst = ExecInstValues{54, 126, 0}

func (e ExecInstEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecInstEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecInstEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecInst)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecInst, unknown enumeration value %d", e)
}

func (*ExecInstEnum) EncodedLength() int64 {
	return 1
}

func (*ExecInstEnum) ParticipateDontInitiateSinceVersion() uint16 {
	return 0
}

func (e *ExecInstEnum) ParticipateDontInitiateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ParticipateDontInitiateSinceVersion()
}

func (*ExecInstEnum) ParticipateDontInitiateDeprecated() uint16 {
	return 0
}

func (*ExecInstEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (e *ExecInstEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NonRepresentableSinceVersion()
}

func (*ExecInstEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
