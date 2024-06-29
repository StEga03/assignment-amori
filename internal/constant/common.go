package constant

import "time"

const (
	ApplicationName = "finanxier-api"
	ServiceTypeHTTP = "http"
)

const (
	LangID   = "id"
	LangENUS = "en-us"
)

// Default value for commonly used data types.
const (
	DefaultKey string = "DEFAULT"
	// DefaultString : the zero value of type string
	DefaultString string = ""
	// DefaultFloat64 : the zero value of type float64
	DefaultFloat64 float64 = 0
	// DefaultFloat32 : the zero value of type float32
	DefaultFloat32 float32 = 0
	// DefaultInt64 : the zero value of type int64
	DefaultInt64 int64 = 0
	// DefaultInt32 : the zero value of type int32
	DefaultInt32 = 0
	// DefaultInt : the zero value of type int
	DefaultInt int = 0
	// DefaultInt : the zero value of type int8
	DefaultInt8 int8 = 0
	// DefaultUInt64 : the zero value of type uint64
	DefaultUInt64 uint64 = 0

	// bool values.
	BoolTrue  bool = true
	BoolFalse bool = false

	// DefaultIntOne : the one value of type int
	DefaultIntOne = 1
)

const (
	CustomCommonDateFormat   = "02 January 2006"
	CustomDateFormat         = "2006-01-02T15:04:05"
	CustomAbsoluteDateFormat = "2006-01-02"
	CustomDateWOSecondFormat = "2006-01-02T15:04+07:00"
	CustomDateTimeFormat     = "2006-01-02 15:04:05"
	CustomDateISO8601        = "20060102T150405Z"
)

const (
	WIBTimeDifferenceHrs = 7
)

// Custom RFC3339 at Asia/Jakarta Time Zone
const (
	CustomRFC3339 = "2006-01-02T15:04:05+07:00"
)

const (
	DefaultLimit  int = 10
	DefaultOffset int = 0
)

const (
	// LockerTTL defines the time-to-live for lockers.
	LockerTTL = 2 * time.Minute

	// CacheTTL defines the time-to-live for cache entries.
	CacheTTL = 2 * time.Hour

	// SessionTTL defines the time-to-live for user sessions.
	SessionTTL = 24 * time.Hour
)

const (
	ContentTypeText = "text"
)
