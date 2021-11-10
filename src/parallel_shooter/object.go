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
	phase bool
	image_l *ebiten.Image
	image_d *ebiten.Image
}

type common interface {
	Update() error
	Draw(img *ebiten.Image) error
	getx() int
	gety() int
	getArea() *Area
}

