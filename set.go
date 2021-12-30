package typed

// Set is a wrapper around map[T]struct{} and works the same way
// as "traditional" Go sets, but with much improved readability.
type Set[T comparable] map[T]struct{}

func MakeSet[T comparable](values ...T) Set[T] {
	s := make(Set[T], len(values))
	for _, v := range values {
		s.Add(v)
	}
	return s
}

// -----------------------------------------------------
// Basic Operations

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

// -----------------------------------------------------
// Simple Iterator

func (s Set[T]) ForEach(f func(v T)) {
	for k, _ := range s {
		f(k)
	}
}

// -----------------------------------------------------
// Multi-Set Operations

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := make(Set[T])

	// find shortest set (best to iterate over)
	shortest, longest := s, other
	if len(s) > len(other) {
		shortest, longest = other, s
	}

	// populate intersection Set
	for k, _ := range shortest {
		if longest.Has(k) {
			intersection.Add(k)
		}
	}
	return intersection
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T])
	for k, _ := range s {
		union.Add(k)
	}
	for k, _ := range other {
		union.Add(k)
	}
	return union
}

func (s Set[T]) Subtract(other Set[T]) Set[T] {
	complement := make(Set[T])
	for k, _ := range s {
		if !other.Has(k) {
			complement.Add(k)
		}
	}
	return complement
}
