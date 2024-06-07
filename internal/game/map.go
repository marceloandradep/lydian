package game

import (
	"github.com/marceloandradep/lydian/internal/math"
	"github.com/marceloandradep/lydian/internal/rendering"
	"math/rand"
)

func createMap(size float64) []*rendering.Triangle3D {
	squareSize := 100.0
	halfSize := size / 2
	numSquares := int(size / squareSize)

	x := -halfSize
	z := -halfSize

	vertexList := make([]*math.Vector, 0)
	triangleList := make([]*rendering.Triangle3D, 0)

	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			y := rand.Float64() * 50
			vertexList = append(vertexList, math.NewVector3(x, y, z))
			x += squareSize
		}
		z += squareSize
		x = -halfSize
	}

	for i := 0; i < numSquares-1; i++ {
		for j := 0; j < numSquares-1; j++ {
			bottomLeft := i*numSquares + j
			bottomRight := bottomLeft + 1
			topLeft := bottomLeft + numSquares
			topRight := topLeft + 1

			triangleList = append(triangleList, rendering.NewTriangle3D(vertexList, bottomLeft, topLeft, topRight, true, 0x303030ff))
			triangleList = append(triangleList, rendering.NewTriangle3D(vertexList, bottomLeft, topRight, bottomRight, true, 0x303030ff))
		}
	}

	return triangleList
}
