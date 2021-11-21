package parallel_shooter

import (
	"time"
)

type Boss1Strategy struct {
	Strategy
	game *Game
	status StrategyStatus
	ef *Boss1Factory
	phase Phase
	enemies []*Character
}

func NewBoss1Strategy(g *Game,ef *Boss1Factory, p Phase) *Boss1Strategy {
	s := &Boss1Strategy{}
	s.game = g
	s.status = INIT
	s.ef = ef
	s.phase = p
	return s
}

func (s *Boss1Strategy) getStatus() (StrategyStatus) {
	return s.status
}

func (s *Boss1Strategy) run(ch chan bool) {
	s.status = RUNNING
	s.enemies = append(s.enemies, s.ef.NewBoss1(300,50,s.phase))
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

func (s *Boss1Strategy) isAnyoneAlive() bool {
	for _,e := range s.enemies {
		if s.game.isObjectAlive(e) {
			return true
		}
	}
	return false
}

func (s *Boss1Strategy) allOutOfScreen() bool {
	for _,e := range s.enemies {
		if !s.game.outOfScreen(e.getArea()) {
			return false
		}
	}
	return true
}
