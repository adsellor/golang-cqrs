package command

func NewCommandMessage(command string, args ...string) CommandMessage {
	return CommandMessage{
		Command: command,
		Args:    args,
	}
}

type CommandMessage struct {
	Command string
	Args    []string
}
