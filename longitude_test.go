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
	// 1970-01-01 - 5 degrees east longitude
	{358.29314, -0.05825, 2440587.98611, 281.06956},
	// 2000-01-01 - Toronto (-79.38)
	{357.74642, -0.0769, 2451545.2205, 280.59957},
	// 2004-04-01 - prime meridian
	{87.18073, 1.91415, 2453097, 12.038440},
}

func TestEclipticLongitude(t *testing.T) {
	for _, tt := range dataEclipticLongitude {
		v := EclipticLongitude(tt.anomalyIn, tt.centerIn, tt.jDateIn)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
