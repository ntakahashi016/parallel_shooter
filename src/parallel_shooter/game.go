package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{
	input    *Input
	message  string
	player   *Charactor
	playerImage *ebiten.Image
}

func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
		message: "Hello World!",
	}
	g.player = NewCharactor(50,50,10,10,true,10,10)
	g.playerImage = ebiten.NewImage(g.player.height, g.player.width)
	return g, nil
}

func (g *Game) Update() error {
	if dir, ok := g.input.Dir(); ok {
		g.message = dir.String()
		g.player.command(dir)
	}
	// if key, ok := g.input.Key(); ok {
	// 	g.player.command(key)
	// }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x00, 0x00, 0xff})
	g.player.Draw(g.playerImage)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.getx()), float64(g.player.gety()))
	screen.DrawImage(g.playerImage,op)
	ebitenutil.DebugPrintAt(screen, g.message, 0, 0)
	if _, ok := g.input.Key(); ok {
		ebitenutil.DebugPrintAt(screen, "shot!", 0, 20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

