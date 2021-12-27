package typed

import "fmt"

// Stack is a standard LIFO stack backed by a slice. Queue specific methods
// include Push and Pop. Stacks may be declared as empty or allocated with
// `make` as a regular slice.
type Stack[T any] []T

// Push puts an element onto the top of the stack.
func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

// Pop returns the top element from the stack. If the stack is empty an
// error will be returned. Else the error will always be nil
func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var nope T
		return nope, fmt.Errorf("empty queue")
	}
	element := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return element, nil
}