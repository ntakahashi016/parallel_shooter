package parallel_shooter

import (
	"time"
)

type StrategyStatus int
const (
	INIT StrategyStatus = iota
	RUNNING
	DONE
	CLEAR
)

type Strategy struct {
	game *Game
	status StrategyStatus
	ef *Enemy1Factory
	phase Phase
	enemies []*Character
}

func NewStrategy(g *Game,ef *Enemy1Factory, p Phase) *Strategy {
	s := &Strategy{}
	s.game = g
	s.status = INIT
	s.ef = ef
	s.phase = p
	return s
}

func (s *Strategy) getStatus() (StrategyStatus) {
	return s.status
}

func (s *Strategy) run(ch chan bool) {
	s.status = RUNNING
	s.enemies = append(s.enemies, s.ef.NewEnemy1(200,100,s.phase))
	s.enemies = append(s.enemies, s.ef.NewEnemy1(150,100,s.phase))
	s.enemies = append(s.enemies, s.ef.NewEnemy1(100,100,s.phase))
	for _,e := range s.enemies {
		s.game.setObject(e)
		time.Sleep(time.Second * 3)
	}
	for s.isAnyoneAlive() {
		if s.allOutOfScreen() {
			s.status = DONE
			ch <- true
			return
		}
	time.Sleep(time.Second)
	}
	s .status = CLEAR
	ch <- true
	return
}

func (s *Strategy) isAnyoneAlive() bool {
	for _,e := range s.enemies {
		if s.game.isObjectAlive(e) {
			return true
		}
	}
	return false
}

func (s *Strategy) allOutOfScreen() bool {
	for _,e := range s.enemies {
		if !s.game.outOfScreen(e.getArea()) {
			return false
		}
	}
	return true
}
