package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const width = 640
const height = 480

type Game struct{
	objects map[interface{}]*ebiten.Image
	phase bool
	clear bool
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.phase = true
	g.clear = false
	g.objects = map[interface{}]*ebiten.Image{}
	p := NewPlayer(160,200,10,10,true,10,10, g, NewInput())
	pImg := ebiten.NewImage(p.height, p.width)
	g.objects[p] = pImg
	o := &Object{game:g, x:100, y:100, height:10, width:10, phase: g.phase, image_l: ebiten.NewImage(10,10), image_d: ebiten.NewImage(10,10)}
	e := NewCharacter(o, 100, 100)
	eImg := ebiten.NewImage(e.height, e.width)
	g.objects[e] = eImg
	return g, nil
}

func (g *Game) Update() Mode {
	for o, _ := range g.objects {
		c := o.(common)
		c.Update()
		if g.outOfScreen(c.getx(), c.gety()) {
			g.deleteObject(c)
		}
	}
	if g.clear { return RESULT }
	return GAME
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.phase {
		screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	} else {
		screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	}

	for o, i := range g.objects {
		c := o.(common)
		c.Draw(i)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.getx()), float64(c.gety()))
		screen.DrawImage(i,op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var w int = width
	var h int = height
	return w, h
}

func (g *Game) setObject(o interface{}, i *ebiten.Image) {
	g.objects[o] = i
}

func (g *Game) deleteObject(o interface{}) {
	delete(g.objects, o)
}

func (g *Game) outOfScreen(x,y int) bool {
	if x < 0 || width <= x { return true }
	if y < 0 || height <= y { return true }
	return false
}

func (g *Game) getEnemy() *Character {
	for k,_ := range g.objects {
		switch k.(type) {
		case *Character:
			return k.(*Character)
		}
	}
	return nil
}

func (g *Game) phaseShift() {
	g.phase = !g.phase
}

func (g *Game) checkGameClear() {
	for k,_ := range g.objects {
		switch k.(type) {
		case *Character:
			g.clear = false
		}
	}
	g.clear = true
}
