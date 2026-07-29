// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrdStatusEnum byte
type OrdStatusValues struct {
	New              OrdStatusEnum
	PartiallyFilled  OrdStatusEnum
	Filled           OrdStatusEnum
	Canceled         OrdStatusEnum
	PendingCancel    OrdStatusEnum
	Rejected         OrdStatusEnum
	PendingNew       OrdStatusEnum
	Expired          OrdStatusEnum
	NonRepresentable OrdStatusEnum
	NullValue        OrdStatusEnum
}

var OrdStatus = OrdStatusValues{48, 49, 50, 52, 54, 56, 65, 67, 126, 0}

func (o OrdStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdStatus, unknown enumeration value %d", o)
}

func (*OrdStatusEnum) EncodedLength() int64 {
	return 1
}

func (*OrdStatusEnum) NewSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) NewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NewSinceVersion()
}

func (*OrdStatusEnum) NewDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) PartiallyFilledSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) PartiallyFilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PartiallyFilledSinceVersion()
}

func (*OrdStatusEnum) PartiallyFilledDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) FilledSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) FilledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.FilledSinceVersion()
}

func (*OrdStatusEnum) FilledDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) CanceledSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) CanceledInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CanceledSinceVersion()
}

func (*OrdStatusEnum) CanceledDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) PendingCancelSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) PendingCancelInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PendingCancelSinceVersion()
}

func (*OrdStatusEnum) PendingCancelDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) RejectedSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) RejectedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RejectedSinceVersion()
}

func (*OrdStatusEnum) RejectedDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) PendingNewSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) PendingNewInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PendingNewSinceVersion()
}

func (*OrdStatusEnum) PendingNewDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) ExpiredSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) ExpiredInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpiredSinceVersion()
}

func (*OrdStatusEnum) ExpiredDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NonRepresentableSinceVersion()
}

func (*OrdStatusEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
