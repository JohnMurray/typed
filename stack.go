package typed

import "fmt"

// Stack is a standard LIFO stack backed by a slice. An initial capacity may
// be specified in `NewStack`, but there is not fixed capacity.
// thread-safe
type Stack[T any] []T

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var nope T
		return nope, fmt.Errorf("empty queue")
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element, nil
}