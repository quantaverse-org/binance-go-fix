// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type SessionRejectReasonEnum uint8
type SessionRejectReasonValues struct {
	InvalidTagNumber                          SessionRejectReasonEnum
	RequiredTagMissing                        SessionRejectReasonEnum
	TagNotDefinedForThisMessageType           SessionRejectReasonEnum
	UndefinedTag                              SessionRejectReasonEnum
	ValueIsIncorrect                          SessionRejectReasonEnum
	IncorrectDataFormatForValue               SessionRejectReasonEnum
	SignatureProblem                          SessionRejectReasonEnum
	SendingTimeAccuracyProblem                SessionRejectReasonEnum
	TagAppearsMoreThanOnce                    SessionRejectReasonEnum
	TagSpecifiedOutOfRequiredOrder            SessionRejectReasonEnum
	RepeatingGroupFieldsOutOfOrder            SessionRejectReasonEnum
	IncorrectNumInGroupCountForRepeatingGroup SessionRejectReasonEnum
	Other                                     SessionRejectReasonEnum
	NonRepresentable                          SessionRejectReasonEnum
	NullValue                                 SessionRejectReasonEnum
}

var SessionRejectReason = SessionRejectReasonValues{0, 1, 2, 3, 5, 6, 8, 10, 13, 14, 15, 16, 99, 254, 255}

func (s SessionRejectReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(s)); err != nil {
		return err
	}
	return nil
}

func (s *SessionRejectReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(s)); err != nil {
		return err
	}
	return nil
}

func (s SessionRejectReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(SessionRejectReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on SessionRejectReason, unknown enumeration value %d", s)
}

func (*SessionRejectReasonEnum) EncodedLength() int64 {
	return 1
}

func (*SessionRejectReasonEnum) InvalidTagNumberSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) InvalidTagNumberInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.InvalidTagNumberSinceVersion()
}

func (*SessionRejectReasonEnum) InvalidTagNumberDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) RequiredTagMissingSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) RequiredTagMissingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RequiredTagMissingSinceVersion()
}

func (*SessionRejectReasonEnum) RequiredTagMissingDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) TagNotDefinedForThisMessageTypeSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) TagNotDefinedForThisMessageTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TagNotDefinedForThisMessageTypeSinceVersion()
}

func (*SessionRejectReasonEnum) TagNotDefinedForThisMessageTypeDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) UndefinedTagSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) UndefinedTagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.UndefinedTagSinceVersion()
}

func (*SessionRejectReasonEnum) UndefinedTagDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) ValueIsIncorrectSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) ValueIsIncorrectInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.ValueIsIncorrectSinceVersion()
}

func (*SessionRejectReasonEnum) ValueIsIncorrectDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) IncorrectDataFormatForValueSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) IncorrectDataFormatForValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.IncorrectDataFormatForValueSinceVersion()
}

func (*SessionRejectReasonEnum) IncorrectDataFormatForValueDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) SignatureProblemSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) SignatureProblemInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SignatureProblemSinceVersion()
}

func (*SessionRejectReasonEnum) SignatureProblemDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) SendingTimeAccuracyProblemSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) SendingTimeAccuracyProblemInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SendingTimeAccuracyProblemSinceVersion()
}

func (*SessionRejectReasonEnum) SendingTimeAccuracyProblemDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) TagAppearsMoreThanOnceSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) TagAppearsMoreThanOnceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TagAppearsMoreThanOnceSinceVersion()
}

func (*SessionRejectReasonEnum) TagAppearsMoreThanOnceDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) TagSpecifiedOutOfRequiredOrderSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) TagSpecifiedOutOfRequiredOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TagSpecifiedOutOfRequiredOrderSinceVersion()
}

func (*SessionRejectReasonEnum) TagSpecifiedOutOfRequiredOrderDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) RepeatingGroupFieldsOutOfOrderSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) RepeatingGroupFieldsOutOfOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.RepeatingGroupFieldsOutOfOrderSinceVersion()
}

func (*SessionRejectReasonEnum) RepeatingGroupFieldsOutOfOrderDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) IncorrectNumInGroupCountForRepeatingGroupSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) IncorrectNumInGroupCountForRepeatingGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.IncorrectNumInGroupCountForRepeatingGroupSinceVersion()
}

func (*SessionRejectReasonEnum) IncorrectNumInGroupCountForRepeatingGroupDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) OtherSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) OtherInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OtherSinceVersion()
}

func (*SessionRejectReasonEnum) OtherDeprecated() uint16 {
	return 0
}

func (*SessionRejectReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (s *SessionRejectReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NonRepresentableSinceVersion()
}

func (*SessionRejectReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
