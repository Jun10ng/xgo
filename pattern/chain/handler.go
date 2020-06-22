package chain

// handler represents a processor or filter etc.
type Handler interface {
	Handle(content interface{}) error
}
