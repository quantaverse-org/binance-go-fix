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

type MarketDataIncrementalTrade struct {
	TransactTime  int64
	PriceExponent int8
	QtyExponent   int8
	MDEntries     []MarketDataIncrementalTradeMDEntries
	Symbol        []uint8
}
type MarketDataIncrementalTradeMDEntries struct {
	TradeID       int64
	MDEntryPx     int64
	MDEntrySize   int64
	AggressorSide SideEnum
}

func (m *MarketDataIncrementalTrade) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.QtyExponent); err != nil {
		return err
	}
	var MDEntriesBlockLength uint16 = 25
	if err := _m.WriteUint16(_w, MDEntriesBlockLength); err != nil {
		return err
	}
	var MDEntriesNumInGroup uint32 = uint32(len(m.MDEntries))
	if err := _m.WriteUint32(_w, MDEntriesNumInGroup); err != nil {
		return err
	}
	for i := range m.MDEntries {
		if err := m.MDEntries[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(m.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, m.Symbol); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalTrade) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.TransactTimeInActingVersion(actingVersion) {
		m.TransactTime = m.TransactTimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.TransactTime); err != nil {
			return err
		}
	}
	if !m.PriceExponentInActingVersion(actingVersion) {
		m.PriceExponent = m.PriceExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.PriceExponent); err != nil {
			return err
		}
	}
	if !m.QtyExponentInActingVersion(actingVersion) {
		m.QtyExponent = m.QtyExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &m.QtyExponent); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.MDEntriesInActingVersion(actingVersion) {
		var MDEntriesBlockLength uint16
		if err := _m.ReadUint16(_r, &MDEntriesBlockLength); err != nil {
			return err
		}
		var MDEntriesNumInGroup uint32
		if err := _m.ReadUint32(_r, &MDEntriesNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntries) < int(MDEntriesNumInGroup) {
			m.MDEntries = make([]MarketDataIncrementalTradeMDEntries, MDEntriesNumInGroup)
		}
		m.MDEntries = m.MDEntries[:MDEntriesNumInGroup]
		for i := range m.MDEntries {
			if err := m.MDEntries[i].Decode(_m, _r, actingVersion, uint(MDEntriesBlockLength)); err != nil {
				return err
			}
		}
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
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketDataIncrementalTrade) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TransactTimeInActingVersion(actingVersion) {
		if m.TransactTime < m.TransactTimeMinValue() || m.TransactTime > m.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on m.TransactTime (%v < %v > %v)", m.TransactTimeMinValue(), m.TransactTime, m.TransactTimeMaxValue())
		}
	}
	if m.PriceExponentInActingVersion(actingVersion) {
		if m.PriceExponent < m.PriceExponentMinValue() || m.PriceExponent > m.PriceExponentMaxValue() {
			return fmt.Errorf("Range check failed on m.PriceExponent (%v < %v > %v)", m.PriceExponentMinValue(), m.PriceExponent, m.PriceExponentMaxValue())
		}
	}
	if m.QtyExponentInActingVersion(actingVersion) {
		if m.QtyExponent < m.QtyExponentMinValue() || m.QtyExponent > m.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on m.QtyExponent (%v < %v > %v)", m.QtyExponentMinValue(), m.QtyExponent, m.QtyExponentMaxValue())
		}
	}
	for i := range m.MDEntries {
		if err := m.MDEntries[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(m.Symbol[:]) {
		return errors.New("m.Symbol failed UTF-8 validation")
	}
	return nil
}

func MarketDataIncrementalTradeInit(m *MarketDataIncrementalTrade) {
	return
}

func (m *MarketDataIncrementalTradeMDEntries) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.TradeID); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	if err := m.AggressorSide.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalTradeMDEntries) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.TradeIDInActingVersion(actingVersion) {
		m.TradeID = m.TradeIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.TradeID); err != nil {
			return err
		}
	}
	if !m.MDEntryPxInActingVersion(actingVersion) {
		m.MDEntryPx = m.MDEntryPxNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MDEntryPx); err != nil {
			return err
		}
	}
	if !m.MDEntrySizeInActingVersion(actingVersion) {
		m.MDEntrySize = m.MDEntrySizeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.MDEntrySize); err != nil {
			return err
		}
	}
	if m.AggressorSideInActingVersion(actingVersion) {
		if err := m.AggressorSide.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataIncrementalTradeMDEntries) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.TradeIDInActingVersion(actingVersion) {
		if m.TradeID < m.TradeIDMinValue() || m.TradeID > m.TradeIDMaxValue() {
			return fmt.Errorf("Range check failed on m.TradeID (%v < %v > %v)", m.TradeIDMinValue(), m.TradeID, m.TradeIDMaxValue())
		}
	}
	if m.MDEntryPxInActingVersion(actingVersion) {
		if m.MDEntryPx < m.MDEntryPxMinValue() || m.MDEntryPx > m.MDEntryPxMaxValue() {
			return fmt.Errorf("Range check failed on m.MDEntryPx (%v < %v > %v)", m.MDEntryPxMinValue(), m.MDEntryPx, m.MDEntryPxMaxValue())
		}
	}
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue() {
			return fmt.Errorf("Range check failed on m.MDEntrySize (%v < %v > %v)", m.MDEntrySizeMinValue(), m.MDEntrySize, m.MDEntrySizeMaxValue())
		}
	}
	if err := m.AggressorSide.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func MarketDataIncrementalTradeMDEntriesInit(m *MarketDataIncrementalTradeMDEntries) {
	return
}

func (*MarketDataIncrementalTrade) SbeBlockLength() (blockLength uint16) {
	return 10
}

func (*MarketDataIncrementalTrade) SbeTemplateId() (templateId uint16) {
	return 205
}

func (*MarketDataIncrementalTrade) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataIncrementalTrade) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalTrade) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MarketDataIncrementalTrade) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataIncrementalTrade) TransactTimeId() uint16 {
	return 60
}

func (*MarketDataIncrementalTrade) TransactTimeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTrade) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TransactTimeSinceVersion()
}

func (*MarketDataIncrementalTrade) TransactTimeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTrade) TransactTimeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTrade) TransactTimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalTrade) TransactTimeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalTrade) TransactTimeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalTrade) PriceExponentId() uint16 {
	return 25054
}

func (*MarketDataIncrementalTrade) PriceExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTrade) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceExponentSinceVersion()
}

func (*MarketDataIncrementalTrade) PriceExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTrade) PriceExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTrade) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalTrade) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalTrade) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalTrade) QtyExponentId() uint16 {
	return 25055
}

func (*MarketDataIncrementalTrade) QtyExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTrade) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QtyExponentSinceVersion()
}

func (*MarketDataIncrementalTrade) QtyExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTrade) QtyExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTrade) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalTrade) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalTrade) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDId() uint16 {
	return 1003
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTradeMDEntries) TradeIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TradeIDSinceVersion()
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTradeMDEntries) TradeIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalTradeMDEntries) TradeIDNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTradeMDEntries) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalTradeMDEntries) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTradeMDEntries) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalTradeMDEntries) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalTradeMDEntries) AggressorSideId() uint16 {
	return 2446
}

func (*MarketDataIncrementalTradeMDEntries) AggressorSideSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTradeMDEntries) AggressorSideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AggressorSideSinceVersion()
}

func (*MarketDataIncrementalTradeMDEntries) AggressorSideDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTradeMDEntries) AggressorSideMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTrade) MDEntriesId() uint16 {
	return 268
}

func (*MarketDataIncrementalTrade) MDEntriesSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTrade) MDEntriesInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesSinceVersion()
}

func (*MarketDataIncrementalTrade) MDEntriesDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalTradeMDEntries) SbeBlockLength() (blockLength uint) {
	return 25
}

func (*MarketDataIncrementalTradeMDEntries) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalTrade) SymbolMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalTrade) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalTrade) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MarketDataIncrementalTrade) SymbolDeprecated() uint16 {
	return 0
}

func (MarketDataIncrementalTrade) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataIncrementalTrade) SymbolHeaderLength() uint64 {
	return 1
}
