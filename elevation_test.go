package sunrise

import (
	"testing"
	"time"
)

func abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}

// Sunrise is defined to be when the Sun is 50 arc minutes below the horizon.
// This is due to atmospheric refraction and measuring the position of the top
// rather than the centre of the Sun.
// https://en.wikipedia.org/wiki/Sunrise#Angle
var sunriseElevation = -50.0 / 60.0

var dataElevation = []struct {
	inLatitude  float64
	inLongitude float64
	inElevation float64
	inYear      int
	inMonth     time.Month
	inDay       int
	outFirst    time.Time
	outSecond   time.Time
}{
	// 1970-01-01 - prime meridian
	{
		0, 0, sunriseElevation,
		1970, time.January, 1,
		time.Date(1970, time.January, 1, 5, 59, 54, 0, time.UTC),
		time.Date(1970, time.January, 1, 18, 07, 07, 0, time.UTC),
	},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{
		43.65, -79.38, sunriseElevation,
		2000, time.January, 1,
		time.Date(2000, time.January, 1, 12, 51, 00, 0, time.UTC),
		time.Date(2000, time.January, 1, 21, 50, 36, 0, time.UTC),
	},
	// 2004-04-01 - (52° N, 5° E)
	{
		52, 5, sunriseElevation,
		2004, time.April, 1,
		time.Date(2004, time.April, 1, 5, 13, 40, 0, time.UTC),
		time.Date(2004, time.April, 1, 18, 13, 27, 0, time.UTC),
	},
	// 2020-06-15 - Igloolik, Canada
	{
		69.3321443, -81.6781126, sunriseElevation,
		2020, time.June, 25,
		time.Time{},
		time.Time{},
	},
	// 2022-08-27 - London, end of Shabbat
	{
		51.5072, -0.1276, -8.5,
		2022, time.August, 27,
		time.Date(2022, time.August, 27, 4, 11, 31, 0, time.UTC),
		time.Date(2022, time.August, 27, 19, 53, 6, 0, time.UTC),
	},
	// 2022-06-21 - London, highest point around noon
	{
		51.5072, -0.1276, 61.93,
		2022, time.June, 21,
		time.Date(2022, time.June, 21, 12, 0, 12, 0, time.UTC),
		time.Date(2022, time.June, 21, 12, 4, 16, 0, time.UTC),
	},
	// 2022-06-21 - London, too high, never reached
	{
		51.5072, -0.1276, 61.94,
		2022, time.June, 21,
		time.Time{},
		time.Time{},
	},
	// 2022-06-21 - London, too low, never reached
	{
		51.5072, -0.1276, -16,
		2022, time.June, 21,
		time.Time{},
		time.Time{},
	},
	// 2022-11-26 - New York City, end of Shabbat
	{
		40.7128, -74.006, -8.5,
		2022, time.November, 26,
		time.Date(2022, time.November, 26, 11, 11, 19, 0, time.UTC),
		time.Date(2022, time.November, 26, 22, 15, 46, 0, time.UTC),
	},
}

func TestTimeOfElevation(t *testing.T) {
	for _, tt := range dataElevation {
		vFirst, vSecond := TimeOfElevation(tt.inLatitude, tt.inLongitude, tt.inElevation, tt.inYear, tt.inMonth, tt.inDay)
		if abs(vFirst.Unix()-tt.outFirst.Unix()) > 2 {
			t.Fatalf("%s != %s", vFirst.String(), tt.outFirst.String())
		}
		if abs(vSecond.Unix()-tt.outSecond.Unix()) > 2 {
			t.Fatalf("%s != %s", vSecond.String(), tt.outSecond.String())
		}
	}
}
