package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"lydian/internal/rendering"
)

func DrawObject(screen *ebiten.Image, clipper Clipper, object rendering.Object3D) {
	if !object.IsVisible() {
		return
	}
	for _, poly := range object.GetPolygonList() {
		if !poly.IsVisible() {
			continue
		}
		p0 := poly.VertexAt(0)
		p1 := poly.VertexAt(1)
		p2 := poly.VertexAt(2)
		DrawClippedTriangle(screen, clipper, int(p0.X), int(p0.Y), int(p1.X), int(p1.Y), int(p2.X), int(p2.Y), colornames.Red)
	}
}
