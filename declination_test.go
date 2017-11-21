package sunrise

import (
	"testing"
)

var dataDeclination = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - prime meridian
	{281.08372, -22.97753},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{280.59957, -23.01689},
	// 2004-04-01 - (52° N, 5° E)
	{12.02474, 4.75374},
}

func TestDeclination(t *testing.T) {
	for _, tt := range dataDeclination {
		v := Declination(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
