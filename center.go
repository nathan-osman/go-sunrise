package sunrise

import (
	"math"
)

// This module calculates the Equation of Center for the sun from Earth.


func GetEquationOfCenter(anomaly float64) float64 {
	// The numbers being multiplied below are coefficients for the Equation of Center for Earth
	var anomalyInRad float64
	anomalyInRad = anomaly * (math.Pi/180)
	return 1.9148 * (math.Sin(anomalyInRad)) + 0.0200 * math.Sin(2 * anomalyInRad) + 0.0003 * math.Sin(3 * anomalyInRad)
}