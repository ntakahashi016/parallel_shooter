package parallel_shooter

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	common
	Object
	Characteristic
	CharacterAttr
	rand *rand.Rand
}

func NewCharacter(o Object, ca CharacterAttr) *Character {
	c := &Character{
		Object:        o,
		CharacterAttr: ca,
	}
	source := rand.NewSource(time.Now().UnixNano())
	c.rand = rand.New(source)
	return c
}

func (c *Character) command(cmd Command) error {
	switch cmd {
	case DirUp:
		c.y = c.y - 1
	case DirLeft:
		c.x = c.x - 1
	case DirRight:
		c.x = c.x + 1
	case DirDown:
		c.y = c.y + 1
	case KeySpace:
		o := Object{game: c.game, x: c.x, y: c.y, height: 5, width: 5, phase: c.phase, images: c.shotImages}
		shot := newShot(o, 1, NewVector(0, -5))
		enemies := c.game.getPlayers()
		for _, e := range enemies {
			shot.addEnemy(e)
		}
		c.game.setObject(shot)
		// case KeyCtrl:
		// 	c.game.phaseShift()
	}
	return nil
}

func (c *Character) Update() error {
	return nil
}

func (c *Character) run() {
	if c.game.outOfScreen(c.getArea()) {
		c.game.deleteObject(c)
		return
	}
	var cmd Command
	switch c.rand.Intn(5) {
	case 0:
		cmd = DirUp
	case 1:
		cmd = DirRight
	case 2:
		cmd = DirDown
	case 3:
		cmd = DirLeft
	case 4:
		cmd = KeySpace
	}
	c.command(cmd)
}

func (c *Character) Draw(img *ebiten.Image) error {
	return nil
}

func (c *Character) getx() int { return c.x }
func (c *Character) gety() int { return c.y }
func (c *Character) getArea() *Area {
	return NewArea(NewPoint(c.x, c.y), NewPoint(c.x+c.width, c.y+c.height))
}
func (c *Character) getPhase() Phase { return c.phase }

func (c *Character) hit(damage int) {
	c.hp -= damage
	if c.hp <= 0 {
		c.destroy()
	}
}

func (c *Character) destroy() {
	c.game.deleteObject(c)
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
