package rendering

import "github.com/marceloandradep/lydian/internal/math"

type Triangle3D struct {
	Indices  [3]int
	Color    uint32
	Is2Sided bool

	vertices []*math.Vector
}

func NewTriangle3D(vertices []*math.Vector, p0, p1, p2 int, is2Sided bool, clr uint32) *Triangle3D {
	return &Triangle3D{
		Indices:  [3]int{p0, p1, p2},
		Color:    clr,
		Is2Sided: is2Sided,
		vertices: vertices,
	}
}

func (t *Triangle3D) Vertices() (*math.Vector, *math.Vector, *math.Vector) {
	return t.vertices[t.Indices[0]], t.vertices[t.Indices[1]], t.vertices[t.Indices[2]]
}
