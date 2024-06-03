package camera

import (
	"lydian/internal/geometry"
	m "lydian/internal/math"
	"math"
)

type UVNCamera struct {
	BaseCamera
	target  m.Vector
	u, v, n m.Vector
}

func NewUVN(pos m.Vector, target m.Vector, nearClipZ float64, farClipZ float64, fov float64, viewPortSize geometry.Dimension) *UVNCamera {
	uvn := &UVNCamera{
		target: target,
		u:      *m.NewVector(1, 0, 0),
		v:      *m.NewVector(0, 1, 0),
		n:      *m.NewVector(0, 0, 1),
	}

	uvn.init(pos, nearClipZ, farClipZ, fov, viewPortSize)

	return uvn
}

func (c *UVNCamera) SetTarget(target m.Vector) {
	c.target = target
}

func (c *UVNCamera) SetPolarCoordinates(elevation, heading float64) {
	phi := m.DegreesToRadians(elevation)
	theta := m.DegreesToRadians(heading)

	cosPhi := math.Cos(phi)
	sinPhi := math.Sin(phi)

	cosTheta := math.Cos(theta)
	sinTheta := math.Sin(theta)

	c.target = *m.NewVector(-cosPhi*sinTheta+c.pos.X, sinPhi+c.pos.Y, cosPhi*cosTheta+c.pos.Z)
}

func (c *UVNCamera) Update() {
	c.computeWorldToCamera()
	c.computeCameraToPerspective()
	c.computePerspectiveToScreen()
}

func (c *UVNCamera) computeWorldToCamera() {
	n := c.target.Sub(c.pos)
	v := m.NewVector(0, 1, 0)
	u := c.v.Cross(n)
	v = n.Cross(u)

	u = u.Normalize()
	v = v.Normalize()
	n = n.Normalize()

	c.u = *u
	c.v = *v
	c.n = *n

	translation := m.TranslationMatrix(-c.pos.X, -c.pos.Y, -c.pos.Z)
	uvn := m.Matrix{
		M: [4][4]float64{
			{c.u.X, c.v.X, c.n.X, 0},
			{c.u.Y, c.v.Y, c.n.Y, 0},
			{c.u.Z, c.v.Z, c.n.Z, 0},
			{0, 0, 0, 1},
		},
	}
	c.worldToCamera = *translation.MultiplyMatrix(uvn)
}
