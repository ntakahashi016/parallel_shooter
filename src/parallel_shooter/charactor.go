package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type common interface {
	Update() error
	Draw(img *ebiten.Image) error
	getx() int
	gety() int
}

type Charactor struct {
	common
	Object
	hp int
	score int
	value int
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

func (c *Charactor) Update() error {
	return nil
}

func (c *Charactor) Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	return nil
}

func (c *Charactor) getx() int { return c.x }
func (c *Charactor) gety() int { return c.y }
