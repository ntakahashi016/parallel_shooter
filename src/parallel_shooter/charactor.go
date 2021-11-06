package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Charactor struct {
	Object
	hp int
	score int
	value int
	Controllable
}

func NewCharactor(x,y,h,w int, p bool, hp, v int) *Charactor {
	c := &Charactor{
		Object: Object{
			x: x,
			y: y,
			height: h,
			width: w,
			phase: p,
		},
		hp: hp,
		value: v,
		score: 0,
	}
	return c
}

func (c *Charactor) command(dir Dir) error {
	switch dir {
	case DirUp:
		c.y = c.y - 1
	case DirLeft:
		c.x = c.x - 1
	case DirRight:
		c.x = c.x + 1
	case DirDown:
		c.y = c.y + 1
	}
	return nil
}

func (c *Charactor) Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	return nil
}

func (c *Charactor) getx() int { return c.x }
func (c *Charactor) gety() int { return c.y }
