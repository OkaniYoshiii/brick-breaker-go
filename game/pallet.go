package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pallet struct {
	*ebiten.Image
	*ebiten.DrawImageOptions

	Speed int
}

func NewPallet(width, height, speed int, color color.Color) Pallet {
	pallet := Pallet{
		Image:            ebiten.NewImage(width, height),
		DrawImageOptions: &ebiten.DrawImageOptions{},
		Speed:            5,
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
