package parallel_shooter

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	x int
	y int
}

func NewObject(x int,y int) *Object {
	object := Object{}
	object.x = x
	object.y = y
	return &object
}

func (o *Object) Update(dir Dir) error {
	switch dir {
	case DirUp:
		o.y = o.y - 1
	case DirLeft:
		o.x = o.x - 1
	case DirRight:
		o.x = o.x + 1
	case DirDown:
		o.y = o.y + 1
	}
	return nil
}

func (o *Object) Draw(objectImage *ebiten.Image) error {
	objectImage.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	return nil
}

func (o *Object) getx() int { return o.x }
func (o *Object) gety() int { return o.y }
