//  Created by paincompiler on 5/5/16

package alGobra

import (
	"fmt"
	"testing"
)

type strSliceMondad []string

func (m strSliceMondad) Do(fn func(interface{}, Monad) Monad) Monad {
	monad := strSliceMondad{}
	for _, val := range m {
		next := fn(val, m).(strSliceMondad)
		monad = append(monad, next...)
	}
	return monad
}

func (m strSliceMondad) Return(a interface{}) Monad {
	return strSliceMondad{a.(string)}
}

var (
	repeatDouble = func(i interface{}, m Monad) Monad {
		return m.Return(i.(string) + i.(string))
	}

	repeatTriple = func(i interface{}, m Monad) Monad {
		return m.Return(i.(string) + i.(string) + i.(string))
	}

	nothing = func(i interface{}, m Monad) Monad {
		switch m.(type) {
		case strSliceMondad:
			return strSliceMondad(nil)
		default:
			return Nil
		}
	}
)

func TestAlgobra(t *testing.T) {

	if fmt.Sprint(strSliceMondad{"A", "B"}.Do(repeatDouble).Do(repeatTriple)) != "[BBBBBB CCCCCC]" {
		t.Error("String slice chain failed")
	}

	if fmt.Sprint(strSliceMondad{"C", "D"}.Do(repeatTriple).Do(nothing).Do(repeatDouble)) != "[]" {
		t.Error("Nil string slice in chain failed")
	}

	if fmt.Sprint(Some("E").Do(repeatDouble).Do(repeatTriple)) != "Some(AAAAAA)" {
		t.Error("Some chain failed")
	}

	if fmt.Sprint(Some("F").Do(nothing).Do(repeatTriple)) != "Nil" {
		t.Error("Nil some in chain failed")
	}
	return
}
