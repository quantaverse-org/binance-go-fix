// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"math"
)

type GroupSize32Encoding struct {
	BlockLength uint16
	NumInGroup  uint32
}

func (g *GroupSize32Encoding) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, g.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, g.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (g *GroupSize32Encoding) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if !g.BlockLengthInActingVersion(actingVersion) {
		g.BlockLength = g.BlockLengthNullValue()
	} else {
		if err := _m.ReadUint16(_r, &g.BlockLength); err != nil {
			return err
		}
	}
	if !g.NumInGroupInActingVersion(actingVersion) {
		g.NumInGroup = g.NumInGroupNullValue()
	} else {
		if err := _m.ReadUint32(_r, &g.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (g *GroupSize32Encoding) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if g.BlockLengthInActingVersion(actingVersion) {
		if g.BlockLength < g.BlockLengthMinValue() || g.BlockLength > g.BlockLengthMaxValue() {
			return fmt.Errorf("Range check failed on g.BlockLength (%v < %v > %v)", g.BlockLengthMinValue(), g.BlockLength, g.BlockLengthMaxValue())
		}
	}
	if g.NumInGroupInActingVersion(actingVersion) {
		if g.NumInGroup < g.NumInGroupMinValue() || g.NumInGroup > g.NumInGroupMaxValue() {
			return fmt.Errorf("Range check failed on g.NumInGroup (%v < %v > %v)", g.NumInGroupMinValue(), g.NumInGroup, g.NumInGroupMaxValue())
		}
	}
	return nil
}

func GroupSize32EncodingInit(g *GroupSize32Encoding) {
	return
}

func (*GroupSize32Encoding) EncodedLength() int64 {
	return 6
}

func (*GroupSize32Encoding) BlockLengthMinValue() uint16 {
	return 0
}

func (*GroupSize32Encoding) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*GroupSize32Encoding) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (*GroupSize32Encoding) BlockLengthSinceVersion() uint16 {
	return 0
}

func (g *GroupSize32Encoding) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.BlockLengthSinceVersion()
}

func (*GroupSize32Encoding) BlockLengthDeprecated() uint16 {
	return 0
}

func (*GroupSize32Encoding) NumInGroupMinValue() uint32 {
	return 0
}

func (*GroupSize32Encoding) NumInGroupMaxValue() uint32 {
	return 2147483647
}

func (*GroupSize32Encoding) NumInGroupNullValue() uint32 {
	return math.MaxUint32
}

func (*GroupSize32Encoding) NumInGroupSinceVersion() uint16 {
	return 0
}

func (g *GroupSize32Encoding) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.NumInGroupSinceVersion()
}

func (*GroupSize32Encoding) NumInGroupDeprecated() uint16 {
	return 0
}
