package query

type Query struct {
	QueryId string
	Query   func()
	Args    []string
}

func NewQuery(queryId string, query func(), args ...string) *Query {
	return &Query{
		QueryId: queryId,
		Query:   query,
		Args:    args,
	}
}
