package entities

import (
	"github.com/OkaniYoshiii/brick-breaker-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func (gO *GameObject) ImgX() float64 {
	return utils.ImgX(gO.DrawImageOptions)
}

func (gO *GameObject) ImgY() float64 {
	return utils.ImgY(gO.DrawImageOptions)
}
