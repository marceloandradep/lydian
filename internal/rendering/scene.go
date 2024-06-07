package rendering

import "github.com/marceloandradep/lydian/internal/math"

type Scene struct {
	triangleList []*Triangle3D
}

func NewScene() *Scene {
	return &Scene{
		triangleList: make([]*Triangle3D, 0),
	}
}

func (s *Scene) Clear() {
	s.triangleList = s.triangleList[:0]
}

func (s *Scene) AddToScene(o *Object3D) {
	pos := o.GetWorldPos()
	translation := math.TranslationMatrix(pos.X, pos.Y, pos.Z)
	o.Transform(translation, LocalToTransformed, true)
	for _, t := range o.TriangleList() {
		s.triangleList = append(s.triangleList, t)
	}
}
func (s *Scene) TriangleList() []*Triangle3D {
	return s.triangleList
}
