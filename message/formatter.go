package message

import (
	"strconv"
	"time"
)

func FormatBool(value bool) string {
	if value {
		return "Y"
	}
	return "N"
}

func FormatInt(value int64) string {
	return strconv.FormatInt(value, 10)
}

func FormatUint(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func FormatFloat(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func FormatTimestampMs(value time.Time) string {
	return value.UTC().Format("20060102-15:04:05.000")
}

func FormatTag(value Tag) string {
	return FormatUint(uint64(value))
}

func FormatMsgType(value MsgType) string {
	return string(value)
}
