package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/assignment-amori/internal/constant"
)

func ParseStringToTime(str string, format string) (time.Time, error) {
	if len(str) == constant.DefaultInt {
		return time.Time{}, errors.New("empty input string to StringToTimeParse function")
	}

	timeParsed, err := time.Parse(format, str)
	if err != nil {
		return time.Time{}, err
	}

	return timeParsed, nil
}

func FormatToString(t time.Time, format string) string {
	if t.IsZero() {
		return constant.DefaultString
	}
	return t.Format(format)
}

func EndOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 999999999, dateTime.Location())
}

func StartOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
}

func GetUntilDateFromRRule(rrule string) (time.Time, error) {
	var result time.Time

	re := regexp.MustCompile(`UNTIL=(\d+T\d+Z);`)
	match := re.FindStringSubmatch(rrule)
	if len(match) > 1 {
		untilDateStr := match[1]
		untilDate, err := time.Parse(constant.CustomDateISO8601, untilDateStr)
		if err == nil {
			return untilDate, nil
		}
		return result, err
	}
	return result, fmt.Errorf("UNTIL not found")
}

func ConvertTime12To24(time string, ampm string) (string, error) {
	parts := RegexSplitTime.Split(time, -1)

	if len(parts) < 2 {
		return "", errors.New("invalid time format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", errors.New("invalid time format")
	}

	if hours == 12 {
		hours = 0
	}
	if ampm == "PM" {
		hours += 12
	}

	// Reconstruct the time string in 24-hour format
	var time24 string
	if len(parts) == 3 {
		// If seconds are included
		time24 = fmt.Sprintf("%02d:%s:%s", hours, parts[1], parts[2])
	} else {
		// If only hours and minutes are included
		time24 = fmt.Sprintf("%02d:%s", hours, parts[1])
	}

	return time24, nil
}

func NormalizeTime(time string) string {
	parts := RegexSplitTime.Split(time, -1)

	// Ensure hours, minutes, and seconds are present
	if len(parts) == 2 {
		parts = append(parts, "00") // Default seconds to "00" if not present
	}

	// Pad hours and minutes to 2 digits if necessary
	hours := PadLeft(parts[0], 2, "0")
	minutes := PadLeft(parts[1], 2, "0")
	seconds := PadLeft(parts[2], 2, "0")

	return fmt.Sprintf("%s:%s:%s", hours, minutes, seconds)
}

func NormalizeAMPM(ampm string) string {
	normalizedAMPM := RegexNonAPM.ReplaceAllString(ampm, "")
	return strings.ToUpper(normalizedAMPM)
}
