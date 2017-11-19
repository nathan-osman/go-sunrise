package sunrise

import (
	"testing"
)

var dataTestEclipticLongitude = []struct {
	anomalyIn float64
	centerIn  float64
	jDateIn   float64
	out       float64
}{
	// April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html)
	// The problem was slightly modified to take Julian leap seconds into account in anomaly.go. Expect small deviations from the above site.
	{87.19521, 1.914164, 2453097, 12.05293},
	// Prime meridian on January 1, 1990
	{281.55531, -1.88359, 2447892.5, 202.56997},
	// Toronto. January 1, 2000
	{281.18486, -1.88579, 2451544.5, 202.22912},
}

func TestEclipticLongitude(t *testing.T) {
	for _, tt := range dataTestEclipticLongitude {
		v := Round(EclipticLongitude(tt.anomalyIn, tt.centerIn, tt.jDateIn), DefaultPlaces)
		if v != tt.out {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
