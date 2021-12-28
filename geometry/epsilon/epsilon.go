package epsilon

// see https://github.com/downflux/go-geometry/epsilon

import (
	"math"
)

const (
	epsilon = 1e-12
)

func AlmostEqual(a float64, b float64) bool {
	if (a == math.Inf(-1) && b == math.Inf(-1)) || (a == math.Inf(0) && b == math.Inf(0)) {
		return true
	}
	return math.Abs(a-b) < epsilon
}
