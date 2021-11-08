package parallel_shooter

type Point struct {
	x int
	y int
}

func NewPoint(x,y int) *Point {
	p := &Point{}
	p.x = x
	p.y = y
	return p
}

func (p *Point) direction(other *Point) *Point {
	x := other.x - p.x
	switch  {
	case x > 0:
		x=1
	case x < 0:
		x=-1
	default:
		x=0
	}
	y := other.y - p.y
	switch {
	case y > 0:
		y=1
	case y < 0:
		y=-1
	default:
		y=0
	}
	return NewPoint(x,y)
}
