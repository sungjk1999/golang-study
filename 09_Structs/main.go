package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

type Circle struct {
	center Point
	radius float32
}

func (circle Circle) circumference() float32 {
	return 2 * math.Pi * circle.radius
}

func (circle *Circle) modify(newRadius float32) {
	circle.radius = newRadius
}

type Circle_anonymous struct {
	Point
	radius float32
}

func main() {

	circle := Circle{Point{3, 4}, 5}
	circle.modify(1.5)
	fmt.Println(circle)
	fmt.Println(circle.center.x, circle.center.y, circle.radius)

	_circle := Circle_anonymous{Point{3, 4}, 5}
	fmt.Println(_circle)
	fmt.Println(_circle.x, _circle.y, _circle.radius)

}
