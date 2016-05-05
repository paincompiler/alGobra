//  Created by paincompiler on 5/5/16

package alGobra

import "fmt"

type Monad interface {
	Do(fn func(interface{}, Monad) Monad) Monad
	Return(value interface{}) Monad
}

type Maybe struct {
	value *interface{}
}

var Nil = Maybe{nil}

func Some(a interface{}) Monad {
	return Maybe{&a}
}

func (m Maybe) Do(fn func(interface{}, Monad) Monad) Monad {
	if m == Nil {
		return Nil
	}
	return fn(*m.value, m)
}

func (m Maybe) Return(a interface{}) Monad {
	return Some(a)
}

func (m Maybe) String() (str string) {
	if m == Nil {
		str = "Nil"
	} else {
		str = fmt.Sprintf("Some(%v)", *m.value)
	}
	return
}
