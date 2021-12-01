package parallel_shooter

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
		s.enemies = append(s.enemies, s.ef.NewBoss1(300,50,DARK_PHASE))
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
