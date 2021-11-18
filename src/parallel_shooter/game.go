package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"sync"
)

const width = 640
const height = 480

type Phase int
const (
	Light Phase = iota
	Dark
)

type ImageSet struct {
	light *ebiten.Image
	dark *ebiten.Image
	gray *ebiten.Image
}

type Game struct{
	objects []interface{}
	mu sync.Mutex
	phase Phase
	clear bool
	pf *PlayerFactory
	ef *Enemy1Factory
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.phase = Light
	g.clear = false
	g.objects = []interface{}{}
	g.pf,_ = NewPlayerFactory(g)
	g.objects = append(g.objects, g.pf.NewPlayer())
	g.ef,_ = NewEnemy1Factory(g)
	g.objects = append(g.objects, g.ef.NewEnemy1d())
	g.objects = append(g.objects, g.ef.NewEnemy1l())
	return g, nil
}

func (g *Game) Update() Mode {
	for _,v := range g.objects {
		c := v.(common)
		go c.run()
	}
	if g.clear { return RESULT }
	return GAME
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.phase == Light {
		screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	} else {
		screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	}

	for _, v := range g.objects {
		g.mu.Lock()
		c := v.(common)
		// c.Draw(c.getImage())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.getx()), float64(c.gety()))
		screen.DrawImage(c.getImage(),op)
		g.mu.Unlock()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var w int = width
	var h int = height
	return w, h
}

func (g *Game) setObject(o interface{}) {
	g.objects = append(g.objects, o)
}

func (g *Game) deleteObject(o interface{}) {
	g.mu.Lock()
	defer g.mu.Unlock()
	newObjects := []interface{}{}
	for _,v := range g.objects {
		if v == o {
			continue
		}
		newObjects = append(newObjects, v)
	}
	g.objects = newObjects
}

func (g *Game) outOfScreen(a *Area) bool {
	if (a.p2.x < 0 || width <= a.p1.x) || (a.p2.y < 0 || height <= a.p1.y) {
		return true
	}
	return false
}

func (g *Game) insideOfScreen(a *Area) bool {
	if (a.p1.x >= 0 && width > a.p2.x) && (a.p1.y >= 0 && height > a.p2.y) {
		return true
	}
	return false
}

func (g *Game) repointOnScreen(a *Area) Point {
	p := Point{x: a.p1.x, y: a.p1.y}
	if a.p1.x < 0 {
		p.x = 0
	} else if width <= a.p2.x {
		p.x = width - (a.p2.x - a.p1.x)
	}
	if a.p1.y < 0 {
		p.y = 0
	} else if height <= a.p2.y {
		p.y = height - (a.p2.y - a.p1.y)
	}
	return p
}

func (g *Game) getEnemies() []*Character {
	var enemies []*Character
	for _,v := range g.objects {
		switch v.(type) {
		case *Character:
			enemies = append(enemies, v.(*Character))
		}
	}
	return enemies
}

func (g *Game) getPlayers() []*Player {
	var players []*Player
	for _,v := range g.objects {
		switch v.(type) {
		case *Player:
			players = append(players, v.(*Player))
		}
	}
	return players
}

func (g *Game) getPhase() Phase { return g.phase }

func (g *Game) phaseShift() {
	if g.phase == Light {
		g.phase = Dark
	} else {
		g.phase = Light
	}
	players := g.getPlayers()
	for _,p := range players {
		p.setPhase(g.phase)
	}
}

func (g *Game) checkGameClear() {
	var flag bool
	flag = true
	for _,v := range g.objects {
		switch v.(type) {
		case *Character:
			flag = false
		}
	}
	g.clear = flag
}
