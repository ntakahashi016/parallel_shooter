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

type Common interface {
	run()

	Draw(img *ebiten.Image) error
	X() int
	Y() int
	Area() *Area
	Phase() Phase
	Image() *ebiten.Image
}

