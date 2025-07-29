package scene

import (
	"image/color"

	"github.com/OkaniYoshiii/brick-breaker-go/game"
	"github.com/OkaniYoshiii/brick-breaker-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

func First() *FirstLevel {
	scene := func() *FirstLevel {
		width, height := 1000, 1000

		lvl := &FirstLevel{
			Image:            ebiten.NewImage(width, height),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		}

		return lvl
	}()

	scene.bounds = func() Bounds {
		width, height := 500, 500

		bounds := Bounds{
			Image:            ebiten.NewImage(width, height),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		}

		tx := float64(scene.Bounds().Dx()/2 - width/2)
		ty := float64(scene.Bounds().Dy()/2 - height/2)

		bounds.GeoM.Translate(tx, ty)

		return bounds
	}()

	// Pallet creation
	scene.pallet = func() game.Pallet {
		width := 150
		height := 25
		pallet := game.NewPallet(width, height, 5, color.RGBA{255, 0, 0, 255})

		pallet.GeoM.Translate(float64(scene.bounds.Bounds().Dx())/2-float64(width)/2, float64(scene.bounds.Bounds().Dy()-height-15))

		return pallet
	}()

	// Ball creation
	scene.ball = func() game.Ball {
		radius := 10
		size := radius * 2
		speed := 5

		ball := game.NewBall(radius, speed, color.RGBA{0, 0, 0, 255})

		ball.GeoM.Translate(float64(scene.bounds.Bounds().Dx()/2-radius), scene.pallet.ImgY()-float64(size)-10)

		return ball
	}()

	return scene
}

type Bounds struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func (b *Bounds) ImgX() float64 {
	return utils.ImgX(b.DrawImageOptions)
}

func (b *Bounds) ImgY() float64 {
	return utils.ImgY(b.DrawImageOptions)
}

type FirstLevel struct {
	*ebiten.Image
	*ebiten.DrawImageOptions

	IsStarted bool

	bounds Bounds
	ball   game.Ball
	pallet game.Pallet
}

func (lvl *FirstLevel) Update() error {
	if !lvl.IsStarted && ebiten.IsKeyPressed(ebiten.KeySpace) {
		lvl.IsStarted = true

		dirX, dirY := utils.RandDirection(180)

		lvl.ball.Direction.X = dirX
		lvl.ball.Direction.Y = dirY
	}

	if !lvl.IsStarted {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		lvl.pallet.GeoM.Translate(-float64(lvl.pallet.Speed), 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		lvl.pallet.GeoM.Translate(float64(lvl.pallet.Speed), 0)
	}

	lvl.ball.GeoM.Translate(lvl.ball.Direction.X*float64(lvl.ball.Speed), lvl.ball.Direction.Y*float64(lvl.ball.Speed))

	lvl.ball.BounceInsideImg(lvl.bounds.Image)

	return nil
}

func (lvl *FirstLevel) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
	lvl.Fill(color.RGBA{100, 255, 100, 255})
	lvl.bounds.Fill(color.RGBA{155, 255, 155, 255})

	lvl.bounds.DrawImage(lvl.pallet.Image, lvl.pallet.DrawImageOptions)
	lvl.bounds.DrawImage(lvl.ball.Image, lvl.ball.DrawImageOptions)

	lvl.DrawImage(lvl.bounds.Image, lvl.bounds.DrawImageOptions)

	screen.DrawImage(lvl.Image, lvl.DrawImageOptions)
}
