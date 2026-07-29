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

type InstrumentList struct {
	RelatedSym      []InstrumentListRelatedSym
	InstrumentReqID []uint8
}
type InstrumentListRelatedSym struct {
	PriceExponent         int8
	QtyExponent           int8
	MinTradeVol           int64
	MaxTradeVol           int64
	MinQtyIncrement       int64
	MarketMinTradeVol     int64
	MarketMaxTradeVol     int64
	MarketMinQtyIncrement int64
	StartPriceRange       int64
	EndPriceRange         int64
	MinPriceIncrement     int64
	Symbol                []uint8
	Currency              []uint8
}

func (i *InstrumentList) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := i.RangeCheck(i.SbeSchemaVersion(), i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	var RelatedSymBlockLength uint16 = 74
	if err := _m.WriteUint16(_w, RelatedSymBlockLength); err != nil {
		return err
	}
	var RelatedSymNumInGroup uint16 = uint16(len(i.RelatedSym))
	if err := _m.WriteUint16(_w, RelatedSymNumInGroup); err != nil {
		return err
	}
	for idx := range i.RelatedSym {
		if err := i.RelatedSym[idx].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(i.InstrumentReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.InstrumentReqID); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}

	if i.RelatedSymInActingVersion(actingVersion) {
		var RelatedSymBlockLength uint16
		if err := _m.ReadUint16(_r, &RelatedSymBlockLength); err != nil {
			return err
		}
		var RelatedSymNumInGroup uint16
		if err := _m.ReadUint16(_r, &RelatedSymNumInGroup); err != nil {
			return err
		}
		if cap(i.RelatedSym) < int(RelatedSymNumInGroup) {
			i.RelatedSym = make([]InstrumentListRelatedSym, RelatedSymNumInGroup)
		}
		i.RelatedSym = i.RelatedSym[:RelatedSymNumInGroup]
		for idx := range i.RelatedSym {
			if err := i.RelatedSym[idx].Decode(_m, _r, actingVersion, uint(RelatedSymBlockLength)); err != nil {
				return err
			}
		}
	}

	if i.InstrumentReqIDInActingVersion(actingVersion) {
		var InstrumentReqIDLength uint8
		if err := _m.ReadUint8(_r, &InstrumentReqIDLength); err != nil {
			return err
		}
		if cap(i.InstrumentReqID) < int(InstrumentReqIDLength) {
			i.InstrumentReqID = make([]uint8, InstrumentReqIDLength)
		}
		i.InstrumentReqID = i.InstrumentReqID[:InstrumentReqIDLength]
		if err := _m.ReadBytes(_r, i.InstrumentReqID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := i.RangeCheck(actingVersion, i.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (i *InstrumentList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	for idx := range i.RelatedSym {
		if err := i.RelatedSym[idx].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(i.InstrumentReqID[:]) {
		return errors.New("i.InstrumentReqID failed UTF-8 validation")
	}
	return nil
}

func InstrumentListInit(i *InstrumentList) {
	return
}

func (i *InstrumentListRelatedSym) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt8(_w, i.PriceExponent); err != nil {
		return err
	}
	if err := _m.WriteInt8(_w, i.QtyExponent); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MinTradeVol); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MaxTradeVol); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MinQtyIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MarketMinTradeVol); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MarketMaxTradeVol); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MarketMinQtyIncrement); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.StartPriceRange); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.EndPriceRange); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, i.MinPriceIncrement); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(i.Symbol))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.Symbol); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(i.Currency))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, i.Currency); err != nil {
		return err
	}
	return nil
}

func (i *InstrumentListRelatedSym) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !i.PriceExponentInActingVersion(actingVersion) {
		i.PriceExponent = i.PriceExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &i.PriceExponent); err != nil {
			return err
		}
	}
	if !i.QtyExponentInActingVersion(actingVersion) {
		i.QtyExponent = i.QtyExponentNullValue()
	} else {
		if err := _m.ReadInt8(_r, &i.QtyExponent); err != nil {
			return err
		}
	}
	if !i.MinTradeVolInActingVersion(actingVersion) {
		i.MinTradeVol = i.MinTradeVolNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MinTradeVol); err != nil {
			return err
		}
	}
	if !i.MaxTradeVolInActingVersion(actingVersion) {
		i.MaxTradeVol = i.MaxTradeVolNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MaxTradeVol); err != nil {
			return err
		}
	}
	if !i.MinQtyIncrementInActingVersion(actingVersion) {
		i.MinQtyIncrement = i.MinQtyIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MinQtyIncrement); err != nil {
			return err
		}
	}
	if !i.MarketMinTradeVolInActingVersion(actingVersion) {
		i.MarketMinTradeVol = i.MarketMinTradeVolNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MarketMinTradeVol); err != nil {
			return err
		}
	}
	if !i.MarketMaxTradeVolInActingVersion(actingVersion) {
		i.MarketMaxTradeVol = i.MarketMaxTradeVolNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MarketMaxTradeVol); err != nil {
			return err
		}
	}
	if !i.MarketMinQtyIncrementInActingVersion(actingVersion) {
		i.MarketMinQtyIncrement = i.MarketMinQtyIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MarketMinQtyIncrement); err != nil {
			return err
		}
	}
	if !i.StartPriceRangeInActingVersion(actingVersion) {
		i.StartPriceRange = i.StartPriceRangeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.StartPriceRange); err != nil {
			return err
		}
	}
	if !i.EndPriceRangeInActingVersion(actingVersion) {
		i.EndPriceRange = i.EndPriceRangeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.EndPriceRange); err != nil {
			return err
		}
	}
	if !i.MinPriceIncrementInActingVersion(actingVersion) {
		i.MinPriceIncrement = i.MinPriceIncrementNullValue()
	} else {
		if err := _m.ReadInt64(_r, &i.MinPriceIncrement); err != nil {
			return err
		}
	}
	if actingVersion > i.SbeSchemaVersion() && blockLength > i.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-i.SbeBlockLength()))
	}

	if i.SymbolInActingVersion(actingVersion) {
		var SymbolLength uint8
		if err := _m.ReadUint8(_r, &SymbolLength); err != nil {
			return err
		}
		if cap(i.Symbol) < int(SymbolLength) {
			i.Symbol = make([]uint8, SymbolLength)
		}
		i.Symbol = i.Symbol[:SymbolLength]
		if err := _m.ReadBytes(_r, i.Symbol); err != nil {
			return err
		}
	}

	if i.CurrencyInActingVersion(actingVersion) {
		var CurrencyLength uint8
		if err := _m.ReadUint8(_r, &CurrencyLength); err != nil {
			return err
		}
		if cap(i.Currency) < int(CurrencyLength) {
			i.Currency = make([]uint8, CurrencyLength)
		}
		i.Currency = i.Currency[:CurrencyLength]
		if err := _m.ReadBytes(_r, i.Currency); err != nil {
			return err
		}
	}
	return nil
}

func (i *InstrumentListRelatedSym) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if i.PriceExponentInActingVersion(actingVersion) {
		if i.PriceExponent < i.PriceExponentMinValue() || i.PriceExponent > i.PriceExponentMaxValue() {
			return fmt.Errorf("Range check failed on i.PriceExponent (%v < %v > %v)", i.PriceExponentMinValue(), i.PriceExponent, i.PriceExponentMaxValue())
		}
	}
	if i.QtyExponentInActingVersion(actingVersion) {
		if i.QtyExponent < i.QtyExponentMinValue() || i.QtyExponent > i.QtyExponentMaxValue() {
			return fmt.Errorf("Range check failed on i.QtyExponent (%v < %v > %v)", i.QtyExponentMinValue(), i.QtyExponent, i.QtyExponentMaxValue())
		}
	}
	if i.MinTradeVolInActingVersion(actingVersion) {
		if i.MinTradeVol != i.MinTradeVolNullValue() && (i.MinTradeVol < i.MinTradeVolMinValue() || i.MinTradeVol > i.MinTradeVolMaxValue()) {
			return fmt.Errorf("Range check failed on i.MinTradeVol (%v < %v > %v)", i.MinTradeVolMinValue(), i.MinTradeVol, i.MinTradeVolMaxValue())
		}
	}
	if i.MaxTradeVolInActingVersion(actingVersion) {
		if i.MaxTradeVol != i.MaxTradeVolNullValue() && (i.MaxTradeVol < i.MaxTradeVolMinValue() || i.MaxTradeVol > i.MaxTradeVolMaxValue()) {
			return fmt.Errorf("Range check failed on i.MaxTradeVol (%v < %v > %v)", i.MaxTradeVolMinValue(), i.MaxTradeVol, i.MaxTradeVolMaxValue())
		}
	}
	if i.MinQtyIncrementInActingVersion(actingVersion) {
		if i.MinQtyIncrement != i.MinQtyIncrementNullValue() && (i.MinQtyIncrement < i.MinQtyIncrementMinValue() || i.MinQtyIncrement > i.MinQtyIncrementMaxValue()) {
			return fmt.Errorf("Range check failed on i.MinQtyIncrement (%v < %v > %v)", i.MinQtyIncrementMinValue(), i.MinQtyIncrement, i.MinQtyIncrementMaxValue())
		}
	}
	if i.MarketMinTradeVolInActingVersion(actingVersion) {
		if i.MarketMinTradeVol != i.MarketMinTradeVolNullValue() && (i.MarketMinTradeVol < i.MarketMinTradeVolMinValue() || i.MarketMinTradeVol > i.MarketMinTradeVolMaxValue()) {
			return fmt.Errorf("Range check failed on i.MarketMinTradeVol (%v < %v > %v)", i.MarketMinTradeVolMinValue(), i.MarketMinTradeVol, i.MarketMinTradeVolMaxValue())
		}
	}
	if i.MarketMaxTradeVolInActingVersion(actingVersion) {
		if i.MarketMaxTradeVol != i.MarketMaxTradeVolNullValue() && (i.MarketMaxTradeVol < i.MarketMaxTradeVolMinValue() || i.MarketMaxTradeVol > i.MarketMaxTradeVolMaxValue()) {
			return fmt.Errorf("Range check failed on i.MarketMaxTradeVol (%v < %v > %v)", i.MarketMaxTradeVolMinValue(), i.MarketMaxTradeVol, i.MarketMaxTradeVolMaxValue())
		}
	}
	if i.MarketMinQtyIncrementInActingVersion(actingVersion) {
		if i.MarketMinQtyIncrement != i.MarketMinQtyIncrementNullValue() && (i.MarketMinQtyIncrement < i.MarketMinQtyIncrementMinValue() || i.MarketMinQtyIncrement > i.MarketMinQtyIncrementMaxValue()) {
			return fmt.Errorf("Range check failed on i.MarketMinQtyIncrement (%v < %v > %v)", i.MarketMinQtyIncrementMinValue(), i.MarketMinQtyIncrement, i.MarketMinQtyIncrementMaxValue())
		}
	}
	if i.StartPriceRangeInActingVersion(actingVersion) {
		if i.StartPriceRange != i.StartPriceRangeNullValue() && (i.StartPriceRange < i.StartPriceRangeMinValue() || i.StartPriceRange > i.StartPriceRangeMaxValue()) {
			return fmt.Errorf("Range check failed on i.StartPriceRange (%v < %v > %v)", i.StartPriceRangeMinValue(), i.StartPriceRange, i.StartPriceRangeMaxValue())
		}
	}
	if i.EndPriceRangeInActingVersion(actingVersion) {
		if i.EndPriceRange != i.EndPriceRangeNullValue() && (i.EndPriceRange < i.EndPriceRangeMinValue() || i.EndPriceRange > i.EndPriceRangeMaxValue()) {
			return fmt.Errorf("Range check failed on i.EndPriceRange (%v < %v > %v)", i.EndPriceRangeMinValue(), i.EndPriceRange, i.EndPriceRangeMaxValue())
		}
	}
	if i.MinPriceIncrementInActingVersion(actingVersion) {
		if i.MinPriceIncrement != i.MinPriceIncrementNullValue() && (i.MinPriceIncrement < i.MinPriceIncrementMinValue() || i.MinPriceIncrement > i.MinPriceIncrementMaxValue()) {
			return fmt.Errorf("Range check failed on i.MinPriceIncrement (%v < %v > %v)", i.MinPriceIncrementMinValue(), i.MinPriceIncrement, i.MinPriceIncrementMaxValue())
		}
	}
	if !utf8.Valid(i.Symbol[:]) {
		return errors.New("i.Symbol failed UTF-8 validation")
	}
	if !utf8.Valid(i.Currency[:]) {
		return errors.New("i.Currency failed UTF-8 validation")
	}
	return nil
}

func InstrumentListRelatedSymInit(i *InstrumentListRelatedSym) {
	i.MinTradeVol = math.MinInt64
	i.MaxTradeVol = math.MinInt64
	i.MinQtyIncrement = math.MinInt64
	i.MarketMinTradeVol = math.MinInt64
	i.MarketMaxTradeVol = math.MinInt64
	i.MarketMinQtyIncrement = math.MinInt64
	i.StartPriceRange = math.MinInt64
	i.EndPriceRange = math.MinInt64
	i.MinPriceIncrement = math.MinInt64
	return
}

func (*InstrumentList) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*InstrumentList) SbeTemplateId() (templateId uint16) {
	return 201
}

func (*InstrumentList) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*InstrumentList) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*InstrumentList) SbeSemanticType() (semanticType []byte) {
	return []byte("y")
}

func (*InstrumentList) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*InstrumentListRelatedSym) PriceExponentId() uint16 {
	return 25054
}

func (*InstrumentListRelatedSym) PriceExponentSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) PriceExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.PriceExponentSinceVersion()
}

func (*InstrumentListRelatedSym) PriceExponentDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) PriceExponentMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) PriceExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*InstrumentListRelatedSym) PriceExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*InstrumentListRelatedSym) PriceExponentNullValue() int8 {
	return math.MinInt8
}

func (*InstrumentListRelatedSym) QtyExponentId() uint16 {
	return 25055
}

func (*InstrumentListRelatedSym) QtyExponentSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) QtyExponentInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.QtyExponentSinceVersion()
}

func (*InstrumentListRelatedSym) QtyExponentDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) QtyExponentMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) QtyExponentMinValue() int8 {
	return math.MinInt8 + 1
}

func (*InstrumentListRelatedSym) QtyExponentMaxValue() int8 {
	return math.MaxInt8
}

func (*InstrumentListRelatedSym) QtyExponentNullValue() int8 {
	return math.MinInt8
}

func (*InstrumentListRelatedSym) MinTradeVolId() uint16 {
	return 562
}

func (*InstrumentListRelatedSym) MinTradeVolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MinTradeVolSinceVersion()
}

func (*InstrumentListRelatedSym) MinTradeVolDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MinTradeVolMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MinTradeVolMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MinTradeVolMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MinTradeVolNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MaxTradeVolId() uint16 {
	return 1140
}

func (*InstrumentListRelatedSym) MaxTradeVolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MaxTradeVolSinceVersion()
}

func (*InstrumentListRelatedSym) MaxTradeVolDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MaxTradeVolMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MaxTradeVolMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MaxTradeVolMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MaxTradeVolNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MinQtyIncrementId() uint16 {
	return 25039
}

func (*InstrumentListRelatedSym) MinQtyIncrementSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MinQtyIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MinQtyIncrementSinceVersion()
}

func (*InstrumentListRelatedSym) MinQtyIncrementDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MinQtyIncrementMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MinQtyIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MinQtyIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MinQtyIncrementNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MarketMinTradeVolId() uint16 {
	return 25040
}

func (*InstrumentListRelatedSym) MarketMinTradeVolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MarketMinTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MarketMinTradeVolSinceVersion()
}

func (*InstrumentListRelatedSym) MarketMinTradeVolDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MarketMinTradeVolMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MarketMinTradeVolMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MarketMinTradeVolMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MarketMinTradeVolNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolId() uint16 {
	return 25041
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MarketMaxTradeVolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MarketMaxTradeVolSinceVersion()
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MarketMaxTradeVolMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MarketMaxTradeVolNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementId() uint16 {
	return 25042
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MarketMinQtyIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MarketMinQtyIncrementSinceVersion()
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MarketMinQtyIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MarketMinQtyIncrementNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) StartPriceRangeId() uint16 {
	return 2551
}

func (*InstrumentListRelatedSym) StartPriceRangeSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) StartPriceRangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.StartPriceRangeSinceVersion()
}

func (*InstrumentListRelatedSym) StartPriceRangeDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) StartPriceRangeMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) StartPriceRangeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) StartPriceRangeMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) StartPriceRangeNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) EndPriceRangeId() uint16 {
	return 2552
}

func (*InstrumentListRelatedSym) EndPriceRangeSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) EndPriceRangeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.EndPriceRangeSinceVersion()
}

func (*InstrumentListRelatedSym) EndPriceRangeDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) EndPriceRangeMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) EndPriceRangeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) EndPriceRangeMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) EndPriceRangeNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) MinPriceIncrementId() uint16 {
	return 969
}

func (*InstrumentListRelatedSym) MinPriceIncrementSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) MinPriceIncrementInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.MinPriceIncrementSinceVersion()
}

func (*InstrumentListRelatedSym) MinPriceIncrementDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) MinPriceIncrementMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) MinPriceIncrementMinValue() int64 {
	return math.MinInt64 + 1
}

func (*InstrumentListRelatedSym) MinPriceIncrementMaxValue() int64 {
	return math.MaxInt64
}

func (*InstrumentListRelatedSym) MinPriceIncrementNullValue() int64 {
	return math.MinInt64
}

func (*InstrumentListRelatedSym) SymbolMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) SymbolSinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.SymbolSinceVersion()
}

func (*InstrumentListRelatedSym) SymbolDeprecated() uint16 {
	return 0
}

func (InstrumentListRelatedSym) SymbolCharacterEncoding() string {
	return "UTF-8"
}

func (InstrumentListRelatedSym) SymbolHeaderLength() uint64 {
	return 1
}

func (*InstrumentListRelatedSym) CurrencyMetaAttribute(meta int) string {
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

func (*InstrumentListRelatedSym) CurrencySinceVersion() uint16 {
	return 0
}

func (i *InstrumentListRelatedSym) CurrencyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.CurrencySinceVersion()
}

func (*InstrumentListRelatedSym) CurrencyDeprecated() uint16 {
	return 0
}

func (InstrumentListRelatedSym) CurrencyCharacterEncoding() string {
	return "UTF-8"
}

func (InstrumentListRelatedSym) CurrencyHeaderLength() uint64 {
	return 1
}

func (*InstrumentList) RelatedSymId() uint16 {
	return 146
}

func (*InstrumentList) RelatedSymSinceVersion() uint16 {
	return 0
}

func (i *InstrumentList) RelatedSymInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.RelatedSymSinceVersion()
}

func (*InstrumentList) RelatedSymDeprecated() uint16 {
	return 0
}

func (*InstrumentListRelatedSym) SbeBlockLength() (blockLength uint) {
	return 74
}

func (*InstrumentListRelatedSym) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*InstrumentList) InstrumentReqIDMetaAttribute(meta int) string {
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

func (*InstrumentList) InstrumentReqIDSinceVersion() uint16 {
	return 0
}

func (i *InstrumentList) InstrumentReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= i.InstrumentReqIDSinceVersion()
}

func (*InstrumentList) InstrumentReqIDDeprecated() uint16 {
	return 0
}

func (InstrumentList) InstrumentReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (InstrumentList) InstrumentReqIDHeaderLength() uint64 {
	return 1
}
