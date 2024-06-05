package rendering

import (
	"lydian/internal/math"
	"lydian/internal/rendering/camera"
)

type RenderList struct {
	PolygonList []PolyFace
}

func (r *RenderList) Reset() {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].reset()
	}
}

func (r *RenderList) Update(cam camera.Camera) {
	r.modelToWorldCameraTransform()
	r.cull(cam)
	// r.removeBackFaces(cam)
	r.worldToCameraTransform(cam)
	r.cameraToPerspectiveTransform(cam)
	r.perspectiveToScreenTransform(cam)
}

func (r *RenderList) cull(cam camera.Camera) {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].cull(cam)
	}
}

func (r *RenderList) removeBackFaces(cam camera.Camera) {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].removeBackFace(cam)
	}
}

func (r *RenderList) modelToWorldCameraTransform() {
	worldPos := *math.NewVector(0, 0, 0)
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].modelToWorldTransform(worldPos, LocalToTransformed)
	}
}

func (r *RenderList) worldToCameraTransform(cam camera.Camera) {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].worldToCameraTransform(cam)
	}
}

func (r *RenderList) cameraToPerspectiveTransform(cam camera.Camera) {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].cameraToPerspectiveTransform(cam)
	}
}

func (r *RenderList) perspectiveToScreenTransform(cam camera.Camera) {
	for i := 0; i < len(r.PolygonList); i++ {
		r.PolygonList[i].perspectiveToScreenTransform(cam)
	}
}
