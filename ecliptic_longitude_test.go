package sunrise

import (
	"testing"
)

var dataTestUpdatePerihelion = []struct {
	in float64
	out float64
}{
	// January 1 2017. Example from the wikipedia article.
	{2457784.958333, 102.98437},
	// April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html)
	// The problem was slightly modified to take Julian leap seconds into account in anomaly.go. Expect small deviations from the above site.
	{2453097,102.94356},
	// Prime meridian, January 1, 1990
	{2447892.5, 102.89826},
	// Toronto, January 1, 2000
	{2451544.5, 102.93005},
}

var dataTestEclipticLongitude = []struct {
	anomalyIn float64
	centerIn float64
	jDateIn float64
	out float64
}{
	// April 1, 2004. 5 degrees east longitude. This is a test problem from (http://aa.quae.nl/en/reken/zonpositie.html)
	// The problem was slightly modified to take Julian leap seconds into account in anomaly.go. Expect small deviations from the above site.
	{87.19521, 1.914164, 2453097, 12.05293},
	// Prime meridian on January 1, 1990
	{281.55531, -1.88359, 2447892.5, 202.56998},
	// Toronto. January 1, 2000
	{281.18486, -1.88579, 2451544.5, 202.22912},
}

func TestUpdatePerihelion(t *testing.T) {
	for _, tt := range dataTestUpdatePerihelion {
		v := Round(UpdatePerihelion(tt.in), places)
		if v != tt.out {
			t.Fatalf("%f != %f. Of type %t and %t", v, tt.out, v, tt.out)
		}
	}
}

func TestEclipticLongitude(t *testing.T) {
	for _, tt := range dataTestEclipticLongitude {
		v := Round(EclipticLongitude(tt.anomalyIn, tt.centerIn, tt.jDateIn), places)
		if v != tt.out {
			t.Log(tt)
			t.Fatalf("%f != %f. Of type %t and %t", v, tt.out, v, tt.out)
		}
	}
}

