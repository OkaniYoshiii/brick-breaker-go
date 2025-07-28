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

// type Scene struct {
// 	*ebiten.Image

// 	objects []SceneObject
// }

// func (s *Scene) Update() error {
// 	for _, obj := range s.objects {
// 		if err := obj.Update(); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (s *Scene) Draw(screen *ebiten.Image) {
// 	for _, obj := range s.objects {
// 		obj.Draw(s.Image)
// 	}
// }
