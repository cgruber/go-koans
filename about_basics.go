package go_koans

import (
	"github.com/fluentassert/verify"
	"testing"
)

func aboutBasics(t *testing.T) {
	verify.Obj(true).Equal(__boolean__).Require(t) // what is truth?
	verify.Obj(true).NotEqual(__bool__).Require(t) // in it there is nothing false

	// precision is in the eye of the beholder
	var i int = __int__
	verify.Number(i).Equal(1.0000000000000000000000000000000000000).Require(t)

	k := __int__ //short assignment can be used, as well
	verify.Number(k).Equal(1.0000000000000000000000000000000000000).Require(t)

	verify.Number(5 % 2).Equal(__int__).Require(t)
	verify.Number(5 * 2).Equal(__int__).Require(t)
	verify.Number(5 ^ 2).Equal(__int__).Require(t)

	var x int
	verify.Number(x).Equal(__int__).Require(t) // zero values are valued in Go

	var f float32
	verify.Number(f).Equal(__float32__).Require(t) // for types of all types

	var s string
	verify.String(s).Equal(__string__).Require(t) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	verify.Number(c.x).Equal(__int__).Require(t)     // and types within composite types
	verify.Number(c.f).Equal(__float32__).Require(t) // which match the other types
	verify.String(c.s).Equal(__string__).Require(t)  // in a typical way
}
