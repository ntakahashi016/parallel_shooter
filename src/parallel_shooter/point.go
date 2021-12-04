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
