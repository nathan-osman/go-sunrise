package sunrise

import (
	"testing"
)

var dataHourAngle = []struct {
	inLatitude          float64
	inEclipticLongitude float64
	out                 float64
}{
	// 1970-01-01 - prime meridian
	{0, 281.08372, 90.9018},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{43.65, 280.59957, 67.44918},
	// 2004-04-01 - (52° N, 5° E)
	{52, 12.02474, 97.47282},
}

func TestHourAngle(t *testing.T) {
	for _, tt := range dataHourAngle {
		v := HourAngle(tt.inLatitude, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
