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
	MODE_INIT Mode = iota
	MODE_TITLE
	MODE_GAME
	MODE_RESULT
)

type Manager struct {
	mode Mode
	pre_mode Mode
	next_mode Mode
	current interface{}
}

func NewManager() (*Manager, error) {
	m := &Manager{}
	m.mode = MODE_INIT
	m.pre_mode = MODE_INIT
	m.next_mode = MODE_INIT
	m.current, _ = NewTitle()
	return m, nil
}

func (m *Manager) Update() error {
	var score int
	switch m.next_mode {
	case MODE_INIT:
		m.next_mode = MODE_TITLE
	case MODE_TITLE:
		if m.pre_mode != m.next_mode {
			m.pre_mode = m.mode
			m.mode = m.next_mode
			m.current,_ = NewTitle()
		}
		m.next_mode = m.current.(updatable).Update()
	case MODE_GAME:
		if m.pre_mode != m.next_mode {
			m.pre_mode = m.mode
			m.mode = m.next_mode
			m.current,_ = NewGame()
		}
		m.next_mode = m.current.(updatable).Update()
	case MODE_RESULT:
		if m.pre_mode != m.next_mode {
			m.pre_mode = m.mode
			m.mode = m.next_mode
			if m.pre_mode == MODE_GAME {
				game := m.current.(*Game)
				score = game.getScore()
			}
			m.current,_ = NewResult()
			result := m.current.(*Result)
			result.setScore(score)
		}
		m.next_mode = m.current.(updatable).Update()
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
