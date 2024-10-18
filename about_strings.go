package go_koans

import (
	"fmt"
	"github.com/fluentassert/verify"
	"testing"
)

func aboutStrings(t *testing.T) {

	verify.String("a" + __string__).Equal("abc").Require(t) // string concatenation need not be difficult
	verify.Number(len("abc")).Equal(__int__).Require(t)     // and bounds are thoroughly checked

	verify.Number("abc"[0]).Equal(__byte__).Require(t) // their contents are reminiscent of C

	verify.String("smith"[2:]).Equal(__string__).Require(t)  // slicing may omit the end point
	verify.String("smith"[:4]).Equal(__string__).Require(t)  // or the beginning
	verify.String("smith"[2:4]).Equal(__string__).Require(t) // or neither
	verify.String("smith"[:]).Equal(__string__).Require(t)   // or both

	verify.String("smith").Equal(__string__).Require(t)  // they can be compared directly
	verify.String("smith").Lesser(__string__).Require(t) // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	verify.String(string(bytes)).Equal(__string__).Require(t) // strings can be created from byte-slices

	bytes[0] = 'z'
	verify.String(string(bytes)).Equal(__string__).Require(t) // byte-slices can be mutated, although strings cannot

	verify.String(fmt.Sprintf("hello %s", __string__)).Equal("hello world").Require(t) // our old friend sprintf returns
	verify.String(fmt.Sprintf("hello \"%s\"", "world")).Equal(__string__).Require(t)   // quoting is familiar
	verify.String(fmt.Sprintf("hello %q", "world")).Equal(__string__).Require(t)       // although it can be done more easily

	verify.String(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589)).Equal(__string__).Require(t) // "the root of all evil" is actually a misquotation, by the way
}
