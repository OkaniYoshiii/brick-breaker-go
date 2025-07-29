package entities

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
