package visitor

import "fmt"

// We can add new service for different nodes by using visitor pattern, add new service class(like VisitotA or VisitorB), decoupling the origin Node class from new business logic and follow segregation principles.

// the Vistor interface
type Visitor interface {
	Visit(Node)
}

// and an implementation
type VisitorA struct{}

func (v VisitorA) Visit(node Node) {
	switch node.(type) {
	case ConcreteNodeX:
		fmt.Println("VisitorA on Node X")
	case ConcreteNodeY:
		fmt.Println("VisitorA on Node Y")
	}
}

type VisitorB struct{}

func (v VisitorB) Visit(node Node) {
	switch node.(type) {
	case ConcreteNodeX:
		fmt.Println("VisitorB on Node X")
	case ConcreteNodeY:
		fmt.Println("VisitorB on Node Y")
	}
}

// the Node interface
type Node interface {
	Accept(Visitor)
}
type ConcreteNodeX struct{}

func (n ConcreteNodeX) Accept(visitor Visitor) {
	visitor.Visit(n)
}

type ConcreteNodeY struct{}

func (n ConcreteNodeY) Accept(visitor Visitor) {
	visitor.Visit(n)
}
