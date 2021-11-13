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
}

func NewGame() (*Game, error) {
	g := &Game{}
	g.phase = Light
	g.clear = false
	g.objects = []interface{}{}
	var o Object
	var ca CharacterAttr
	playerImageSet := &ImageSet{}
	playerImageSet.light = ebiten.NewImage(10,10)
	playerImageSet.dark = ebiten.NewImage(10,10)
	playerImageSet.gray = ebiten.NewImage(10,10)
	playerImageSet.light.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	playerImageSet.dark.Fill(color.RGBA{0x00, 0xff, 0xff, 0xff})
	playerImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	playerShotImageSet := &ImageSet{}
	playerShotImageSet.light = ebiten.NewImage(5,5)
	playerShotImageSet.dark = ebiten.NewImage(5,5)
	playerShotImageSet.gray = ebiten.NewImage(5,5)
	playerShotImageSet.light.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	playerShotImageSet.dark.Fill(color.RGBA{0x00, 0xff, 0xff, 0xff})
	playerShotImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	o = Object{game:g, x: (int)(width/2), y: height-20, height: 10, width: 10, phase: g.phase,images: playerImageSet}
	ca = CharacterAttr{hp: 10,score: 0, value: 0, shotImages: playerShotImageSet}
	g.objects = append(g.objects, NewPlayer(o, ca, NewInput()))
	enemyImageSet := &ImageSet{}
	enemyImageSet.light = ebiten.NewImage(10,10)
	enemyImageSet.dark = ebiten.NewImage(10,10)
	enemyImageSet.gray = ebiten.NewImage(10,10)
	enemyImageSet.light.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	enemyImageSet.dark.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})
	enemyImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	enemyShotImageSet := &ImageSet{}
	enemyShotImageSet.light = ebiten.NewImage(5,5)
	enemyShotImageSet.dark = ebiten.NewImage(5,5)
	enemyShotImageSet.gray = ebiten.NewImage(5,5)
	enemyShotImageSet.light.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	enemyShotImageSet.dark.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})
	enemyShotImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	o = Object{game:g, x:100, y:100, height:10, width:10, phase: Dark, images: enemyImageSet}
	ca = CharacterAttr{hp: 10, score: 0, value: 100, shotImages: enemyShotImageSet}
	g.objects = append(g.objects, NewCharacter(o, ca))
	o = Object{game:g, x:200, y:100, height:10, width:10, phase: Light, images: enemyImageSet}
	ca = CharacterAttr{hp: 10, score: 0, value: 100, shotImages: enemyShotImageSet}
	g.objects = append(g.objects, NewCharacter(o, ca))
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

func (g *Game) outOfScreen(x,y int) bool {
	if x < 0 || width <= x { return true }
	if y < 0 || height <= y { return true }
	return false
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
