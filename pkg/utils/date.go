package utils

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

// DateComponents represent a date as an array of three strings.
type DateComponents [3]string

// NumericDate represents a date as an array of three integers. The date can
// be in the form of [day, month, year] or [month, day, year].
type NumericDate [3]int

type DateOrder int

func (do DateOrder) String() string {
	switch do {
	case DateOrderUnknown:
		return "unknown"
	case DateOrderDaysFirst:
		return "days first"
	case DateOrderMonthsFirst:
		return "months first"
	default:
		panic("unknown date order")
	}
}

const (
	DateOrderUnknown DateOrder = iota
	DateOrderDaysFirst
	DateOrderMonthsFirst
)

// checkAbove12 takes a slice of slices of integers (representing numeric dates)
// and determines if the day comes before the month or vice versa.
func checkAbove12(numericDates []NumericDate) (DateOrder, error) {
	for _, date := range numericDates {
		if len(date) < 3 {
			return DateOrderUnknown, errors.New("invalid date format")
		}

		if date[0] > 12 {
			return DateOrderDaysFirst, nil
		}
		if date[1] > 12 {
			return DateOrderMonthsFirst, nil
		}
	}

	return DateOrderUnknown, nil
}

// checkDecreasing examines an array of numeric dates to determine if the days
// come before the month or vice versa, based on whether a set of numbers
// during the same year decreases at some point.
func checkDecreasing(numericDates []NumericDate) (DateOrder, error) {
	if len(numericDates) == 0 {
		return DateOrderUnknown, nil
	}

	sort.Slice(numericDates, func(i, j int) bool {
		return numericDates[i][2] < numericDates[j][2]
	})

	var daysDecreasing, monthsDecreasing bool

	for i := 1; i < len(numericDates); i++ {
		if len(numericDates[i]) < 3 || len(numericDates[i-1]) < 3 {
			return DateOrderUnknown, errors.New("invalid date format")
		}

		// Only compare within the same year
		if numericDates[i][2] != numericDates[i-1][2] {
			continue
		}

		if numericDates[i][0] < numericDates[i-1][0] {
			daysDecreasing = true
		}
		if numericDates[i][1] < numericDates[i-1][1] {
			monthsDecreasing = true
		}
	}

	if daysDecreasing {
		return DateOrderDaysFirst, nil
	} else if monthsDecreasing {
		return DateOrderMonthsFirst, nil
	}

	return DateOrderUnknown, nil
}

// checkChangeFrequency analyzes an array of numeric dates to determine if the days come
// before the month or vice versa by looking at which number changes more frequently.
func checkChangeFrequency(numericDates []NumericDate) (DateOrder, error) {
	if len(numericDates) < 2 {
		return DateOrderUnknown, nil
	}

	var firstTotal, secondTotal int

	for i := 1; i < len(numericDates); i++ {
		if len(numericDates[i]) < 3 || len(numericDates[i-1]) < 3 {
			return DateOrderUnknown, errors.New("invalid date format")
		}

		firstDiff := abs(numericDates[i][0] - numericDates[i-1][0])
		secondDiff := abs(numericDates[i][1] - numericDates[i-1][1])

		firstTotal += firstDiff
		secondTotal += secondDiff
	}

	if firstTotal > secondTotal {
		return DateOrderDaysFirst, nil
	}
	if firstTotal < secondTotal {
		return DateOrderMonthsFirst, nil
	}

	return DateOrderUnknown, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Takes an array of numeric dates and tries to understand if the days come
// before the month or the other way around by running the dates through various
// checks.
func InferDateOrder(numericDates []NumericDate) (DateOrder, error) {
	dateOrder, err := checkAbove12(numericDates)
	if err != nil {
		return DateOrderUnknown, err
	}

	if dateOrder != DateOrderUnknown {
		return dateOrder, nil
	}

	dateOrder, err = checkDecreasing(numericDates)
	if err != nil {
		return DateOrderUnknown, err
	}

	if dateOrder != DateOrderUnknown {
		return dateOrder, nil
	}

	return checkChangeFrequency(numericDates)
}

// OrderDateComponents reorders the date components to ensure the longest number
// (year) is at the end.
func OrderDateComponents(date string) (DateComponents, error) {
	parts := RegexSplitDate.Split(date, -1)

	if len(parts) != 3 {
		return DateComponents{}, errors.New("invalid date components")
	}

	// Find the index of the longest part
	maxLength := len(parts[0])
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > maxLength {
			maxLength = len(parts[i])
		}
	}

	// Rearrange so that the longest part is at the end
	if len(parts[2]) == maxLength {
		return DateComponents{parts[0], parts[1], parts[2]}, nil
	} else if len(parts[1]) == maxLength {
		return DateComponents{parts[0], parts[2], parts[1]}, nil
	}

	return DateComponents{parts[1], parts[2], parts[0]}, nil
}

// Takes `year`, `month` and `day` as strings and pads them to `4`, `2`, `2`
// digits respectively.
func NormalizeDate(year string, month string, day string) (string, string, string) {
	return PadLeft(year, 4, "2000"), PadLeft(month, 2, "0"), PadLeft(day, 2, "0")
}

// Convert DateComponents into NumericDate
func DateComponentsToNumericDate(dc DateComponents) (NumericDate, error) {
	if len(dc) != 3 {
		return NumericDate{}, errors.New("invalid date components")
	}

	numDate := NumericDate{}

	for i, comp := range dc {
		num, err := strconv.Atoi(comp)
		if err != nil {
			return NumericDate{}, fmt.Errorf("failed converting DateComponents into NumericDate: %w", err)
		}

		numDate[i] = num
	}

	return numDate, nil
}
