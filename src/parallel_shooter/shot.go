package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shot struct {
	common
	Object
	dir int
	speed int
	atack int
	enemies []interface{}
}

func newShot(x,y,h,w int, p bool, d,spd,a int) *Shot{
	s := &Shot{}
	s.x = x
	s.y = y
	s.height = h
	s.width = w
	s.phase = p
	s.dir = d
	s.speed = spd
	s.atack = a
	s.enemies = nil
	return s
}

func (s *Shot)Update() error {
	prev_y := s.y
	s.y -= s.speed
	hitArea := NewArea(NewPoint(s.x,prev_y),NewPoint(s.x+s.width,s.y))
	for _,o := range s.enemies {
		e,_ := o.(common)
		if hitArea.isHit(e.getArea()) {
			s.speed = 0
			s = nil
			o = nil
		}
	}
	return nil
}
func (s *Shot)Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	return nil
}

func (s *Shot) getx() int { return s.x }
func (s *Shot) gety() int { return s.y }
func (s *Shot) getArea() *Area { return NewArea(NewPoint(s.x, s.y), NewPoint(s.x+s.width, s.y+s.height)) }

func (s *Shot) addEnemy(e *Charactor) {
	s.enemies = append(s.enemies, e)
}

func (s *Shot) deletEnemy(e *Charactor) {
	for _,v := range s.enemies {
		if v == e {
			v = nil
		}
	}
}
