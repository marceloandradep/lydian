package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func DrawClippedTriangle(screen *ebiten.Image, clipper Clipper, x0, y0, x1, y1, x2, y2 int, clr color.Color) {
	DrawClippedLine(screen, clipper, x0, y0, x1, y1, clr)
	DrawClippedLine(screen, clipper, x1, y1, x2, y2, clr)
	DrawClippedLine(screen, clipper, x2, y2, x0, y0, clr)
}
