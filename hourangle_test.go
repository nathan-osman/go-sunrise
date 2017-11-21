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
	{0, 281.08372, -77.94743},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{43.65, 280.59957, -78.471634},
	// 2004-04-01 - (52° N, 5° E)
	{52, 12.02474, 11.05801},
}

func TestHourAngle(t *testing.T) {
	for _, tt := range dataHourAngle {
		v := HourAngle(tt.inLatitude, tt.inEclipticLongitude)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
