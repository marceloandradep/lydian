package camera

import (
	"lydian/internal/geometry"
	m "lydian/internal/math"
	"math"
)

type Camera interface {
	SetPosition(x, y, z float64)
	GetPosition() *m.Vector
	IsCulled(sphere m.Sphere) bool
	Update()
	WorldToCameraMatrix() *m.Matrix
	CameraToPerspectiveMatrix() *m.Matrix
	PerspectiveToScreenMatrix() *m.Matrix
}

type BaseCamera struct {
	pos                 m.Vector
	nearClipZ           float64
	farClipZ            float64
	fov                 float64
	viewPortSize        geometry.Dimension
	viewPlaneSize       geometry.Dimension
	viewDistance        float64
	aspectRatio         float64
	rightClipPlane      m.Plane
	leftClipPlane       m.Plane
	topClipPlane        m.Plane
	bottomClipPlane     m.Plane
	worldToCamera       m.Matrix
	cameraToPerspective m.Matrix
	perspectiveToScreen m.Matrix
}

func (c *BaseCamera) GetPosition() *m.Vector {
	return &c.pos
}

func (c *BaseCamera) SetPosition(x, y, z float64) {
	c.pos.X = x
	c.pos.Y = y
	c.pos.Z = z
}

func (c *BaseCamera) WorldToCameraMatrix() *m.Matrix {
	return &c.worldToCamera
}

func (c *BaseCamera) CameraToPerspectiveMatrix() *m.Matrix {
	return &c.cameraToPerspective
}

func (c *BaseCamera) PerspectiveToScreenMatrix() *m.Matrix {
	return &c.perspectiveToScreen
}

func (c *BaseCamera) IsCulled(sphere m.Sphere) bool {
	// no clipping available yet
	if (sphere.Center.Z - sphere.Radius) < 0 {
		return true
	}

	if ((sphere.Center.Z - sphere.Radius) > c.farClipZ) || ((sphere.Center.Z + sphere.Radius) < c.nearClipZ) {
		return true
	}

	xLimit := (0.5 * c.viewPlaneSize.Width * sphere.Center.Z) / c.viewDistance
	if ((sphere.Center.X - sphere.Radius) > xLimit) || ((sphere.Center.X + sphere.Radius) < -xLimit) {
		return true
	}

	yLimit := (0.5 * c.viewPlaneSize.Height * sphere.Center.Z) / c.viewDistance
	if ((sphere.Center.Y - sphere.Radius) > yLimit) || ((sphere.Center.Y + sphere.Radius) < -xLimit) {
		return true
	}

	return false
}

func (c *BaseCamera) init(pos m.Vector, nearClipZ float64, farClipZ float64, fov float64, viewPortSize geometry.Dimension) {
	c.pos = pos
	c.nearClipZ = nearClipZ
	c.farClipZ = farClipZ
	c.fov = fov
	c.viewPortSize = viewPortSize

	c.aspectRatio = c.viewPortSize.Width / c.viewPortSize.Height

	c.viewPlaneSize = geometry.Dimension{
		Width:  2.0,
		Height: 2.0 / c.aspectRatio,
	}

	halfFov := m.DegreesToRadians(c.fov / 2)
	tanHalfFov := math.Tan(halfFov)
	halfViewPlaneWidth := c.viewPlaneSize.Width / 2

	c.viewDistance = (halfViewPlaneWidth) / tanHalfFov

	origin := *m.NewVector(0, 0, 0)
	if c.fov == 90 {
		c.rightClipPlane = *m.NewPlane(origin, *m.NewVector(1, 0, -1), true)
		c.leftClipPlane = *m.NewPlane(origin, *m.NewVector(-1, 0, -1), true)
		c.topClipPlane = *m.NewPlane(origin, *m.NewVector(0, 1, -1), true)
		c.bottomClipPlane = *m.NewPlane(origin, *m.NewVector(0, -1, -1), true)
	} else {
		c.rightClipPlane = *m.NewPlane(origin, *m.NewVector(c.viewDistance, 0, -halfViewPlaneWidth), true)
		c.leftClipPlane = *m.NewPlane(origin, *m.NewVector(-c.viewDistance, 0, -halfViewPlaneWidth), true)
		c.topClipPlane = *m.NewPlane(origin, *m.NewVector(0, c.viewDistance, -halfViewPlaneWidth), true)
		c.bottomClipPlane = *m.NewPlane(origin, *m.NewVector(0, -c.viewDistance, -halfViewPlaneWidth), true)
	}
}

func (c *BaseCamera) computeCameraToPerspective() {
	c.cameraToPerspective = m.Matrix{
		M: [4][4]float64{
			{c.viewDistance, 0, 0, 0},
			{0, c.viewDistance * c.aspectRatio, 0, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0},
		},
	}
}

func (c *BaseCamera) computePerspectiveToScreen() {
	alpha := (c.viewPortSize.Width - 1) / 2
	beta := (c.viewPortSize.Height - 1) / 2
	c.perspectiveToScreen = m.Matrix{
		M: [4][4]float64{
			{alpha, 0, 0, 0},
			{0, -beta, 0, 0},
			{alpha, beta, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
