package query

type Query struct {
	Query func()
	Args  []string
}

func NewQuery(query func(), args ...string) *Query {
	return &Query{
		Query: query,
		Args:  args,
	}
}
