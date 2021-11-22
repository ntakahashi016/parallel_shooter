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

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
)

func (c Command) Vector() (x, y int) {
	switch c {
	case DirUp:
		return 0, -1
	case DirRight:
		return 1, 0
	case DirDown:
		return 0, 1
	case DirLeft:
		return -1, 0
	}
	panic("not reach")
}

type Input struct {
	touches       []ebiten.TouchID
	touchState    touchState
	touchID       ebiten.TouchID
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      Command
}

func NewInput() *Input {
	return &Input{}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vecToDir(dx, dy int) (Command, bool) {
	if abs(dx) < 4 && abs(dy) < 4 {
		return 0, false
	}
	if abs(dx) < abs(dy) {
		if dy < 0 {
			return DirUp, true
		}
		return DirDown, true
	}
	if dx < 0 {
		return DirLeft, true
	}
	return DirRight, true
}

func (i *Input) Update() {
	i.touches = ebiten.AppendTouchIDs(i.touches[:0])
	switch i.touchState {
	case touchStateNone:
		if len(i.touches) == 1 {
			i.touchID = i.touches[0]
			x, y := ebiten.TouchPosition(i.touches[0])
			i.touchInitPosX = x
			i.touchInitPosY = y
			i.touchLastPosX = x
			i.touchLastPosY = y
			i.touchState = touchStatePressing
		}
	case touchStatePressing:
		if len(i.touches) >= 2 {
			break
		}
		if len(i.touches) == 1 {
			if i.touches[0] != i.touchID {
				i.touchState = touchStateInvalid
			} else {
				x, y := ebiten.TouchPosition(i.touches[0])
				i.touchLastPosX = x
				i.touchLastPosY = y
			}
			break
		}
		if len(i.touches) == 0 {
			dx := i.touchLastPosX - i.touchInitPosX
			dy := i.touchLastPosY - i.touchInitPosY
			d, ok := vecToDir(dx, dy)
			if !ok {
				i.touchState = touchStateNone
				break
			}
			i.touchDir = d
			i.touchState = touchStateSettled
		}
	case touchStateSettled:
		i.touchState = touchStateNone
	case touchStateInvalid:
		if len(i.touches) == 0 {
			i.touchState = touchStateNone
		}
	}
}

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 1
		interval = 1
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func (i *Input) getCommand() (Command, bool) {
	if repeatingKeyPressed(ebiten.KeyArrowUp) {
		return DirUp, true
	}
	if repeatingKeyPressed(ebiten.KeyArrowLeft) {
		return DirLeft, true
	}
	if repeatingKeyPressed(ebiten.KeyArrowRight) {
		return DirRight, true
	}
	if repeatingKeyPressed(ebiten.KeyArrowDown) {
		return DirDown, true
	}
	if repeatingKeyPressed(ebiten.KeySpace) {
		return KeySpace, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyControl) {
		return KeyCtrl, true
	}
	return 0, false
}

