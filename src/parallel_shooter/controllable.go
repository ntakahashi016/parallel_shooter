package parallel_shooter

type Controllable interface {
	command(i Input) error
}
