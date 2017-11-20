package sunrise

import (
	"testing"
)

var dataSolarMeanAnomaly = []struct {
	in  float64
	out float64
}{
	// April 1, 2004 - 5 degrees east longitude
	{2453097, 87.18073},
	// January 1, 1990 - prime meridian
	{2447893, -1.88312},
	// January 1, 2000 - Toronto
	{2451545.2205, 357.74642},
}

func TestSolarMeanAnomaly(t *testing.T) {
	for _, tt := range dataSolarMeanAnomaly {
		v := SolarMeanAnomaly(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
