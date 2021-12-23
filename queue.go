package typed

import "sync/atomic"

// Queue is a wrapper type for `chan T` that exposes a simple, thread-safe,
// bounded FIFO queue implementation. As it stands, this cannot currently
// be resized from the initial capacity specified in `NewQueue`
type Queue[T any] struct {
	c      chan T
	length uint64
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		c: make(chan T, capacity),
	}
}

func (q *Queue[T]) Close() {
	close(q.c)
}

// -----------------------------------------------------
// Basic Operations

func (q *Queue[T]) Length() uint64 {
	return atomic.LoadUint64(&q.length)
}

func (q *Queue[T]) Push(element T) {
	q.c <- element
	atomic.AddUint64(&q.length, uint64(1))
}

func (q *Queue[T]) TryPush(element T) bool {
	select {
	case q.c <- element:
		atomic.AddUint64(&q.length, uint64(1))
		return true
	default:
		return false
	}
}

func (q *Queue[T]) Pop() T {
	result := <-q.c
	atomic.AddUint64(&q.length, ^uint64(0))
	return result
}

func (q *Queue[T]) TryPop() (T, bool) {
	select {
	case element := <-q.c:
		atomic.AddUint64(&q.length, ^uint64(0))
		return element, true
	default:
		var nope T
		return nope, false
	}
}
