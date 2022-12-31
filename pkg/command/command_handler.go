package command

type CommandHandler struct {
	CommandID string
	Command   func(args []interface{})
	Args      []any
}

func NewHandler(commandId string, command func(args []interface{})) *CommandHandler {
	return &CommandHandler{
		CommandID: commandId,
		Command:   command,
	}
}

func (ch *CommandHandler) Execute(args ...any) {
	ch.Command(args)
}
