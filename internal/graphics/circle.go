package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

func DrawCircleOutline(img *ebiten.Image, cx, cy, radius float64, clr color.Color) {
	for i := 0.0; i < 360.0; i += 1.0 {
		radians := i * (math.Pi / 180.0)
		x := cx + radius*math.Cos(radians)
		y := cy + radius*math.Sin(radians)
		img.Set(int(x), int(y), clr)
	}
}
