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

type MarketDataIncrementalBookTicker struct {
	LastBookUpdateID int64
	PriceExponent    int8
	QtyExponent      int8
	MDEntriesBids    []MarketDataIncrementalBookTickerMDEntriesBids
	MDEntriesAsks    []MarketDataIncrementalBookTickerMDEntriesAsks
	Symbol           []uint8
}
type MarketDataIncrementalBookTickerMDEntriesBids struct {
	MDEntryPx   int64
	MDEntrySize int64
}
type MarketDataIncrementalBookTickerMDEntriesAsks struct {
	MDEntryPx   int64
	MDEntrySize int64
}

func (m *MarketDataIncrementalBookTicker) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.LastBookUpdateID); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, m.QtyExponent); err != nil {
		return err
	}
	var MDEntriesBidsBlockLength uint8 = 16
	if err := _m.WriteUint8(_w, MDEntriesBidsBlockLength); err != nil {
		return err
	}
	var MDEntriesBidsNumInGroup uint8 = uint8(len(m.MDEntriesBids))
	if err := _m.WriteUint8(_w, MDEntriesBidsNumInGroup); err != nil {
		return err
	}
	for i := range m.MDEntriesBids {
		if err := m.MDEntriesBids[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	var MDEntriesAsksBlockLength uint8 = 16
	if err := _m.WriteUint8(_w, MDEntriesAsksBlockLength); err != nil {
		return err
	}
	var MDEntriesAsksNumInGroup uint8 = uint8(len(m.MDEntriesAsks))
	if err := _m.WriteUint8(_w, MDEntriesAsksNumInGroup); err != nil {
		return err
	}
	for i := range m.MDEntriesAsks {
		if err := m.MDEntriesAsks[i].Encode(_m, _w); err != nil {
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

func (m *MarketDataIncrementalBookTicker) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.LastBookUpdateIDInActingVersion(actingVersion) {
		m.LastBookUpdateID = m.LastBookUpdateIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.LastBookUpdateID); err != nil {
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

	if m.MDEntriesBidsInActingVersion(actingVersion) {
		var MDEntriesBidsBlockLength uint8
		if err := _m.ReadUint8(_r, &MDEntriesBidsBlockLength); err != nil {
			return err
		}
		var MDEntriesBidsNumInGroup uint8
		if err := _m.ReadUint8(_r, &MDEntriesBidsNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntriesBids) < int(MDEntriesBidsNumInGroup) {
			m.MDEntriesBids = make([]MarketDataIncrementalBookTickerMDEntriesBids, MDEntriesBidsNumInGroup)
		}
		m.MDEntriesBids = m.MDEntriesBids[:MDEntriesBidsNumInGroup]
		for i := range m.MDEntriesBids {
			if err := m.MDEntriesBids[i].Decode(_m, _r, actingVersion, uint(MDEntriesBidsBlockLength)); err != nil {
				return err
			}
		}
	}

	if m.MDEntriesAsksInActingVersion(actingVersion) {
		var MDEntriesAsksBlockLength uint8
		if err := _m.ReadUint8(_r, &MDEntriesAsksBlockLength); err != nil {
			return err
		}
		var MDEntriesAsksNumInGroup uint8
		if err := _m.ReadUint8(_r, &MDEntriesAsksNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntriesAsks) < int(MDEntriesAsksNumInGroup) {
			m.MDEntriesAsks = make([]MarketDataIncrementalBookTickerMDEntriesAsks, MDEntriesAsksNumInGroup)
		}
		m.MDEntriesAsks = m.MDEntriesAsks[:MDEntriesAsksNumInGroup]
		for i := range m.MDEntriesAsks {
			if err := m.MDEntriesAsks[i].Decode(_m, _r, actingVersion, uint(MDEntriesAsksBlockLength)); err != nil {
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

func (m *MarketDataIncrementalBookTicker) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LastBookUpdateIDInActingVersion(actingVersion) {
		if m.LastBookUpdateID < m.LastBookUpdateIDMinValue() || m.LastBookUpdateID > m.LastBookUpdateIDMaxValue() {
			return fmt.Errorf("Range check failed on m.LastBookUpdateID (%v < %v > %v)", m.LastBookUpdateIDMinValue(), m.LastBookUpdateID, m.LastBookUpdateIDMaxValue())
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
	for i := range m.MDEntriesBids {
		if err := m.MDEntriesBids[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	for i := range m.MDEntriesAsks {
		if err := m.MDEntriesAsks[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(m.Symbol[:]) {
		return errors.New("m.Symbol failed UTF-8 validation")
	}
	return nil
}

func MarketDataIncrementalBookTickerInit(m *MarketDataIncrementalBookTicker) {
	return
}

func (m *MarketDataIncrementalBookTickerMDEntriesBids) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalBookTickerMDEntriesBids) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataIncrementalBookTickerMDEntriesBids) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MDEntryPxInActingVersion(actingVersion) {
		if m.MDEntryPx < m.MDEntryPxMinValue() || m.MDEntryPx > m.MDEntryPxMaxValue() {
			return fmt.Errorf("Range check failed on m.MDEntryPx (%v < %v > %v)", m.MDEntryPxMinValue(), m.MDEntryPx, m.MDEntryPxMaxValue())
		}
	}
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize != m.MDEntrySizeNullValue() && (m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDEntrySize (%v < %v > %v)", m.MDEntrySizeMinValue(), m.MDEntrySize, m.MDEntrySizeMaxValue())
		}
	}
	return nil
}

func MarketDataIncrementalBookTickerMDEntriesBidsInit(m *MarketDataIncrementalBookTickerMDEntriesBids) {
	m.MDEntrySize = math.MinInt64
	return
}

func (m *MarketDataIncrementalBookTickerMDEntriesAsks) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalBookTickerMDEntriesAsks) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketDataIncrementalBookTickerMDEntriesAsks) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MDEntryPxInActingVersion(actingVersion) {
		if m.MDEntryPx < m.MDEntryPxMinValue() || m.MDEntryPx > m.MDEntryPxMaxValue() {
			return fmt.Errorf("Range check failed on m.MDEntryPx (%v < %v > %v)", m.MDEntryPxMinValue(), m.MDEntryPx, m.MDEntryPxMaxValue())
		}
	}
	if m.MDEntrySizeInActingVersion(actingVersion) {
		if m.MDEntrySize != m.MDEntrySizeNullValue() && (m.MDEntrySize < m.MDEntrySizeMinValue() || m.MDEntrySize > m.MDEntrySizeMaxValue()) {
			return fmt.Errorf("Range check failed on m.MDEntrySize (%v < %v > %v)", m.MDEntrySizeMinValue(), m.MDEntrySize, m.MDEntrySizeMaxValue())
		}
	}
	return nil
}

func MarketDataIncrementalBookTickerMDEntriesAsksInit(m *MarketDataIncrementalBookTickerMDEntriesAsks) {
	m.MDEntrySize = math.MinInt64
	return
}

func (*MarketDataIncrementalBookTicker) SbeBlockLength() (blockLength uint16) {
	return 10
}

func (*MarketDataIncrementalBookTicker) SbeTemplateId() (templateId uint16) {
	return 206
}

func (*MarketDataIncrementalBookTicker) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataIncrementalBookTicker) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalBookTicker) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MarketDataIncrementalBookTicker) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDId() uint16 {
	return 25044
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) LastBookUpdateIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastBookUpdateIDSinceVersion()
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalBookTicker) LastBookUpdateIDNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalBookTicker) PriceExponentId() uint16 {
	return 25054
}

func (*MarketDataIncrementalBookTicker) PriceExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceExponentSinceVersion()
}

func (*MarketDataIncrementalBookTicker) PriceExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTicker) PriceExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTicker) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalBookTicker) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalBookTicker) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalBookTicker) QtyExponentId() uint16 {
	return 25055
}

func (*MarketDataIncrementalBookTicker) QtyExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QtyExponentSinceVersion()
}

func (*MarketDataIncrementalBookTicker) QtyExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTicker) QtyExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTicker) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalBookTicker) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalBookTicker) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalBookTicker) MDEntriesBidsId() uint16 {
	return 25047
}

func (*MarketDataIncrementalBookTicker) MDEntriesBidsSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) MDEntriesBidsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesBidsSinceVersion()
}

func (*MarketDataIncrementalBookTicker) MDEntriesBidsDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataIncrementalBookTickerMDEntriesBids) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalBookTicker) MDEntriesAsksId() uint16 {
	return 25048
}

func (*MarketDataIncrementalBookTicker) MDEntriesAsksSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) MDEntriesAsksInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesAsksSinceVersion()
}

func (*MarketDataIncrementalBookTicker) MDEntriesAsksDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataIncrementalBookTickerMDEntriesAsks) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalBookTicker) SymbolMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalBookTicker) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalBookTicker) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MarketDataIncrementalBookTicker) SymbolDeprecated() uint16 {
	return 0
}

func (MarketDataIncrementalBookTicker) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataIncrementalBookTicker) SymbolHeaderLength() uint64 {
	return 1
}
