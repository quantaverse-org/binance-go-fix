// Generated SBE (Simple Binary Encoding) message codec

package fixsbe

import (
	"fmt"
	"io"
	"math"
)

type GroupSize8Encoding struct {
	BlockLength uint16
	NumInGroup  uint8
}

func (g *GroupSize8Encoding) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, g.BlockLength); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, g.NumInGroup); err != nil {
		return err
	}
	return nil
}

func (g *GroupSize8Encoding) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
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
		if err := _m.ReadUint8(_r, &g.NumInGroup); err != nil {
			return err
		}
	}
	return nil
}

func (g *GroupSize8Encoding) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func GroupSize8EncodingInit(g *GroupSize8Encoding) {
	return
}

func (*GroupSize8Encoding) EncodedLength() int64 {
	return 3
}

func (*GroupSize8Encoding) BlockLengthMinValue() uint16 {
	return 0
}

func (*GroupSize8Encoding) BlockLengthMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*GroupSize8Encoding) BlockLengthNullValue() uint16 {
	return math.MaxUint16
}

func (*GroupSize8Encoding) BlockLengthSinceVersion() uint16 {
	return 0
}

func (g *GroupSize8Encoding) BlockLengthInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.BlockLengthSinceVersion()
}

func (*GroupSize8Encoding) BlockLengthDeprecated() uint16 {
	return 0
}

func (*GroupSize8Encoding) NumInGroupMinValue() uint8 {
	return 0
}

func (*GroupSize8Encoding) NumInGroupMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*GroupSize8Encoding) NumInGroupNullValue() uint8 {
	return math.MaxUint8
}

func (*GroupSize8Encoding) NumInGroupSinceVersion() uint16 {
	return 0
}

func (g *GroupSize8Encoding) NumInGroupInActingVersion(actingVersion uint16) bool {
	return actingVersion >= g.NumInGroupSinceVersion()
}

func (*GroupSize8Encoding) NumInGroupDeprecated() uint16 {
	return 0
}
