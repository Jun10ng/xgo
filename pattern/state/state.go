package state

type State interface {
	Handle(*Context)
}

// Context delegate the request to its State member.
type Context struct {
	name  string
	state State
}

func (c *Context) Handle() {
	c.state.Handle(c)
}

func (c *Context) SetState(state State) {
	c.state = state
}

func NewContext(name string) *Context {
	return &Context{
		name: name,
	}
}

// StateManager contains a state's name - state map, which is called when runtime
// flyweight pattern
type StateManager struct {
	states map[string]State
}

func (s *StateManager) Add(name string, state State) {
	if _, exist := s.states[name]; exist {
		return
	} else {
		s.states[name] = state
	}
}

func (s *StateManager) Get(name string) State {
	if state, exist := s.states[name]; exist {
		return state
	} else {
		return nil
	}
}

// State pattern with golang has anohter implementation by using the feature of golang -- function is value.
// We can store functions in a map, use state or event as map's key.
