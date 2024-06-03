package rendering

type ColorMode int

const (
	Color8Bit ColorMode = iota
	ColorRBG16
	ColorRGB24
)

type ShadeMode int

const (
	ShadeModePure ShadeMode = iota
	ShadeModeFlat
	ShadeModeGouraud
	ShadeModePhong
)

type TransformType int

const (
	LocalOnly TransformType = iota
	TransformedOnly
	LocalToTransformed
)
