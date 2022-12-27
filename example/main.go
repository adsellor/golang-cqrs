package main

import (
	"github.com/adsellor/golang-cqrs/internal/cmd"
	"github.com/adsellor/golang-cqrs/internal/command"
	"github.com/adsellor/golang-cqrs/internal/query"
)

type contextKey string

func printLine() {
	println("hello there from myCommand")
}

func printLine2() {
	println("hello there from myCommand 2")
}

func printEvent() {
	println("hello event")
}

func main() {

	// create new facade
	localCommand := *command.NewCommand()

	// create new command
	myCommand := *command.NewCommand()

	// create command handler
	localCommandHandler := *command.NewHandler(localCommand.CommandId, printLine)
	// my command handler
	myCommandHandler := *command.NewHandler(myCommand.CommandId, printLine2)

	// create new facade
	facade := cmd.NewFacade([]command.Command{localCommand, myCommand}, []command.CommandHandler{myCommandHandler, localCommandHandler}, []query.Query{})
	facade.CommandBus.Execute(&localCommand)
	facade.CommandBus.Execute(&localCommand)

}
