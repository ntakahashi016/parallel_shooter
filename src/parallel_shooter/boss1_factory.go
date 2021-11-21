package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	BOSS1HEIGHT int = 50
	BOSS1WIDTH int = 50
)

type Boss1Factory struct {
	game *Game
	imageSet ImageSet
	shotImageSet ImageSet
}

func NewBoss1Factory(g *Game) *Boss1Factory {
	ef := &Boss1Factory{}
	ef.game = g
	enemyImageSet := ImageSet{}
	enemyImageSet.light = ebiten.NewImage(BOSS1HEIGHT,BOSS1WIDTH)
	enemyImageSet.dark = ebiten.NewImage(BOSS1HEIGHT,BOSS1WIDTH)
	enemyImageSet.gray = ebiten.NewImage(BOSS1HEIGHT,BOSS1WIDTH)
	enemyImageSet.light.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	enemyImageSet.dark.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})
	enemyImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	ef.imageSet = enemyImageSet
	enemyShotImageSet := ImageSet{}
	enemyShotImageSet.light = ebiten.NewImage(5,5)
	enemyShotImageSet.dark = ebiten.NewImage(5,5)
	enemyShotImageSet.gray = ebiten.NewImage(5,5)
	enemyShotImageSet.light.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	enemyShotImageSet.dark.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})
	enemyShotImageSet.gray.Fill(color.RGBA{0x88, 0x88, 0x88, 0xff})
	ef.shotImageSet = enemyShotImageSet
	return ef
}

func (ef *Boss1Factory) NewBoss1(x,y int, p Phase) *Character {
	o := Object{game:ef.game, x:x, y:y, height:BOSS1HEIGHT, width:BOSS1WIDTH, phase: p, images: &ef.imageSet}
	ca := CharacterAttr{hp: 100, score: 0, value: 1000, shotImages: &ef.shotImageSet}
	return NewCharacter(o, ca)
}

