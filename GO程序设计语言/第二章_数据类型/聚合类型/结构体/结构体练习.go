package main

import "fmt"

type Point struct {
	X, Y int
}

type PointZ struct {
	X, Y, Z int
}

type Person struct {
	name string
	Point
	PointZ
}

func (p Point) String() string {
	return fmt.Sprintf("Point{%d, %d}", p.X, p.Y)
}

func zeroPoint(point Point) {
	point.X = 0
	point.Y = 0
}

func zeroPoint2(point *Point) {
	point.X = 0
	point.Y = 0
}

func main() {
	p := Person{"hdy", Point{1, 2}, PointZ{3, 4, 5}}
	fmt.Println(p.name, p.Point.X, p.Point.Y, p.PointZ.X, p.PointZ.Y, p.Z)
}
