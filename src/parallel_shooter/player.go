package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	common
	Object
	Characteristic
	Character
	input *Input
}

func NewPlayer(x,y,h,w int, p bool, hp, v int, g *Game, i *Input) *Player {
	pl := &Player{}
	pl.x = x
	pl.y = y
	pl.height = h
	pl.width = w
	pl.phase = p
	pl.hp = hp
	pl.value = v
	pl.score = 0
	pl.game = g
	pl.input = i
	return pl
}

func (p *Player) command(cmd Command) error {
	switch cmd {
	case DirUp:
		p.y = p.y - 1
	case DirLeft:
		p.x = p.x - 1
	case DirRight:
		p.x = p.x + 1
	case DirDown:
		p.y = p.y + 1
	case KeySpace:
		shot := newShot(p.x,p.y,5,5,p.phase,0,5,1, ebiten.NewImage(5,5), p.game)
		e := p.game.getEnemy()
		if e!=nil {
			shot.addEnemy(e)
		}
		p.game.setObject(shot, shot.getImage())
	case KeyCtrl:
		p.game.phaseShift()
	}
	return nil
}

func (p *Player) Update() error {
	if cmd, ok := p.input.getCommand(); ok {
		p.command(cmd)
	}
	return nil
}

func (p *Player) Draw(img *ebiten.Image) error {
	img.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	return nil
}

func (p *Player) getx() int { return p.x }
func (p *Player) gety() int { return p.y }
func (p *Player) getArea() *Area { return NewArea(NewPoint(p.x, p.y), NewPoint(p.x+p.width, p.y+p.height)) }
func (p *Player) getPhase() bool { return p.phase }
func (p *Player) setPhase(phase bool) { p.phase = phase }
func (p *Player) hit(damage int) {
	p.hp -= damage
	if p.hp <= 0 {
		p.destroy()
	}
}
func (p *Player) destroy() {
	p.game.deleteObject(p)
}
