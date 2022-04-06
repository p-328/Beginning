package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle {radius: 5}
	fmt.Println("Area:", c.getArea())
	t := Triangle { base: 9, height: 10 }
	fmt.Println("Area:", t.getArea())
}
type Shape interface {
	getArea() float64
}
type Side4 interface {
	getPerimeter() float64
}
type Circle struct {
	radius float64
}
type Triangle struct {
	base float64
	height float64
}
func (c *Circle) getArea() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
func (t *Triangle) getArea() float64 {
	formula := t.base * t.height
	return formula/2
}
