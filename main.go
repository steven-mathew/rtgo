package main

import (
	"fmt"
	"math"

	"github.com/golang/geo/r3"
	"github.com/steven-mathew/rtgo/internal/pkg/camera"
	"github.com/steven-mathew/rtgo/internal/pkg/color"
	"github.com/steven-mathew/rtgo/internal/pkg/geo"
)

func main() {
	aspectRatio := 16. / 9.
	imageWidth := 400
	imageHeight := int64(float64(imageWidth) / aspectRatio)

	cam := camera.NewCamera(
		r3.Vector{X: 0, Y: 0, Z: 0},
		r3.Vector{X: 0, Y: 0, Z: 10},
		r3.Vector{X: 0, Y: 1, Z: 0},
		math.Pi/6,
	)

	cam.Focus(r3.Vector{X: 0, Y: 0, Z: 2}, 1)

	fmt.Printf("P3\n %d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			ray := cam.CastRay(u, v)
			color := func(ray geo.Ray) color.Color {
				unitDirection := ray.Direction.Normalize()
				t := 0.5 * (unitDirection.Y + 1.)
				u := r3.Vector{X: 1, Y: 1, Z: 1}.Mul(1 - t)
				v := r3.Vector{X: 0.5, Y: 0.7, Z: 1.0}.Mul(t)
				return u.Add(v)
			}(ray)

			px := int(255.999 * color.X)
			py := int(255.999 * color.Y)
			pz := int(255.999 * color.Z)

			fmt.Println(px, py, pz)
		}

	}

}
