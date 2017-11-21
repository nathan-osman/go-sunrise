package sunrise

import (
	"math"
)

const degree = math.Pi / 180

// Declination calculates the one of the angles required to locate a point on
// the celestial sphere in the equatorial coordinate system. The ecliptic
// longitude parameter must be in degrees.
func Declination(eclipticLongitude float64) float64 {
	return math.Asin(math.Sin(eclipticLongitude*degree)*math.Sin(0.40911)) / degree
}
