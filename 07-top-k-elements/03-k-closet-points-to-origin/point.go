package main

import "strconv"

type Point struct {
	x, y int
}

func InitPoint(x, y int) Point {
	p := Point{}
	p.x = x
	p.y = y
	return p
}

func (p Point) Less(other Point) bool {
	return p.DistFromOrigin() < other.DistFromOrigin()
}

func (p Point) String() string {
	return "[" + strconv.Itoa(p.x) + ", " + strconv.Itoa(p.y) + "]"
}

func (p Point) DistFromOrigin() int {
	return (p.x * p.x) + (p.y * p.y)
}
