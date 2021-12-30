package geo

import "github.com/golang/geo/r3"

// A parametrized curve (in our case, a straight line) which starts from
// an origin with direction in R3.
type Ray struct {
	Origin    r3.Vector
	Direction r3.Vector
}

// At returns the vector (at t) in R3 for a ray parametrized by r = a + tb,
// where a is `r.Origin` and b is `r.Direction`.
func (r Ray) At(t float64) r3.Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}
