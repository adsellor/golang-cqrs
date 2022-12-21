package main

import (
	"context"
	"time"

	"github.com/adsellor/golang-cqrs/internal/command"
	"github.com/adsellor/golang-cqrs/internal/event"
)

type contextKey string

func printLine() {
	println("hello there")
}

func printEvent() {
	println("hello event")
}

func main() {
	var commandId contextKey = "commandId"
	ctx := context.WithValue(context.Background(), commandId, "1")

	localCommandId := ctx.Value(commandId).(string)
	commandBus := command.NewCommandBus()
	printCommand := command.NewCommand(printLine)
	commandHandler := command.NewHandler(printCommand.Command)

	commandBus.AddHandler(localCommandId, commandHandler)
	commandBus.Execute(localCommandId)

	// create new event bus
	eventBus := event.NewEventBus()
	// create new event
	myEvent := event.NewEvent(localCommandId, printEvent)
	// create new event handler
	event.NewEventHandler(myEvent, eventBus)
	// subscribe to event
	// publish event
	eventBus.Publish(myEvent)
	// wait for the event to be handled before terminating the program
	time.Sleep(1 * time.Second)
}
