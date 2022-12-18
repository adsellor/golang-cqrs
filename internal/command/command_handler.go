package command

type CommandHandler struct {
	Command func()
	Args    []string
}

func (ch *CommandHandler) Execute() {
	ch.Command()
}

func NewHandler(command func(), args ...string) *CommandHandler {
	return &CommandHandler{
		Command: command,
		Args:    args,
	}
}
