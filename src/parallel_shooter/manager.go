package parallel_shooter

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type drawable interface {
	Draw(screen *ebiten.Image)
}

type updatable interface {
	Update() Mode
}

type Mode int

const (
	MODE_TITLE Mode = iota
	MODE_GAME
	MODE_RESULT
)

type Manager struct {
	mode Mode
	pre_mode Mode
	current interface{}
}

func NewManager() (*Manager, error) {
	m := &Manager{}
	m.mode = MODE_TITLE
	m.current, _ = NewTitle()
	return m, nil
}

func (m *Manager) Update() error {
	next := m.current.(updatable).Update()
	switch m.mode {
	case MODE_TITLE:
		switch next {
		case MODE_GAME:
			m.current,_ = NewGame()
		}
	case MODE_GAME:
		switch next {
		case MODE_RESULT:
			m.current,_ = NewResult()
		}
	case MODE_RESULT:
		switch next {
		case MODE_TITLE:
			m.current,_ = NewTitle()
		}
	}
	m.mode = next
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	game := m.current.(drawable)
	game.Draw(screen)
}

func (m *Manager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var w int = width
	var h int = height
	return w, h
}
