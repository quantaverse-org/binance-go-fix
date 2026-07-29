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

type LimitResponse struct {
	LimitIndicators []LimitResponseLimitIndicators
	ReqID           []uint8
}
type LimitResponseLimitIndicators struct {
	LimitType                    LimitTypeEnum
	LimitCount                   uint64
	LimitMax                     uint64
	LimitResetInterval           uint32
	LimitResetIntervalResolution LimitResetIntervalResolutionEnum
}

func (l *LimitResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	var LimitIndicatorsBlockLength uint8 = 22
	if err := _m.WriteUint8(_w, LimitIndicatorsBlockLength); err != nil {
		return err
	}
	var LimitIndicatorsNumInGroup uint8 = uint8(len(l.LimitIndicators))
	if err := _m.WriteUint8(_w, LimitIndicatorsNumInGroup); err != nil {
		return err
	}
	for i := range l.LimitIndicators {
		if err := l.LimitIndicators[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, uint8(len(l.ReqID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ReqID); err != nil {
		return err
	}
	return nil
}

func (l *LimitResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.LimitIndicatorsInActingVersion(actingVersion) {
		var LimitIndicatorsBlockLength uint8
		if err := _m.ReadUint8(_r, &LimitIndicatorsBlockLength); err != nil {
			return err
		}
		var LimitIndicatorsNumInGroup uint8
		if err := _m.ReadUint8(_r, &LimitIndicatorsNumInGroup); err != nil {
			return err
		}
		if cap(l.LimitIndicators) < int(LimitIndicatorsNumInGroup) {
			l.LimitIndicators = make([]LimitResponseLimitIndicators, LimitIndicatorsNumInGroup)
		}
		l.LimitIndicators = l.LimitIndicators[:LimitIndicatorsNumInGroup]
		for i := range l.LimitIndicators {
			if err := l.LimitIndicators[i].Decode(_m, _r, actingVersion, uint(LimitIndicatorsBlockLength)); err != nil {
				return err
			}
		}
	}

	if l.ReqIDInActingVersion(actingVersion) {
		var ReqIDLength uint8
		if err := _m.ReadUint8(_r, &ReqIDLength); err != nil {
			return err
		}
		if cap(l.ReqID) < int(ReqIDLength) {
			l.ReqID = make([]uint8, ReqIDLength)
		}
		l.ReqID = l.ReqID[:ReqIDLength]
		if err := _m.ReadBytes(_r, l.ReqID); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *LimitResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	for i := range l.LimitIndicators {
		if err := l.LimitIndicators[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(l.ReqID[:]) {
		return errors.New("l.ReqID failed UTF-8 validation")
	}
	return nil
}

func LimitResponseInit(l *LimitResponse) {
	return
}

func (l *LimitResponseLimitIndicators) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := l.LimitType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, l.LimitCount); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, l.LimitMax); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, l.LimitResetInterval); err != nil {
		return err
	}
	if err := l.LimitResetIntervalResolution.Encode(_m, _w); err != nil {
		return err
	}
	return nil
}

func (l *LimitResponseLimitIndicators) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if l.LimitTypeInActingVersion(actingVersion) {
		if err := l.LimitType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.LimitCountInActingVersion(actingVersion) {
		l.LimitCount = l.LimitCountNullValue()
	} else {
		if err := _m.ReadUint64(_r, &l.LimitCount); err != nil {
			return err
		}
	}
	if !l.LimitMaxInActingVersion(actingVersion) {
		l.LimitMax = l.LimitMaxNullValue()
	} else {
		if err := _m.ReadUint64(_r, &l.LimitMax); err != nil {
			return err
		}
	}
	if !l.LimitResetIntervalInActingVersion(actingVersion) {
		l.LimitResetInterval = l.LimitResetIntervalNullValue()
	} else {
		if err := _m.ReadUint32(_r, &l.LimitResetInterval); err != nil {
			return err
		}
	}
	if l.LimitResetIntervalResolutionInActingVersion(actingVersion) {
		if err := l.LimitResetIntervalResolution.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}
	return nil
}

func (l *LimitResponseLimitIndicators) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if err := l.LimitType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if l.LimitCountInActingVersion(actingVersion) {
		if l.LimitCount < l.LimitCountMinValue() || l.LimitCount > l.LimitCountMaxValue() {
			return fmt.Errorf("Range check failed on l.LimitCount (%v < %v > %v)", l.LimitCountMinValue(), l.LimitCount, l.LimitCountMaxValue())
		}
	}
	if l.LimitMaxInActingVersion(actingVersion) {
		if l.LimitMax < l.LimitMaxMinValue() || l.LimitMax > l.LimitMaxMaxValue() {
			return fmt.Errorf("Range check failed on l.LimitMax (%v < %v > %v)", l.LimitMaxMinValue(), l.LimitMax, l.LimitMaxMaxValue())
		}
	}
	if l.LimitResetIntervalInActingVersion(actingVersion) {
		if l.LimitResetInterval != l.LimitResetIntervalNullValue() && (l.LimitResetInterval < l.LimitResetIntervalMinValue() || l.LimitResetInterval > l.LimitResetIntervalMaxValue()) {
			return fmt.Errorf("Range check failed on l.LimitResetInterval (%v < %v > %v)", l.LimitResetIntervalMinValue(), l.LimitResetInterval, l.LimitResetIntervalMaxValue())
		}
	}
	if err := l.LimitResetIntervalResolution.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	return nil
}

func LimitResponseLimitIndicatorsInit(l *LimitResponseLimitIndicators) {
	l.LimitResetInterval = math.MaxUint32
	return
}

func (*LimitResponse) SbeBlockLength() (blockLength uint16) {
	return 0
}

func (*LimitResponse) SbeTemplateId() (templateId uint16) {
	return 121
}

func (*LimitResponse) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LimitResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*LimitResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("XLR")
}

func (*LimitResponse) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LimitResponseLimitIndicators) LimitTypeId() uint16 {
	return 25004
}

func (*LimitResponseLimitIndicators) LimitTypeSinceVersion() uint16 {
	return 0
}

func (l *LimitResponseLimitIndicators) LimitTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitTypeSinceVersion()
}

func (*LimitResponseLimitIndicators) LimitTypeDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitTypeMetaAttribute(meta int) string {
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

func (*LimitResponseLimitIndicators) LimitCountId() uint16 {
	return 25005
}

func (*LimitResponseLimitIndicators) LimitCountSinceVersion() uint16 {
	return 0
}

func (l *LimitResponseLimitIndicators) LimitCountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitCountSinceVersion()
}

func (*LimitResponseLimitIndicators) LimitCountDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitCountMetaAttribute(meta int) string {
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

func (*LimitResponseLimitIndicators) LimitCountMinValue() uint64 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitCountMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*LimitResponseLimitIndicators) LimitCountNullValue() uint64 {
	return math.MaxUint64
}

func (*LimitResponseLimitIndicators) LimitMaxId() uint16 {
	return 25006
}

func (*LimitResponseLimitIndicators) LimitMaxSinceVersion() uint16 {
	return 0
}

func (l *LimitResponseLimitIndicators) LimitMaxInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitMaxSinceVersion()
}

func (*LimitResponseLimitIndicators) LimitMaxDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitMaxMetaAttribute(meta int) string {
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

func (*LimitResponseLimitIndicators) LimitMaxMinValue() uint64 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitMaxMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*LimitResponseLimitIndicators) LimitMaxNullValue() uint64 {
	return math.MaxUint64
}

func (*LimitResponseLimitIndicators) LimitResetIntervalId() uint16 {
	return 25007
}

func (*LimitResponseLimitIndicators) LimitResetIntervalSinceVersion() uint16 {
	return 0
}

func (l *LimitResponseLimitIndicators) LimitResetIntervalInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitResetIntervalSinceVersion()
}

func (*LimitResponseLimitIndicators) LimitResetIntervalDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitResetIntervalMetaAttribute(meta int) string {
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

func (*LimitResponseLimitIndicators) LimitResetIntervalMinValue() uint32 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitResetIntervalMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*LimitResponseLimitIndicators) LimitResetIntervalNullValue() uint32 {
	return math.MaxUint32
}

func (*LimitResponseLimitIndicators) LimitResetIntervalResolutionId() uint16 {
	return 25008
}

func (*LimitResponseLimitIndicators) LimitResetIntervalResolutionSinceVersion() uint16 {
	return 0
}

func (l *LimitResponseLimitIndicators) LimitResetIntervalResolutionInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitResetIntervalResolutionSinceVersion()
}

func (*LimitResponseLimitIndicators) LimitResetIntervalResolutionDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) LimitResetIntervalResolutionMetaAttribute(meta int) string {
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

func (*LimitResponse) LimitIndicatorsId() uint16 {
	return 25003
}

func (*LimitResponse) LimitIndicatorsSinceVersion() uint16 {
	return 0
}

func (l *LimitResponse) LimitIndicatorsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.LimitIndicatorsSinceVersion()
}

func (*LimitResponse) LimitIndicatorsDeprecated() uint16 {
	return 0
}

func (*LimitResponseLimitIndicators) SbeBlockLength() (blockLength uint) {
	return 22
}

func (*LimitResponseLimitIndicators) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*LimitResponse) ReqIDMetaAttribute(meta int) string {
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

func (*LimitResponse) ReqIDSinceVersion() uint16 {
	return 0
}

func (l *LimitResponse) ReqIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ReqIDSinceVersion()
}

func (*LimitResponse) ReqIDDeprecated() uint16 {
	return 0
}

func (LimitResponse) ReqIDCharacterEncoding() string {
	return "UTF-8"
}

func (LimitResponse) ReqIDHeaderLength() uint64 {
	return 1
}
