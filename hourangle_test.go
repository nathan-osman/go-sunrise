package sunrise

import (
	"testing"
)

var dataHourAngle = []struct {
	inLatitude    float64
	inDeclination float64
	out           float64
}{
	// 1970-01-01 - 5 degrees east longitude
	{0, -22.97889, 90.90178},
	// 2000-01-01 - Toronto (-79.38)
	{43.65, -23.01707, 67.44891},
	// 2004-04-01 - prime meridian
	{0, 4.75912, 90.83309},
}

func TestHourAngle(t *testing.T) {
	for _, tt := range dataHourAngle {
		v := HourAngle(tt.inLatitude, tt.inDeclination)
		if Round(v, DefaultPlaces) != Round(tt.out, DefaultPlaces) {
			t.Fatalf("%f != %f", v, tt.out)
		}
	}
}
