package sunrise

import (
	"math"
	"time"
)

// SolarAzimuth calculates the solar azimuth angle at a given moment
// at the specified location.
func SolarAzimuth(latitude, longitude float64, when time.Time) float64 {
	var (
		d                 = MeanSolarNoon(longitude, when.Year(), when.Month(), when.Day())
		solarAnomaly      = SolarMeanAnomaly(d)
		equationOfCenter  = EquationOfCenter(solarAnomaly)
		eclipticLongitude = EclipticLongitude(solarAnomaly, equationOfCenter, d)
		solarTransit      = SolarTransit(d, solarAnomaly, eclipticLongitude)
		declination       = Declination(eclipticLongitude)
		frac              = TimeToJulianDay(when) - solarTransit
		hourAngle         = 2 * math.Pi * frac
		elevation         = Elevation(latitude, longitude, when)
		firstPart         = math.Sin(declination*Degree) * math.Cos(latitude*Degree)
		secondPart        = math.Cos(declination*Degree) * math.Sin(latitude*Degree) * math.Cos(hourAngle)
		azimuth           = math.Acos((firstPart-secondPart)/math.Cos(elevation*Degree)) / Degree
	)
	if hourAngle >= 0 {
		azimuth = 360 - azimuth
	}
	return azimuth
}
