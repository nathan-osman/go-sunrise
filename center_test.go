package sunrise

import (
	"testing"
)

var dataEquationOfCenter = []struct {
	in  float64
	out float64
}{
	// April 1, 2004 12:00:00 UTC - 5 degrees east longitude
	{87.18073, 1.91415},
	// January 1, 2000 12:00:00 EST - Toronto
	{357.73443, -0.07731},
	// January 1, 1990 12:00:00 UTC - prime meridian
	{358.11688, -0.06426},
}

func TestEquationOfCenter(t *testing.T) {
	for _, tt := range dataEquationOfCenter {
		v := EquationOfCenter(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
