package main

import (
	"image/color"
	"log"

	"github.com/OkaniYoshiii/brick-breaker-go/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	aspectRatio  = 16.0 / 9.0
	ScreenWidth  = 1920
	ScreenHeight = int(float64(ScreenWidth) / aspectRatio)

	pallet game.Pallet
	ball   game.Ball
)

func init() {
	// Pallet creation
	pallet = func() game.Pallet {
		width := 150
		height := 25
		pallet := game.NewPallet(width, height, 5, color.RGBA{255, 0, 0, 255})

		pallet.GeoM.Translate(float64(ScreenWidth)/2-float64(width)/2, float64(ScreenHeight-height-15))

		return pallet
	}()

	// Ball creation
	ball = func() game.Ball {
		radius := 10
		size := radius * 2
		speed := 5

		ball := game.NewBall(radius, speed, color.RGBA{0, 0, 0, 255})

		ball.GeoM.Translate(float64(ScreenWidth/2-radius), pallet.Y()-float64(size)-10)

		return ball
	}()
}

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		pallet.GeoM.Translate(-float64(pallet.Speed), 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		pallet.GeoM.Translate(float64(pallet.Speed), 0)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})

	screen.DrawImage(pallet.Image, pallet.DrawImageOptions)
	screen.DrawImage(ball.Image, ball.DrawImageOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	width := 700
	height := int(float64(width) / aspectRatio)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Brick Breaker Go")
	ebiten.SetWindowPosition(0, 0)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
