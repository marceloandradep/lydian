package math

import "math"

type RotationAxis int

const (
	XAxis RotationAxis = iota
	YAxis
	ZAxis
)

type RotationSequence int

const (
	RotationXYZ RotationSequence = iota
	RotationYXZ
	RotationXZY
	RotationYZX
	RotationZYX
	RotationZXY
)

type Matrix struct {
	M [4][4]float64
}

func (m *Matrix) MultiplyMatrix(n Matrix) *Matrix {
	return &Matrix{
		M: [4][4]float64{
			{m.M[0][0]*n.M[0][0] + m.M[0][1]*n.M[1][0] + m.M[0][2]*n.M[2][0] + m.M[0][3]*n.M[3][0], m.M[0][0]*n.M[0][1] + m.M[0][1]*n.M[1][1] + m.M[0][2]*n.M[2][1] + m.M[0][3]*n.M[3][1], m.M[0][0]*n.M[0][2] + m.M[0][1]*n.M[1][2] + m.M[0][2]*n.M[2][2] + m.M[0][3]*n.M[3][2], m.M[0][0]*n.M[0][3] + m.M[0][1]*n.M[1][3] + m.M[0][2]*n.M[2][3] + m.M[0][3]*n.M[3][3]},
			{m.M[1][0]*n.M[0][0] + m.M[1][1]*n.M[1][0] + m.M[1][2]*n.M[2][0] + m.M[1][3]*n.M[3][0], m.M[1][0]*n.M[0][1] + m.M[1][1]*n.M[1][1] + m.M[1][2]*n.M[2][1] + m.M[1][3]*n.M[3][1], m.M[1][0]*n.M[0][2] + m.M[1][1]*n.M[1][2] + m.M[1][2]*n.M[2][2] + m.M[1][3]*n.M[3][2], m.M[1][0]*n.M[0][3] + m.M[1][1]*n.M[1][3] + m.M[1][2]*n.M[2][3] + m.M[1][3]*n.M[3][3]},
			{m.M[2][0]*n.M[0][0] + m.M[2][1]*n.M[1][0] + m.M[2][2]*n.M[2][0] + m.M[2][3]*n.M[3][0], m.M[2][0]*n.M[0][1] + m.M[2][1]*n.M[1][1] + m.M[2][2]*n.M[2][1] + m.M[2][3]*n.M[3][1], m.M[2][0]*n.M[0][2] + m.M[2][1]*n.M[1][2] + m.M[2][2]*n.M[2][2] + m.M[2][3]*n.M[3][2], m.M[2][0]*n.M[0][3] + m.M[2][1]*n.M[1][3] + m.M[2][2]*n.M[2][3] + m.M[2][3]*n.M[3][3]},
			{m.M[3][0]*n.M[0][0] + m.M[3][1]*n.M[1][0] + m.M[3][2]*n.M[2][0] + m.M[3][3]*n.M[3][0], m.M[3][0]*n.M[0][1] + m.M[3][1]*n.M[1][1] + m.M[3][2]*n.M[2][1] + m.M[3][3]*n.M[3][1], m.M[3][0]*n.M[0][2] + m.M[3][1]*n.M[1][2] + m.M[3][2]*n.M[2][2] + m.M[3][3]*n.M[3][2], m.M[3][0]*n.M[0][3] + m.M[3][1]*n.M[1][3] + m.M[3][2]*n.M[2][3] + m.M[3][3]*n.M[3][3]},
		},
	}
}

func (m *Matrix) MultiplyVertex(v Vector) *Vector {
	return &Vector{
		X: v.X*m.M[0][0] + v.Y*m.M[1][0] + v.Z*m.M[2][0] + v.W*m.M[3][0],
		Y: v.X*m.M[0][1] + v.Y*m.M[1][1] + v.Z*m.M[2][1] + v.W*m.M[3][1],
		Z: v.X*m.M[0][2] + v.Y*m.M[1][2] + v.Z*m.M[2][2] + v.W*m.M[3][2],
		W: v.X*m.M[0][3] + v.Y*m.M[1][3] + v.Z*m.M[2][3] + v.W*m.M[3][3],
	}
}

func IdentityMatrix() *Matrix {
	return &Matrix{
		M: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func TranslationMatrix(dx, dy, dz float64) *Matrix {
	return &Matrix{
		M: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{dx, dy, dz, 1},
		},
	}
}

func ScaleMatrix(sx, sy, sz float64) *Matrix {
	return &Matrix{
		M: [4][4]float64{
			{sx, 0, 0, 0},
			{0, sy, 0, 0},
			{0, 0, sz, 0},
			{0, 0, 0, 1},
		},
	}
}

func RotationMatrix(degrees float64, axis RotationAxis) *Matrix {
	rad := DegreesToRadians(degrees)
	switch axis {
	case XAxis:
		return &Matrix{
			M: [4][4]float64{
				{1, 0, 0, 0},
				{0, math.Cos(rad), math.Sin(rad), 0},
				{0, -math.Sin(rad), math.Cos(rad), 0},
				{0, 0, 0, 1},
			},
		}
	case YAxis:
		return &Matrix{
			M: [4][4]float64{
				{math.Cos(rad), 0, -math.Sin(rad), 0},
				{0, 1, 0, 0},
				{math.Sin(rad), 0, math.Cos(rad), 0},
				{0, 0, 0, 1},
			},
		}
	case ZAxis:
		return &Matrix{
			M: [4][4]float64{
				{math.Cos(rad), math.Sin(rad), 0, 0},
				{-math.Sin(rad), math.Cos(rad), 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		}
	default:
		return nil
	}
}

func RotationSequenceMatrix(dx, dy, dz float64, seq RotationSequence) *Matrix {
	xRot := RotationMatrix(dx, XAxis)
	yRot := RotationMatrix(dy, YAxis)
	zRot := RotationMatrix(dz, ZAxis)
	switch seq {
	case RotationXYZ:
		return xRot.MultiplyMatrix(*yRot.MultiplyMatrix(*zRot))
	case RotationYXZ:
		return yRot.MultiplyMatrix(*xRot.MultiplyMatrix(*zRot))
	case RotationXZY:
		return xRot.MultiplyMatrix(*zRot.MultiplyMatrix(*yRot))
	case RotationYZX:
		return yRot.MultiplyMatrix(*zRot.MultiplyMatrix(*xRot))
	case RotationZYX:
		return zRot.MultiplyMatrix(*yRot.MultiplyMatrix(*xRot))
	case RotationZXY:
		return zRot.MultiplyMatrix(*xRot.MultiplyMatrix(*yRot))
	default:
		return nil
	}
}
