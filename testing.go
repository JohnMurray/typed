package typed

import "testing"

// Must provides a way to wrap two-tuple functions calls where it _must_ succeed. The
// result is returned without the error. If the error is non-nil, the function panics.
//
// Example:
//     func doThing() (string, error) { ... }
//     func TestDothing() {
//         var result string = Must(doThing())
//     }
func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

// MustT behaves similar to `Must`, but takes a `testing.TB` and will make a call to
// Fatalf when an error is encountered.
func MustT[T any](t testing.TB, result T, err error) T {
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	return result
}