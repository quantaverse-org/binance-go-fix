package message

import (
	"testing"
	"time"
)

func TestFormatBool(t *testing.T) {
	if got := FormatBool(true); got != "Y" {
		t.Fatalf("FormatBool(true) = %q, want %q", got, "Y")
	}
	if got := FormatBool(false); got != "N" {
		t.Fatalf("FormatBool(false) = %q, want %q", got, "N")
	}
}

func TestFormatNumbers(t *testing.T) {
	if got := FormatInt(-42); got != "-42" {
		t.Fatalf("FormatInt() = %q, want %q", got, "-42")
	}
	if got := FormatUint(42); got != "42" {
		t.Fatalf("FormatUint() = %q, want %q", got, "42")
	}
	if got := FormatFloat(6000.346); got != "6000.346" {
		t.Fatalf("FormatFloat() = %q, want %q", got, "6000.346")
	}
}

func TestFormatTimestampMs(t *testing.T) {
	local := time.FixedZone("UTC+8", 8*60*60)
	value := time.Date(2001, time.December, 17, 17, 30, 47, 123456000, local)

	got := FormatTimestampMs(value)
	want := "20011217-09:30:47.123"
	if got != want {
		t.Fatalf("FormatTimestampMs() = %q, want %q", got, want)
	}

	parsed, err := ParseTimestamp(got)
	if err != nil {
		t.Fatalf("ParseTimestamp() error = %v", err)
	}
	if parsed.Nanosecond() != 123000000 {
		t.Fatalf("ParseTimestamp().Nanosecond() = %d, want %d", parsed.Nanosecond(), 123000000)
	}
}

func TestFormatTagAndMsgType(t *testing.T) {
	if got := FormatTag(TagSendingTime); got != "52" {
		t.Fatalf("FormatTag() = %q, want %q", got, "52")
	}
	if got := FormatMsgType(MsgTypeLogon); got != "A" {
		t.Fatalf("FormatMsgType() = %q, want %q", got, "A")
	}
}
