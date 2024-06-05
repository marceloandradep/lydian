package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"lydian/internal/rendering"
)

func DrawRenderList(screen *ebiten.Image, clipper Clipper, list *rendering.RenderList) {
	for _, poly := range list.PolygonList {
		if !poly.IsVisible() {
			continue
		}
		p0 := poly.VertexAt(0)
		p1 := poly.VertexAt(1)
		p2 := poly.VertexAt(2)
		DrawClippedTriangle(screen, clipper, int(p0.X), int(p0.Y), int(p1.X), int(p1.Y), int(p2.X), int(p2.Y), colornames.Red)
	}
}
