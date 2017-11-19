package sunrise

import (
	"testing"
)

var dataEquationOfCenter = []struct {
	in  float64
	out float64
}{
	// // April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html).
	// The problem was slightly modified to take Julian leap seconds into account in anomaly.go. Expect small deviations from the above site.
	{87.19521, 1.914164},
	// Prime meridian on January 1, 1990
	{281.55531, -1.88359},
	// Toronto. January 1, 2000
	{281.18486, -1.88579},
}

func TestGetEquationOfCenter(t *testing.T) {
	for _, tt := range dataEquationOfCenter {
		v := GetEquationOfCenter(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f. Of type %t and %t", v, tt.out, v, tt.out)
		}
	}
}
