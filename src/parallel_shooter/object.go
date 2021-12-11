package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	game *Game
	point *Point
	height int
	width int
	phase Phase
	images *ImageSet
}

type Common interface {
	Update()

	Draw(img *ebiten.Image) error
	X() int
	Y() int
	Area() *Area
	Phase() Phase
	Image() *ebiten.Image
	Center() *Point
}

