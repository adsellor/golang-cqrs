package command

import "errors"

type CommandBus struct {
	handlers map[string]*CommandHandler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]*CommandHandler),
	}
}

func (cb *CommandBus) GetHandler(commandId string) *CommandHandler {
	return cb.handlers[commandId]
}

func (cb *CommandBus) AddHandler(commandId string, handler *CommandHandler) {
	_handler := cb.GetHandler(commandId)

	if _handler != nil {
		panic(errors.New("cannot override existing command handler"))
	}
	cb.handlers[commandId] = handler
}

func (cb *CommandBus) Execute(commandId string) {
	handler := cb.GetHandler(commandId)
	if handler != nil {
		handler.Execute()
	} else {
		panic(errors.New("command not found"))
	}
}
