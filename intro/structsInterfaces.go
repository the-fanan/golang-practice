package main

import (
	"fmt"
	"math"
)
//structs
type Circle struct {
	x, y, r float64
}
//methods
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

type Rectangle struct {
	x,y, l, b float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.b
}
//interface
type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
//embeded types
type MultiShape struct {
	shapes []Shape
}
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	var c Circle 
	d := new(Circle)//returns a pointer to the struct -- uncommon convention
	e := Circle{x: 0, y: 0, r: 3}
	f := Circle{0, 0, 2}//assigns values in order in chich they are defined in the struct
	g := &Circle{1, 2, 5}//also returns a pointer to the struct
	fmt.Println(c,d,e,f, g)
	fmt.Println(e.area())
	r := Rectangle{l: 3, b: 4}
	fmt.Println(r.area())
	fmt.Println(totalArea(&e, &r))
	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0, 0, 5},
			&Rectangle{0, 0, 10, 10},
		},
	}

	fmt.Println(multiShape.area())
}