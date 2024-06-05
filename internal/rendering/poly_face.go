package rendering

import (
	"lydian/internal/math"
	"lydian/internal/rendering/camera"
)

type PolyFace struct {
	active, clipped, backface bool
	is2Sided                  bool
	transparent               bool
	color                     uint32
	colorMode                 ColorMode
	shadeMode                 ShadeMode
	vertexList                [3]math.Vector
	transformedVertexList     [3]math.Vector
}

func NewPolyFace(v0, v1, v2 math.Vector, color uint32) *PolyFace {
	return &PolyFace{
		active:                true,
		color:                 color,
		vertexList:            [3]math.Vector{v0, v1, v2},
		transformedVertexList: [3]math.Vector{v0, v1, v2},
	}
}

func (p *PolyFace) IsVisible() bool {
	return p.active && !(p.clipped || p.backface)
}

func (p *PolyFace) VertexAt(index int) *math.Vector {
	return &p.transformedVertexList[index]
}

func (p *PolyFace) transform(mt math.Matrix, transformType TransformType) {
	if p.active && !(p.clipped || p.backface) {
		switch transformType {
		case LocalOnly:
			p.vertexList[0] = *mt.MultiplyVertex(p.vertexList[0])
			p.vertexList[1] = *mt.MultiplyVertex(p.vertexList[1])
			p.vertexList[2] = *mt.MultiplyVertex(p.vertexList[2])
		case TransformedOnly:
			p.transformedVertexList[0] = *mt.MultiplyVertex(p.transformedVertexList[0])
			p.transformedVertexList[1] = *mt.MultiplyVertex(p.transformedVertexList[1])
			p.transformedVertexList[2] = *mt.MultiplyVertex(p.transformedVertexList[2])
		case LocalToTransformed:
			p.transformedVertexList[0] = *mt.MultiplyVertex(p.vertexList[0])
			p.transformedVertexList[1] = *mt.MultiplyVertex(p.vertexList[1])
			p.transformedVertexList[2] = *mt.MultiplyVertex(p.vertexList[2])
		}
	}
}

func (p *PolyFace) modelToWorldTransform(worldPos math.Vector, transformType TransformType) {
	if p.active && !(p.clipped || p.backface) {
		switch transformType {
		case LocalOnly:
			p.vertexList[0] = *p.vertexList[0].Add(worldPos)
			p.vertexList[1] = *p.vertexList[1].Add(worldPos)
			p.vertexList[2] = *p.vertexList[2].Add(worldPos)
		case TransformedOnly:
			p.transformedVertexList[0] = *p.transformedVertexList[0].Add(worldPos)
			p.transformedVertexList[1] = *p.transformedVertexList[1].Add(worldPos)
			p.transformedVertexList[2] = *p.transformedVertexList[2].Add(worldPos)
		case LocalToTransformed:
			p.transformedVertexList[0] = *p.vertexList[0].Add(worldPos)
			p.transformedVertexList[1] = *p.vertexList[1].Add(worldPos)
			p.transformedVertexList[2] = *p.vertexList[2].Add(worldPos)
		}
	}
}

func (p *PolyFace) worldToCameraTransform(cam camera.Camera) {
	if p.active && !(p.clipped || p.backface) {
		mt := cam.WorldToCameraMatrix()
		p.transformedVertexList[0] = *mt.MultiplyVertex(p.vertexList[0])
		p.transformedVertexList[1] = *mt.MultiplyVertex(p.vertexList[1])
		p.transformedVertexList[2] = *mt.MultiplyVertex(p.vertexList[2])
	}
}

func (p *PolyFace) cameraToPerspectiveTransform(camera camera.Camera) {
	if p.active && !(p.clipped || p.backface) {
		mt := camera.CameraToPerspectiveMatrix()
		p.transformedVertexList[0] = *mt.MultiplyVertex(p.transformedVertexList[0]).DeHomogenize()
		p.transformedVertexList[1] = *mt.MultiplyVertex(p.transformedVertexList[1]).DeHomogenize()
		p.transformedVertexList[2] = *mt.MultiplyVertex(p.transformedVertexList[2]).DeHomogenize()
	}
}

func (p *PolyFace) perspectiveToScreenTransform(cam camera.Camera) {
	if p.active && !(p.clipped || p.backface) {
		mt := cam.PerspectiveToScreenMatrix()
		p.transformedVertexList[0] = *mt.MultiplyVertex(p.transformedVertexList[0]).DeHomogenize()
		p.transformedVertexList[1] = *mt.MultiplyVertex(p.transformedVertexList[1]).DeHomogenize()
		p.transformedVertexList[2] = *mt.MultiplyVertex(p.transformedVertexList[2]).DeHomogenize()
	}
}

func (p *PolyFace) shouldTestBackFaceRemoval() bool {
	return p.active && !(p.clipped || p.is2Sided || p.backface)
}

func (p *PolyFace) reset() {
	p.clipped = false
	p.backface = false
}

func (p *PolyFace) cull(cam camera.Camera) {
	if !(cam.IsVertexWithinFrustum(p.transformedVertexList[0]) &&
		cam.IsVertexWithinFrustum(p.transformedVertexList[1]) &&
		cam.IsVertexWithinFrustum(p.transformedVertexList[2])) {
		p.clipped = true
	}
}

func (p *PolyFace) removeBackFace(cam camera.Camera) {
	if !p.shouldTestBackFaceRemoval() {
		return
	}

	p0 := p.transformedVertexList[0]
	p1 := p.transformedVertexList[1]
	p2 := p.transformedVertexList[2]

	u := p1.Sub(p0)
	v := p2.Sub(p0)

	normal := u.Cross(v)
	view := cam.GetPosition().Sub(p0)

	dot := normal.Dot(view)

	if dot <= 0 {
		p.backface = true
	}
}
