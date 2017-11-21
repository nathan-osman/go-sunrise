package sunrise

import "testing"

var dataSolarTransit = []struct {
	inSolarNoon         float64
	inSolarAnomaly      float64
	inEclipticLongitude float64
	out                 float64
}{
	// April 1, 2004 12:00:00 UTC - 5 degrees east longitude
	{1552.01468889, 87.19521, 12.05293, 2453096.9393},
	// January 1, 2000 12:00:00 EST - Toronto
	{2447893, 281.55531, 202.56998, 4899438.48992},
	// January 1, 1990 12:00:00 UTC - prime meridian
	{2451545.2205, 201.18486, 202.22912, 4903090.71375},
}

func TestSolarTransit(t *testing.T) {
	for _, tt := range dataSolarTransit {
		v := SolarTransit(tt.inSolarNoon, tt.inSolarAnomaly, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
