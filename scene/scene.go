package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type SceneObject interface {
	Update() error
	Draw(img *ebiten.Image)
}
