package graphics

import "image/color"

func RGBA(clr uint32) *color.RGBA {
	red := uint8(clr & 0xff000000 >> 24)
	green := uint8(clr & 0x00ff0000 >> 16)
	blue := uint8(clr & 0x0000ff00 >> 8)
	alpha := uint8(clr & 0x000000ff)
	return &color.RGBA{
		R: red,
		G: green,
		B: blue,
		A: alpha,
	}
}
