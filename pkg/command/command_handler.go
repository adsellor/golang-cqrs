package command

type CommandHandler struct {
	CommandID string
	Command   func()
	Args      []any
}

func NewHandler(commandId string, command func()) *CommandHandler {
	return &CommandHandler{
		CommandID: commandId,
		Command:   command,
	}
}

func (ch *CommandHandler) Execute() {
	ch.Command()
}
