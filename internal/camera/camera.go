package camera

import (
	"lydian/internal/geometry"
	"lydian/internal/math"
	"lydian/internal/rendering"
	m "math"
)

type cameraType int

const (
	cameraTypeEuler cameraType = iota
	cameraTypeUVN
)

type Camera struct {
	camType             cameraType
	pos                 *math.Vector
	rotation            *math.Vector
	target              *math.Vector
	rotSeq              math.RotationSequence
	nearClipZ           float64
	farClipZ            float64
	fov                 float64
	viewPortSize        *geometry.Dimension
	viewPlaneSize       *geometry.Dimension
	viewDistance        float64
	aspectRatio         float64
	rightClipPlane      *math.Plane
	leftClipPlane       *math.Plane
	topClipPlane        *math.Plane
	bottomClipPlane     *math.Plane
	worldToCamera       *math.Matrix
	cameraToPerspective *math.Matrix
	perspectiveToScreen *math.Matrix
	localVertexList     []*math.Vector
	triangleList        []*rendering.Triangle3D
}

func NewEuler(pos *math.Vector, rotation *math.Vector, rotSeq math.RotationSequence, nearClipZ float64, farClipZ float64, fov float64, viewPortSize *geometry.Dimension) *Camera {
	c := &Camera{
		camType:      cameraTypeEuler,
		pos:          pos,
		rotation:     rotation,
		rotSeq:       rotSeq,
		nearClipZ:    nearClipZ,
		farClipZ:     farClipZ,
		fov:          fov,
		viewPortSize: viewPortSize,
	}

	c.init()
	return c
}

func NewUVN(pos *math.Vector, target *math.Vector, nearClipZ float64, farClipZ float64, fov float64, viewPortSize *geometry.Dimension) *Camera {
	c := &Camera{
		camType:      cameraTypeUVN,
		pos:          pos,
		target:       target,
		nearClipZ:    nearClipZ,
		farClipZ:     farClipZ,
		fov:          fov,
		viewPortSize: viewPortSize,
	}

	c.init()
	return c
}

func (c *Camera) SetPos(newPos *math.Vector) {
	c.pos.Set(newPos)
}

func (c *Camera) Rotate(rotation *math.Vector) {
	c.rotation.Set(rotation)
}

func (c *Camera) SetTarget(target *math.Vector) {
	c.target = target
}

func (c *Camera) SetPolarCoordinates(elevation, heading float64) {
	phi := math.DegreesToRadians(elevation)
	theta := math.DegreesToRadians(heading)

	cosPhi := m.Cos(phi)
	sinPhi := m.Sin(phi)

	cosTheta := m.Cos(theta)
	sinTheta := m.Sin(theta)

	c.target = math.NewVector3(-cosPhi*sinTheta+c.pos.X, sinPhi+c.pos.Y, cosPhi*cosTheta+c.pos.Z)
}

func (c *Camera) Clear() {
	c.localVertexList = c.localVertexList[:0]
	c.triangleList = c.triangleList[:0]
}

func (c *Camera) Update() {
	c.computeWorldToCamera()
	c.computeCameraToPerspective()
	c.computePerspectiveToScreen()
}

func (c *Camera) AddTriangleToCamera(t *rendering.Triangle3D) {
	index := len(c.localVertexList)

	if c.isBackFaced(t) {
		return
	}

	p0, p1, p2 := t.Vertices()
	t0 := math.NewVector3(p0.X, p0.Y, p0.Z)
	t1 := math.NewVector3(p1.X, p1.Y, p1.Z)
	t2 := math.NewVector3(p2.X, p2.Y, p2.Z)

	c.worldToCameraTransform(t0)
	c.worldToCameraTransform(t1)
	c.worldToCameraTransform(t2)

	c.localVertexList = append(c.localVertexList, t0)
	c.localVertexList = append(c.localVertexList, t1)
	c.localVertexList = append(c.localVertexList, t2)

	triangleCopy := rendering.NewTriangle3D(c.localVertexList, index, index+1, index+2, t.Is2Sided, t.Color)

	if c.culled(triangleCopy) {
		return
	}

	c.triangleList = append(c.triangleList, triangleCopy)
}

func (c *Camera) AddSceneToCamera(s *rendering.Scene) {
	for _, t := range s.TriangleList() {
		c.AddTriangleToCamera(t)
	}
}

func (c *Camera) ProjectTriangles() {
	visited := make([]bool, len(c.localVertexList))
	for _, t := range c.triangleList {
		for _, k := range t.Indices {
			if visited[k] {
				continue
			}
			visited[k] = true
			c.cameraToPerspectiveTransform(c.localVertexList[k])
			c.perspectiveToScreenTransform(c.localVertexList[k])
		}
	}
}

func (c *Camera) TriangleList() []*rendering.Triangle3D {
	return c.triangleList
}

func (c *Camera) init() {
	c.localVertexList = make([]*math.Vector, 0)
	c.triangleList = make([]*rendering.Triangle3D, 0)

	c.aspectRatio = c.viewPortSize.Width / c.viewPortSize.Height

	c.viewPlaneSize = &geometry.Dimension{
		Width:  2.0,
		Height: 2.0 / c.aspectRatio,
	}

	halfFov := math.DegreesToRadians(c.fov / 2)
	tanHalfFov := m.Tan(halfFov)
	halfViewPlaneWidth := c.viewPlaneSize.Width / 2

	c.viewDistance = (halfViewPlaneWidth) / tanHalfFov

	origin := math.NewVector3(0, 0, 0)
	if c.fov == 90 {
		c.rightClipPlane = math.NewPlane(origin, math.NewVector3(1, 0, -1), true)
		c.leftClipPlane = math.NewPlane(origin, math.NewVector3(-1, 0, -1), true)
		c.topClipPlane = math.NewPlane(origin, math.NewVector3(0, 1, -1), true)
		c.bottomClipPlane = math.NewPlane(origin, math.NewVector3(0, -1, -1), true)
	} else {
		c.rightClipPlane = math.NewPlane(origin, math.NewVector3(c.viewDistance, 0, -halfViewPlaneWidth), true)
		c.leftClipPlane = math.NewPlane(origin, math.NewVector3(-c.viewDistance, 0, -halfViewPlaneWidth), true)
		c.topClipPlane = math.NewPlane(origin, math.NewVector3(0, c.viewDistance, -halfViewPlaneWidth), true)
		c.bottomClipPlane = math.NewPlane(origin, math.NewVector3(0, -c.viewDistance, -halfViewPlaneWidth), true)
	}
}

func (c *Camera) culled(t *rendering.Triangle3D) bool {
	p0, p1, p2 := t.Vertices()

	if !c.contains(p0) && !c.contains(p1) && !c.contains(p2) {
		return true
	}

	if p0.Z <= 0 || p1.Z <= 0 || p2.Z <= 0 {
		return true
	}

	return false
}

func (c *Camera) contains(p *math.Vector) bool {
	if (p.Z > c.farClipZ) || (p.Z < c.nearClipZ) {
		return false
	}

	xLim := (0.5 * c.viewPlaneSize.Width * p.Z) / c.viewDistance
	if (p.X > xLim) || (p.X < -xLim) {
		return false
	}

	yLim := (0.5 * c.viewPortSize.Height * p.Z) / c.viewDistance
	if (p.Y > yLim) || (p.Y < -yLim) {
		return false
	}

	return true
}

func (c *Camera) isBackFaced(t *rendering.Triangle3D) bool {
	if t.Is2Sided {
		return false
	}

	p0, p1, p2 := t.Vertices()

	u := p1.Sub(p0)
	v := p2.Sub(p0)

	normal := u.Cross(v)
	view := c.pos.Sub(p0)

	dot := normal.Dot(view)

	if dot <= 0 {
		return true
	}

	return false
}

func (c *Camera) computeWorldToCamera() {
	switch c.camType {
	case cameraTypeEuler:
		c.computeEulerTransformMatrix()
	case cameraTypeUVN:
		c.computeUVNTransformMatrix()
	}
}

func (c *Camera) computeEulerTransformMatrix() {
	translation := math.TranslationMatrix(-c.pos.X, -c.pos.Y, -c.pos.Z)
	rotation := math.RotationSequenceMatrix(c.rotation.X, c.rotation.Y, c.rotation.Z, c.rotSeq)
	c.worldToCamera = translation.MultiplyMatrix(rotation)
}

func (c *Camera) computeUVNTransformMatrix() {
	n := c.target.Sub(c.pos)
	v := math.NewVector3(0, 1, 0)
	u := v.Cross(n)
	v = n.Cross(u)

	u.Normalize()
	v.Normalize()
	n.Normalize()

	translation := math.TranslationMatrix(-c.pos.X, -c.pos.Y, -c.pos.Z)
	uvn := &math.Matrix{
		M: [4][4]float64{
			{u.X, v.X, n.X, 0},
			{u.Y, v.Y, n.Y, 0},
			{u.Z, v.Z, n.Z, 0},
			{0, 0, 0, 1},
		},
	}

	c.worldToCamera = translation.MultiplyMatrix(uvn)
}

func (c *Camera) computeCameraToPerspective() {
	c.cameraToPerspective = &math.Matrix{
		M: [4][4]float64{
			{c.viewDistance, 0, 0, 0},
			{0, c.viewDistance * c.aspectRatio, 0, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0},
		},
	}
}

func (c *Camera) computePerspectiveToScreen() {
	alpha := (c.viewPortSize.Width - 1) / 2
	beta := (c.viewPortSize.Height - 1) / 2
	c.perspectiveToScreen = &math.Matrix{
		M: [4][4]float64{
			{alpha, 0, 0, 0},
			{0, -beta, 0, 0},
			{alpha, beta, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func (c *Camera) transform(mt *math.Matrix, p *math.Vector) {
	p.Set(mt.MultiplyVertex(p).DeHomogenize())
}

func (c *Camera) worldToCameraTransform(p *math.Vector) {
	c.transform(c.worldToCamera, p)
}

func (c *Camera) cameraToPerspectiveTransform(p *math.Vector) {
	c.transform(c.cameraToPerspective, p)
}

func (c *Camera) perspectiveToScreenTransform(p *math.Vector) {
	c.transform(c.perspectiveToScreen, p)
}
