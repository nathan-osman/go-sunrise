package sunrise

import (
	"testing"
)

var dataEquationOfCenter = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - prime meridian
	{358.30683, -0.05778},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{357.74642, -0.0769},
	// 2004-04-01 - (52° N, 5° E)
	{87.16704, 1.91414},
}

func TestEquationOfCenter(t *testing.T) {
	for _, tt := range dataEquationOfCenter {
		v := EquationOfCenter(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
