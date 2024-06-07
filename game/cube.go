package game

import (
	"github.com/marceloandradep/lydian/math"
	"github.com/marceloandradep/lydian/rendering"
)

func Cubes(n int) ([]*rendering.Object3D, error) {
	cubes := make([]*rendering.Object3D, n)

	template, err := rendering.Load("/Users/marcelopereira/GolandProjects/lydian/game/resources/cube.plg", math.NewVector3(0, 0, 0), math.NewVector3(1, 1, 1), math.NewVector3(0, 0, 0))
	if err != nil {
		return nil, err
	}

	padding := 100.0
	cubeSize := 25.0
	totalSize := float64(n) * (padding + cubeSize)

	scale := math.NewVector3(5, 5, 5)
	rot := math.NewVector3(0, 0, 0)

	x := float64(-totalSize) / 2
	y := (float64(cubeSize) / 2) + 20
	for i := 0; i < n; i++ {
		pos := math.NewVector3(x, y, 200)
		cube := template.Copy(pos, scale, rot)
		cubes[i] = cube
		x += cubeSize + padding
	}

	return cubes, nil
}
