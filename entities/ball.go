package entities

import (
	"image/color"
	"log"
	"math"
	"slices"

	"github.com/OkaniYoshiii/brick-breaker-go/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	Circle

	Direction Vector
	Speed     int
}

func NewBall(radius, speed int, color color.Color) Ball {
	size := radius * 2

	ball := Ball{
		Circle: Circle{
			GameObject: GameObject{
				Image:            ebiten.NewImage(size, size),
				DrawImageOptions: &ebiten.DrawImageOptions{},
			},
			Radius: radius,
		},
		Speed: speed,
	}

	vector.DrawFilledCircle(
		ball.Image,
		float32(radius),
		float32(radius),
		float32(radius),
		color,
		false,
	)

	return ball
}

func (b *Ball) Rotate(angleInDeg int) {
	angleInRad := utils.DegToRad(angleInDeg)

	width := b.Bounds().Dx()
	height := b.Bounds().Dy()

	halfW := float64(width / 2)
	halfH := float64(height / 2)

	oldX := b.X()
	oldY := b.Y()
	b.GeoM.Translate(-b.X()-halfW, -b.Y()-halfH)
	b.GeoM.Rotate(angleInRad)
	b.GeoM.Translate(oldX+halfW, oldY+halfH)
}

func (b *Ball) BounceInside(gameObject GameObject) {
	// Left collision
	if b.X()-float64(b.Radius) < gameObject.ImgX() {
		b.Direction.X *= -1
	}

	// Right collision
	if b.X()+float64(b.Radius) > gameObject.ImgX()+float64(gameObject.Bounds().Dx()) {
		b.Direction.X *= -1
	}

	// Top collision
	if b.Y()-float64(b.Radius) < gameObject.ImgY() {
		b.Direction.Y *= -1
	}

	// Bottom collision
	if b.Y()+float64(b.Radius) > gameObject.ImgY()+float64(gameObject.Bounds().Dy()) {
		b.Direction.Y *= -1
	}
}

func (b *Ball) BounceOn(gameObject GameObject) {
	hitEdges := b.CollisionEdges(gameObject)

	if slices.Contains(hitEdges, "left") {
		b.Direction.X = math.Abs(b.Direction.X) * -1.0
	}

	if slices.Contains(hitEdges, "right") {
		b.Direction.X = math.Abs(b.Direction.X)
	}

	if slices.Contains(hitEdges, "top") {
		b.Direction.Y = math.Abs(b.Direction.Y) * -1.0
	}

	if slices.Contains(hitEdges, "bottom") {
		b.Direction.Y = math.Abs(b.Direction.Y)
	}
}

func (b *Ball) CollisionEdges(gameObject GameObject) []string {
	hitEdges := []string{"none", "none"}

	if !b.CollidesWith(gameObject) {
		return hitEdges
	}

	if b.X()-float64(b.Radius) <= gameObject.ImgX() {
		hitEdges = append(hitEdges, "left")
	}

	if b.X()+float64(b.Radius) >= gameObject.ImgX()+float64(gameObject.Bounds().Dx()) {
		hitEdges = append(hitEdges, "right")
	}

	if b.Y()-float64(b.Radius) <= gameObject.ImgY() {
		hitEdges = append(hitEdges, "top")
	}

	if b.Y()+float64(b.Radius) >= gameObject.ImgY()+float64(gameObject.Bounds().Dy()) {
		hitEdges = append(hitEdges, "bottom")
	}

	return hitEdges
}

func (b *Ball) CollidesWith(gameObject GameObject) bool {
	isBetweenRightAndLeftEdge := b.X()+float64(b.Radius) > gameObject.ImgX() && b.X()-float64(b.Radius) < gameObject.ImgX()+float64(gameObject.Bounds().Dx())
	isBetweenTopAndBottomEdge := b.Y()+float64(b.Radius) > gameObject.ImgY() && b.Y()-float64(b.Radius) < gameObject.ImgY()+float64(gameObject.Bounds().Dy())

	log.Println(isBetweenTopAndBottomEdge)

	if isBetweenRightAndLeftEdge && isBetweenTopAndBottomEdge {
		return true
	}

	return false
}
