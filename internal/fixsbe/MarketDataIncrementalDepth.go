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

type MarketDataIncrementalDepth struct {
	FirstBookUpdateID int64
	LastBookUpdateID  int64
	PriceExponent     int8
	QtyExponent       int8
	MDEntriesBids     []MarketDataIncrementalDepthMDEntriesBids
	MDEntriesAsks     []MarketDataIncrementalDepthMDEntriesAsks
	Symbol            []uint8
}
type MarketDataIncrementalDepthMDEntriesBids struct {
	MDEntryPx   int64
	MDEntrySize int64
}
type MarketDataIncrementalDepthMDEntriesAsks struct {
	MDEntryPx   int64
	MDEntrySize int64
}

func (m *MarketDataIncrementalDepth) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, m.FirstBookUpdateID); err != nil {
		return err
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
	var MDEntriesBidsNumInGroup uint16 = uint16(len(m.MDEntriesBids))
	if err := _m.WriteUint16(_w, MDEntriesBidsNumInGroup); err != nil {
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
	var MDEntriesAsksNumInGroup uint16 = uint16(len(m.MDEntriesAsks))
	if err := _m.WriteUint16(_w, MDEntriesAsksNumInGroup); err != nil {
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

func (m *MarketDataIncrementalDepth) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.FirstBookUpdateIDInActingVersion(actingVersion) {
		m.FirstBookUpdateID = m.FirstBookUpdateIDNullValue()
	} else {
		if err := _m.ReadInt64(_r, &m.FirstBookUpdateID); err != nil {
			return err
		}
	}
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
		var MDEntriesBidsNumInGroup uint16
		if err := _m.ReadUint16(_r, &MDEntriesBidsNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntriesBids) < int(MDEntriesBidsNumInGroup) {
			m.MDEntriesBids = make([]MarketDataIncrementalDepthMDEntriesBids, MDEntriesBidsNumInGroup)
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
		var MDEntriesAsksNumInGroup uint16
		if err := _m.ReadUint16(_r, &MDEntriesAsksNumInGroup); err != nil {
			return err
		}
		if cap(m.MDEntriesAsks) < int(MDEntriesAsksNumInGroup) {
			m.MDEntriesAsks = make([]MarketDataIncrementalDepthMDEntriesAsks, MDEntriesAsksNumInGroup)
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

func (m *MarketDataIncrementalDepth) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.FirstBookUpdateIDInActingVersion(actingVersion) {
		if m.FirstBookUpdateID < m.FirstBookUpdateIDMinValue() || m.FirstBookUpdateID > m.FirstBookUpdateIDMaxValue() {
			return fmt.Errorf("Range check failed on m.FirstBookUpdateID (%v < %v > %v)", m.FirstBookUpdateIDMinValue(), m.FirstBookUpdateID, m.FirstBookUpdateIDMaxValue())
		}
	}
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

func MarketDataIncrementalDepthInit(m *MarketDataIncrementalDepth) {
	return
}

func (m *MarketDataIncrementalDepthMDEntriesBids) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalDepthMDEntriesBids) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MarketDataIncrementalDepthMDEntriesBids) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MarketDataIncrementalDepthMDEntriesBidsInit(m *MarketDataIncrementalDepthMDEntriesBids) {
	m.MDEntrySize = math.MinInt64
	return
}

func (m *MarketDataIncrementalDepthMDEntriesAsks) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataIncrementalDepthMDEntriesAsks) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MarketDataIncrementalDepthMDEntriesAsks) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MarketDataIncrementalDepthMDEntriesAsksInit(m *MarketDataIncrementalDepthMDEntriesAsks) {
	m.MDEntrySize = math.MinInt64
	return
}

func (*MarketDataIncrementalDepth) SbeBlockLength() (blockLength uint16) {
	return 18
}

func (*MarketDataIncrementalDepth) SbeTemplateId() (templateId uint16) {
	return 207
}

func (*MarketDataIncrementalDepth) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataIncrementalDepth) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalDepth) SbeSemanticType() (semanticType []byte) {
	return []byte("X")
}

func (*MarketDataIncrementalDepth) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDId() uint16 {
	return 25043
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) FirstBookUpdateIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.FirstBookUpdateIDSinceVersion()
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepth) FirstBookUpdateIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepth) FirstBookUpdateIDNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDId() uint16 {
	return 25044
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) LastBookUpdateIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastBookUpdateIDSinceVersion()
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepth) LastBookUpdateIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepth) LastBookUpdateIDNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepth) PriceExponentId() uint16 {
	return 25054
}

func (*MarketDataIncrementalDepth) PriceExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceExponentSinceVersion()
}

func (*MarketDataIncrementalDepth) PriceExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepth) PriceExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepth) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalDepth) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalDepth) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalDepth) QtyExponentId() uint16 {
	return 25055
}

func (*MarketDataIncrementalDepth) QtyExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QtyExponentSinceVersion()
}

func (*MarketDataIncrementalDepth) QtyExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepth) QtyExponentMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepth) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataIncrementalDepth) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataIncrementalDepth) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepthMDEntriesBids) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepthMDEntriesBids) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataIncrementalDepthMDEntriesAsks) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataIncrementalDepth) MDEntriesBidsId() uint16 {
	return 25047
}

func (*MarketDataIncrementalDepth) MDEntriesBidsSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) MDEntriesBidsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesBidsSinceVersion()
}

func (*MarketDataIncrementalDepth) MDEntriesBidsDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesBids) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataIncrementalDepthMDEntriesBids) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalDepth) MDEntriesAsksId() uint16 {
	return 25048
}

func (*MarketDataIncrementalDepth) MDEntriesAsksSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) MDEntriesAsksInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesAsksSinceVersion()
}

func (*MarketDataIncrementalDepth) MDEntriesAsksDeprecated() uint16 {
	return 0
}

func (*MarketDataIncrementalDepthMDEntriesAsks) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataIncrementalDepthMDEntriesAsks) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataIncrementalDepth) SymbolMetaAttribute(meta int) string {
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

func (*MarketDataIncrementalDepth) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MarketDataIncrementalDepth) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MarketDataIncrementalDepth) SymbolDeprecated() uint16 {
	return 0
}

func (MarketDataIncrementalDepth) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataIncrementalDepth) SymbolHeaderLength() uint64 {
	return 1
}
