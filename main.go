package main

import (
	"log"

	"github.com/OkaniYoshiii/brick-breaker-go/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	aspectRatio  = 16.0 / 9.0
	ScreenWidth  = 1920
	ScreenHeight = int(float64(ScreenWidth) / aspectRatio)
)

type Game struct {
	Scene scene.Scene
}

func (g *Game) Update() error {
	return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
	// screen.DrawImage(g.Scene.Image, nil)
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

	game := &Game{
		Scene: scene.First(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
