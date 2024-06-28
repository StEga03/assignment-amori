// Based on Kero package timeutil.
package timeutils

import (
	"sync"
	"time"

	"github.com/assignment-amori/internal/constant"
)

// TimeLocations represent many time location.
type TimeLocations struct {
	WIB *time.Location
}

var tl *TimeLocations
var tlOnce sync.Once

// NewTimeLocation is a function that will initialize some time location based on server time.
func NewTimeLocation() error {
	var err error

	// Using once because when run go cover test it will run concurrently, so doesn't need to call this many time.
	tlOnce.Do(func() {
		tl = new(TimeLocations)
		err = initLocationWIB()
		if err != nil {
			return
		}

	})

	return err
}

func initLocationWIB() (err error) {
	tl.WIB, err = time.LoadLocation("Asia/Jakarta")
	return err
}

// GetTimeLocation is a function that will grab time location based on global var,
// that fulfilled when app initialize.
func GetTimeLocation() *TimeLocations {
	return tl
}

func GetWIBCurrentTime() time.Time {
	wibTime, _ := convertTimeToWIBTime(time.Now())
	return wibTime
}

func ConvertUTCtoWIB(t time.Time) (time.Time, error) {
	wibTime, err := convertTimeToWIBTime(t)
	if err != nil {
		return time.Time{}, err
	}
	return subtractHoursFromTime(wibTime, constant.WIBTimeDifferenceHrs), nil
}

func convertTimeToWIBTime(t time.Time) (time.Time, error) {
	loc := GetTimeLocation().WIB
	return t.In(loc), nil
}

func subtractHoursFromTime(t time.Time, duration time.Duration) time.Time {
	return t.Add(-duration * time.Hour)
}

func GetDateFromTime(t time.Time) time.Time {
	curDate, _ := time.Parse(constant.CustomAbsoluteDateFormat, t.Format(constant.CustomAbsoluteDateFormat))
	return curDate
}
