package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func CalculateShape(sh Shape) {
	area := sh.Area()
	perimeter := sh.Perimeter()

	fmt.Println(area, perimeter)

}

func main() {
	CalculateShape(&Rectangle{Width: 2, Height: 3})
	CalculateShape(&Circle{Radius: 3})
}
