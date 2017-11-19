package sunrise

import (
	"testing"
)

var dataArgumentOfPerihelion = []struct {
	in  float64
	out float64
}{
	// January 1 2017 - example from the wikipedia article
	{2457784.958333, 102.98437},
	// April 1, 2004 - 5 degrees east longitude; this is a test problem from
	// (http://aa.quae.nl/en/reken/zonpositie.html) - the problem was slightly
	// modified to take Julian leap seconds into account in anomaly.go; expect
	// small deviations from the above site
	{2453097, 102.94356},
	// Prime meridian - January 1, 1990
	{2447892.5, 102.89825},
	// Toronto - January 1, 2000
	{2451544.5, 102.93005},
}

func TestArgumentOfPerihelion(t *testing.T) {
	for _, tt := range dataArgumentOfPerihelion {
		v := ArgumentOfPerihelion(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
