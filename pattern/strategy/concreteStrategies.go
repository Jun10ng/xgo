package strategy

import (
	"errors"
	"reflect"
)

// this file contains concrete strategies that written by yourself.
// you can write each concrete Strategy in a separate file,
// Note : you need to write init() to register them in echo file.

func init() {

	var i int
	var str string
	var intSty IntSty
	var strSty StrSty

	Register(reflect.TypeOf(i), intSty)
	Register(reflect.TypeOf(str), strSty)
}

//IntSty is used to deal with Int
type IntSty struct {
	// data interface{}
}

func (s IntSty) doStrategy(content interface{}) error {
	if _, ok := content.(int); ok {
		return nil
	}
	return errors.New("content is not int type")
}

//StrSty is used to deal with String
type StrSty struct {
	// data interface{}
}

func (s StrSty) doStrategy(content interface{}) error {
	if _, ok := content.(string); ok {
		return nil
	}
	return errors.New("content is not string type")
}
