package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Key int

const (
	KeySpace Key = iota
)

type Dir int

const (
	DirUp Dir = iota
	DirRight
	DirDown
	DirLeft
)

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type touchState int

const (
	touchStateNone touchState = iota
	touchStatePressing
	touchStateSettled
	touchStateInvalid
)

func (d Dir) String() string {
	switch d {
	case DirUp:
		return "Up"
	case DirRight:
		return "Right"
	case DirDown:
		return "Down"
	case DirLeft:
		return "Left"
	}
	panic("not reach")
}

func (d Dir) Vector() (x, y int) {
	switch d {
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
	mouseState    mouseState
	mouseInitPosX int
	mouseInitPosY int
	mouseDir      Dir

	touches       []ebiten.TouchID
	touchState    touchState
	touchID       ebiten.TouchID
	touchInitPosX int
	touchInitPosY int
	touchLastPosX int
	touchLastPosY int
	touchDir      Dir
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

func vecToDir(dx, dy int) (Dir, bool) {
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
	switch i.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			i.mouseInitPosX = x
			i.mouseInitPosY = y
			i.mouseState = mouseStatePressing
		}
	case mouseStatePressing:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			dx := x - i.mouseInitPosX
			dy := y - i.mouseInitPosY
			d, ok := vecToDir(dx, dy)
			if !ok {
				i.mouseState = mouseStateNone
				break
			}
			i.mouseDir = d
			i.mouseState = mouseStateSettled
		}
	case mouseStateSettled:
		i.mouseState = mouseStateNone
	}

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
		delay    = 10
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

func (i *Input) Dir() (Dir, bool) {
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
	if i.mouseState == mouseStateSettled {
		return i.mouseDir, true
	}
	if i.touchState == touchStateSettled {
		return i.touchDir, true
	}
	return 0, false
}

func (i *Input) Key() (Key, bool) {
	if repeatingKeyPressed(ebiten.KeySpace) {
		return KeySpace, true
	}
	return 0, false
}
