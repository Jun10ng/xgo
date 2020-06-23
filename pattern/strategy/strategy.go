package strategy

import (
	"errors"
	"reflect"
)

// Strategy has a doStrategy method
type Strategy interface {
	doStrategy(content interface{}) error
}

//Strategies is used to register, it is a global var
var Strategies = make(map[reflect.Type]Strategy)

// Register return an error if the strategy to deal with type already exists
func Register(typ reflect.Type, sty Strategy) error {
	if _, exists := Strategies[typ]; exists {
		errors.New("Strategy already registered")
	}
	Strategies[typ] = sty
	return nil
}
