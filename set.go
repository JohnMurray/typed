package typed

type Set[T comparable] struct {
	set map[T]struct{}
}

func MakeSet[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

func MakeSetValues[T comparable](values ...T) *Set[T] {
	s := &Set[T]{
		set: make(map[T]struct{}, len(values)),
	}
	for _, v := range values {
		s.Add(v)
	}
	return s
}

// -----------------------------------------------------
// Basic Operations

func (s *Set[T]) Add(element T) {
	s.set[element] = struct{}{}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.set[element]
	return ok
}

func (s *Set[T]) Remove(element T) {
	delete(s.set, element)
}

func (s *Set[T]) Length() int {
	return len(s.set)
}

// -----------------------------------------------------
// Simple Iterator

func (s *Set[T]) ForEach(f func(v T)) {
	for k, _ := range s.set {
		f(k)
	}
}

// -----------------------------------------------------
// Multi-Set Operations

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	intersection := MakeSet[T]()

	// find shortest set (best to iterate over)
	shortest, longest := s, other
	if len(s.set) > len(other.set) {
		shortest, longest = other, s
	}

	// populate intersection Set
	for k, _ := range shortest.set {
		if longest.Has(k) {
			intersection.Add(k)
		}
	}
	return intersection
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	union := MakeSet[T]()
	for k, _ := range s.set {
		union.Add(k)
	}
	for k, _ := range other.set {
		union.Add(k)
	}
	return union
}

func (s *Set[T]) Subtract(other *Set[T]) *Set[T] {
	complement := MakeSet[T]()
	for k, _ := range s.set {
		if !other.Has(k) {
			complement.Add(k)
		}
	}
	return complement
}
