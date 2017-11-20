package sunrise

import (
	"time"
)

// J2000 is the Julian date for January 1, 2000, 12:00:00 TT.
const J2000 = 2451545

const (
	secondsInADay      = 86400
	unixEpochJulianDay = 2440587.5
)

// TimeToJulianDay converts a time.Time into a Julian day.
func TimeToJulianDay(t time.Time) float64 {
	return float64(t.Unix())/secondsInADay + unixEpochJulianDay
}

// JulianDayToTime converts a Julian day into a time.Time.
func JulianDayToTime(d float64) time.Time {
	return time.Unix(int64((d-unixEpochJulianDay)*secondsInADay), 0)
}
