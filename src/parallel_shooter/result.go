package parallel_shooter

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// var mplusNormalFont font.Face

type Result struct {
	input *Input
}

func NewResult() (*Result, error) {
	r := &Result{}
	r.input = NewInput()
	return r, nil
}

func (r *Result) Update() Mode {
	if cmd, ok := r.input.getCommand(); ok {
		if cmd == KeySpace { return TITLE }
	}
	return RESULT
}

func (r *Result) Draw(screen *ebiten.Image) {
	text.Draw(screen, "Congratulations!", mplusNormalFont, 100, 300, color.White )
	text.Draw(screen, "Press Space key to return to title", mplusNormalFont, 200, 400, color.White )
}

func init () {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

