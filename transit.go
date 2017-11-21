package sunrise

import "math"

// SolarTransit calculates the Julian data for the local true solar transit.
func SolarTransit(d, solarAnomaly, eclipticLongitude float64) float64 {
	return 2451545.5 + d + 0.005*math.Sin(solarAnomaly) -
		0.0069*math.Sin(2*eclipticLongitude)
}
