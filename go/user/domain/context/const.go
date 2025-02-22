package context

type contextKey int

const (
	UserRepository contextKey = iota
	Validator
)
