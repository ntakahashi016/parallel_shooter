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

func (sm *StageManager) run() {
	results := []bool{}
	ch := make(chan bool)
	for _,s := range sm.strategies {
		strategy,_ := s.(Strategy)
		go strategy.run(ch)
		res := <-ch
		results = append(results, res)
	}
	result := true
	for _,r := range results {
		if !r {
			result = false
		}
	}
	if result {
		sm.game.stageClear()
	}
}
