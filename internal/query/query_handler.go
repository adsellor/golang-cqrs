package query

type QueryHandler struct {
	CommandId string
	query     func()
	Args      []string
}

func (ch *QueryHandler) Execute() {
	ch.query()
}

func NewHandler(commandId string, query func(), args ...string) *QueryHandler {
	return &QueryHandler{
		CommandId: commandId,
		query:     query,
		Args:      args,
	}
}
