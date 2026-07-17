package message

import (
	"fmt"
	"strings"
	"time"
)

type MsgBuilder struct {
	message *Message
}

func NewMsgBuilder() *MsgBuilder {
	msg := new(Message)
	msg.SetField(TagBeginString, "FIX.4.4")
	return &MsgBuilder{message: msg}
}

func (mb *MsgBuilder) MsgType(msgTy MsgType) *MsgBuilder {
	mb.message.SetField(TagMsgType, string(msgTy))
	return mb
}

func (mb *MsgBuilder) SenderCompID(senderCompID string) *MsgBuilder {
	mb.message.SetField(TagSenderCompID, senderCompID)
	return mb
}

func (mb *MsgBuilder) TargetCompID(targetCompID string) *MsgBuilder {
	mb.message.SetField(TagTargetCompID, targetCompID)
	return mb
}

func (mb *MsgBuilder) MsgSeqNum(msgSeqNum uint32) *MsgBuilder {
	mb.message.SetField(TagMsgSeqNum, FormatUint(uint64(msgSeqNum)))
	return mb
}

func (mb *MsgBuilder) SendingTime(timestamp time.Time) *MsgBuilder {
	mb.message.SetField(TagSendingTime, FormatTimestampMs(timestamp))
	return mb
}

func (mb *MsgBuilder) SetField(tag Tag, value string) *MsgBuilder {
	mb.message.SetField(tag, value)
	return mb
}

func (mb *MsgBuilder) AddField(tag Tag, value string) *MsgBuilder {
	mb.message.AddField(tag, value)
	return mb
}

func (mb *MsgBuilder) Build() (*Message, error) {
	requiredTags := []Tag{
		TagBeginString,
		TagMsgType,
		TagSenderCompID,
		TagTargetCompID,
		TagMsgSeqNum,
		TagSendingTime,
	}
	for _, tag := range requiredTags {
		if !mb.message.HasField(tag) {
			return nil, fmt.Errorf("missing field: %d", tag)
		}
	}

	bodyLength := mb.calculateBodyLength()
	mb.message.SetField(TagBodyLength, FormatUint(uint64(bodyLength)))

	checksum := mb.message.CalculateChecksum()
	mb.message.SetField(TagCheckSum, fmt.Sprintf("%03d", checksum))

	mb.message.rawMessage = buildRawMessage(mb.message.fields)

	return mb.message, nil
}

func (mb *MsgBuilder) calculateBodyLength() int {
	var length int
	for _, field := range mb.message.fields {
		if field.tag == TagBeginString || field.tag == TagBodyLength || field.tag == TagCheckSum {
			continue
		}
		length += len(FormatTag(field.tag)) + 1 + len(field.value) + 1
	}
	return length
}

func buildRawMessage(fields []Field) string {
	var builder strings.Builder

	buildField := func(field Field) {
		builder.WriteString(FormatTag(field.tag))
		builder.WriteByte('=')
		builder.WriteString(field.value)
		builder.WriteByte(SOH)
	}

	buildOnly := func(tag Tag) {
		for _, field := range fields {
			if field.tag == tag {
				buildField(field)
				break
			}
		}
	}

	buildOnly(TagBeginString)
	buildOnly(TagBodyLength)
	buildOnly(TagMsgType)
	for _, field := range fields {
		if field.tag != TagBeginString &&
			field.tag != TagBodyLength &&
			field.tag != TagMsgType &&
			field.tag != TagCheckSum {
			buildField(field)
		}
	}
	buildOnly(TagCheckSum)

	return builder.String()
}
