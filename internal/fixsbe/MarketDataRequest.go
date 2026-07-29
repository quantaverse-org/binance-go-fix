// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type MarketDataRequest struct {
	SubscriptionRequestType SubscriptionRequestTypeEnum
	MarketDepth             uint16
	AggregatedBook          BoolEnumEnum
	RelatedSym              []MarketDataRequestRelatedSym
	MDEntryTypes            []MarketDataRequestMDEntryTypes
	MDReqID                 []uint8
}
type MarketDataRequestRelatedSym struct {
	Symbol []uint8
}
type MarketDataRequestMDEntryTypes struct {
	MDEntryType MdEntryTypeEnum
}

func (m *MarketDataRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := m.SubscriptionRequestType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.MarketDepth); err != nil {
		return err
	}
	if err := m.AggregatedBook.Encode(_m, _w); err != nil {
		return err
	}
	var RelatedSymBlockLength uint16 = 0
	if err := _m.WriteUint16(_w, RelatedSymBlockLength); err != nil {
		return err
	}
	var RelatedSymNumInGroup uint16 = uint16(len(m.RelatedSym))
	if err := _m.WriteUint16(_w, RelatedSymNumInGroup); err != nil {
		return err
	}
	for i := range m.RelatedSym {
		if err := m.RelatedSym[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	var MDEntryTypesBlockLength uint8 = 1
	if err := _m.WriteUint8(_w, MDEntryTypesBlockLength); err != nil {
		return err
	}
	var MDEntryTypesNumInGroup uint8 = uint8(len(m.MDEntryTypes))
	if err := _m.WriteUint8(_w, MDEntryTypesNumInGroup); err != nil {
		return err
	}
	for i := range m.MDEntryTypes {
		if err := m.MDEntryTypes[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(m.MDReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.MDReqID); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if m.SubscriptionRequestTypeInActingVersion(actingVersion) {
		if err := m.SubscriptionRequestType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !m.MarketDepthInActingVersion(actingVersion) {
		m.MarketDepth = m.MarketDepthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.MarketDepth); err != nil {
			return err
		}
	}
	if m.AggregatedBookInActingVersion(actingVersion) {
		if err := m.AggregatedBook.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.RelatedSymInActingVersion(actingVersion) {
		var RelatedSymBlockLength uint16
		if err := _m.ReadUint16(_r, &RelatedSymBlockLength); err != nil {
			return err
		}
		var RelatedSymNumInGroup uint16
		if err := _m.ReadUint16(_r, &RelatedSymNumInGroup); err != nil {
			return err
		}
		if cap(m.RelatedSym) < int(RelatedSymNumInGroup) {
			m.RelatedSym = make([]MarketDataRequestRelatedSym, RelatedSymNumInGroup)
		}
		m.RelatedSym = m.RelatedSym[:RelatedSymNumInGroup]
		for i := range m.RelatedSym {
			if err := m.RelatedSym[i].Decode(_m, _r, actingVersion, uint(RelatedSymBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.MDEntryTypesInActingVersion(actingVersion) {
		var MDEntryTypesBlockLength uint8
		if err := _m.ReadUint8(_r, &MDEntryTypesBlockLength); err != nil {
			return err
		}
		var MDEntryTypesNumInGroup uint8
		if err := _m.ReadUint8(_r, &MDEntryTypesNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntryTypes) < int(MDEntryTypesNumInGroup) {
			m.MDEntryTypes = make([]MarketDataRequestMDEntryTypes, MDEntryTypesNumInGroup)
		}
		m.MDEntryTypes = m.MDEntryTypes[:MDEntryTypesNumInGroup]
		for i := range m.MDEntryTypes {
			if err := m.MDEntryTypes[i].Decode(_m, _r, actingVersion, uint(MDEntryTypesBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.MDReqIDInActingVersion(actingVersion) {
		var MDReqIDLength uint8
		if err := _m.ReadUint8(_r, &MDReqIDLength); err != nil {
			return err
		}
		if cap(m.MDReqID) < int(MDReqIDLength) {
			m.MDReqID = make([]uint8, MDReqIDLength)
		}
		m.MDReqID = m.MDReqID[:MDReqIDLength]
		if err := _m.ReadBytes(_r, m.MDReqID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := m.SubscriptionRequestType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if m.MarketDepthInActingVersion(actingVersion) {
		if m.MarketDepth != m.MarketDepthNullValue() && (m.MarketDepth < m.MarketDepthMinValue() || m.MarketDepth > m.MarketDepthMaxValue()) {
			return fmt.Errorf("Range check failed on m.MarketDepth (%v < %v > %v)", m.MarketDepthMinValue(), m.MarketDepth, m.MarketDepthMaxValue())
		}
	}
	if err := m.AggregatedBook.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	for i := range m.RelatedSym {
		if err := m.RelatedSym[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for i := range m.MDEntryTypes {
		if err := m.MDEntryTypes[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(m.MDReqID[:]) {
		return errors.New("m.MDReqID failed UTF-8 validation")
	}
	return nil
}

func MarketDataRequestInit(m *MarketDataRequest) {
	m.MarketDepth = math.MaxUint16
	return
}

func (m *MarketDataRequestRelatedSym) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, uint8(len(m.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Symbol); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataRequestRelatedSym) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(m.Symbol) < int(SymbolLength) {
			m.Symbol = make([]uint8, SymbolLength)
		}
		m.Symbol = m.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, m.Symbol); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataRequestRelatedSym) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(m.Symbol[:]) {
		return errors.New("m.Symbol failed UTF-8 validation")
	}
	return nil
}

func MarketDataRequestRelatedSymInit(m *MarketDataRequestRelatedSym) {
	return
}

func (m *MarketDataRequestMDEntryTypes) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := m.MDEntryType.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataRequestMDEntryTypes) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if m.MDEntryTypeInActingVersion(actingVersion) {
		if err := m.MDEntryType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataRequestMDEntryTypes) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := m.MDEntryType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MarketDataRequestMDEntryTypesInit(m *MarketDataRequestMDEntryTypes) {
	return
}

func (*MarketDataRequest) SbeBlockLength() (blockLength uint16) {
	return 4
}

func (*MarketDataRequest) SbeTemplateId() (templateId uint16) {
	return 202
}

func (*MarketDataRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("V")
}

func (*MarketDataRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataRequest) SubscriptionRequestTypeId() uint16 {
	return 263
}

func (*MarketDataRequest) SubscriptionRequestTypeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) SubscriptionRequestTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SubscriptionRequestTypeSinceVersion()
}

func (*MarketDataRequest) SubscriptionRequestTypeDeprecated() uint16 {
	return 0
}

func (*MarketDataRequest) SubscriptionRequestTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataRequest) MarketDepthId() uint16 {
	return 264
}

func (*MarketDataRequest) MarketDepthSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) MarketDepthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketDepthSinceVersion()
}

func (*MarketDataRequest) MarketDepthDeprecated() uint16 {
	return 0
}

func (*MarketDataRequest) MarketDepthMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*MarketDataRequest) MarketDepthMinValue() uint16 {
	return 0
}

func (*MarketDataRequest) MarketDepthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketDataRequest) MarketDepthNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketDataRequest) AggregatedBookId() uint16 {
	return 266
}

func (*MarketDataRequest) AggregatedBookSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) AggregatedBookInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggregatedBookSinceVersion()
}

func (*MarketDataRequest) AggregatedBookDeprecated() uint16 {
	return 0
}

func (*MarketDataRequest) AggregatedBookMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "optional"
	}
	return ""
}

func (*MarketDataRequestRelatedSym) SymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataRequestRelatedSym) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestRelatedSym) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MarketDataRequestRelatedSym) SymbolDeprecated() uint16 {
	return 0
}

func (MarketDataRequestRelatedSym) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataRequestRelatedSym) SymbolHeaderLength() uint64 {
	return 1
}

func (*MarketDataRequestMDEntryTypes) MDEntryTypeId() uint16 {
	return 269
}

func (*MarketDataRequestMDEntryTypes) MDEntryTypeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequestMDEntryTypes) MDEntryTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypeSinceVersion()
}

func (*MarketDataRequestMDEntryTypes) MDEntryTypeDeprecated() uint16 {
	return 0
}

func (*MarketDataRequestMDEntryTypes) MDEntryTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataRequest) RelatedSymId() uint16 {
	return 146
}

func (*MarketDataRequest) RelatedSymSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) RelatedSymInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.RelatedSymSinceVersion()
}

func (*MarketDataRequest) RelatedSymDeprecated() uint16 {
	return 0
}

func (*MarketDataRequestRelatedSym) SbeBlockLength() (blockLength uint) {
	return 0
}

func (*MarketDataRequestRelatedSym) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataRequest) MDEntryTypesId() uint16 {
	return 267
}

func (*MarketDataRequest) MDEntryTypesSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) MDEntryTypesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryTypesSinceVersion()
}

func (*MarketDataRequest) MDEntryTypesDeprecated() uint16 {
	return 0
}

func (*MarketDataRequestMDEntryTypes) SbeBlockLength() (blockLength uint) {
	return 1
}

func (*MarketDataRequestMDEntryTypes) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataRequest) MDReqIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*MarketDataRequest) MDReqIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataRequest) MDReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDReqIDSinceVersion()
}

func (*MarketDataRequest) MDReqIDDeprecated() uint16 {
	return 0
}

func (MarketDataRequest) MDReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataRequest) MDReqIDHeaderLength() uint64 {
	return 1
}
