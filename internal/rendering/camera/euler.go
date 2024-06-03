package camera

import (
	"lydian/internal/geometry"
	m "lydian/internal/math"
	"math"
)

type EulerCamera struct {
	pos                 m.Vector
	rotation            m.Vector
	rotSeq              m.RotationSequence
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

func NewEuler(pos m.Vector, rotation m.Vector, rotSeq m.RotationSequence, nearClipZ float64, farClipZ float64, fov float64, viewPortSize geometry.Dimension) *EulerCamera {
	euler := &EulerCamera{
		pos:          pos,
		rotation:     rotation,
		rotSeq:       rotSeq,
		nearClipZ:    nearClipZ,
		farClipZ:     farClipZ,
		viewPortSize: viewPortSize,
		fov:          fov,
	}

	euler.aspectRatio = viewPortSize.Width / viewPortSize.Height

	euler.viewPlaneSize = geometry.Dimension{
		Width:  2.0,
		Height: 2.0 / euler.aspectRatio,
	}

	halfFov := m.DegreesToRadians(fov / 2)
	tanHalfFov := math.Tan(halfFov)
	halfViewPlaneWidth := euler.viewPlaneSize.Width / 2

	euler.viewDistance = (halfViewPlaneWidth) / tanHalfFov

	origin := *m.NewVector(0, 0, 0)
	if fov == 90 {
		euler.rightClipPlane = *m.NewPlane(origin, *m.NewVector(1, 0, -1), true)
		euler.leftClipPlane = *m.NewPlane(origin, *m.NewVector(-1, 0, -1), true)
		euler.topClipPlane = *m.NewPlane(origin, *m.NewVector(0, 1, -1), true)
		euler.bottomClipPlane = *m.NewPlane(origin, *m.NewVector(0, -1, -1), true)
	} else {
		euler.rightClipPlane = *m.NewPlane(origin, *m.NewVector(euler.viewDistance, 0, -halfViewPlaneWidth), true)
		euler.leftClipPlane = *m.NewPlane(origin, *m.NewVector(-euler.viewDistance, 0, -halfViewPlaneWidth), true)
		euler.topClipPlane = *m.NewPlane(origin, *m.NewVector(0, euler.viewDistance, -halfViewPlaneWidth), true)
		euler.bottomClipPlane = *m.NewPlane(origin, *m.NewVector(0, -euler.viewDistance, -halfViewPlaneWidth), true)
	}

	return euler
}

func (c *EulerCamera) Update() {
	c.computeWorldToCamera()
	c.computeCameraToPerspective()
	c.computePerspectiveToScreen()
}

func (c *EulerCamera) computeWorldToCamera() {
	translation := m.TranslationMatrix(-c.pos.X, -c.pos.Y, -c.pos.Z)
	rotation := m.RotationSequenceMatrix(c.rotation.X, c.rotation.Y, c.rotation.Z, c.rotSeq)
	c.worldToCamera = *translation.MultiplyMatrix(*rotation)
}

func (c *EulerCamera) computeCameraToPerspective() {
	c.cameraToPerspective = m.Matrix{
		M: [4][4]float64{
			{c.viewDistance, 0, 0, 0},
			{0, c.viewDistance * c.aspectRatio, 0, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0},
		},
	}
}

func (c *EulerCamera) computePerspectiveToScreen() {
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

func (c *EulerCamera) WorldToCameraMatrix() *m.Matrix {
	return &c.worldToCamera
}

func (c *EulerCamera) CameraToPerspectiveMatrix() *m.Matrix {
	return &c.cameraToPerspective
}

func (c *EulerCamera) PerspectiveToScreenMatrix() *m.Matrix {
	return &c.perspectiveToScreen
}

func (c *EulerCamera) IsCulled(sphere m.Sphere) bool {
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

func (c *EulerCamera) GetPosition() *m.Vector {
	return &c.pos
}

func (c *EulerCamera) SetPosition(x, y, z float64) {
	c.pos.X = x
	c.pos.Y = y
	c.pos.Z = z
}

func (c *EulerCamera) SetRotation(x, y, z float64) {
	c.rotation.X = x
	c.rotation.Y = y
	c.rotation.Z = z
}
