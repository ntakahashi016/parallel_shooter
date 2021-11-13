package parallel_shooter

import (
	// "image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Characteristic interface {
	hit(int)
}

type Character struct {
	common
	Object
	Characteristic
	hp int
	score int
	value int
}

func NewCharacter(object *Object, hp,v int) *Character {
	c := &Character{
		Object: *object,
	}
	c.hp = hp
	c.value = v
	c.score = 0
	return c
}

func (c *Character) Update() error {
	return nil
}

func (c *Character) Draw(img *ebiten.Image) error {
	return nil
}

func (c *Character) getx() int { return c.x }
func (c *Character) gety() int { return c.y }
func (c *Character) getArea() *Area { return NewArea(NewPoint(c.x, c.y), NewPoint(c.x+c.width, c.y+c.height)) }
func (c *Character) getPhase() Phase { return c.phase }

func (c *Character) hit(damage int) {
	c.hp -= damage
	if c.hp <= 0 {
		c.destroy()
	}
}

func (c *Character) destroy() {
	c.game.deleteObject(c)
	c.game.checkGameClear()
}

func (c *Character) getImage() *ebiten.Image {
	var i *ebiten.Image
	gPhase := c.game.getPhase()
	if c.phase == gPhase {
		switch gPhase {
		case Light:
			i = c.images.light
		case Dark:
			i = c.images.dark
		}
	} else {
		i = c.images.gray
	}
	return i
}
