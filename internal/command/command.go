package command

type Command struct {
	CommandId string
	Command   func()
	Args      []string
}

func NewCommand(commandId string, command func(), args ...string) *Command {
	return &Command{
		CommandId: commandId,
		Command:   command,
		Args:      args,
	}
}
