package sunrise

import "testing"

var dataSolarTransit = []struct {
	inSolarNoon         float64
	inSolarAnomaly      float64
	inEclipticLongitude float64
	out                 float64
}{
	// 1970-01-01 - 5 degrees east longitude
	{2440587.98611, 358.29314, 281.06956, 4892133.48545},
	// 2000-01-01 - Toronto (-79.38)
	{2451545.2205, 357.74642, 280.59957, 4903090.71229},
	// 2004-04-01 - prime meridian
	{2453097, 87.18073, 12.038440, 4904642.50248},
}

func TestSolarTransit(t *testing.T) {
	for _, tt := range dataSolarTransit {
		v := SolarTransit(tt.inSolarNoon, tt.inSolarAnomaly, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
