package scene

import (
	"image/color"

	"github.com/OkaniYoshiii/brick-breaker-go/entities"
	"github.com/OkaniYoshiii/brick-breaker-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayArea struct {
	*ebiten.Image
	*ebiten.DrawImageOptions
}

func (pA *PlayArea) ImgX() float64 {
	return utils.ImgX(pA.DrawImageOptions)
}

func (pA *PlayArea) ImgY() float64 {
	return utils.ImgY(pA.DrawImageOptions)
}

type FirstLevel struct {
	*ebiten.Image
	*ebiten.DrawImageOptions

	IsStarted bool

	PlayArea PlayArea
	ball     entities.Ball
	pallet   entities.Pallet
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

	lvl.ball.BounceInside(entities.GameObject(lvl.PlayArea))
	lvl.ball.BounceOn(lvl.pallet.GameObject)

	lvl.ball.GeoM.Translate(lvl.ball.Direction.X*float64(lvl.ball.Speed), lvl.ball.Direction.Y*float64(lvl.ball.Speed))

	return nil
}

func (lvl *FirstLevel) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
	lvl.Fill(color.RGBA{100, 255, 100, 255})
	lvl.PlayArea.Fill(color.RGBA{155, 255, 155, 255})

	screen.DrawImage(lvl.Image, lvl.DrawImageOptions)
	screen.DrawImage(lvl.PlayArea.Image, lvl.PlayArea.DrawImageOptions)
	screen.DrawImage(lvl.pallet.Image, lvl.pallet.DrawImageOptions)
	screen.DrawImage(lvl.ball.Image, lvl.ball.DrawImageOptions)
}

func First() *FirstLevel {
	scene := func() *FirstLevel {
		width, height := 1280, 720

		lvl := &FirstLevel{
			Image:            ebiten.NewImage(width, height),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		}

		return lvl
	}()

	scene.PlayArea = func() PlayArea {
		width, height := scene.Bounds().Dx()-150, scene.Bounds().Dy()-150

		bounds := PlayArea{
			Image:            ebiten.NewImage(width, height),
			DrawImageOptions: &ebiten.DrawImageOptions{},
		}

		tx := float64(scene.Bounds().Dx()/2 - width/2)
		ty := float64(scene.Bounds().Dy()/2 - height/2)

		bounds.GeoM.Translate(tx, ty)

		return bounds
	}()

	// Pallet creation
	scene.pallet = func() entities.Pallet {
		width := 300
		height := width / 6
		pallet := entities.NewPallet(width, height, 5, color.RGBA{255, 0, 0, 255})

		tx := float64(scene.PlayArea.ImgX()+float64(scene.PlayArea.Bounds().Dx()))/2 - float64(width)/2
		ty := float64(scene.PlayArea.ImgY() + float64(scene.PlayArea.Bounds().Dy()-height-15))

		pallet.GeoM.Translate(tx, ty)

		return pallet
	}()

	// Ball creation
	scene.ball = func() entities.Ball {
		radius := 15
		size := radius * 2
		speed := 9

		ball := entities.NewBall(radius, speed, color.RGBA{0, 0, 0, 255})

		tx := float64(scene.PlayArea.Bounds().Dx()/2 - radius)
		ty := scene.pallet.ImgY() - float64(size) - 10

		ball.GeoM.Translate(tx, ty)

		return ball
	}()

	return scene
}
