package game

import "github.com/hajimehoshi/ebiten/v2"

type GameObject struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func (gO *GameObject) X() float64 {
	return gO.GeoM.Element(0, 2)
}

func (gO *GameObject) Y() float64 {
	return gO.GeoM.Element(1, 2)
}
