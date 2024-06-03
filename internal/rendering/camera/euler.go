package camera

import (
	"lydian/internal/geometry"
	m "lydian/internal/math"
)

type EulerCamera struct {
	BaseCamera
	rotation m.Vector
	rotSeq   m.RotationSequence
}

func NewEuler(pos m.Vector, rotation m.Vector, rotSeq m.RotationSequence, nearClipZ float64, farClipZ float64, fov float64, viewPortSize geometry.Dimension) *EulerCamera {
	euler := &EulerCamera{
		rotation: rotation,
		rotSeq:   rotSeq,
	}

	euler.init(pos, nearClipZ, farClipZ, fov, viewPortSize)

	return euler
}

func (c *EulerCamera) SetRotation(x, y, z float64) {
	c.rotation.X = x
	c.rotation.Y = y
	c.rotation.Z = z
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
