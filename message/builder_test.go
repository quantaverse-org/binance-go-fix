package message

import (
	"strings"
	"testing"
)

func TestMsgBuilderBuild(t *testing.T) {
	message, err := NewMsgBuilder().
		MsgType(MsgTypeHeartbeat).
		SenderCompID("SENDER").
		TargetCompID("TARGET").
		MsgSeqNum(1).
		SetField(TagSendingTime, "20011217-09:30:47.123").
		Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	bodyLength, err := message.GetRequiredField(TagBodyLength)
	if err != nil {
		t.Fatalf("GetRequiredField(TagBodyLength) error = %v", err)
	}
	if bodyLength != "55" {
		t.Fatalf("BodyLength = %q, want %q", bodyLength, "55")
	}

	checksum, err := message.GetRequiredField(TagCheckSum)
	if err != nil {
		t.Fatalf("GetRequiredField(TagCheckSum) error = %v", err)
	}
	if len(checksum) != 3 {
		t.Fatalf("CheckSum length = %d, want %d", len(checksum), 3)
	}

	wantTags := []Tag{
		TagBeginString,
		TagMsgType,
		TagSenderCompID,
		TagTargetCompID,
		TagMsgSeqNum,
		TagSendingTime,
		TagBodyLength,
		TagCheckSum,
	}
	if len(message.fields) != len(wantTags) {
		t.Fatalf("field count = %d, want %d", len(message.fields), len(wantTags))
	}
	for i, want := range wantTags {
		if message.fields[i].tag != want {
			t.Fatalf("field[%d].tag = %d, want %d", i, message.fields[i].tag, want)
		}
	}

	wantPrefix := "8=FIX.4.4\x019=55\x0134=1\x0135=0\x0149=SENDER\x0156=TARGET\x0152=20011217-09:30:47.123\x01"
	if !strings.HasPrefix(message.rawMessage, wantPrefix) {
		t.Fatalf("rawMessage = %q, want prefix %q", message.rawMessage, wantPrefix)
	}
	if !strings.HasSuffix(message.rawMessage, "10="+checksum+"\x01") {
		t.Fatalf("rawMessage = %q, want checksum suffix %q", message.rawMessage, "10="+checksum+"\x01")
	}

	if err := message.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
}

func TestMsgBuilderBuildMissingRequiredField(t *testing.T) {
	_, err := NewMsgBuilder().Build()
	if err == nil {
		t.Fatal("Build() error = nil, want error")
	}
}
