package message

import (
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  time.Time
	}{
		{
			name:  "seconds",
			value: "20011217-09:30:47",
			want:  time.Date(2001, time.December, 17, 9, 30, 47, 0, time.UTC),
		},
		{
			name:  "milliseconds",
			value: "20011217-09:30:47.123",
			want:  time.Date(2001, time.December, 17, 9, 30, 47, 123000000, time.UTC),
		},
		{
			name:  "microseconds",
			value: "20011217-09:30:47.123456",
			want:  time.Date(2001, time.December, 17, 9, 30, 47, 123456000, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTimestamp(tt.value)
			if err != nil {
				t.Fatalf("ParseTimestamp() error = %v", err)
			}
			if !got.Equal(tt.want) {
				t.Fatalf("ParseTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseTimestampInvalid(t *testing.T) {
	if _, err := ParseTimestamp("20011217-09:30:47.12"); err == nil {
		t.Fatal("ParseTimestamp() error = nil, want error")
	}
}
