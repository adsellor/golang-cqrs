package command

type CommandBus struct {
	handlers map[string]*CommandHandler
	commands map[string]*Command
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		commands: make(map[string]*Command),
		handlers: make(map[string]*CommandHandler),
	}
}

func (cb *CommandBus) RegisterHandler(commandId string, handler *CommandHandler) {
	cb.handlers[commandId] = handler
}

func (cb *CommandBus) GetHandler(commandId string) *CommandHandler {
	return cb.handlers[commandId]
}

func (cb *CommandBus) StoreCommand(command *Command) {
	cb.commands[command.CommandId] = command
}

func (cb *CommandBus) GetCommandId(command *Command) string {
	return command.CommandId
}

func (cb *CommandBus) Execute(command *Command) {
	cb.GetHandler(command.CommandId).Execute()
}
