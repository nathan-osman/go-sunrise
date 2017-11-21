package sunrise

import (
	"testing"
)

var dataEclipticLongitude = []struct {
	anomalyIn float64
	centerIn  float64
	jDateIn   float64
	out       float64
}{
	// 1970-01-01 - prime meridian
	{358.30683, -0.05778, 2440588, 281.08372},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{357.74642, -0.0769, 2451545.2205, 280.59957},
	// 2004-04-01 - (52° N, 5° E)
	{87.16704, 1.91414, 2453096.98611, 12.02474},
}

func TestEclipticLongitude(t *testing.T) {
	for _, tt := range dataEclipticLongitude {
		v := EclipticLongitude(tt.anomalyIn, tt.centerIn, tt.jDateIn)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
