package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	game *Game
	x int
	y int
	height int
	width int
	phase Phase
	images *ImageSet
}

type common interface {
	Update() error
	Draw(img *ebiten.Image) error
	getx() int
	gety() int
	getArea() *Area
	getPhase() Phase
	getImage() *ebiten.Image
}

