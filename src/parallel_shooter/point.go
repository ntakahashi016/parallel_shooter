package parallel_shooter

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x,y float64) *Point {
	p := &Point{}
	p.x = x
	p.y = y
	return p
}

func (p *Point) X() float64 { return p.x }
func (p *Point) Y() float64 { return p.y }

func (p *Point) equal(other *Point) bool {
	return p.x == other.X() && p.y == other.Y()
}

func (p *Point) direction(other *Point) float64 {
	radian := math.Atan2(other.Y() - p.Y(), other.X() - p.X())
	return radian
}

func (p *Point) offset(v Vector) *Point {
	return NewPoint(p.X() + v.X(), p.Y() + v.Y())
}
