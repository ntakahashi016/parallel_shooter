package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Shot struct {
	Common
	Object
	velocity Vector
	attack   int
	enemies  []interface{}
}

func newShot(o Object, a int, v Vector) *Shot {
	s := &Shot{
		Object: o,
	}
	s.velocity = v
	s.attack = a
	s.enemies = []interface{}{}
	return s
}

func (s *Shot) Update() {
	if s.game.outOfScreen(s.Area()) {
		s.game.deleteObject(s)
		return
	}
	prev_y := s.y
	s.y += int(s.velocity.Y())
	s.x += int(s.velocity.X())
	hitArea := NewArea(NewPoint(s.x, prev_y), NewPoint(s.x+s.width, s.y))
	for _, o := range s.enemies {
		e, _ := o.(Common)
		if e.Phase() == s.phase && hitArea.isHit(e.Area()) {
			e.(Characteristic).hit(s.attack)
			s.destroy()
		}
	}
}

func (s *Shot) Draw(img *ebiten.Image) error {
	return nil
}

func (s *Shot) X() int { return s.x }
func (s *Shot) Y() int { return s.y }
func (s *Shot) Area() *Area {
	return NewArea(NewPoint(s.x, s.y), NewPoint(s.x+s.width, s.y+s.height))
}

func (s *Shot) addEnemy(e interface{}) {
	s.enemies = append(s.enemies, e)
}

func (s *Shot) deletEnemy(e interface{}) {
	for _, v := range s.enemies {
		if v == e {
			v = nil
		}
	}
}

func (s *Shot) Image() *ebiten.Image {
	var i *ebiten.Image
	gPhase := s.game.getPhase()
	if s.phase == gPhase {
		switch gPhase {
		case LIGHT_PHASE:
			i = s.images.light
		case DARK_PHASE:
			i = s.images.dark
		}
	} else {
		i = s.images.gray
	}
	return i
}

func (s *Shot) destroy() {
	s.game.deleteObject(s)
}

func (s *Shot) setCenter(a *Area) {
	s.x = (a.p2.x - a.p1.x) / 2 + a.p1.x - s.width / 2
	s.y = (a.p2.y - a.p1.y) / 2 + a.p1.y - s.height / 2
}
