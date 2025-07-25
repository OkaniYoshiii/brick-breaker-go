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
		pallet := game.NewPallet(150, 25, 5, color.RGBA{255, 0, 0, 255})

		pallet.GeoM.Translate(150, 150)

		return pallet
	}()

	// Ball creation
	ball = func() game.Ball {
		ball := game.NewBall(10, 5, color.RGBA{0, 0, 0, 255})

		ball.GeoM.Translate(150, 250)

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
	ebiten.SetWindowSize(ebiten.Monitor().Size())
	ebiten.SetWindowTitle("Brick Breaker Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
