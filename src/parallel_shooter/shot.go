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
	return s
}

func (s *Shot)Update() error {
	s.y--
	return nil
}
func (s *Shot)Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
	return nil
}

func (s *Shot) getx() int { return s.x }
func (s *Shot) gety() int { return s.y }
