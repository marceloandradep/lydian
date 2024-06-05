package math

import "math"

const (
	piOver180 = math.Pi / 180
)

func DegreesToRadians(degrees float64) float64 {
	return degrees * piOver180
}
