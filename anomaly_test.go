package sunrise

import (
	"testing"
)

var dataSolarMeanAnomaly = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - 5 degrees east longitude
	{2440587.98611, 358.29314},
	// 2000-01-01 - Toronto (-79.38)
	{2451545.2205, 357.74642},
	// 2016-04-01 - prime meridian
	{2457480, 87.06676},
}

func TestSolarMeanAnomaly(t *testing.T) {
	for _, tt := range dataSolarMeanAnomaly {
		v := SolarMeanAnomaly(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
