package parallel_shooter

type StageManager struct {
	game *Game
	strategies []interface{}
}

func NewStageManager(g *Game, s []interface{}) *StageManager {
	sm := &StageManager{}
	sm.game = g
	sm.strategies = s
	return sm
}

func (sm *StageManager) Update() {
	results := []StrategyStatus{}
	pre_status := STRATEGY_INIT
	for _,s := range sm.strategies {
		strategy,_ := s.(Strategy)
		status := strategy.Status()
		if pre_status == STRATEGY_CLEAR || pre_status == STRATEGY_INIT {
			status = strategy.run()
			pre_status = status
		}
		results = append(results, status)
	}
	clear := true
	for _,r := range results {
		clear = clear && r == STRATEGY_CLEAR
	}
	if clear {
		sm.game.stageClear()
	}
}
