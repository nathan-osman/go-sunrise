package sunrise

import (
	"math"
)

func EclipticLongitude(solarAnomaly float64, equationOfCenter float64, jDate float64) float64 {
	return math.Mod(solarAnomaly+equationOfCenter+180+ArgumentOfPerihelion(jDate), 360)
}
