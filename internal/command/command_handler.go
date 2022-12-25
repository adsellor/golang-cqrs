package command

type CommandHandler struct {
	CommandID string
	Command   func()
	Args      []string
}

func (ch *CommandHandler) Execute() {
	ch.Command()
}

func NewHandler(commandId string, command func(), args ...string) *CommandHandler {
	return &CommandHandler{
		CommandID: commandId,
		Command:   command,
		Args:      args,
	}
}
