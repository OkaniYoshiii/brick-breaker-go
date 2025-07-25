package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pallet struct {
	*ebiten.Image
	*ebiten.DrawImageOptions

	Speed int
}

type Ball struct {
	*ebiten.Image
	*ebiten.DrawImageOptions

	Speed int
}

var pallet Pallet
var ball Ball

func init() {
	// Pallet creation
	pallet = func() Pallet {
		pallet := Pallet{
			Image:            ebiten.NewImage(120, 25),
			DrawImageOptions: &ebiten.DrawImageOptions{},
			Speed:            5,
		}

		pallet.Fill(color.RGBA{255, 0, 0, 255})
		pallet.GeoM.Translate(150, 150)

		return pallet
	}()

	// Ball creation
	ball = func() Ball {
		size := 25
		posX := 125
		posY := 125

		ball := Ball{
			Image:            ebiten.NewImage(size, size),
			DrawImageOptions: &ebiten.DrawImageOptions{},
			Speed:            5,
		}

		ball.GeoM.Translate(float64(posX), float64(posY))
		vector.DrawFilledCircle(ball.Image, float32(size/2), float32(size/2), float32(size/2), color.RGBA{0, 0, 0, 255}, true)

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
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(ebiten.Monitor().Size())
	ebiten.SetWindowTitle("Brick Breaker Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
