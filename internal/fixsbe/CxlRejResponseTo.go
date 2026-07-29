// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type CxlRejResponseToEnum byte
type CxlRejResponseToValues struct {
	OrderCancelRequest CxlRejResponseToEnum
	NonRepresentable   CxlRejResponseToEnum
	NullValue          CxlRejResponseToEnum
}

var CxlRejResponseTo = CxlRejResponseToValues{49, 126, 0}

func (c CxlRejResponseToEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(c)); err != nil {
		return err
	}
	return nil
}

func (c *CxlRejResponseToEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(c)); err != nil {
		return err
	}
	return nil
}

func (c CxlRejResponseToEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(CxlRejResponseTo)
	for idx := 0; idx < value.NumField(); idx++ {
		if c == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on CxlRejResponseTo, unknown enumeration value %d", c)
}

func (*CxlRejResponseToEnum) EncodedLength() int64 {
	return 1
}

func (*CxlRejResponseToEnum) OrderCancelRequestSinceVersion() uint16 {
	return 0
}

func (c *CxlRejResponseToEnum) OrderCancelRequestInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.OrderCancelRequestSinceVersion()
}

func (*CxlRejResponseToEnum) OrderCancelRequestDeprecated() uint16 {
	return 0
}

func (*CxlRejResponseToEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (c *CxlRejResponseToEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.NonRepresentableSinceVersion()
}

func (*CxlRejResponseToEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
