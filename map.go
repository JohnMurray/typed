package typed

import (
	"constraints"
	"sort"
)

type Map[K comparable, V any] map[K]V

// Return the set of Keys in a Map as a slice
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

// Return the keys from the map as a set of type Set[K]
func (m Map[K, V]) KeySet() Set[K] {
	return MakeSet(m.Keys()...)
}