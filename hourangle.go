package sunrise

import (
	"math"
)

// HourAngle calculates the second of the two angles required to locate a point
// on the celestial sphere in the equatorial coordinate system.
func HourAngle(latitude, eclipticLongitude float64) float64 {
	eclipticLongitudeRad := eclipticLongitude * Degree
	return math.Atan2(
		math.Sin(eclipticLongitudeRad)*math.Cos(0.40909),
		math.Cos(eclipticLongitudeRad),
	) / Degree
}
