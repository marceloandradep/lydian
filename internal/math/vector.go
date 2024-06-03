package math

import "math"

type Vector struct {
	X, Y, Z, W float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z, 1}
}

func (v *Vector) Add(u Vector) *Vector {
	return NewVector(v.X+u.X, v.Y+u.Y, v.Z+u.Z)
}

func (v *Vector) Sub(u Vector) *Vector {
	return NewVector(v.X-u.X, v.Y-u.Y, v.Z-u.Z)
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Normalize() *Vector {
	length := v.Length()
	return NewVector(v.X/length, v.Y/length, v.Z/length)
}

func (v *Vector) Dot(u *Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v *Vector) Cross(u *Vector) *Vector {
	return NewVector(v.Y*u.Z-v.Z*u.Y, v.Z*u.X-v.X*u.Z, v.X*u.Y-v.Y*u.X)
}

func (v *Vector) DeHomogenize() *Vector {
	return NewVector(v.X/v.W, v.Y/v.W, v.Z/v.W)
}
