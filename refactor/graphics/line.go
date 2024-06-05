package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func DrawLine(screen *ebiten.Image, x0, y0, x1, y1 int, clr color.Color) {
	x := x0
	y := y0

	dx := x1 - x0
	dy := y1 - y0

	var xIncrement, yIncrement int

	if dx >= 0 {
		xIncrement = 1
	} else {
		xIncrement = -1
		dx = -dx
	}

	if dy >= 0 {
		yIncrement = 1
	} else {
		yIncrement = -1
		dy = -dy
	}

	dx2 := dx << 1
	dy2 := dy << 1

	if dx > dy {
		err := dy2 - dx
		for idx := 0; idx <= dx; idx++ {
			screen.Set(x, y, clr)
			if err >= 0 {
				err -= dx2
				y += yIncrement
			}
			err += dy2
			x += xIncrement
		}
	} else {
		err := dx2 - dy
		for idx := 0; idx <= dy; idx++ {
			screen.Set(x, y, clr)
			if err >= 0 {
				err -= dy2
				x += xIncrement
			}
			err += dx2
			y += yIncrement
		}
	}
}

func DrawClippedLine(screen *ebiten.Image, clipper Clipper, x0, y0, x1, y1 int, clr color.Color) {
	x0, y0, x1, y1, visible := clipper.ClipLine(x0, y0, x1, y1)
	if visible {
		DrawLine(screen, x0, y0, x1, y1, clr)
	}
}
