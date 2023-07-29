// 24. Разработать программу нахождения расстояния между двумя
// точками, которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) DistanceTo(p2 Point) float64 {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(4, 5)

	fmt.Println(p1.DistanceTo(p2))
}
