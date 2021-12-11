package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Player struct {
	Common
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
		o1 := Object{game: p.game, point: p.point, height: 5, width: 5, phase: p.phase, images: p.shotImages}
		s1 := newShot(o1, 1, NewVector(math.Cos(p.direction)*5, math.Sin(p.direction)*5))
		o2 := Object{game: p.game, point: p.point, height: 5, width: 5, phase: p.phase, images: p.shotImages}
		s2 := newShot(o2, 1, NewVector(math.Cos(p.direction - 1.0/12.0 * math.Pi)*5, math.Sin(p.direction - 1.0/12.0 * math.Pi)*5))
		o3 := Object{game: p.game, point: p.point, height: 5, width: 5, phase: p.phase, images: p.shotImages}
		s3 := newShot(o3, 1, NewVector(math.Cos(p.direction + 1.0/12.0 * math.Pi)*5, math.Sin(p.direction + 1.0/12.0 * math.Pi)*5))
		shots := []*Shot{s1,s2,s3}
		for _,shot := range shots {
			shot.setCenter(p.Center())
			enemies := p.game.getEnemies()
			for _, e := range enemies {
				shot.addEnemy(e)
			}
			p.game.setObject(shot)
		}
	case KeyCtrl:
		p.game.phaseShift()
	}
}

func (p *Player) move(v Vector) {
	np := NewPoint(p.point.X() + v.X(), p.point.Y() + v.Y())
	a := NewArea(np, NewPoint(np.X()+float64(p.width-1), np.Y()+float64(p.height-1)))
	if !p.point.equal(np) {
		p.direction = math.Atan2(v.Y(), v.X())
	}
	if p.game.insideOfScreen(a) {
		p.point = np
	} else {
		rp := p.game.repointOnScreen(a)
		p.point = rp
	}
}

func (p *Player) Update() {
	commands := p.input.getCommands()
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

func (p *Player) X() int { return int(p.point.X()) }
func (p *Player) Y() int { return int(p.point.Y()) }
func (p *Player) Area() *Area {
	return NewArea(NewPoint(p.point.X(), p.point.Y()), NewPoint(p.point.X()+float64(p.width-1), p.point.Y()+float64(p.height-1)))
}
func (p *Player) Phase() Phase      { return p.phase }
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

func (p *Player) Image() *ebiten.Image {
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

func (p *Player) Center() *Point {
	a := p.Area()
	x := (a.p2.x - a.p1.x) / 2 + a.p1.x
	y := (a.p2.y - a.p1.y) / 2 + a.p1.y
	return NewPoint(x,y)
}
