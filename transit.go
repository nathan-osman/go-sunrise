package sunrise

import "math"

func SolarTransit(solarNoon, solarAnomaly, eclipticLongitude float64) float64 {
	return 2451545.5 + solarNoon + 0.005*math.Sin(solarAnomaly) - 0.0069*math.Sin(2*eclipticLongitude)
}
