package command

import "github.com/google/uuid"

type Command struct {
	CommandId string
}

func NewCommand() *Command {
	commandId := uuid.New().String()
	return &Command{
		CommandId: commandId,
	}
}
