package math

import "math"

type Vector struct {
	X, Y, Z, W float64
}

func NewVector3(x, y, z float64) *Vector {
	return &Vector{x, y, z, 1}
}

func NewVector4(x, y, z, w float64) *Vector {
	return &Vector{x, y, z, w}
}

func (v *Vector) Copy() *Vector {
	return &Vector{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
		W: v.W,
	}
}

func (v *Vector) Set(u *Vector) {
	v.X = u.X
	v.Y = u.Y
	v.Z = u.Z
}

func (v *Vector) Add(u *Vector) *Vector {
	return &Vector{
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}

func (v *Vector) Sub(u *Vector) *Vector {
	return &Vector{
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Normalize() {
	length := v.Length()
	v.X /= length
	v.Y /= length
	v.Z /= length
}

func (v *Vector) Dot(u *Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v *Vector) Cross(u *Vector) *Vector {
	return NewVector3(v.Y*u.Z-v.Z*u.Y, v.Z*u.X-v.X*u.Z, v.X*u.Y-v.Y*u.X)
}

func (v *Vector) DeHomogenize() *Vector {
	v.X /= v.W
	v.Y /= v.W
	v.Z /= v.W
	v.W = 1
	return v
}
