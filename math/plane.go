package math

type Plane struct {
	P0     *Vector
	Normal *Vector
}

func NewPlane(p0 *Vector, n *Vector, normalize bool) *Plane {
	if normalize {
		n.Normalize()
	}
	return &Plane{
		P0:     p0,
		Normal: n,
	}
}
