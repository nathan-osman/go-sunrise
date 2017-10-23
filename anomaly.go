package sunrise

import (
	"math"
)

func GetSolarAnomaly(solarNoon float64) float64 {
	return math.Mod(357.5291+0.98560028*solarNoon, 360)
}
