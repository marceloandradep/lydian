package camera

import m "lydian/internal/math"

type Camera interface {
	SetPosition(x, y, z float64)
	SetRotation(x, y, z float64)
	GetPosition() *m.Vector
	IsCulled(sphere m.Sphere) bool
	Update()
	WorldToCameraMatrix() *m.Matrix
	CameraToPerspectiveMatrix() *m.Matrix
	PerspectiveToScreenMatrix() *m.Matrix
}
