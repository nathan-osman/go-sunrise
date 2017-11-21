package sunrise

import (
	"math"
)

// HourAngle calculates the second of the two angles required to locate a point
// on the celestial sphere in the equatorial coordinate system.
func HourAngle(latitude, eclipticLongitude float64) float64 {
	var (
		declination = math.Asin(
			math.Sin(eclipticLongitude*Degree)*0.39779,
		) / Degree
		numerator   = -0.01449 - math.Sin(latitude*Degree)*math.Sin(declination*Degree)
		denominator = math.Cos(latitude*Degree) * math.Cos(declination*Degree)
	)
	return math.Acos(numerator/denominator) / Degree
}
