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
	// 1970-01-01 - 5 degrees east longitude
	{5, 1970, 1, 1, 2440587.98611},
	// 2000-01-01 - Toronto (-79.38)
	{-79.38, 2000, 1, 1, 2451545.2205},
	// 2016-04-01 - prime meridian
	{0, 2016, 4, 1, 2457480},
}

func TestMeanSolarNoon(t *testing.T) {
	for _, tt := range dataTestMeanSolarNoon {
		v := MeanSolarNoon(tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
