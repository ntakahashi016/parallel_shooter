package parallel_shooter

import (
	"math"
)

type Boss1Strategy struct {
	Strategy
	game *Game
	status StrategyStatus
	ef *Boss1Factory
	enemies []*Character
}

func NewBoss1Strategy(g *Game, ef *Boss1Factory) *Boss1Strategy {
	s := &Boss1Strategy{}
	s.game = g
	s.status = STRATEGY_INIT
	s.ef = ef
	return s
}

func (s *Boss1Strategy) run() StrategyStatus {
	switch s.status {
	case STRATEGY_INIT:
		ec := NewEnemyCommander()
		ec.addMotion(NewMotion(NewVector(0,0),math.MaxInt64))
		s.enemies = append(s.enemies, s.ef.NewBoss1(300,50,DARK_PHASE,ec))
		for _,e := range s.enemies {
			s.game.setObject(e)
		}
		s.status = STRATEGY_RUNNING
	case STRATEGY_RUNNING:
		if !s.isAnyoneAlive() {
			s.status = STRATEGY_CLEAR
		}
	}
	return s.status
}

func (s *Boss1Strategy) isAnyoneAlive() bool {
	result := false
	for _,e := range s.enemies {
		result = result || s.game.isObjectAlive(e)
	}
	return result
}

func (s *Boss1Strategy) Status() StrategyStatus { return s.status }
