// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ExecTypeEnum byte
type ExecTypeValues struct {
	New              ExecTypeEnum
	Canceled         ExecTypeEnum
	Replaced         ExecTypeEnum
	Rejected         ExecTypeEnum
	Trade            ExecTypeEnum
	Expired          ExecTypeEnum
	NonRepresentable ExecTypeEnum
	NullValue        ExecTypeEnum
}

var ExecType = ExecTypeValues{48, 52, 53, 56, 70, 67, 126, 0}

func (e ExecTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExecTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExecTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExecType)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExecType, unknown enumeration value %d", e)
}

func (*ExecTypeEnum) EncodedLength() int64 {
	return 1
}

func (*ExecTypeEnum) NewSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) NewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NewSinceVersion()
}

func (*ExecTypeEnum) NewDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) CanceledSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) CanceledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.CanceledSinceVersion()
}

func (*ExecTypeEnum) CanceledDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) ReplacedSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) ReplacedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ReplacedSinceVersion()
}

func (*ExecTypeEnum) ReplacedDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectedSinceVersion()
}

func (*ExecTypeEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) TradeSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) TradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.TradeSinceVersion()
}

func (*ExecTypeEnum) TradeDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) ExpiredSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) ExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExpiredSinceVersion()
}

func (*ExecTypeEnum) ExpiredDeprecated() uint16 {
	return 0
}

func (*ExecTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (e *ExecTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NonRepresentableSinceVersion()
}

func (*ExecTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
