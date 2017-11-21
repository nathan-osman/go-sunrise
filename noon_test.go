package sunrise

import (
	"testing"
	"time"
)

var dataTestMeanSolarNoon = []struct {
	inLongitude float64
	inYear      int
	inMonth     time.Month
	inDay       int
	out         float64
}{
	// 1970-01-01 - prime meridian
	{0, 1970, 1, 1, 2440588},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{-79.38, 2000, 1, 1, 2451545.2205},
	// 2004-04-01 - (52° N, 5° E)
	{5, 2004, 4, 1, 2453096.98611},
}

func TestMeanSolarNoon(t *testing.T) {
	for _, tt := range dataTestMeanSolarNoon {
		v := MeanSolarNoon(tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
