package camera

import (
	"math"
	"math/rand"

	"github.com/golang/geo/r3"
	"github.com/steven-mathew/rtgo/internal/pkg/geom"
)

type Camera struct {
	// Location of the camera
	Origin r3.Vector

	// Direction that the camera is facing
	Direction r3.Vector

	// Direction of "up" for the screen, must be orthogonal to `Direction`
	Up r3.Vector

	// Field of view in the longer direction as an angle in radians
	// TODO: find bounds
	VerticalFov float64

	// Apertrue radius for depth-of-field
	Aperture float64

	// FocalDistance, exists iff `Aperture` is non-zero
	FocalDistance float64
}

// NewCamera creates a camera placed at `from` and points towards `at`.
// Here, `up` specifies a vector which lies in the plane orthogonal to the view
// direction.
func NewCamera(
	from, at, up r3.Vector,
	vfov float64,
) Camera {
	direction := from.Sub(at).Normalize()
	u := direction.Mul(up.Dot(direction))

	up = up.Sub(u).Normalize()

	return Camera{
		Origin:        from,
		Direction:     direction,
		Up:            up,
		VerticalFov:   vfov,
		Aperture:      0.,
		FocalDistance: 0.,
	}
}

// Focus modifies the camera to focus on a position, with depth-of-field
func (c *Camera) Focus(focalPoint r3.Vector, aperture float64) {
	c.FocalDistance = focalPoint.Sub(c.Origin).Dot(c.Direction)
	c.Aperture = aperture
}

// CastRay returns a Ray, where s, t in [-1, 1]
func (c Camera) CastRay(s, t float64) geom.Ray {
	d := 1. / math.Tan(c.VerticalFov/2.0)
	right := c.Direction.Cross(c.Up).Normalize()

	origin := c.Origin
	newDirection := c.Direction.Mul(d).
		Add(right.Mul(s)).
		Add(c.Up.Mul(t)).
		Normalize()

	if c.Aperture > 0 {
		focalPoint := c.Origin.Add(newDirection).Mul(c.FocalDistance)
		s = -1 + rand.Float64()*2.
		t = -1 + rand.Float64()*2.

		u := right.Mul(s).Add(c.Up.Mul(t)).Mul(c.Aperture)
		origin = origin.Add(u)

		newDirection = focalPoint.Add(origin.Mul(-1))
	}

	return geom.Ray{
		Origin:    c.Origin,
		Direction: newDirection.Normalize(),
	}
}
