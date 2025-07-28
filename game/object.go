package game

import "github.com/hajimehoshi/ebiten/v2"

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func (gO *GameObject) ImgX() float64 {
	return gO.GeoM.Element(0, 2)
}

func (gO *GameObject) ImgY() float64 {
	return gO.GeoM.Element(1, 2)
}
