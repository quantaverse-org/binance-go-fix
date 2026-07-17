package message

import (
	"fmt"
	"strconv"
	"time"
)

func ParseBool(value string) (bool, error) {
	switch value {
	case "Y":
		return true, nil
	case "N":
		return false, nil
	default:
		return false, fmt.Errorf("failed to parse %s as bool", value)
	}
}

func ParseInt(value string) (int64, error) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as int", value)
	}
	return v, nil
}

func ParseUint(value string) (uint64, error) {
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as uint", value)
	}
	return v, nil
}

func ParseFloat(value string) (float64, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s as float", value)
	}
	return v, nil
}

// ParseTimestamp
// Supported UTCTIMESTAMP formats:
// 20011217-09:30:47 - seconds
// 20011217-09:30:47.123 - milliseconds
// 20011217-09:30:47.123456 - microseconds (always used in messages from the exchange)
func ParseTimestamp(value string) (time.Time, error) {
	var layout string
	switch len(value) {
	case len("20011217-09:30:47"):
		layout = "20060102-15:04:05"
	case len("20011217-09:30:47.123"):
		layout = "20060102-15:04:05.000"
	case len("20011217-09:30:47.123456"):
		layout = "20060102-15:04:05.000000"
	default:
		return time.Time{}, fmt.Errorf("failed to parse %s as timestamp", value)
	}

	t, err := time.ParseInLocation(layout, value, time.UTC)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse %s as timestamp", value)
	}
	return t, nil
}

func ParseTag(value string) (Tag, error) {
	v, err := ParseUint(value)
	if err != nil {
		return 0, err
	}
	return Tag(v), nil
}

func ParseMsgType(value string) (MsgType, error) {
	v := MsgType(value)
	if v.IsValid() {
		return v, nil
	}
	return v, fmt.Errorf("invalid message type: %s", value)
}
