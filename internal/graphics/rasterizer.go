package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marceloandradep/lydian/internal/camera"
)

func Rasterize(screen *ebiten.Image, clipper Clipper, c *camera.Camera) {
	tList := c.TriangleList()
	for i := 0; i < len(tList); i++ {
		t := tList[i]
		p0, p1, p2 := t.Vertices()
		DrawClippedTriangle(screen, clipper, int(p0.X), int(p0.Y), int(p1.X), int(p1.Y), int(p2.X), int(p2.Y), RGBA(t.Color))
	}
}
