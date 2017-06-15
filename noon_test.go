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
	// Prime meridian on January 1, 1990
	{0, 1990, 1, 1, 2447893},
	// Toronto on January 1, 2000
	{-79.38, 2000, 1, 1, 2451545.2205},
}

func TestMeanSolarNoon(t *testing.T) {
	for _, tt := range dataTestMeanSolarNoon {
		v := MeanSolarNoon(tt.inLongitude, tt.inYear, tt.inMonth, tt.inDay)
		if v != tt.out {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
