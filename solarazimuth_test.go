package sunrise

import (
	"math"
	"testing"
	"time"
)

func TestSolarAzimuth(t *testing.T) {
	tests := []struct {
		name        string
		inLatitude  float64
		inLongitude float64
		inTime      time.Time
		outAzimuth  float64
	}{
		{
			name:        "2023-04-27 Kyiv 1:00 am",
			inLatitude:  50.45466,
			inLongitude: 30.5238,
			inTime:      time.Date(2023, time.April, 27, 1, 0, 0, 0, time.UTC),
			outAzimuth:  46.23,
		},
		{
			name:        "1970-01-01 Prime meridian 5:59 am",
			inLatitude:  0,
			inLongitude: 0,
			inTime:      time.Date(1970, time.January, 1, 5, 59, 0, 0, time.UTC),
			outAzimuth:  113.04,
		},
		{
			name:        "2000-01-14 Toronto 3:42 pm",
			inLatitude:  43.65,
			inLongitude: -79.38,
			inTime:      time.Date(2000, time.January, 14, 15, 42, 0, 0, time.UTC),
			outAzimuth:  154.01,
		},
		{
			name:        "2024-02-29 London 6:05 pm",
			inLatitude:  51.5072,
			inLongitude: -0.1276,
			inTime:      time.Date(2024, time.February, 29, 18, 0, 0, 0, time.UTC),
			outAzimuth:  263.74,
		},
		{ // Sun directly overhead, elevation 89.9°
			name:        "2023-09-23 AUTUMNAL EQUINOX 7:25 am",
			inLatitude:  0,
			inLongitude: 66.8928,
			inTime:      time.Date(2023, time.September, 23, 7, 25, 0, 0, time.UTC),
			outAzimuth:  6.38,
		},
		{
			name:        "2022-11-26 New York City 11:11 am",
			inLatitude:  40.7128,
			inLongitude: -74.006,
			inTime:      time.Date(2022, time.November, 26, 11, 11, 0, 0, time.UTC),
			outAzimuth:  110.41,
		},
	}
	for _, tt := range tests {
		if out := SolarAzimuth(tt.inLatitude, tt.inLongitude, tt.inTime); math.Abs(out-tt.outAzimuth) > 2 {
			t.Fatalf("%s: %f != %f", tt.name, out, tt.outAzimuth)
		}
	}
}
