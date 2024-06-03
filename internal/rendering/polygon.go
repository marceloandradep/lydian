package rendering

import (
	"lydian/internal/math"
)

type Polygon3D struct {
	active, clipped, backface bool
	is2Sided                  bool
	transparent               bool
	color                     uint32
	colorMode                 ColorMode
	shadeMode                 ShadeMode
	vertexList                []math.Vector
	indices                   [3]int
}

func (p *Polygon3D) shouldTestBackFaceRemoval() bool {
	return p.active && !(p.clipped || p.is2Sided || p.backface)
}

func (p *Polygon3D) IsVisible() bool {
	return p.active && !(p.clipped || p.backface)
}

func (p *Polygon3D) VertexAt(index int) *math.Vector {
	return &p.vertexList[p.indices[index]]
}
