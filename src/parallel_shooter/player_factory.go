package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerFactory struct {
	Factory
	game *Game
	imageSet ImageSet
	shotImageSet ImageSet
}

func NewPlayerFactory(g *Game) *PlayerFactory {
	pf := &PlayerFactory{}
	pf.game = g
	playerImageSet := ImageSet{}
	playerImageSet.light = ebiten.NewImage(10,10)
	playerImageSet.dark = ebiten.NewImage(10,10)
	playerImageSet.gray = ebiten.NewImage(10,10)
	playerImageSet.light.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	playerImageSet.dark.Fill(color.RGBA{0x00, 0xff, 0xff, 0xff})
	playerImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	pf.imageSet = playerImageSet
	playerShotImageSet := ImageSet{}
	playerShotImageSet.light = ebiten.NewImage(5,5)
	playerShotImageSet.dark = ebiten.NewImage(5,5)
	playerShotImageSet.gray = ebiten.NewImage(5,5)
	playerShotImageSet.light.Fill(color.RGBA{0x00, 0x00, 0xff, 0xff})
	playerShotImageSet.dark.Fill(color.RGBA{0x00, 0xff, 0xff, 0xff})
	playerShotImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	pf.shotImageSet = playerShotImageSet
	return pf
}

func (pf *PlayerFactory) NewObject() *Player {
	o := Object{game:pf.game, point: NewPoint(float64(width)/2, float64(height-20)), height: 10, width: 10, phase: pf.game.phase,images: &pf.imageSet}
	ca := CharacterAttr{hp: 10,score: 0, value: 0, shotImages: &pf.shotImageSet}
	return NewPlayer(o, ca, NewInput())
}

