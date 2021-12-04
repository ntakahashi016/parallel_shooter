package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const width = 640
const height = 480

type Phase int
const (
	LIGHT_PHASE Phase = iota
	DARK_PHASE
)

type ImageSet struct {
	light *ebiten.Image
	dark *ebiten.Image
	gray *ebiten.Image
}

type Game struct{
	objects []interface{}
	phase Phase
	clear bool
	pf *PlayerFactory
	ef *Enemy1Factory
	sm *StageManager
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.phase = DARK_PHASE
	g.clear = false
	g.objects = []interface{}{}
	g.pf = NewPlayerFactory(g)
	g.objects = append(g.objects, g.pf.NewObject())
	s1 := NewEnemy1Strategy(g,NewEnemy1Factory(g))
	s2 := NewEnemy1Strategy(g,NewEnemy1Factory(g))
	s3 := NewBoss1Strategy(g,NewBoss1Factory(g))
	g.sm = NewStageManager(g,[]interface{}{s1,s2,s3})
	return g, nil
}

func (g *Game) Update() Mode {
	g.sm.Update()
	for _,v := range g.objects {
		c := v.(Common)
		c.run()
	}
	if g.clear { return MODE_RESULT }
	return MODE_GAME
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.phase == LIGHT_PHASE {
		screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	} else {
		screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	}

	for _, v := range g.objects {
		c := v.(Common)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.X()), float64(c.Y()))
		screen.DrawImage(c.Image(),op)
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

func (g *Game) setObjects(os []interface{}) {
	g.objects = append(g.objects, os...)
}

func (g *Game) deleteObject(o interface{}) {
	newObjects := []interface{}{}
	for _,v := range g.objects {
		if v != o {
			newObjects = append(newObjects, v)
		}
	}
	g.objects = newObjects
}

func (g *Game) isObjectAlive(o interface{}) bool {
	for _,v := range g.objects {
		if v == o {
			return true
		}
	}
	return false
}

func (g *Game) outOfScreen(a *Area) bool {
	return (a.p2.x < 0 || width <= a.p1.x || a.p2.y < 0 || height <= a.p1.y)
}

func (g *Game) insideOfScreen(a *Area) bool {
	return (a.p1.x >= 0 && width > a.p2.x && a.p1.y >= 0 && height > a.p2.y)
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
	var targets []*Character
	for _,v := range g.objects {
		switch v.(type) {
		case *Character:
			targets = append(targets, v.(*Character))
		}
	}
	return targets
}

func (g *Game) getPlayers() []*Player {
	var targets []*Player
	for _,v := range g.objects {
		switch v.(type) {
		case *Player:
			targets = append(targets, v.(*Player))
		}
	}
	return targets
}

func (g *Game) getPhase() Phase { return g.phase }

func (g *Game) phaseShift() {
	if g.phase == LIGHT_PHASE {
		g.phase = DARK_PHASE
	} else {
		g.phase = LIGHT_PHASE
	}
	players := g.getPlayers()
	for _,p := range players {
		p.setPhase(g.phase)
	}
}

func (g *Game) stageClear() {
	g.clear = true
}

func (g *Game) gameover() {
	g.clear = true
}

