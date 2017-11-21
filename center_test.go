package sunrise

import (
	"testing"
)

var dataEquationOfCenter = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - 5 degrees east longitude
	{358.29314, -0.05825},
	// 2000-01-01 - Toronto (-79.38)
	{357.74642, -0.0769},
	// 2016-04-01 - prime meridian
	{87.06676, 1.91404},
}

func TestEquationOfCenter(t *testing.T) {
	for _, tt := range dataEquationOfCenter {
		v := EquationOfCenter(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
