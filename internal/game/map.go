package game

import (
	"lydian/internal/math"
	"lydian/internal/rendering"
)

func createMap(size float64) *rendering.RenderList {
	halfSize := size / 2
	numSquares := int(size / 10)

	x := -halfSize
	z := 0.0

	polygonList := make([]rendering.PolyFace, 0)

	for i := 0; i < numSquares; i++ {
		for j := 0; j < numSquares; j++ {
			v0 := math.NewVector(x, -10, z)
			v1 := math.NewVector(x+10, -10, z)
			v2 := math.NewVector(x+10, -10, z+10)
			v3 := math.NewVector(x, -10, z+10)

			polygonList = append(polygonList, *rendering.NewPolyFace(*v0, *v1, *v2, 0xffffffff))
			polygonList = append(polygonList, *rendering.NewPolyFace(*v0, *v2, *v3, 0xffffffff))

			x += 10
		}
		z += 10
		x = -halfSize
	}

	/*v0 := math.NewVector(-5, -5, 100)
	v1 := math.NewVector(5, -5, 100)
	v2 := math.NewVector(0, 5, 100)

	polygonList = append(polygonList, *rendering.NewPolyFace(*v2, *v1, *v0, 0xffffffff))*/

	return &rendering.RenderList{
		PolygonList: polygonList,
	}
}
