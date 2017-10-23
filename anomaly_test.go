package sunrise

import (
	"testing"
)

var dataSolarAnomaly = []struct {
	in  float64
	out float64
}{
	// April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html)
	// The problem was slightly modified to take Julian leap seconds into account in anomaly.go. Expect small deviations from the above site.
	{1552.01468889, 87.19521},
	// Prime meridian on January 1, 1990
	{2447893, 281.55531},
	// Toronto. January 1, 2000
	{2451545.2205, 281.18486},
}

func TestSolarAnomaly(t *testing.T) {
	for _, tt := range dataSolarAnomaly {
		v := GetSolarAnomaly(tt.in)
		if Round(v, places) != Round(tt.out, places) {
			t.Fatalf("%f != %f. %t and %t", v, tt.out, v, tt.out)
		}
	}
}
