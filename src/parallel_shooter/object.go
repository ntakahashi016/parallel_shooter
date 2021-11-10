package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	x int
	y int
	height int
	width int
	phase bool
}

type common interface {
	Update() error
	Draw(img *ebiten.Image) error
	getx() int
	gety() int
	getArea() *Area
	hit(int)
}

