package rendering

import "lydian/refactor/math"

type Triangle3D struct {
	vertices []*math.Vector
	Indices  [3]int
}

func NewTriangle3D(vertices []*math.Vector, p0, p1, p2 int) *Triangle3D {
	return &Triangle3D{
		vertices: vertices,
		Indices:  [3]int{p0, p1, p2},
	}
}

func (t *Triangle3D) Vertices() (*math.Vector, *math.Vector, *math.Vector) {
	return t.vertices[t.Indices[0]], t.vertices[t.Indices[1]], t.vertices[t.Indices[2]]
}
