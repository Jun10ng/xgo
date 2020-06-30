package visitor

import "testing"

func TestElem(t *testing.T) {
	va := new(VisitorA)
	vb := new(VisitorB)

	nx := new(ConcreteNodeX)
	ny := new(ConcreteNodeY)

	nx.Accept(va)
	ny.Accept(va)
	nx.Accept(vb)
	ny.Accept(vb)

}
