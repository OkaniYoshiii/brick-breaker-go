package game

import (
	"image/color"

	"github.com/OkaniYoshiii/brick-breaker-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	Circle

	Direction Vector
	Speed     int
}

func NewBall(radius, speed int, color color.Color) Ball {
	size := radius * 2

	ball := Ball{
		Circle: Circle{
			GameObject: GameObject{
				Image:            ebiten.NewImage(size, size),
				DrawImageOptions: &ebiten.DrawImageOptions{},
			},
			Radius: radius,
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

func (b *Ball) Rotate(angleInDeg int) {
	angleInRad := utils.DegToRad(angleInDeg)

	width := b.Bounds().Dx()
	height := b.Bounds().Dy()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	oldX := b.X()
	oldY := b.Y()
	b.GeoM.Translate(-b.X()-halfW, -b.Y()-halfH)
	b.GeoM.Rotate(angleInRad)
	b.GeoM.Translate(oldX+halfW, oldY+halfH)
}

func (b *Ball) BounceInsideImg(img *ebiten.Image) {
	// Left collision
	if b.X()-float64(b.Radius) < 0 {
		b.Direction.X *= -1
	}

	// Right collision
	if b.X()+float64(b.Radius) > float64(img.Bounds().Dx()) {
		b.Direction.X *= -1
	}

	// Top collision
	if b.Y()-float64(b.Radius) < 0 {
		b.Direction.Y *= -1
	}

	// Bottom collision
	if b.Y()+float64(b.Radius) > float64(img.Bounds().Dy()) {
		b.Direction.Y *= -1
	}
}
