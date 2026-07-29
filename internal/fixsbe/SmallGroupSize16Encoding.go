// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"math"
)

type SmallGroupSize16Encoding struct {
	BlockLength uint8
	NumInGroup  uint16
}

func (s *SmallGroupSize16Encoding) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, s.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (s *SmallGroupSize16Encoding) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !s.BlockLengthInActingVersion(actingVersion) {
		s.BlockLength = s.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.BlockLength); err != nil {
			return err
		}
	}
	if !s.NumInGroupInActingVersion(actingVersion) {
		s.NumInGroup = s.NumInGroupNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (s *SmallGroupSize16Encoding) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.BlockLengthInActingVersion(actingVersion) {
		if s.BlockLength < s.BlockLengthMinValue() || s.BlockLength > s.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on s.BlockLength (%v < %v > %v)", s.BlockLengthMinValue(), s.BlockLength, s.BlockLengthMaxValue())
		}
	}
	if s.NumInGroupInActingVersion(actingVersion) {
		if s.NumInGroup < s.NumInGroupMinValue() || s.NumInGroup > s.NumInGroupMaxValue() {
			return fmt.Errorf("Range check failed on s.NumInGroup (%v < %v > %v)", s.NumInGroupMinValue(), s.NumInGroup, s.NumInGroupMaxValue())
		}
	}
	return nil
}

func SmallGroupSize16EncodingInit(s *SmallGroupSize16Encoding) {
	return
}

func (*SmallGroupSize16Encoding) EncodedLength() int64 {
	return 3
}

func (*SmallGroupSize16Encoding) BlockLengthMinValue() uint8 {
	return 0
}

func (*SmallGroupSize16Encoding) BlockLengthMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SmallGroupSize16Encoding) BlockLengthNullValue() uint8 {
	return math.MaxUint8
}

func (*SmallGroupSize16Encoding) BlockLengthSinceVersion() uint16 {
	return 0
}

func (s *SmallGroupSize16Encoding) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BlockLengthSinceVersion()
}

func (*SmallGroupSize16Encoding) BlockLengthDeprecated() uint16 {
	return 0
}

func (*SmallGroupSize16Encoding) NumInGroupMinValue() uint16 {
	return 0
}

func (*SmallGroupSize16Encoding) NumInGroupMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SmallGroupSize16Encoding) NumInGroupNullValue() uint16 {
	return math.MaxUint16
}

func (*SmallGroupSize16Encoding) NumInGroupSinceVersion() uint16 {
	return 0
}

func (s *SmallGroupSize16Encoding) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NumInGroupSinceVersion()
}

func (*SmallGroupSize16Encoding) NumInGroupDeprecated() uint16 {
	return 0
}
