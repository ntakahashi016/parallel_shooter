package parallel_shooter

type StrategyStatus int
const (
	STRATEGY_INIT StrategyStatus = iota
	STRATEGY_RUNNING
	STRATEGY_CLEAR
)

type Strategy interface {
	run() StrategyStatus
	isAnyoneAlive() bool
	Status() StrategyStatus
}
