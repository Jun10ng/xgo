package memento

type StateType string

//
type Memento interface {
	GetState() StateType
	SetState(s StateType)
}

type Originator interface {
	SetState(s StateType)
	CreateMemento() Memento
	Restore(m Memento)
}

//
type TextMemento struct {
	// state stores the text of the previous version
	// you can use a stack<StateType> to store previous versions
	state StateType
}

func (t TextMemento) GetState() StateType {
	return t.state
}

func (t TextMemento) SetState(s StateType) {
	t.state = s
}

//
type TextOriginator struct {
	state StateType
}

func NewTextOri(s StateType) TextOriginator {
	return TextOriginator{
		state: s,
	}
}
// 
func (t *TextOriginator) CreateMemento() TextMemento{
	return TextMemento{
		state: t.state,
	}
}

func (t *TextOriginator)SetState(s StateType)  {
	t.state = s
}

func (t *TextOriginator)Restore(m Memento){
	t.state = m.GetState()
}

