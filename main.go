package main

import (
	"log"

	"github.com/OkaniYoshiii/brick-breaker-go/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// Pixel perfect dimensions
	LogicalWidth  = 1280
	LogicalHeight = 720
)

type Game struct {
	Scene scene.Scene
}

func (g *Game) Update() error {
	return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return LogicalWidth, LogicalHeight
}

func main() {
	width := 700
	height := int(float64(width) / float64(16.0/9.0))

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
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
