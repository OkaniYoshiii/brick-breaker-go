package utils

import (
	"math"
	"math/rand/v2"
)

// Converts angle in degrees to radians
func DegToRad(angle int) float64 {
	return float64(angle) * (math.Pi / 180.0)
}

func RandDirection(maxInDeg float64) (float64, float64) {
	if maxInDeg < 0 || maxInDeg > 360 {
		panic("Argument 'maxInDeg' must be positive and must not be superior to 360")
	}

	ratio := maxInDeg / 360
	angle := rand.Float64() * (math.Pi * 2) * ratio

	dirX := math.Cos(angle)
	dirY := -math.Sin(angle)

	return dirX, dirY
}
