package cmd

import (
	"github.com/adsellor/golang-cqrs/pkg/command"
	"github.com/adsellor/golang-cqrs/pkg/query"
)

type Facade struct {
	Commands        map[string]command.Command
	CommandHandlers map[string]command.CommandHandler
	CommandBus      *command.CommandBus
	QueryBus        *query.QueryBus
}

func NewFacade(commands []command.Command, commandHandlers []command.CommandHandler, queries []query.Query) *Facade {
	facade := &Facade{
		Commands:        make(map[string]command.Command),
		CommandHandlers: make(map[string]command.CommandHandler),
		CommandBus:      command.NewCommandBus(),
		QueryBus:        query.NewQueryBus(),
	}
	facade.addCommands(commands)
	facade.addCommandHandlers(commandHandlers)
	facade.registerCommands()
	return facade
}

func (f *Facade) registerCommands() {
	for _, command := range f.Commands {
		commandHandler := f.CommandHandlers[command.CommandId]
		f.CommandBus.RegisterHandler(command.CommandId, &commandHandler)
	}
}

func (f *Facade) addCommands(commands []command.Command) {
	for _, command := range commands {
		f.Commands[command.CommandId] = command
	}
}

func (f *Facade) addCommandHandlers(commandHandlers []command.CommandHandler) {
	for _, commandHandler := range commandHandlers {
		f.CommandHandlers[commandHandler.CommandID] = commandHandler
	}
}
