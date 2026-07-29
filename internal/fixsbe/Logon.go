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

type Logon struct {
	EncryptMethod       uint8
	HeartBtInt          uint32
	ResetSeqNumFlag     BoolEnumEnum
	MessageHandling     MessageHandlingEnum
	ResponseMode        ResponseModeEnum
	ExecutionReportType ExecutionReportTypeEnum
	DropCopyFlag        BoolEnumEnum
	RecvWindow          uint32
	SenderCompId        []uint8
	TargetCompId        []uint8
	RawData             []uint8
	Username            []uint8
}

func (l *Logon) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, l.EncryptMethod); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, l.HeartBtInt); err != nil {
		return err
	}
	if err := l.ResetSeqNumFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.MessageHandling.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.ResponseMode.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.ExecutionReportType.Encode(_m, _w); err != nil {
		return err
	}
	if err := l.DropCopyFlag.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, l.RecvWindow); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(l.SenderCompId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.SenderCompId); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(l.TargetCompId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.TargetCompId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.RawData))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.RawData); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.Username))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.Username); err != nil {
		return err
	}
	return nil
}

func (l *Logon) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.EncryptMethodInActingVersion(actingVersion) {
		l.EncryptMethod = l.EncryptMethodNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.EncryptMethod); err != nil {
			return err
		}
	}
	if !l.HeartBtIntInActingVersion(actingVersion) {
		l.HeartBtInt = l.HeartBtIntNullValue()
	} else {
		if err := _m.ReadUint32(_r, &l.HeartBtInt); err != nil {
			return err
		}
	}
	if l.ResetSeqNumFlagInActingVersion(actingVersion) {
		if err := l.ResetSeqNumFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.MessageHandlingInActingVersion(actingVersion) {
		if err := l.MessageHandling.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.ResponseModeInActingVersion(actingVersion) {
		if err := l.ResponseMode.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.ExecutionReportTypeInActingVersion(actingVersion) {
		if err := l.ExecutionReportType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if l.DropCopyFlagInActingVersion(actingVersion) {
		if err := l.DropCopyFlag.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !l.RecvWindowInActingVersion(actingVersion) {
		l.RecvWindow = l.RecvWindowNullValue()
	} else {
		if err := _m.ReadUint32(_r, &l.RecvWindow); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.SenderCompIdInActingVersion(actingVersion) {
		var SenderCompIdLength uint8
		if err := _m.ReadUint8(_r, &SenderCompIdLength); err != nil {
			return err
		}
		if cap(l.SenderCompId) < int(SenderCompIdLength) {
			l.SenderCompId = make([]uint8, SenderCompIdLength)
		}
		l.SenderCompId = l.SenderCompId[:SenderCompIdLength]
		if err := _m.ReadBytes(_r, l.SenderCompId); err != nil {
			return err
		}
	}

	if l.TargetCompIdInActingVersion(actingVersion) {
		var TargetCompIdLength uint8
		if err := _m.ReadUint8(_r, &TargetCompIdLength); err != nil {
			return err
		}
		if cap(l.TargetCompId) < int(TargetCompIdLength) {
			l.TargetCompId = make([]uint8, TargetCompIdLength)
		}
		l.TargetCompId = l.TargetCompId[:TargetCompIdLength]
		if err := _m.ReadBytes(_r, l.TargetCompId); err != nil {
			return err
		}
	}

	if l.RawDataInActingVersion(actingVersion) {
		var RawDataLength uint16
		if err := _m.ReadUint16(_r, &RawDataLength); err != nil {
			return err
		}
		if cap(l.RawData) < int(RawDataLength) {
			l.RawData = make([]uint8, RawDataLength)
		}
		l.RawData = l.RawData[:RawDataLength]
		if err := _m.ReadBytes(_r, l.RawData); err != nil {
			return err
		}
	}

	if l.UsernameInActingVersion(actingVersion) {
		var UsernameLength uint16
		if err := _m.ReadUint16(_r, &UsernameLength); err != nil {
			return err
		}
		if cap(l.Username) < int(UsernameLength) {
			l.Username = make([]uint8, UsernameLength)
		}
		l.Username = l.Username[:UsernameLength]
		if err := _m.ReadBytes(_r, l.Username); err != nil {
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

func (l *Logon) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.EncryptMethodInActingVersion(actingVersion) {
		if l.EncryptMethod != l.EncryptMethodNullValue() && (l.EncryptMethod < l.EncryptMethodMinValue() || l.EncryptMethod > l.EncryptMethodMaxValue()) {
			return fmt.Errorf("Range check failed on l.EncryptMethod (%v < %v > %v)", l.EncryptMethodMinValue(), l.EncryptMethod, l.EncryptMethodMaxValue())
		}
	}
	if l.HeartBtIntInActingVersion(actingVersion) {
		if l.HeartBtInt < l.HeartBtIntMinValue() || l.HeartBtInt > l.HeartBtIntMaxValue() {
			return fmt.Errorf("Range check failed on l.HeartBtInt (%v < %v > %v)", l.HeartBtIntMinValue(), l.HeartBtInt, l.HeartBtIntMaxValue())
		}
	}
	if err := l.ResetSeqNumFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.MessageHandling.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.ResponseMode.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.ExecutionReportType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := l.DropCopyFlag.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if l.RecvWindowInActingVersion(actingVersion) {
		if l.RecvWindow != l.RecvWindowNullValue() && (l.RecvWindow < l.RecvWindowMinValue() || l.RecvWindow > l.RecvWindowMaxValue()) {
			return fmt.Errorf("Range check failed on l.RecvWindow (%v < %v > %v)", l.RecvWindowMinValue(), l.RecvWindow, l.RecvWindowMaxValue())
		}
	}
	if !utf8.Valid(l.SenderCompId[:]) {
		return errors.New("l.SenderCompId failed UTF-8 validation")
	}
	if !utf8.Valid(l.TargetCompId[:]) {
		return errors.New("l.TargetCompId failed UTF-8 validation")
	}
	if !utf8.Valid(l.RawData[:]) {
		return errors.New("l.RawData failed UTF-8 validation")
	}
	if !utf8.Valid(l.Username[:]) {
		return errors.New("l.Username failed UTF-8 validation")
	}
	return nil
}

func LogonInit(l *Logon) {
	l.EncryptMethod = math.MaxUint8
	l.RecvWindow = math.MaxUint32
	return
}

func (*Logon) SbeBlockLength() (blockLength uint16) {
	return 14
}

func (*Logon) SbeTemplateId() (templateId uint16) {
	return 20008
}

func (*Logon) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Logon) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*Logon) SbeSemanticType() (semanticType []byte) {
	return []byte("A")
}

func (*Logon) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Logon) EncryptMethodId() uint16 {
	return 98
}

func (*Logon) EncryptMethodSinceVersion() uint16 {
	return 0
}

func (l *Logon) EncryptMethodInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.EncryptMethodSinceVersion()
}

func (*Logon) EncryptMethodDeprecated() uint16 {
	return 0
}

func (*Logon) EncryptMethodMetaAttribute(meta int) string {
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

func (*Logon) EncryptMethodMinValue() uint8 {
	return 0
}

func (*Logon) EncryptMethodMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Logon) EncryptMethodNullValue() uint8 {
	return math.MaxUint8
}

func (*Logon) HeartBtIntId() uint16 {
	return 108
}

func (*Logon) HeartBtIntSinceVersion() uint16 {
	return 0
}

func (l *Logon) HeartBtIntInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.HeartBtIntSinceVersion()
}

func (*Logon) HeartBtIntDeprecated() uint16 {
	return 0
}

func (*Logon) HeartBtIntMetaAttribute(meta int) string {
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

func (*Logon) HeartBtIntMinValue() uint32 {
	return 0
}

func (*Logon) HeartBtIntMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Logon) HeartBtIntNullValue() uint32 {
	return math.MaxUint32
}

func (*Logon) ResetSeqNumFlagId() uint16 {
	return 141
}

func (*Logon) ResetSeqNumFlagSinceVersion() uint16 {
	return 0
}

func (l *Logon) ResetSeqNumFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ResetSeqNumFlagSinceVersion()
}

func (*Logon) ResetSeqNumFlagDeprecated() uint16 {
	return 0
}

func (*Logon) ResetSeqNumFlagMetaAttribute(meta int) string {
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

func (*Logon) MessageHandlingId() uint16 {
	return 25035
}

func (*Logon) MessageHandlingSinceVersion() uint16 {
	return 0
}

func (l *Logon) MessageHandlingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.MessageHandlingSinceVersion()
}

func (*Logon) MessageHandlingDeprecated() uint16 {
	return 0
}

func (*Logon) MessageHandlingMetaAttribute(meta int) string {
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

func (*Logon) ResponseModeId() uint16 {
	return 25036
}

func (*Logon) ResponseModeSinceVersion() uint16 {
	return 0
}

func (l *Logon) ResponseModeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ResponseModeSinceVersion()
}

func (*Logon) ResponseModeDeprecated() uint16 {
	return 0
}

func (*Logon) ResponseModeMetaAttribute(meta int) string {
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

func (*Logon) ExecutionReportTypeId() uint16 {
	return 25045
}

func (*Logon) ExecutionReportTypeSinceVersion() uint16 {
	return 0
}

func (l *Logon) ExecutionReportTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ExecutionReportTypeSinceVersion()
}

func (*Logon) ExecutionReportTypeDeprecated() uint16 {
	return 0
}

func (*Logon) ExecutionReportTypeMetaAttribute(meta int) string {
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

func (*Logon) DropCopyFlagId() uint16 {
	return 9406
}

func (*Logon) DropCopyFlagSinceVersion() uint16 {
	return 0
}

func (l *Logon) DropCopyFlagInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.DropCopyFlagSinceVersion()
}

func (*Logon) DropCopyFlagDeprecated() uint16 {
	return 0
}

func (*Logon) DropCopyFlagMetaAttribute(meta int) string {
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

func (*Logon) RecvWindowId() uint16 {
	return 25000
}

func (*Logon) RecvWindowSinceVersion() uint16 {
	return 0
}

func (l *Logon) RecvWindowInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.RecvWindowSinceVersion()
}

func (*Logon) RecvWindowDeprecated() uint16 {
	return 0
}

func (*Logon) RecvWindowMetaAttribute(meta int) string {
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

func (*Logon) RecvWindowMinValue() uint32 {
	return 0
}

func (*Logon) RecvWindowMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*Logon) RecvWindowNullValue() uint32 {
	return math.MaxUint32
}

func (*Logon) SenderCompIdMetaAttribute(meta int) string {
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

func (*Logon) SenderCompIdSinceVersion() uint16 {
	return 0
}

func (l *Logon) SenderCompIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SenderCompIdSinceVersion()
}

func (*Logon) SenderCompIdDeprecated() uint16 {
	return 0
}

func (Logon) SenderCompIdCharacterEncoding() string {
	return "UTF-8"
}

func (Logon) SenderCompIdHeaderLength() uint64 {
	return 1
}

func (*Logon) TargetCompIdMetaAttribute(meta int) string {
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

func (*Logon) TargetCompIdSinceVersion() uint16 {
	return 0
}

func (l *Logon) TargetCompIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TargetCompIdSinceVersion()
}

func (*Logon) TargetCompIdDeprecated() uint16 {
	return 0
}

func (Logon) TargetCompIdCharacterEncoding() string {
	return "UTF-8"
}

func (Logon) TargetCompIdHeaderLength() uint64 {
	return 1
}

func (*Logon) RawDataMetaAttribute(meta int) string {
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

func (*Logon) RawDataSinceVersion() uint16 {
	return 0
}

func (l *Logon) RawDataInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.RawDataSinceVersion()
}

func (*Logon) RawDataDeprecated() uint16 {
	return 0
}

func (Logon) RawDataCharacterEncoding() string {
	return "UTF-8"
}

func (Logon) RawDataHeaderLength() uint64 {
	return 2
}

func (*Logon) UsernameMetaAttribute(meta int) string {
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

func (*Logon) UsernameSinceVersion() uint16 {
	return 0
}

func (l *Logon) UsernameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.UsernameSinceVersion()
}

func (*Logon) UsernameDeprecated() uint16 {
	return 0
}

func (Logon) UsernameCharacterEncoding() string {
	return "UTF-8"
}

func (Logon) UsernameHeaderLength() uint64 {
	return 2
}
