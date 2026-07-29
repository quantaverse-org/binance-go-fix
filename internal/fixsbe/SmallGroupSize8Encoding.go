// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"math"
)

type SmallGroupSize8Encoding struct {
	BlockLength uint8
	NumInGroup  uint8
}

func (s *SmallGroupSize8Encoding) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, s.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (s *SmallGroupSize8Encoding) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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
		if err := _m.ReadUint8(_r, &s.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (s *SmallGroupSize8Encoding) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func SmallGroupSize8EncodingInit(s *SmallGroupSize8Encoding) {
	return
}

func (*SmallGroupSize8Encoding) EncodedLength() int64 {
	return 2
}

func (*SmallGroupSize8Encoding) BlockLengthMinValue() uint8 {
	return 0
}

func (*SmallGroupSize8Encoding) BlockLengthMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SmallGroupSize8Encoding) BlockLengthNullValue() uint8 {
	return math.MaxUint8
}

func (*SmallGroupSize8Encoding) BlockLengthSinceVersion() uint16 {
	return 0
}

func (s *SmallGroupSize8Encoding) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BlockLengthSinceVersion()
}

func (*SmallGroupSize8Encoding) BlockLengthDeprecated() uint16 {
	return 0
}

func (*SmallGroupSize8Encoding) NumInGroupMinValue() uint8 {
	return 0
}

func (*SmallGroupSize8Encoding) NumInGroupMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SmallGroupSize8Encoding) NumInGroupNullValue() uint8 {
	return math.MaxUint8
}

func (*SmallGroupSize8Encoding) NumInGroupSinceVersion() uint16 {
	return 0
}

func (s *SmallGroupSize8Encoding) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NumInGroupSinceVersion()
}

func (*SmallGroupSize8Encoding) NumInGroupDeprecated() uint16 {
	return 0
}
