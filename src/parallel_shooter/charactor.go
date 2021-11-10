package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Charactor struct {
	common
	Object
	hp int
	score int
	value int
	game *Game
}

func NewCharactor(x,y,h,w int, p bool, hp, v int, g *Game) *Charactor {
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
		game: g,
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
func (c *Charactor) getArea() *Area { return NewArea(NewPoint(c.x, c.y), NewPoint(c.x+c.width, c.y+c.height)) }

func (c *Charactor) hit(damage int) {
	c.hp -= damage
	if c.hp <= 0 {
		c.destroy()
	}
}

func (c *Charactor) destroy() {
	c.game.deleteObject(c)
}
