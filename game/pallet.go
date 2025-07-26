package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pallet struct {
	GameObject

	Direction Vector
	Speed     int
}

func NewPallet(width, height, speed int, color color.Color) Pallet {
	pallet := Pallet{
		GameObject: GameObject{
			Image:            ebiten.NewImage(width, height),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		},
		Speed: 5,
	}

	vector.DrawFilledRect(
		pallet.Image,
		0,
		0,
		float32(width),
		float32(height),
		color,
		false,
	)

	return pallet
}

func (p *Pallet) Width() int {
	return p.Bounds().Size().X
}

func (p *Pallet) Height() int {
	return p.Bounds().Size().Y
}
