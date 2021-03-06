package parallel_shooter

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Characteristic interface {
	hit(int)
}

type CharacterAttr struct {
	hp         int
	score      int
	value      int
	shotImages *ImageSet
}

type Character struct {
	Common
	Object
	Characteristic
	CharacterAttr
	rand *rand.Rand
	direction float64
	commander Commander
}

func NewCharacter(o Object, ca CharacterAttr, cmd Commander) *Character {
	c := &Character{
		Object:        o,
		CharacterAttr: ca,
	}
	source := rand.NewSource(time.Now().UnixNano())
	c.rand = rand.New(source)
	c.commander = cmd
	return c
}

func (c *Character) command(cmd Command) error {
	switch cmd {
	case KeySpace:
		o := Object{game: c.game, point: c.point, height: 5, width: 5, phase: c.phase, images: c.shotImages}
		enemies := c.game.getPlayers()
		e := enemies[0]
		radian := c.Center().direction(e.Center())
		shot := newShot(o, 1, NewVector(math.Cos(radian)*5, math.Sin(radian)*5))
		shot.setCenter(c.Center())
		for _, e := range enemies {
			shot.addEnemy(e)
		}
		c.game.setObject(shot)
	}
	return nil
}

func (c *Character) move(v Vector) {
	np := NewPoint(c.point.X() + v.X(), c.point.Y() + v.Y())
	if !c.point.equal(np) {
		c.direction = math.Atan2(v.Y(), v.X())
	}
	c.point = np
}

func (c *Character) Update() {
	if c.game.outOfScreen(c.Area()) {
		c.game.deleteObject(c)
		return
	}
	var cmd Command
	switch c.rand.Intn(100) {
	case 0:
		cmd = KeySpace
	}
	c.command(cmd)
	c.commander.Update()
	c.move(c.commander.getVector())
}

func (c *Character) Draw(img *ebiten.Image) error {
	return nil
}

func (c *Character) X() int { return int(c.point.X()) }
func (c *Character) Y() int { return int(c.point.Y()) }
func (c *Character) Area() *Area {
	return NewArea(NewPoint(c.point.X(), c.point.Y()), NewPoint(c.point.X()+float64(c.width), c.point.Y()+float64(c.height)))
}
func (c *Character) Phase() Phase { return c.phase }

func (c *Character) hit(damage int) {
	c.hp -= damage
	if c.hp <= 0 {
		c.destroy()
	}
}

func (c *Character) destroy() {
	c.game.deleteObject(c)
}

func (c *Character) Image() *ebiten.Image {
	var i *ebiten.Image
	gPhase := c.game.getPhase()
	if c.phase == gPhase {
		switch gPhase {
		case LIGHT_PHASE:
			i = c.images.light
		case DARK_PHASE:
			i = c.images.dark
		}
	} else {
		i = c.images.gray
	}
	return i
}

func (c *Character) Center() *Point {
	a := c.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
