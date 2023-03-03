package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func CartesianDistance(a, b *Point) float64 {
	return math.Sqrt((a.GetX()-b.GetX())*(a.GetX()-b.GetX()) - (a.GetY()-b.GetY())*(a.GetY()-b.GetY()))
}
func main() {

	point1 := New(6.1, 5.9)
	point2 := New(6.3, 5.8)

	fmt.Println("расстояние:", CartesianDistance(point1, point2))

}
