package parallel_shooter

type StrategyStatus int
const (
	INIT StrategyStatus = iota
	RUNNING
	DONE
	CLEAR
)

type Strategy interface {
	getStatus() (StrategyStatus)
	run(ch chan bool)
	isAnyoneAlive() bool
	allOutOfScreen() bool
}
