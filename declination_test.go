package sunrise

import (
	"testing"
)

var dataDeclination = []struct {
	in  float64
	out float64
}{
	// 1970-01-01 - 5 degrees east longitude
	{281.06956, -22.97889},
	// 2000-01-01 - Toronto (-79.38)
	{280.59957, -23.01707},
	// 2004-04-01 - prime meridian
	{12.038440, 4.75912},
}

func TestDeclination(t *testing.T) {
	for _, tt := range dataDeclination {
		v := Declination(tt.in)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
