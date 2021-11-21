package parallel_shooter

import (
	"log"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const TitleLogoFilePath string = "../images/title_logo.png"
var mplusNormalFont font.Face

type Title struct {
	titleImage *ebiten.Image
	input *Input
}

func NewTitle() (*Title, error) {
	t := &Title{}
	var err error
	t.titleImage, _, err = ebitenutil.NewImageFromFile(TitleLogoFilePath)
	if err != nil {
		log.Fatal(err)
	}
	t.input = NewInput()
	return t, nil
}

func (t *Title) Update() Mode {
	if cmd, ok := t.input.getCommand(); ok {
		if cmd == KeySpace { return MODE_GAME }
	}
	return MODE_TITLE
}

func (t *Title) Draw(screen *ebiten.Image) {
	screen.DrawImage(t.titleImage, nil)
	text.Draw(screen, "Press Space key to start", mplusNormalFont, 200, 400, color.White )
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

