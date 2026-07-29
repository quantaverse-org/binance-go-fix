// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type SubscriptionRequestTypeEnum byte
type SubscriptionRequestTypeValues struct {
	Subscribe        SubscriptionRequestTypeEnum
	Unsubscribe      SubscriptionRequestTypeEnum
	NonRepresentable SubscriptionRequestTypeEnum
	NullValue        SubscriptionRequestTypeEnum
}

var SubscriptionRequestType = SubscriptionRequestTypeValues{49, 50, 126, 0}

func (s SubscriptionRequestTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(s)); err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionRequestTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(s)); err != nil {
		return err
	}
	return nil
}

func (s SubscriptionRequestTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SubscriptionRequestType)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SubscriptionRequestType, unknown enumeration value %d", s)
}

func (*SubscriptionRequestTypeEnum) EncodedLength() int64 {
	return 1
}

func (*SubscriptionRequestTypeEnum) SubscribeSinceVersion() uint16 {
	return 0
}

func (s *SubscriptionRequestTypeEnum) SubscribeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SubscribeSinceVersion()
}

func (*SubscriptionRequestTypeEnum) SubscribeDeprecated() uint16 {
	return 0
}

func (*SubscriptionRequestTypeEnum) UnsubscribeSinceVersion() uint16 {
	return 0
}

func (s *SubscriptionRequestTypeEnum) UnsubscribeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UnsubscribeSinceVersion()
}

func (*SubscriptionRequestTypeEnum) UnsubscribeDeprecated() uint16 {
	return 0
}

func (*SubscriptionRequestTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (s *SubscriptionRequestTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NonRepresentableSinceVersion()
}

func (*SubscriptionRequestTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
