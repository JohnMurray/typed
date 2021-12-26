package typed

import (
	"fmt"
	"log"
	"testing"
)

func B(t *testing.T) {
}

func TestStackCreation(t *testing.T) {
	s1 := make(Stack[string], 0, 10)
	s2 := make(Stack[string], 0, 0)
	var s3 Stack[string]
	if len(s1) != 0 || len(s2) != 0 || len(s3) != 0{
		log.Fatalf("new stack should be zero-length")
	}
	if cap(s1) != 10 {
		log.Fatalf("new stack should have capacity of ten")
	}
	if cap(s2) != 0 || cap(s3) != 0 {
		log.Fatalf("new stack should have zero capacity")
	}
}

func TestStackPushing(t *testing.T) {
	var emptyStack Stack[string]

	for _, input := range []struct{s Stack[string]; name string;}{
		{make(Stack[string], 0, 10), "reserved capacity"},
		{make(Stack[string], 0, 0), "initialized"},
		{emptyStack, "zero value"},
	} {
		t.Run(input.name, func(t *testing.T) {
			for i := 0; i < 10_000; i++ {
				input.s.Push(fmt.Sprintf("value:%d", i))
			}
			if len(input.s) != 10_000 {
				t.Fatalf("expected to have 10k elements")
			}
		})
	}
}

func TestStackPopSimple(t *testing.T) {
	s := make(Stack[int], 0)
	for i := 0; i < 10_000; i++ {
		s.Push(i)
	}
	for i := 0; i < 10_000; i++ {
		MustT(s.Pop())(t)
	}
}

func TestStackPop(t *testing.T) {
	s := make(Stack[int], 0)
	// push and pop with the pop operations slowly draining
	// the queue
	for i := 0; i < 100; i++ {
		s.Push(i)
	}
	for len(s) > 0 {
		for i := 0; i < 100; i++ {
			s.Push(i)
		}
		for i := 0; i < 105; i++ {
			MustT(s.Pop())(t)
		}
	}
}

func TestStackLifoOrdering(t *testing.T) {
	s := make(Stack[int], 0)
	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	// Pop data off and assert the ordering as we do so
	if MustT(s.Pop())(t) != 4 ||
		MustT(s.Pop())(t) != 3 ||
		MustT(s.Pop())(t) != 2 ||
		MustT(s.Pop())(t) != 1 ||
		MustT(s.Pop())(t) != 0 {
		t.Fatalf("stack did not pop in LIFO order")
	}

	// The stack should be empty again
	if len(s) != 0 {
		t.Fatalf("stack should be empty")
	}

	// Try to remove an item from an empty stack
	_, err := s.Pop()
	if err == nil {
		t.Fatalf("Expected error while popping from an empty stack")
	}
}