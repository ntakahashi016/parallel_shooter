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

func (p *Point) X() int { return p.x }
func (p *Point) Y() int { return p.y }

func (p *Point) equal(other *Point) bool {
	return p.x == other.X() && p.y == other.Y()
}
