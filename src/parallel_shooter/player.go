package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Player struct {
	common
	Object
	Characteristic
	CharacterAttr
	input *Input
	velocity Vector
	direction float64
}

func NewPlayer(o Object, ca CharacterAttr, input *Input) *Player {
	pl := &Player{
		Object:        o,
		CharacterAttr: ca,
	}
	pl.input = input
	pl.velocity = NewVector(0,0)
	pl.direction = 1.5 * math.Pi
	return pl
}

func (p *Player) command(cmd Command) {
	switch cmd {
	case KeySpace:
		o := Object{game: p.game, x: p.x, y: p.y, height: 5, width: 5, phase: p.phase, images: p.shotImages}
		shot := newShot(o, 1, NewVector(math.Cos(p.direction)*5, math.Sin(p.direction)*5))
		shot.setCenter(p.getArea())
		enemies := p.game.getEnemies()
		for _, e := range enemies {
			shot.addEnemy(e)
		}
		p.game.setObject(shot)
	case KeyCtrl:
		p.game.phaseShift()
	}
}

func (p *Player) move(v Vector) {
	x := p.x + int(v.X())
	y := p.y + int(v.Y())
	a := NewArea(NewPoint(x, y), NewPoint(x+p.width-1, y+p.height-1))
	if p.x!=x || p.y!=y {
		p.direction = math.Atan2(v.Y(), v.X())
	}
	if p.game.insideOfScreen(a) {
		p.x = x
		p.y = y
	} else {
		rp := p.game.repointOnScreen(a)
		p.x = rp.x
		p.y = rp.y
	}
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) run() {
	commands := p.input.getCommands()
	// for _, command := range commands {
	// 	p.command(command)
	// }
	for len(commands) > 0 {
		command := commands[0]
		commands = commands[1:]
		p.command(command)
	}
	vector := p.input.getVector()
	p.move(vector)
}


func (p *Player) Draw(img *ebiten.Image) error {
	return nil
}

func (p *Player) getx() int { return p.x }
func (p *Player) gety() int { return p.y }
func (p *Player) getArea() *Area {
	return NewArea(NewPoint(p.x, p.y), NewPoint(p.x+p.width-1, p.y+p.height-1))
}
func (p *Player) getPhase() Phase      { return p.phase }
func (p *Player) setPhase(phase Phase) { p.phase = phase }
func (p *Player) hit(damage int) {
	p.hp -= damage
	if p.hp <= 0 {
		p.destroy()
	}
}

func (p *Player) destroy() {
	p.game.deleteObject(p)
	p.game.gameover()
}

func (p *Player) getImage() *ebiten.Image {
	var i *ebiten.Image
	gPhase := p.game.getPhase()
	if p.phase == gPhase {
		switch gPhase {
		case LIGHT_PHASE:
			i = p.images.light
		case DARK_PHASE:
			i = p.images.dark
		}
	} else {
		i = p.images.gray
	}
	return i
}
