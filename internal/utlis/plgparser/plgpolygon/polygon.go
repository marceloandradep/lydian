package plgpolygon

import (
	"errors"
	"strconv"
	"strings"
)

const (
	bitMask2Sided       = 0x1000
	bitMaskColorModeRGB = 0x8000
	bitMaskShadeMode    = 0x6000
)

type ShadeMode int

const (
	Pure    ShadeMode = 0x0000
	Flat    ShadeMode = 0x2000
	Gouraud ShadeMode = 0x4000
	Phong   ShadeMode = 0x6000
)

type PLGPolygon struct {
	SurfaceDescriptor int
	V0, V1, V2        int
}

func New(rawDescriptor string, v0, v1, v2 int) (*PLGPolygon, error) {
	base := 10
	if isHexadecimal(rawDescriptor) {
		rawDescriptor = strings.TrimPrefix(rawDescriptor, "0x")
		base = 16
	}

	surfaceDescriptor, err := strconv.ParseInt(rawDescriptor, base, 32)
	if err != nil {
		return nil, errors.New("invalid descriptor")
	}

	return &PLGPolygon{int(surfaceDescriptor), v0, v1, v2}, nil
}

func (p *PLGPolygon) Is2Sided() bool {
	return p.SurfaceDescriptor&bitMask2Sided == bitMask2Sided
}

func (p *PLGPolygon) IsColorModeRGB() bool {
	return p.SurfaceDescriptor&bitMaskColorModeRGB == bitMaskColorModeRGB
}

func (p *PLGPolygon) ShadeMode() ShadeMode {
	return ShadeMode(p.SurfaceDescriptor & bitMaskShadeMode)
}

func (p *PLGPolygon) GetRGB32Color() uint32 {
	red := ((p.SurfaceDescriptor & 0xf00) >> 8) * 17
	green := ((p.SurfaceDescriptor & 0x0f0) >> 4) * 17
	blue := (p.SurfaceDescriptor & 0x00f) * 17
	return uint32(red | (green << 8) | (blue << 16) | (0xff << 24))
}

func (p *PLGPolygon) Get8BitColorIndex() uint8 {
	return uint8(p.SurfaceDescriptor & 0x00ff)
}

func isHexadecimal(str string) bool {
	return len(str) >= 2 && (strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X"))
}
