package utils

import (
	"strings"
	"time"
)

var defaultTimeFormats = []string{
	time.RFC3339, "2006-01-02 15:04:05 -0700", "2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02",
	"02-Jan-2006 15:04:05", "02-01-2006 15:04:05", "02-01-2006 15:04", "02 Jan 2006  15:04",
	"2006-01-02 15:04:05 -0700 -0700",
}

// TimeValidatorFormats will convert string to time.Time types based on provided formats.
func TimeValidatorFormats(t string, formats []string) (*time.Time, error) {
	var tempDateTime time.Time
	var err error

	for i := 0; i < len(formats); i++ {
		tempDateTime, err = time.Parse(formats[i], t)
		if err != nil {
			continue
		}
		break
	}

	return &tempDateTime, err
}

// TimeValidator will convert string to time.Time types base default time formats.
func TimeValidator(t string) (*time.Time, error) {
	return TimeValidatorFormats(t, defaultTimeFormats)
}

// Between Get string between 2 characters start from `a` to `b`.
// https://www.dotnetperls.com/between-before-after-go
func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
