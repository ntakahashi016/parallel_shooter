package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy1Factory struct {
	game *Game
	imageSet ImageSet
	shotImageSet ImageSet
}

func NewEnemy1Factory(g *Game) (*Enemy1Factory, error) {
	ef := &Enemy1Factory{}
	ef.game = g
	enemyImageSet := ImageSet{}
	enemyImageSet.light = ebiten.NewImage(10,10)
	enemyImageSet.dark = ebiten.NewImage(10,10)
	enemyImageSet.gray = ebiten.NewImage(10,10)
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
	return ef, nil
}

func (ef *Enemy1Factory) NewEnemy1d() *Character {
	o := Object{game:ef.game, x:100, y:100, height:10, width:10, phase: Dark, images: &ef.imageSet}
	ca := CharacterAttr{hp: 10, score: 0, value: 100, shotImages: &ef.shotImageSet}
	return NewCharacter(o, ca)
}
func (ef *Enemy1Factory) NewEnemy1l() *Character {
	o := Object{game:ef.game, x:200, y:100, height:10, width:10, phase: Light, images: &ef.imageSet}
	ca := CharacterAttr{hp: 10, score: 0, value: 100, shotImages: &ef.shotImageSet}
	return NewCharacter(o, ca)
}
