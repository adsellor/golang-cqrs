package main

import (
	"github.com/adsellor/golang-cqrs/internal/cmd"
	"github.com/adsellor/golang-cqrs/internal/command"
	"github.com/adsellor/golang-cqrs/internal/query"
)

var books []Book

type Book struct {
	ID     string
	Title  string
	Author string
}

// create new facade with command and command hanler that add book to books and returns command, command handler, facade,
func NewBookFacade() (*cmd.Facade, *command.Command) {
	books = make([]Book, 0)
	// create command
	addBook := func() {
		books = append(books, Book{
			ID:     "1",
			Title:  "The Lord of the Rings",
			Author: "J.R.R. Tolkien",
		})

	}
	// create book command
	bookCommand := command.NewCommand()
	commandHandler := command.NewHandler(bookCommand.CommandId, addBook)
	facade := cmd.NewFacade([]command.Command{*bookCommand}, []command.CommandHandler{*commandHandler}, []query.Query{})

	return facade, bookCommand
}

func main() {
	book, addBook := NewBookFacade()
	book.CommandBus.Execute(addBook)
	println(books[0].Title)

}
