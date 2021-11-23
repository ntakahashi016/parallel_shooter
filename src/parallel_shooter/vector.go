package parallel_shooter

import "math"

type Vector interface {
	X() float64
	SetX(x float64)
	Y() float64
	SetY(y float64)
	Magnitude() float64

	/**
	 * Vector{10.0, 2.0}.Normalize() => Vector{1.0, 0.2}
	 * Vector{-20.0, 4.0}.Normalize() => Vector{-1.0, 0.2}
	 */
	Normalize() Vector

	/**
	 * Vector{3.0, 10.0}.Diff(Vector{10.0, 2.0}) => Vector{7.0, -8.0}
	 */
	Diff(other Vector) Vector

	Add(other Vector) Vector
}

func NewVector(x, y float64) Vector {
	return &vector{x: x, y: y}
}

type vector struct {
	x float64
	y float64
}

func (v *vector) X() float64     { return v.x }
func (v *vector) SetX(x float64) { v.x = x }
func (v *vector) Y() float64     { return v.y }
func (v *vector) SetY(y float64) { v.y = y }

func (v *vector) Normalize() Vector {
	if math.Abs(v.x) > math.Abs(v.y) {
		return NewVector(v.x/v.x, v.y/v.x)
	} else {
		return NewVector(v.x/v.y, v.y/v.y)
	}
}
func (v *vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *vector) Diff(other Vector) Vector {
	return NewVector(other.X()-v.x, other.Y()-v.y)
}

func (v *vector) Add(other Vector) Vector {
	return NewVector(other.X()+v.x, other.Y()+v.y)
}
