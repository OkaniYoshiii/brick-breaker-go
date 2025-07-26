package utils

import "math"

// Converts angle in degrees to radians
func DegToRad(angle int) float64 {
	return float64(angle) * (math.Pi / 180.0)
}
