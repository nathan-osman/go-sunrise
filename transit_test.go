package sunrise

import "testing"

var dataSolarTransit = []struct {
	inSolarNoon         float64
	inSolarAnomaly      float64
	inEclipticLongitude float64
	out                 float64
}{
	// April 1, 2004 12:00:00 UTC - 5 degrees east longitude
	{2440587.98611, 358.29314, 281.06956, 4892133.48545},
	// January 1, 2000 12:00:00 EST - Toronto
	{2451545.2205, 357.74642, 280.59957, 4903090.71229},
	// January 1, 1990 12:00:00 UTC - prime meridian
	{2457480, 87.06676, 11.96251, 4909025.50254},
}

func TestSolarTransit(t *testing.T) {
	for _, tt := range dataSolarTransit {
		v := SolarTransit(tt.inSolarNoon, tt.inSolarAnomaly, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
