package main

import (
	"context"

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
	var commandId contextKey = "commandId"
	ctx := context.WithValue(context.Background(), commandId, "1")
	ctx2 := context.WithValue(context.Background(), commandId, "2")
	// create new facade

	localCommandId := ctx.Value(commandId).(string)
	localCommandId2 := ctx2.Value(commandId).(string)
	// create new facade
	localCommand := *command.NewCommand(localCommandId, printLine)
	facade := cmd.NewFacade(map[string]command.Command{
		localCommandId:  localCommand,
		localCommandId2: *command.NewCommand(localCommandId2, printLine2),
	}, map[string]command.CommandHandler{
		localCommandId:  *command.NewHandler(localCommandId, printLine),
		localCommandId2: *command.NewHandler(localCommandId2, printLine2),
	}, []query.Query{
		*query.NewQuery("0", printLine),
	})

	facade.CommandBus.Execute(localCommandId2)
	// create new event bus
	// eventBus := event.NewEventBus()
	// // create new event
	// myEvent := event.NewEvent(localCommandId, printEvent)
	// // create new event handler
	// event.NewEventHandler(myEvent, eventBus)
	// // subscribe to event
	// // publish event
	// eventBus.Publish(myEvent)
	// // wait for the event to be handled before terminating the program
	// time.Sleep(1 * time.Second)
}
