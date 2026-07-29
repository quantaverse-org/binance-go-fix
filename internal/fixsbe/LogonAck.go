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

type LogonAck struct {
	EncryptMethod                uint8
	HeartBtInt                   uint32
	SbeSchemaIdVersionDeprecated BoolEnumEnum
	UUID                         []uint8
}

func (l *LogonAck) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
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
	if err := l.SbeSchemaIdVersionDeprecated.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, uint8(len(l.UUID))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.UUID); err != nil {
		return err
	}
	return nil
}

func (l *LogonAck) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
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
	if l.SbeSchemaIdVersionDeprecatedInActingVersion(actingVersion) {
		if err := l.SbeSchemaIdVersionDeprecated.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.UUIDInActingVersion(actingVersion) {
		var UUIDLength uint8
		if err := _m.ReadUint8(_r, &UUIDLength); err != nil {
			return err
		}
		if cap(l.UUID) < int(UUIDLength) {
			l.UUID = make([]uint8, UUIDLength)
		}
		l.UUID = l.UUID[:UUIDLength]
		if err := _m.ReadBytes(_r, l.UUID); err != nil {
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

func (l *LogonAck) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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
	if err := l.SbeSchemaIdVersionDeprecated.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if !utf8.Valid(l.UUID[:]) {
		return errors.New("l.UUID failed UTF-8 validation")
	}
	return nil
}

func LogonAckInit(l *LogonAck) {
	l.EncryptMethod = math.MaxUint8
	return
}

func (*LogonAck) SbeBlockLength() (blockLength uint16) {
	return 6
}

func (*LogonAck) SbeTemplateId() (templateId uint16) {
	return 20009
}

func (*LogonAck) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LogonAck) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*LogonAck) SbeSemanticType() (semanticType []byte) {
	return []byte("A")
}

func (*LogonAck) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LogonAck) EncryptMethodId() uint16 {
	return 98
}

func (*LogonAck) EncryptMethodSinceVersion() uint16 {
	return 0
}

func (l *LogonAck) EncryptMethodInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.EncryptMethodSinceVersion()
}

func (*LogonAck) EncryptMethodDeprecated() uint16 {
	return 0
}

func (*LogonAck) EncryptMethodMetaAttribute(meta int) string {
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

func (*LogonAck) EncryptMethodMinValue() uint8 {
	return 0
}

func (*LogonAck) EncryptMethodMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*LogonAck) EncryptMethodNullValue() uint8 {
	return math.MaxUint8
}

func (*LogonAck) HeartBtIntId() uint16 {
	return 108
}

func (*LogonAck) HeartBtIntSinceVersion() uint16 {
	return 0
}

func (l *LogonAck) HeartBtIntInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.HeartBtIntSinceVersion()
}

func (*LogonAck) HeartBtIntDeprecated() uint16 {
	return 0
}

func (*LogonAck) HeartBtIntMetaAttribute(meta int) string {
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

func (*LogonAck) HeartBtIntMinValue() uint32 {
	return 0
}

func (*LogonAck) HeartBtIntMaxValue() uint32 {
	return math.MaxUint32 - 1
}

func (*LogonAck) HeartBtIntNullValue() uint32 {
	return math.MaxUint32
}

func (*LogonAck) SbeSchemaIdVersionDeprecatedId() uint16 {
	return 25052
}

func (*LogonAck) SbeSchemaIdVersionDeprecatedSinceVersion() uint16 {
	return 0
}

func (l *LogonAck) SbeSchemaIdVersionDeprecatedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.SbeSchemaIdVersionDeprecatedSinceVersion()
}

func (*LogonAck) SbeSchemaIdVersionDeprecatedDeprecated() uint16 {
	return 0
}

func (*LogonAck) SbeSchemaIdVersionDeprecatedMetaAttribute(meta int) string {
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

func (*LogonAck) UUIDMetaAttribute(meta int) string {
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

func (*LogonAck) UUIDSinceVersion() uint16 {
	return 0
}

func (l *LogonAck) UUIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.UUIDSinceVersion()
}

func (*LogonAck) UUIDDeprecated() uint16 {
	return 0
}

func (LogonAck) UUIDCharacterEncoding() string {
	return "UTF-8"
}

func (LogonAck) UUIDHeaderLength() uint64 {
	return 1
}
