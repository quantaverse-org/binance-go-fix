// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type LimitTypeEnum byte
type LimitTypeValues struct {
	OrderLimit        LimitTypeEnum
	MessageLimit      LimitTypeEnum
	SubscriptionLimit LimitTypeEnum
	NonRepresentable  LimitTypeEnum
	NullValue         LimitTypeEnum
}

var LimitType = LimitTypeValues{49, 50, 51, 126, 0}

func (l LimitTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(l)); err != nil {
		return err
	}
	return nil
}

func (l *LimitTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(l)); err != nil {
		return err
	}
	return nil
}

func (l LimitTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(LimitType)
	for idx := 0; idx < value.NumField(); idx++ {
		if l == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on LimitType, unknown enumeration value %d", l)
}

func (*LimitTypeEnum) EncodedLength() int64 {
	return 1
}

func (*LimitTypeEnum) OrderLimitSinceVersion() uint16 {
	return 0
}

func (l *LimitTypeEnum) OrderLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderLimitSinceVersion()
}

func (*LimitTypeEnum) OrderLimitDeprecated() uint16 {
	return 0
}

func (*LimitTypeEnum) MessageLimitSinceVersion() uint16 {
	return 0
}

func (l *LimitTypeEnum) MessageLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.MessageLimitSinceVersion()
}

func (*LimitTypeEnum) MessageLimitDeprecated() uint16 {
	return 0
}

func (*LimitTypeEnum) SubscriptionLimitSinceVersion() uint16 {
	return 0
}

func (l *LimitTypeEnum) SubscriptionLimitInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SubscriptionLimitSinceVersion()
}

func (*LimitTypeEnum) SubscriptionLimitDeprecated() uint16 {
	return 0
}

func (*LimitTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (l *LimitTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.NonRepresentableSinceVersion()
}

func (*LimitTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
