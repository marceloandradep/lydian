package game

import (
	"lydian/internal/math"
	"lydian/internal/rendering"
)

func createMap(size float64) []*rendering.Triangle3D {
	halfSize := size / 2
	numSquares := int(size / 10)

	x := -halfSize
	z := -halfSize

	vertexList := make([]*math.Vector, 0)
	triangleList := make([]*rendering.Triangle3D, 0)

	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			index := len(vertexList)

			vertexList = append(vertexList, math.NewVector3(x, -10, z))
			vertexList = append(vertexList, math.NewVector3(x+10, -10, z))
			vertexList = append(vertexList, math.NewVector3(x+10, -10, z+10))
			vertexList = append(vertexList, math.NewVector3(x, -10, z+10))

			triangleList = append(triangleList, rendering.NewTriangle3D(vertexList, index, index+1, index+2, true, 0x101010ff))
			triangleList = append(triangleList, rendering.NewTriangle3D(vertexList, index, index+2, index+3, true, 0x101010ff))

			x += 10
		}
		z += 10
		x = -halfSize
	}

	return triangleList
}
