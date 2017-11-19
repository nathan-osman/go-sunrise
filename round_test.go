package sunrise

import (
	"strconv"
	"testing"
)

var dataRound = []struct {
	in       float64
	inPlaces int
	out      string
}{
	{1.500, 0, "2.00000"},
	{1.949, 1, "1.90000"},
	{1.999, 2, "2.00000"},
}

func TestRound(t *testing.T) {
	for _, tt := range dataRound {
		v := strconv.FormatFloat(Round(tt.in, tt.inPlaces), 'f', 5, 64)
		if v != tt.out {
			t.Fatalf("%s != %s", v, tt.out)
		}
	}
}
