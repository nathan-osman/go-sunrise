package sunrise

import (
	"testing"
)

var dataHourAngle = []struct {
	inLatitude    float64
	inDeclination float64
	out           float64
}{
	// 1970-01-01 - prime meridian
	{0, -22.97753, 90.9018},
	// 2000-01-01 - Toronto (43.65° N, 79.38° W)
	{43.65, -23.01689, 67.44917},
	// 2004-04-01 - (52° N, 5° E)
	{52, 4.75374, 97.47283},
}

func TestHourAngle(t *testing.T) {
	for _, tt := range dataHourAngle {
		v := HourAngle(tt.inLatitude, tt.inDeclination)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
