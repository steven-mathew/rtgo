package vector

import (
	"math"

	"github.com/steven-mathew/rtgo/geometry/epsilon"
)

const DIMENSION = 3

// V returns a point in R^3.
type V struct {
	X, Y, Z float64
}

func New(x, y, z float64) *V {
	return &V{x, y, z}
}

func (v V) Add(u V) V {
	return V{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v V) Sub(u V) V {
	return V{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v V) Mul(a float64) V {
	return V{a * v.X, a * v.Y, a * v.Z}
}

func (v V) Dot(u V) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

// SqL2Norm returns the Squared L2-Norm of an R^3 vector.
func (v V) SqL2Norm() float64 {
	return v.Dot(v)
}

func (v V) ToUnit() V {
	snorm := v.SqL2Norm()
	if snorm == 0 {
		return V{0, 0, 0}
	}
	return v.Mul(1. / math.Sqrt(snorm))
}

func (v V) Cross(u V) V {
	return V{
		v.Y*u.Z - v.Z*u.Y,
		v.Z*u.X - v.X*u.Z,
		v.X*u.Y - v.Y*u.X,
	}
}

func (v V) IsOrthogonal(u V) bool {
	return epsilon.AlmostEqual(v.Dot(u), 0.)
}

func (v V) AlmostEqual(u V) bool {
	return epsilon.AlmostEqual(v.X, u.X) && epsilon.AlmostEqual(v.X, u.X) && epsilon.AlmostEqual(v.X, u.X)
}
