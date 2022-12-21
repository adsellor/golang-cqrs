package query

import "errors"

type QueryBus struct {
	handlers map[string]*QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]*QueryHandler),
	}
}

func (cb *QueryBus) GetHandler(queryId string) *QueryHandler {
	return cb.handlers[queryId]
}

func (cb *QueryBus) AddHandler(queryId string, handler *QueryHandler) {
	_handler := cb.GetHandler(queryId)

	if _handler != nil {
		panic(errors.New("cannot override existing command handler"))
	}
	cb.handlers[queryId] = handler
}

func (cb *QueryBus) Execute(queryId string) {
	handler := cb.GetHandler(queryId)
	if handler != nil {
		handler.Execute()
	} else {
		panic(errors.New("command not found"))
	}
}
