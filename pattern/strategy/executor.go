package strategy
import (
	"fmt"
	"reflect"
)


type Executor struct{
	s map[reflect.Type]Strategy
}

func NewExecutor(sm map[reflect.Type]Strategy)Executor{
	return Executor{
		s: sm,
	}
}

func(e *Executor)execute(content interface{}) error {
	typ := reflect.TypeOf(content)
	strategy,ok := e.s[typ]
	if ok {
		return strategy.doStrategy(content)
	}else{
		return fmt.Errorf("the strategy for type [%v] is not exists",typ)
	}
}