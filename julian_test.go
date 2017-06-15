package sunrise

import (
	"testing"
	"time"
)

var dataTimeToJulianDay = []struct {
	in  time.Time
	out float64
}{
	{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), unixEpochJulianDay},
	{time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC), 2449718.5},
	{time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC), 2458850},
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
	{unixEpochJulianDay, time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
	{2449718.5, time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC)},
	{2458850, time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)},
}

func TestJulianDayToTime(t *testing.T) {
	for _, tt := range dataJulianDayToTime {
		v := JulianDayToTime(tt.in).UTC()
		if v != tt.out {
			t.Fatalf("%s != %s", v.String(), tt.out.String())
		}
	}
}
