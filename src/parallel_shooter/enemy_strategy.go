package parallel_shooter

type Enemy1Strategy struct {
	Strategy
	game *Game
	status StrategyStatus
	ef *Enemy1Factory
	enemies []*Character
}

func NewEnemy1Strategy(g *Game, ef *Enemy1Factory) *Enemy1Strategy {
	s := &Enemy1Strategy{}
	s.game = g
	s.status = STRATEGY_INIT
	s.ef = ef
	return s
}

func (s *Enemy1Strategy) run() StrategyStatus {
	switch s.status {
	case STRATEGY_INIT:
		s.enemies = append(s.enemies, s.ef.NewObject(200,100,DARK_PHASE))
		s.enemies = append(s.enemies, s.ef.NewObject(150,100,DARK_PHASE))
		s.enemies = append(s.enemies, s.ef.NewObject(100,100,DARK_PHASE))
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

func (s *Enemy1Strategy) isAnyoneAlive() bool {
	result := false
	for _,e := range s.enemies {
		result = result || s.game.isObjectAlive(e)
	}
	return result
}

func (s *Enemy1Strategy) Status() StrategyStatus { return s.status }
