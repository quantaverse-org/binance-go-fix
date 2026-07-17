package message

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const SOH = '\x01'

type Request interface {
	ToMessage(senderCompId string, targetCompId string, seqNum uint32, sendingTime time.Time) (*Message, error)
}

type Response interface {
	FromMessage(m *Message) error
}

type Field struct {
	tag   Tag
	value string
}

type Message struct {
	fields     []Field
	rawMessage string
}

func ParseMessage(s string) (*Message, error) {
	var fields []Field
	elems := strings.Split(s, string(SOH))
	for _, elem := range elems {
		if elem == "" {
			continue
		}

		pair := strings.SplitN(elem, "=", 2)
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid field: %s", elem)
		}
		tagStr, value := pair[0], pair[1]
		tag, err := ParseTag(tagStr)
		if err != nil {
			return nil, err
		}
		fields = append(fields, Field{tag: tag, value: value})
	}

	return &Message{fields: fields, rawMessage: s}, nil
}

func (m *Message) RawMessage() string {
	return m.rawMessage
}

func (m *Message) HasField(tag Tag) bool {
	for _, field := range m.fields {
		if field.tag == tag {
			return true
		}
	}
	return false
}

func (m *Message) GetField(tag Tag) (string, bool) {
	for _, field := range m.fields {
		if field.tag == tag {
			return field.value, true
		}
	}
	return "", false
}

func (m *Message) GetRequiredField(tag Tag) (string, error) {
	for _, field := range m.fields {
		if field.tag == tag {
			return field.value, nil
		}
	}
	return "", errors.New("missing field")
}

func (m *Message) GetFields(tag Tag) []Field {
	fields := make([]Field, 0, len(m.fields))
	for _, field := range m.fields {
		if field.tag == tag {
			fields = append(fields, field)
		}
	}
	return fields
}

func (m *Message) AddField(tag Tag, value string) {
	m.fields = append(m.fields, Field{tag: tag, value: value})
}

func (m *Message) SetField(tag Tag, value string) {
	for i := range m.fields {
		if m.fields[i].tag == tag {
			m.fields[i].value = value
			return
		}
	}
	m.AddField(tag, value)
}

func (m *Message) MsgType() (MsgType, error) {
	v, err := m.GetRequiredField(TagMsgType)
	if err != nil {
		return "", err
	}
	return ParseMsgType(v)
}

func (m *Message) SenderCompId() (string, error) {
	return m.GetRequiredField(TagSenderCompID)
}

func (m *Message) TargetCompId() (string, error) {
	return m.GetRequiredField(TagTargetCompID)
}

func (m *Message) MsgSeqNum() (uint32, error) {
	v, err := m.GetRequiredField(TagMsgSeqNum)
	if err != nil {
		return 0, err
	}
	seqNo, err := ParseUint(v)
	if err != nil {
		return 0, err
	}
	return uint32(seqNo), nil
}

func (m *Message) SendingTime() (time.Time, error) {
	v, err := m.GetRequiredField(TagSendingTime)
	if err != nil {
		return time.Time{}, err
	}
	return ParseTimestamp(v)
}

func (m *Message) Checksum() (uint8, error) {
	v, err := m.GetRequiredField(TagCheckSum)
	if err != nil {
		return 0, err
	}
	checksum, err := ParseUint(v)
	if err != nil {
		return 0, err
	}
	return uint8(checksum), nil
}

func (m *Message) Validate() error {
	actualCkm, err := m.Checksum()
	if err != nil {
		return err
	}
	expectedCkm := m.CalculateChecksum()
	if expectedCkm != actualCkm {
		return fmt.Errorf("checksum mismatch: %d != %d", expectedCkm, actualCkm)
	}
	return nil
}

// CalculateChecksum calculates the CheckSum (10) value by summing every byte before the CheckSum field,
// including start-of-header (SOH) characters, and returning the remainder modulo 256.
func (m *Message) CalculateChecksum() uint8 {
	var sum uint32
	addString := func(s string) {
		for i := 0; i < len(s); i++ {
			sum += uint32(s[i])
		}
	}
	for _, field := range m.fields {
		if field.tag == TagCheckSum {
			break
		}
		addString(FormatTag(field.tag))
		sum += uint32('=')
		addString(field.value)
		sum += uint32(SOH)
	}
	return uint8(sum % 256)
}

func (m *Message) Display() string {
	newMsg := strings.ReplaceAll(m.rawMessage, string(SOH), "|")
	newMsg = strings.TrimSuffix(newMsg, "|")
	return newMsg
}
