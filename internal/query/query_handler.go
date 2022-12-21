package query

type QueryHandler struct {
	query func()
	Args  []string
}

func (ch *QueryHandler) Execute() {
	ch.query()
}

func NewHandler(query func(), args ...string) *QueryHandler {
	return &QueryHandler{
		query: query,
		Args:  args,
	}
}
