package game

import (
	"lydian/internal/math"
	"lydian/internal/rendering"
)

func Cubes(n int) ([]*rendering.Object3D, error) {
	cubes := make([]*rendering.Object3D, n)

	template, err := rendering.Load("/Users/marcelopereira/GolandProjects/lydian/internal/game/resources/cube.plg", math.NewVector3(0, 0, 0), math.NewVector3(1, 1, 1), math.NewVector3(0, 0, 0))
	if err != nil {
		return nil, err
	}

	padding := 100
	cubeSize := 25
	totalSize := n * (padding + cubeSize)

	scale := math.NewVector3(5, 5, 5)
	rot := math.NewVector3(0, 0, 0)

	x := -totalSize / 2
	for i := 0; i < n; i++ {
		pos := math.NewVector3(float64(x), 17, 200)
		cube := template.Copy(pos, scale, rot)
		cubes[i] = cube
		x += cubeSize + padding
	}

	return cubes, nil
}