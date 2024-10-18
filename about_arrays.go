package go_koans

import (
	"fmt"
	"github.com/fluentassert/verify"
	"testing"
)

func aboutArrays(t *testing.T) {
	fruits := [4]string{"apple", "orange", "mango"}

	verify.String(fruits[0]).Equal(__string__).Require(t) // indexes begin at 0
	verify.String(fruits[1]).Equal(__string__).Require(t) // one is indeed the loneliest number
	verify.String(fruits[2]).Equal(__string__).Require(t) // it takes two to ...tango?
	verify.String(fruits[3]).Equal(__string__).Require(t) // there is no spoon, only an empty value

	verify.Number(len(fruits)).Equal(__int__).Require(t) // the length is what the length is
	verify.Number(cap(fruits)).Equal(__int__).Require(t) // it can hold no more

	verify.Obj(fruits).Equal([4]string{}) // comparing arrays is not like comparing apples and oranges

	tasty_fruits := fruits[1:3]                                                 // defining oneself as a variation of another
	verify.String(fmt.Sprintf("%T", tasty_fruits)).Equal(__string__).Require(t) //and get not a simple array as a result
	verify.String(tasty_fruits[0]).Equal(__string__).Require(t)                 // slices of arrays share some data
	verify.String(tasty_fruits[1]).Equal(__string__).Require(t)                 // albeit slightly askewed

	verify.Number(len(tasty_fruits) == __int__).Require(t) // its length is manifest
	verify.Number(cap(tasty_fruits) == __int__).Require(t) // but its capacity is surprising!

	tasty_fruits[0] = "lemon" // are their shared roots truly identical?

	verify.String(fruits[0]).Equal(__string__).Require(t) // has this element remained the same?
	verify.String(fruits[1]).Equal(__string__).Require(t) // how about the second?
	verify.String(fruits[2]).Equal(__string__).Require(t) // surely one of these must have changed
	verify.String(fruits[3]).Equal(__string__).Require(t) // but who can know these things

	veggies := [...]string{"carrot", "pea"}

	verify.Number(len(veggies)).Equal(__int__).Require(t) // array literals need not repeat an obvious length
}
