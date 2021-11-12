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
	TITLE Mode = iota
	GAME
)

type Manager struct {
	mode Mode
	pre_mode Mode
	current interface{}
}

func NewManager() (*Manager, error) {
	m := &Manager{}
	m.mode = TITLE
	m.pre_mode = m.mode
	m.current, _ = NewTitle()
	return m, nil
}

func (m *Manager) Update() error {
	switch m.mode {
	case TITLE:
		if m.pre_mode != m.mode {
			m.pre_mode = m.mode
			m.current,_ = NewTitle()
		}
		m.mode = m.current.(updatable).Update()
	case GAME:
		if m.pre_mode != m.mode {
			m.pre_mode = m.mode
			m.current,_ = NewGame()
		}
		m.mode = m.current.(updatable).Update()
	}
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
