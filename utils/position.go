package utils

import "github.com/hajimehoshi/ebiten/v2"

func ImgX(imgOpts *ebiten.DrawImageOptions) float64 {
	return imgOpts.GeoM.Element(0, 2)
}

func ImgY(imgOpts *ebiten.DrawImageOptions) float64 {
	return imgOpts.GeoM.Element(1, 2)
}
