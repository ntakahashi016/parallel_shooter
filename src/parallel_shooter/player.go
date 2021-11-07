package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	common
	Charactor
	input *Input
	game *Game
}

func NewPlayer(x,y,h,w int, p bool, hp, v int, i *Input, g *Game) *Player {
	pl := &Player{
		Charactor: Charactor{
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
		},
		input: i,
		game: g,
	}
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
		p.game.setObject(newShot(p.x,p.y,5,5,p.phase,0,0,1),ebiten.NewImage(5,5))
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
