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

type MarketDataSnapshot struct {
	LastBookUpdateID int64
	PriceExponent    int8
	QtyExponent      int8
	MDEntriesBids    []MarketDataSnapshotMDEntriesBids
	MDEntriesAsks    []MarketDataSnapshotMDEntriesAsks
	Symbol           []uint8
}
type MarketDataSnapshotMDEntriesBids struct {
	MDEntryPx   int64
	MDEntrySize int64
}
type MarketDataSnapshotMDEntriesAsks struct {
	MDEntryPx   int64
	MDEntrySize int64
}

func (m *MarketDataSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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

func (m *MarketDataSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
			m.MDEntriesBids = make([]MarketDataSnapshotMDEntriesBids, MDEntriesBidsNumInGroup)
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
			m.MDEntriesAsks = make([]MarketDataSnapshotMDEntriesAsks, MDEntriesAsksNumInGroup)
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

func (m *MarketDataSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.LastBookUpdateIDInActingVersion(actingVersion) {
		if m.LastBookUpdateID != m.LastBookUpdateIDNullValue() && (m.LastBookUpdateID < m.LastBookUpdateIDMinValue() || m.LastBookUpdateID > m.LastBookUpdateIDMaxValue()) {
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

func MarketDataSnapshotInit(m *MarketDataSnapshot) {
	m.LastBookUpdateID = math.MinInt64
	return
}

func (m *MarketDataSnapshotMDEntriesBids) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataSnapshotMDEntriesBids) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MarketDataSnapshotMDEntriesBids) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func MarketDataSnapshotMDEntriesBidsInit(m *MarketDataSnapshotMDEntriesBids) {
	return
}

func (m *MarketDataSnapshotMDEntriesAsks) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt64(_w, m.MDEntryPx); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, m.MDEntrySize); err != nil {
		return err
	}
	return nil
}

func (m *MarketDataSnapshotMDEntriesAsks) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MarketDataSnapshotMDEntriesAsks) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	return nil
}

func MarketDataSnapshotMDEntriesAsksInit(m *MarketDataSnapshotMDEntriesAsks) {
	return
}

func (*MarketDataSnapshot) SbeBlockLength() (blockLength uint16) {
	return 10
}

func (*MarketDataSnapshot) SbeTemplateId() (templateId uint16) {
	return 204
}

func (*MarketDataSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketDataSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("W")
}

func (*MarketDataSnapshot) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketDataSnapshot) LastBookUpdateIDId() uint16 {
	return 25044
}

func (*MarketDataSnapshot) LastBookUpdateIDSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) LastBookUpdateIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.LastBookUpdateIDSinceVersion()
}

func (*MarketDataSnapshot) LastBookUpdateIDDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshot) LastBookUpdateIDMetaAttribute(meta int) string {
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

func (*MarketDataSnapshot) LastBookUpdateIDMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshot) LastBookUpdateIDMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshot) LastBookUpdateIDNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshot) PriceExponentId() uint16 {
	return 25054
}

func (*MarketDataSnapshot) PriceExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceExponentSinceVersion()
}

func (*MarketDataSnapshot) PriceExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshot) PriceExponentMetaAttribute(meta int) string {
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

func (*MarketDataSnapshot) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataSnapshot) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataSnapshot) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataSnapshot) QtyExponentId() uint16 {
	return 25055
}

func (*MarketDataSnapshot) QtyExponentSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QtyExponentSinceVersion()
}

func (*MarketDataSnapshot) QtyExponentDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshot) QtyExponentMetaAttribute(meta int) string {
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

func (*MarketDataSnapshot) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*MarketDataSnapshot) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*MarketDataSnapshot) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotMDEntriesBids) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotMDEntriesBids) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotMDEntriesBids) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotMDEntriesBids) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxId() uint16 {
	return 270
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotMDEntriesAsks) MDEntryPxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntryPxSinceVersion()
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntryPxNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeId() uint16 {
	return 271
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshotMDEntriesAsks) MDEntrySizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntrySizeSinceVersion()
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeMetaAttribute(meta int) string {
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

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeMaxValue() int64 {
	return math.MaxInt64
}

func (*MarketDataSnapshotMDEntriesAsks) MDEntrySizeNullValue() int64 {
	return math.MinInt64
}

func (*MarketDataSnapshot) MDEntriesBidsId() uint16 {
	return 25047
}

func (*MarketDataSnapshot) MDEntriesBidsSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) MDEntriesBidsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesBidsSinceVersion()
}

func (*MarketDataSnapshot) MDEntriesBidsDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesBids) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataSnapshotMDEntriesBids) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataSnapshot) MDEntriesAsksId() uint16 {
	return 25048
}

func (*MarketDataSnapshot) MDEntriesAsksSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) MDEntriesAsksInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MDEntriesAsksSinceVersion()
}

func (*MarketDataSnapshot) MDEntriesAsksDeprecated() uint16 {
	return 0
}

func (*MarketDataSnapshotMDEntriesAsks) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*MarketDataSnapshotMDEntriesAsks) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*MarketDataSnapshot) SymbolMetaAttribute(meta int) string {
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

func (*MarketDataSnapshot) SymbolSinceVersion() uint16 {
	return 0
}

func (m *MarketDataSnapshot) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.SymbolSinceVersion()
}

func (*MarketDataSnapshot) SymbolDeprecated() uint16 {
	return 0
}

func (MarketDataSnapshot) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (MarketDataSnapshot) SymbolHeaderLength() uint64 {
	return 1
}
