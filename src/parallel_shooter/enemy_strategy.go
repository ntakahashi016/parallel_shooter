package parallel_shooter

import (
	"time"
)

type Enemy1Strategy struct {
	Strategy
	game *Game
	status StrategyStatus
	ef *Enemy1Factory
	phase Phase
	enemies []*Character
}

func NewEnemy1Strategy(g *Game,ef *Enemy1Factory, p Phase) *Enemy1Strategy {
	s := &Enemy1Strategy{}
	s.game = g
	s.status = INIT
	s.ef = ef
	s.phase = p
	return s
}

func (s *Enemy1Strategy) getStatus() (StrategyStatus) {
	return s.status
}

func (s *Enemy1Strategy) run(ch chan bool) {
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

func (s *Enemy1Strategy) isAnyoneAlive() bool {
	for _,e := range s.enemies {
		if s.game.isObjectAlive(e) {
			return true
		}
	}
	return false
}

func (s *Enemy1Strategy) allOutOfScreen() bool {
	for _,e := range s.enemies {
		if !s.game.outOfScreen(e.getArea()) {
			return false
		}
	}
	return true
}
