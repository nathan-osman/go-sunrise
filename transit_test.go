package sunrise

import "testing"

var dataSolarTransit = []struct {
	inSolarNoon         float64
	inSolarAnomaly      float64
	inEclipticLongitude float64
	out                 float64
}{
	// 1970-01-01 - prime meridian
	{2440588, 358.30683, 281.08372, 4892133.4996},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{2451545.2205, 357.74642, 280.59957, 4903090.71229},
	// 2004-04-01 - (52° N, 5° E)
	{2453096.98611, 87.16704, 12.02474, 4904642.48863},
}

func TestSolarTransit(t *testing.T) {
	for _, tt := range dataSolarTransit {
		v := SolarTransit(tt.inSolarNoon, tt.inSolarAnomaly, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
