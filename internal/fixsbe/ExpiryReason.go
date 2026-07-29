// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type ExpiryReasonEnum uint8
type ExpiryReasonValues struct {
	Rejected                        ExpiryReasonEnum
	ExchangeCanceled                ExpiryReasonEnum
	OcoTrigger                      ExpiryReasonEnum
	OtoPhaseOneExpired              ExpiryReasonEnum
	UnfilledIocQuantityExpired      ExpiryReasonEnum
	UnfilledFokOrderExpired         ExpiryReasonEnum
	InsufficientLiquidity           ExpiryReasonEnum
	ExecutionRulePriceRangeExceeded ExpiryReasonEnum
	NonRepresentable                ExpiryReasonEnum
	NullValue                       ExpiryReasonEnum
}

var ExpiryReason = ExpiryReasonValues{1, 2, 3, 4, 5, 6, 7, 8, 254, 255}

func (e ExpiryReasonEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(e)); err != nil {
		return err
	}
	return nil
}

func (e *ExpiryReasonEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*uint8)(e)); err != nil {
		return err
	}
	return nil
}

func (e ExpiryReasonEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(ExpiryReason)
	for idx := 0; idx < value.NumField(); idx++ {
		if e == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on ExpiryReason, unknown enumeration value %d", e)
}

func (*ExpiryReasonEnum) EncodedLength() int64 {
	return 1
}

func (*ExpiryReasonEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.RejectedSinceVersion()
}

func (*ExpiryReasonEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) ExchangeCanceledSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) ExchangeCanceledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExchangeCanceledSinceVersion()
}

func (*ExpiryReasonEnum) ExchangeCanceledDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) OcoTriggerSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) OcoTriggerInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OcoTriggerSinceVersion()
}

func (*ExpiryReasonEnum) OcoTriggerDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) OtoPhaseOneExpiredSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) OtoPhaseOneExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.OtoPhaseOneExpiredSinceVersion()
}

func (*ExpiryReasonEnum) OtoPhaseOneExpiredDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) UnfilledIocQuantityExpiredSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) UnfilledIocQuantityExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnfilledIocQuantityExpiredSinceVersion()
}

func (*ExpiryReasonEnum) UnfilledIocQuantityExpiredDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) UnfilledFokOrderExpiredSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) UnfilledFokOrderExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.UnfilledFokOrderExpiredSinceVersion()
}

func (*ExpiryReasonEnum) UnfilledFokOrderExpiredDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) InsufficientLiquiditySinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) InsufficientLiquidityInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.InsufficientLiquiditySinceVersion()
}

func (*ExpiryReasonEnum) InsufficientLiquidityDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) ExecutionRulePriceRangeExceededSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) ExecutionRulePriceRangeExceededInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.ExecutionRulePriceRangeExceededSinceVersion()
}

func (*ExpiryReasonEnum) ExecutionRulePriceRangeExceededDeprecated() uint16 {
	return 0
}

func (*ExpiryReasonEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (e *ExpiryReasonEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= e.NonRepresentableSinceVersion()
}

func (*ExpiryReasonEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
