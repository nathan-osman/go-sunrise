package sunrise

import "math"

const places int = 5

// Adapted from user korya on GitHub (https://gist.github.com/DavidVaini/10308388#gistcomment-1391788)
func Round(f float64, places int) (float64) {
	shift := math.Pow(10, float64(places))
	return math.Floor(f * shift + .5) / shift;
}
