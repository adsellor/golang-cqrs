package cmd

import (
	"github.com/adsellor/golang-cqrs/internal/command"
	"github.com/adsellor/golang-cqrs/internal/query"
)

type Facade struct {
	// Commands is map with commandId as key and command as value
	Commands        map[string]command.Command
	CommandHandlers map[string]command.CommandHandler
	CommandBus      *command.CommandBus
	QueryBus        *query.QueryBus
}

func NewFacade(commands map[string]command.Command, commandHandlers map[string]command.CommandHandler, queries []query.Query) *Facade {
	return &Facade{
		Commands:        commands,
		CommandHandlers: commandHandlers,
		CommandBus:      command.NewCommandBus(),
		QueryBus:        query.NewQueryBus(),
	}
}

// register commandHandler with the same command id as the command
func (f *Facade) RegisterCommands() {
	for commandId := range f.Commands {
		commandHandler := f.CommandHandlers[commandId]
		f.CommandBus.AddHandler(commandId, &commandHandler)
	}
}
