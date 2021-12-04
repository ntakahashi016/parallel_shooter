package parallel_shooter

type Motion struct {
	vector Vector
	times int
}

func NewMotion(v Vector, t int) *Motion {
	m := &Motion{}
	m.vector = v
	m.times = t
	return m
}

func (m *Motion) Vector() Vector { return m.vector }
func (m *Motion) Times() int { return m.times }

