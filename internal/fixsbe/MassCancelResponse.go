// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MassCancelResponseEnum byte
type MassCancelResponseValues struct {
	CancelRequestRejected MassCancelResponseEnum
	CancelSymbolOrders    MassCancelResponseEnum
	NonRepresentable      MassCancelResponseEnum
	NullValue             MassCancelResponseEnum
}

var MassCancelResponse = MassCancelResponseValues{48, 49, 126, 0}

func (m MassCancelResponseEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MassCancelResponseEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MassCancelResponseEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MassCancelResponse)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MassCancelResponse, unknown enumeration value %d", m)
}

func (*MassCancelResponseEnum) EncodedLength() int64 {
	return 1
}

func (*MassCancelResponseEnum) CancelRequestRejectedSinceVersion() uint16 {
	return 0
}

func (m *MassCancelResponseEnum) CancelRequestRejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CancelRequestRejectedSinceVersion()
}

func (*MassCancelResponseEnum) CancelRequestRejectedDeprecated() uint16 {
	return 0
}

func (*MassCancelResponseEnum) CancelSymbolOrdersSinceVersion() uint16 {
	return 0
}

func (m *MassCancelResponseEnum) CancelSymbolOrdersInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.CancelSymbolOrdersSinceVersion()
}

func (*MassCancelResponseEnum) CancelSymbolOrdersDeprecated() uint16 {
	return 0
}

func (*MassCancelResponseEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MassCancelResponseEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MassCancelResponseEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
