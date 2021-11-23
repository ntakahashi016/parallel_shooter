package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Command int

const (
	DirUp Command = iota
	DirRight
	DirDown
	DirLeft
	KeySpace
	KeyCtrl
)

type Input struct {
	keys []ebiten.Key
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update() {
	i.keys = inpututil.AppendPressedKeys(i.keys[:0])
}

func (i *Input) getCommands() []Command {
	i.Update()
	var commands []Command
	for _, k := range i.keys {
		var command Command
		switch k {
		case ebiten.KeyArrowUp:
			command = DirUp
		case ebiten.KeyArrowLeft:
			command = DirLeft
		case ebiten.KeyArrowRight:
			command = DirRight
		case ebiten.KeyArrowDown:
			command = DirDown
		case ebiten.KeySpace:
			command = KeySpace
		case ebiten.KeyControl:
			if inpututil.IsKeyJustPressed(ebiten.KeyControl) {
				command = KeyCtrl
			}
		}
		commands = append(commands, command)
	}
	return commands
}

func (i *Input) getVector() Vector {
	i.Update()
	v := NewVector(0, 0)
	for _, k := range i.keys {
		switch k {
		case ebiten.KeyArrowUp:
			v = v.Add(NewVector(0, -1))
		case ebiten.KeyArrowLeft:
			v = v.Add(NewVector(-1, 0))
		case ebiten.KeyArrowRight:
			v = v.Add(NewVector(1, 0))
		case ebiten.KeyArrowDown:
			v = v.Add(NewVector(0, 1))
		}
	}
	return v
}
