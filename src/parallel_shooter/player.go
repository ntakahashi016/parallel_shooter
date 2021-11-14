package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	common
	Object
	Characteristic
	CharacterAttr
	input *Input
}

func NewPlayer(o Object, ca CharacterAttr, i *Input) *Player {
	pl := &Player{
		Object: o,
		CharacterAttr: ca,
	}
	pl.input = i
	return pl
}

func (p *Player) command(cmd Command) error {
	x := p.x
	y := p.y
	switch cmd {
	case DirUp:
		y = p.y - 1
	case DirLeft:
		x = p.x - 1
	case DirRight:
		x = p.x + 1
	case DirDown:
		y = p.y + 1
	case KeySpace:
		o := Object{game: p.game, x: p.x, y: p.y, height: 5, width:5, phase: p.phase, images: p.shotImages}
		shot := newShot(o,0,5,1)
		enemies := p.game.getEnemies()
		for _,e := range enemies {
			shot.addEnemy(e)
		}
		p.game.setObject(shot)
	case KeyCtrl:
		p.game.phaseShift()
	}
	a := NewArea(NewPoint(x,y), NewPoint(x+p.width-1, y+p.height-1))
	if p.game.insideOfScreen(a) {
		p.x = x
		p.y = y
	} else {
		rp := p.game.repointOnScreen(a)
		p.x = rp.x
		p.y = rp.y
	}
	return nil
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) run() {
	if cmd, ok := p.input.getCommand(); ok {
		p.command(cmd)
	}
}

func (p *Player) Draw(img *ebiten.Image) error {
	return nil
}

func (p *Player) getx() int { return p.x }
func (p *Player) gety() int { return p.y }
func (p *Player) getArea() *Area { return NewArea(NewPoint(p.x, p.y), NewPoint(p.x+p.width-1, p.y+p.height-1)) }
func (p *Player) getPhase() Phase { return p.phase }
func (p *Player) setPhase(phase Phase) { p.phase = phase }
func (p *Player) hit(damage int) {
	p.hp -= damage
	if p.hp <= 0 {
		p.destroy()
	}
}

func (p *Player) destroy() {
	p.game.deleteObject(p)
}

func (p *Player) getImage() *ebiten.Image{
	var i *ebiten.Image
	gPhase := p.game.getPhase()
	if p.phase == gPhase {
		switch gPhase {
		case Light:
			i = p.images.light
		case Dark:
			i = p.images.dark
		}
	} else {
		i = p.images.gray
	}
	return i
}
