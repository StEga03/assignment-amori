package timeutils

import "time"

var (
	Now = time.Now
)

// Seconds/Unix Converter.
func GetTimeNowUnix() int64 {
	return Now().Unix()
}
func GetTimeFromUnix(timestamp int64) time.Time {
	t := time.Unix(timestamp, 0)
	return t
}

// Miliseconds/Unix Converter.
func GetTimeNowUnixMili() int64 {
	return Now().UnixMilli()
}
func GetTimeFromUnixMili(timestamp int64) time.Time {
	t := time.UnixMilli(timestamp)
	return t
}

// Microseconds/Unix Converter.
func GetTimeNowUnixMicro() int64 {
	return Now().UnixMicro()
}
func GetTimeFromUnixMicro(timestamp int64) time.Time {
	t := time.UnixMicro(timestamp)
	return t
}
