package sunrise

import (
	"testing"
)

var dataSolarMeanAnomaly = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - prime meridian
	{2440588, 358.30683},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{2451545.2205, 357.74642},
	// 2004-04-01 - (52° N, 5° E)
	{2453096.98611, 87.16704},
}

func TestSolarMeanAnomaly(t *testing.T) {
	for _, tt := range dataSolarMeanAnomaly {
		v := SolarMeanAnomaly(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
