package command

type Command struct {
	Command func()
	Args    []string
}

func NewCommand(command func(), args ...string) *Command {
	return &Command{
		Command: command,
		Args:    args,
	}
}
