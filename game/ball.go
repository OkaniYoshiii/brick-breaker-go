package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	GameObject

	Speed  int
	Radius int
}

func NewBall(radius, speed int, color color.Color) Ball {
	size := radius * 2

	ball := Ball{
		GameObject: GameObject{
			Image:            ebiten.NewImage(size, size),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		},
		Speed: speed,
	}

	vector.DrawFilledCircle(
		ball.Image,
		float32(radius),
		float32(radius),
		float32(radius),
		color,
		false,
	)

	return ball
}
