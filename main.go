package main

import (
	"context"
	"golang/cqrs/pkg/command"
)

// declare event bus var that will be used in the whole package

// create new command
type contextKey string

func main() {
	var commandId contextKey = "commandId"
	ctx := context.WithValue(context.Background(), commandId, "1")

	localCommandId := ctx.Value(commandId).(string)
	commandBus := command.NewCommandBus()
	printCommand := command.NewCommand(func() {
		println("hello there")
	})
	commandHandler := command.NewHandler(printCommand.Command)

	commandBus.AddHandler(localCommandId, commandHandler)
	commandBus.Execute(localCommandId)
}
