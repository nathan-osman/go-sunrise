package sunrise

import (
	"testing"
)

var dataSolarMeanAnomaly = []struct {
	in  float64
	out float64
}{
	// April 1, 2004 12:00:00 UTC - 5 degrees east longitude
	{2453097, 87.18073},
	// January 1, 2000 12:00:00 EST - Toronto
	{2451545.20833, 357.73443},
	// January 1, 1990 12:00:00 UTC - prime meridian
	{2447893, 358.11688},
}

func TestSolarMeanAnomaly(t *testing.T) {
	for _, tt := range dataSolarMeanAnomaly {
		v := SolarMeanAnomaly(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
