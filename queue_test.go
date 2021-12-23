package typed

import "testing"

func TestQueueConstruction(t *testing.T) {
	// Just validate that none of this panics
	NewQueue[int](0)
	NewQueue[string](1024)
	NewQueue[float64](4096)

	type testStruct struct{}
	NewQueue[testStruct](0)
	NewQueue[*testStruct](32768)
}

func TestPushPull(t *testing.T) {
	q := NewQueue[int](4)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	if q.Length() != 4 {
		t.Fatalf("Queue should have length of 4")
	}

	// Try to push a 5th item onto the queue (above capacity)
	if q.TryPush(5) {
		t.Fatalf("Queue should not accept 5th item")
	}
	if q.Length() != 4 {
		t.Fatalf("Queue should have length of 4")
	}

	// pop all values (validate ordering)
	for _, i := range []int{1, 2, 3, 4} {
		if q.Pop() != i {
			t.Fatalf("Queue is not FIFO")
		}
	}
	if q.Length() != 0 {
		t.Fatalf("Queue should be empty")
	}

	// try to pop empty queue
	_, success := q.TryPop()
	if success {
		t.Fatalf("Queue should not have successfully popped while empty")
	}

	// Validate successful try-push and try-pop
	assertPush := func(b bool) {
		if !b {
			t.Fatalf("Failed to push value onto queue")
		}
	}
	assertPush(q.TryPush(5))
	assertPush(q.TryPush(6))
	assertPush(q.TryPush(7))
	assertPush(q.TryPush(8))

	assertPull := func(v int, b bool) int {
		if !b {
			t.Fatalf("failed to pull value off queue")
		}
		return v
	}
	assertPull(q.TryPop())
	assertPull(q.TryPop())
	assertPull(q.TryPop())
	assertPull(q.TryPop())
}
