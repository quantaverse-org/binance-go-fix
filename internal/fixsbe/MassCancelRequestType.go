// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MassCancelRequestTypeEnum byte
type MassCancelRequestTypeValues struct {
	CancelSymbolOrders MassCancelRequestTypeEnum
	NonRepresentable   MassCancelRequestTypeEnum
	NullValue          MassCancelRequestTypeEnum
}

var MassCancelRequestType = MassCancelRequestTypeValues{49, 126, 0}

func (m MassCancelRequestTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassCancelRequestTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassCancelRequestTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassCancelRequestType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassCancelRequestType, unknown enumeration value %d", m)
}

func (*MassCancelRequestTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MassCancelRequestTypeEnum) CancelSymbolOrdersSinceVersion() uint16 {
	return 0
}

func (m *MassCancelRequestTypeEnum) CancelSymbolOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CancelSymbolOrdersSinceVersion()
}

func (*MassCancelRequestTypeEnum) CancelSymbolOrdersDeprecated() uint16 {
	return 0
}

func (*MassCancelRequestTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MassCancelRequestTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MassCancelRequestTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
