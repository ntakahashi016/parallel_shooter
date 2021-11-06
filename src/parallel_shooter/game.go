package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	input    *Input
	message  string
	object   *Object
	objectImage *ebiten.Image
}

func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
		message: "Hello World!",
		object: NewObject(50,50),
		objectImage: ebiten.NewImage(10,10),
	}
	return g, nil
}

func (g *Game) Update() error {
	if dir, ok := g.input.Dir(); ok {
		g.message = dir.String()
		g.object.Update(dir)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	g.object.Draw(g.objectImage)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.object.getx()), float64(g.object.gety()))
	screen.DrawImage(g.objectImage,op)
	ebitenutil.DebugPrintAt(screen, g.message, 0, 0)
	if _, ok := g.input.Key(); ok {
		ebitenutil.DebugPrintAt(screen, "shot!", 0, 20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

