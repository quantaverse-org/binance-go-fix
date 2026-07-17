package message

import (
	"testing"
)

func TestMessageSetFieldUpdatesExistingField(t *testing.T) {
	message := new(Message)
	message.SetField(TagSenderCompID, "old")
	message.SetField(TagSenderCompID, "new")

	got, ok := message.GetField(TagSenderCompID)
	if !ok {
		t.Fatal("GetField() ok = false, want true")
	}
	if got != "new" {
		t.Fatalf("GetField() = %q, want %q", got, "new")
	}
	if len(message.fields) != 1 {
		t.Fatalf("field count = %d, want %d", len(message.fields), 1)
	}
}

func TestMessageCalculateChecksumFromParsedFields(t *testing.T) {
	rawMessage := "8=FIX.4.4\x019=5\x0135=0\x0110=163\x01"
	message, err := ParseMessage(rawMessage)
	if err != nil {
		t.Fatalf("ParseMessage() error = %v", err)
	}

	if got := message.CalculateChecksum(); got != 163 {
		t.Fatalf("CalculateChecksum() = %d, want %d", got, 163)
	}
}

func TestMessageCalculateChecksumFromFields(t *testing.T) {
	message := &Message{
		fields: []Field{
			{tag: TagBeginString, value: "FIX.4.4"},
			{tag: TagBodyLength, value: "5"},
			{tag: TagMsgType, value: "0"},
			{tag: TagCheckSum, value: "163"},
		},
	}

	if got := message.CalculateChecksum(); got != 163 {
		t.Fatalf("CalculateChecksum() = %d, want %d", got, 163)
	}
}
