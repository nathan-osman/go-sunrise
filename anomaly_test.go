package sunrise

import (
	"testing"
	"math"
)

// Adapted from user korya on GitHub (https://gist.github.com/DavidVaini/10308388#gistcomment-1391788)
func Round(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return math.Floor(f * shift + .5) / shift;
}

var dataSolarAnomaly = []struct {
	in float64
	out float64
}{
	// April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html)
	{1552.01468889, 87.19521},
	// Prime meridian on January 1, 1990
	{2447893, 281.55531},
	// Toronto. January 1, 2000
	{2451545.2205,281.18486 },
}


func TestSolarAnomaly(t *testing.T) {
	for _, tt := range dataSolarAnomaly {
		v := GetSolarAnomaly(tt.in)
		if Round(v, 5) != Round(tt.out, 5) {
			t.Fatalf("%f != %f. %t and %t", v, tt.out, v, tt.out)
		}
	}
}




