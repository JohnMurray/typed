# typed
[![Go](https://github.com/JohnMurray/typed/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/JohnMurray/typed/actions/workflows/go.yml)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
 ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/johnmurray/typed)

Typed is a simple and pragmatic set of type wrappers for common use-cases with Go 1.18+. This library
aims to be simple and __usable in production__ environments.

The Go team decided [_not to update libraries_][no_change] in 1.18 alongside the release of generics.
I strongly agree with this move, and so this library aims to provide wrappers for common operations
and types as both an experiment in typing in Go 1.18+ and also a transitionary package as the standard
library evolves over future releases.


## Examples

### `Queue[T]`

```go
// Allocate a queue with a fixed capacity
q := NewQueue[int](100)

// Fill up the queue using synchronous pushes. This is fine
// since we have he capacity and will be immediate.
for i := 0; i < 100; i++ {
  q.Pushes(i)
}

go func() {
  // Because we're at capacity, this will block until the
  // queue is read from and some capacity is freed up
  q.Push(101)
}()

// We can attempt to push async as well. This will return
// 'false' if the queue is full, but returns immediately.
if q.TryPush(102) {
  panic("queue should be full and return false")
}

// Pop all of our items off of the queue (including the one
// pending additional capacity)
for i := 0; i < 101; i++ {
  fmt.Printf("%d\n", q.Pop())
}

go func() {
  // Popping is also synchronous so this will block until more
  // data is pushed onto the queue
  q.Pop()
}()

// We can attempt to pull async. This return the value and a
// true/false value indicating success. If the queue is empty
// then the empty-value for the type + false is returned.
if val, ok := q.TryPop(); ok {
  panic("queue should be empty")
}

// We could also use TryPop to loop over only the current items
// in the queue
val, ok := q.TryPop()
for ok {
  // use value ...
  // then consume the next one
  val, ok = q.TryPop()
}
```

### `Stack[T]`

```go
// Allocate a stack just as you would a slice
s := make(Stack[string], 0, 100)

// Push should always succeed. Since the Stack is range-backed, this will use
// and append and should take care of growing the stack
s.Push("one")
s.Push("two")

// Popping should return the value or, if the Stack is empty, a non-nil error
// response
value, err := s.Pop()
```

### `Set[T]`

### Test Utilities

__`Must[T]`__

__`MustT[T]`__

__`MustB[T]`__

Alias for `MustT[T]`.


  [no_change]: https://github.com/golang/go/issues/48918