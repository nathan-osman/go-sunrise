package sunrise

import (
	"math"
)

const (
	// Argument of perihelion
	w0 float64 = 102.93005
	// Degrees of change in the argument per century
	wt float64 = 0.31795
	// Julian date in 2000
	j2000 float64 = 2451545
)

func UpdatePerihelion(jDate float64) float64 {
	// el is the number of centuries since j2000, when w0 was calculated. 36525 is the number of days in a century
	el := (jDate - j2000) / 36525
	return w0 + wt*el

}

func EclipticLongitude(solarAnomaly float64, equationOfCenter float64, jDate float64) float64 {
	return math.Mod(solarAnomaly+equationOfCenter+180+UpdatePerihelion(jDate), 360)
}
