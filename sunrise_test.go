package sunrise

import (
	"testing"
	"time"
)

var dataSunriseSunset = []struct {
	inLatitude  float64
	inLongitude float64
	inYear      int
	inMonth     time.Month
	inDay       int
	outSunrise  time.Time
	outSunset   time.Time
}{
	// 1970-01-01 - prime meridian
	{
		0, 0,
		1970, time.January, 1,
		time.Date(1970, time.January, 1, 6, 51, 44, 0, time.UTC),
		time.Date(1970, time.January, 1, 17, 15, 18, 0, time.UTC),
	},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{
		43.65, -79.38,
		2000, time.January, 1,
		time.Date(2000, time.January, 1, 12, 6, 55, 0, time.UTC),
		time.Date(2000, time.January, 1, 22, 34, 41, 0, time.UTC),
	},
	// 2004-04-01 - (52° N, 5° E)
	{
		52, 5,
		2004, time.April, 1,
		time.Date(2004, time.April, 1, 12, 27, 48, 0, time.UTC),
		time.Date(2004, time.April, 1, 10, 59, 20, 0, time.UTC),
	},
}

func TestSunriseSunset(t *testing.T) {
	for _, tt := range dataSunriseSunset {
		vSunrise, vSunset := SunriseSunset(tt.inLatitude, tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if vSunrise != tt.outSunrise {
			t.Fatalf("%s != %s", vSunrise.String(), tt.outSunrise.String())
		}
		if vSunset != tt.outSunset {
			t.Fatalf("%s != %s", vSunset.String(), tt.outSunset.String())
		}
	}
}
