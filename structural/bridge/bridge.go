package bridge

import "fmt"

type DrawAPI interface {
	DrawCircle(x, y, radius int)
}

type DrawAPI1 struct {}

func (d *DrawAPI1)DrawCircle(x, y, radius int)  {
	fmt.Printf("DrawAPI1 at (%d, %d): %d\n", x, y, radius)
}

type DrawAPI2 struct {}

func (d *DrawAPI2)DrawCircle(x, y, radius int)  {
	fmt.Printf("DrawAPI2 at (%d, %d): %d\n", x, y, radius)
}

type Shape interface {
	Draw()
}

type DefaultShape struct{
	drawAPI DrawAPI
}

func NewDefaultShape(drawAPI DrawAPI) *DefaultShape {
	return &DefaultShape{
		drawAPI: drawAPI,
	}
}

type CircleShape struct {
	*DefaultShape
	x, y, radius int
}

func NewCircleShape(x, y, radius int, drawAPI DrawAPI) Shape {
	return &CircleShape {
		x: x,
		y: y,
		radius: radius,
		DefaultShape: NewDefaultShape(drawAPI),
	}
}

func (c *CircleShape)Draw() {
	c.drawAPI.DrawCircle(c.x, c.y, c.radius)
}

