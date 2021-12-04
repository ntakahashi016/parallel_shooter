package parallel_shooter

type EnemyCommander struct {
	Commander
	commandStep int
	commands []Command
	motionStep int
	motionIndex int
	motions []*Motion
}

func NewEnemyCommander() *EnemyCommander {
	return &EnemyCommander{}
}

func (ec *EnemyCommander) getCommands() []Command {
	return ec.commands
}

func (ec *EnemyCommander) getVector() Vector {
	return ec.motions[ec.motionIndex].Vector()
}

func (ec *EnemyCommander) addCommands(commands []Command) {
	// ec.commands = append(ec.commands, commands)
}

func (ec *EnemyCommander) addMotion(m *Motion) {
	ec.motions = append(ec.motions, m)
}

func (ec *EnemyCommander) Update() {
	ec.motionStep++
	if ec.motions[ec.motionIndex].Times() < ec.motionStep {
		ec.motionStep = 0
		ec.motionIndex++
	}
}
