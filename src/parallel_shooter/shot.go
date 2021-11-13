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
	attack int
	enemies []interface{}
	image *ebiten.Image
	game *Game
}

func newShot(x,y,h,w int, p Phase, d,spd,a int, i *ebiten.Image, g *Game) *Shot{
	s := &Shot{}
	s.x = x
	s.y = y
	s.height = h
	s.width = w
	s.phase = p
	s.dir = d
	s.speed = spd
	s.attack = a
	s.enemies = []interface{}{}
	s.image = i
	s.game = g
	return s
}

func (s *Shot) run() {
	prev_y := s.y
	s.y -= s.speed
	hitArea := NewArea(NewPoint(s.x,prev_y),NewPoint(s.x+s.width,s.y))
	for _,o := range s.enemies {
		e,_ := o.(common)
		if e.getPhase() == s.phase && hitArea.isHit(e.getArea()) {
			e.(Characteristic).hit(s.attack)
			s.destroy()
		}
	}
}

func (s *Shot)Update() error {
	return nil
}

func (s *Shot)Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	return nil
}

func (s *Shot) getx() int { return s.x }
func (s *Shot) gety() int { return s.y }
func (s *Shot) getArea() *Area { return NewArea(NewPoint(s.x, s.y), NewPoint(s.x+s.width, s.y+s.height)) }

func (s *Shot) addEnemy(e interface{}) {
	s.enemies = append(s.enemies, e)
}

func (s *Shot) deletEnemy(e interface{}) {
	for _,v := range s.enemies {
		if v == e {
			v = nil
		}
	}
}

func (s *Shot) getImage() *ebiten.Image { return s.image }

func (s *Shot) destroy() {
	s.game.deleteObject(s)
}
