package sunrise

import (
	"testing"
	"time"
)

var dataTimeToJulianDay = []struct {
	in  time.Time
	out float64
}{
	// 1970-01-01 00:00:00 UTC - 5 degrees east longitude
	{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), unixEpochJulianDay},
	// 2000-01-01 12:00:00 UTC - Toronto (-79.38)
	{time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC), J2000},
	// 2016-04-01 12:00:00 UTC - prime meridian
	{time.Date(2016, 4, 1, 12, 0, 0, 0, time.UTC), 2457480},
}

func TestTimeToJulianDay(t *testing.T) {
	for _, tt := range dataTimeToJulianDay {
		v := TimeToJulianDay(tt.in)
		if v != tt.out {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}

var dataJulianDayToTime = []struct {
	in  float64
	out time.Time
}{
	// 1970-01-01 00:00:00 UTC - 5 degrees east longitude
	{unixEpochJulianDay, time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
	// 2000-01-01 12:00:00 UTC - Toronto (-79.38)
	{J2000, time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)},
	// 2016-04-01 12:00:00 UTC - prime meridian
	{2457480, time.Date(2016, 4, 1, 12, 0, 0, 0, time.UTC)},
}

func TestJulianDayToTime(t *testing.T) {
	for _, tt := range dataJulianDayToTime {
		v := JulianDayToTime(tt.in).UTC()
		if v != tt.out {
			t.Fatalf("%s != %s", v.String(), tt.out.String())
		}
	}
}
