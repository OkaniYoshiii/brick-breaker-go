package game

type Circle struct {
	GameObject

	Radius int
}

func (c *Circle) X() float64 {
	return c.ImgX() + float64(c.Radius)
}

func (c *Circle) Y() float64 {
	return c.ImgY() + float64(c.Radius)
}

func (c *Circle) IsCollidingWithImg(rect GameObject) bool {
	if c.X()+float64(c.Radius) < rect.ImgX() {
		return false
	}

	if c.X()-float64(c.Radius) > rect.ImgX()+float64(rect.Bounds().Dx()) {
		return false
	}

	if c.Y()+float64(c.Radius) < rect.ImgY() {
		return false
	}

	if c.Y()-float64(c.Radius) > rect.ImgY()+float64(rect.Bounds().Dy()) {
		return false
	}

	return true
}
