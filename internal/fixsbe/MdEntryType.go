// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"reflect"
)

type MdEntryTypeEnum byte
type MdEntryTypeValues struct {
	Bid              MdEntryTypeEnum
	Offer            MdEntryTypeEnum
	Trade            MdEntryTypeEnum
	NonRepresentable MdEntryTypeEnum
	NullValue        MdEntryTypeEnum
}

var MdEntryType = MdEntryTypeValues{48, 49, 50, 126, 0}

func (m MdEntryTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(m)); err != nil {
		return err
	}
	return nil
}

func (m *MdEntryTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(m)); err != nil {
		return err
	}
	return nil
}

func (m MdEntryTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(MdEntryType)
	for idx := 0; idx < value.NumField(); idx++ {
		if m == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on MdEntryType, unknown enumeration value %d", m)
}

func (*MdEntryTypeEnum) EncodedLength() int64 {
	return 1
}

func (*MdEntryTypeEnum) BidSinceVersion() uint16 {
	return 0
}

func (m *MdEntryTypeEnum) BidInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSinceVersion()
}

func (*MdEntryTypeEnum) BidDeprecated() uint16 {
	return 0
}

func (*MdEntryTypeEnum) OfferSinceVersion() uint16 {
	return 0
}

func (m *MdEntryTypeEnum) OfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OfferSinceVersion()
}

func (*MdEntryTypeEnum) OfferDeprecated() uint16 {
	return 0
}

func (*MdEntryTypeEnum) TradeSinceVersion() uint16 {
	return 0
}

func (m *MdEntryTypeEnum) TradeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeSinceVersion()
}

func (*MdEntryTypeEnum) TradeDeprecated() uint16 {
	return 0
}

func (*MdEntryTypeEnum) NonRepresentableSinceVersion() uint16 {
	return 0
}

func (m *MdEntryTypeEnum) NonRepresentableInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.NonRepresentableSinceVersion()
}

func (*MdEntryTypeEnum) NonRepresentableDeprecated() uint16 {
	return 0
}
