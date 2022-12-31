package command

import "github.com/google/uuid"

type Command struct {
	CommandId string
	Args      []interface{}
}

func NewCommand(args ...interface{}) *Command {
	commandId := uuid.New().String()
	return &Command{
		CommandId: commandId,
		Args:      args,
	}
}
