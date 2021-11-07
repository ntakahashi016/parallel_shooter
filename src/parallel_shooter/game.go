package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	player   *Player
	playerImage *ebiten.Image
	objects map[interface{}]*ebiten.Image
	enemy *Charactor
	enemyImage *ebiten.Image
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.objects = map[interface{}]*ebiten.Image{}
	g.player = NewPlayer(160,200,10,10,true,10,10,NewInput(),g)
	g.playerImage = ebiten.NewImage(g.player.height, g.player.width)
	g.objects[g.player] = g.playerImage
	g.enemy = NewCharactor(100,100,10,10,true,10,10)
	g.enemyImage = ebiten.NewImage(g.enemy.height, g.enemy.width)
	g.objects[g.enemy] = g.enemyImage
	return g, nil
}

func (g *Game) Update() error {
	for o, _ := range g.objects {
		c := o.(common)
		c.Update()
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
	return 320, 240
}

func (g *Game) setObject(o interface{}, i *ebiten.Image) {
	g.objects[o] = i
}
