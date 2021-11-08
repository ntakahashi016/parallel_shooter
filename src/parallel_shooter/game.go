package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const width = 320
const height = 240

type Game struct{
	objects map[interface{}]*ebiten.Image
	e *Charactor
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.objects = map[interface{}]*ebiten.Image{}
	p := NewPlayer(160,200,10,10,true,10,10,NewInput(),g)
	pImg := ebiten.NewImage(p.height, p.width)
	g.objects[p] = pImg
	e := NewCharactor(100,100,10,10,true,10,10)
	eImg := ebiten.NewImage(e.height, e.width)
	g.objects[e] = eImg
	g.e = e
	return g, nil
}

func (g *Game) Update() error {
	for o, _ := range g.objects {
		c := o.(common)
		c.Update()
		if g.outOfScreen(c.getx(), c.gety()) {
			c = nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})

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

func (g *Game) outOfScreen(x,y int) bool {
	if x < 0 || width <= x { return true }
	if y < 0 || height <= y { return true }
	return false
}
