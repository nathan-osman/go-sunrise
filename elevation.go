package sunrise

import (
	"math"
	"time"
)

// TimeOfElevation calculates the times of day when the sun is at a given elevation
// above the horizon on a given day at the specified location.
// Returns time.Time{} if there sun does not reach the elevation
func TimeOfElevation(latitude, longitude, elevation float64, year int, month time.Month, day int) (time.Time, time.Time) {
	var (
		d                 = MeanSolarNoon(longitude, year, month, day)
		solarAnomaly      = SolarMeanAnomaly(d)
		equationOfCenter  = EquationOfCenter(solarAnomaly)
		eclipticLongitude = EclipticLongitude(solarAnomaly, equationOfCenter, d)
		solarTransit      = SolarTransit(d, solarAnomaly, eclipticLongitude)
		declination       = Declination(eclipticLongitude)
		// https://solarsena.com/solar-elevation-angle-altitude/
		numerator        = math.Sin(elevation*Degree) - (math.Sin(latitude*Degree) * math.Sin(declination*Degree))
		denominator      = math.Cos(latitude*Degree) * math.Cos(declination*Degree)
		hourAngle        = math.Acos(numerator / denominator)
		secondsAfterNoon = hourAngle / (2 * math.Pi)
		morning          = solarTransit - secondsAfterNoon
		evening          = solarTransit + secondsAfterNoon
	)

	// Check for cases where the sun never reaches the given elevation.
	if math.IsNaN(hourAngle) {
		return time.Time{}, time.Time{}
	}

	return JulianDayToTime(morning), JulianDayToTime(evening)
}
