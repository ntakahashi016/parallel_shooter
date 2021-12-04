package parallel_shooter

type Commander interface {
	Update()
	getCommands() []Command
	getVector() Vector
}
