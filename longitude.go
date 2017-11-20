package sunrise

import (
	"math"
)

// EclipticLongitude calculates the angular distance of the earth along the
// ecliptic.
func EclipticLongitude(meanAnomaly, equationOfCenter, d float64) float64 {
	return math.Mod(meanAnomaly+equationOfCenter+180+ArgumentOfPerihelion(d), 360)
}
