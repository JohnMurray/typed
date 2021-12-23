package typed

import (
	"constraints"
	"sort"
)

// Given a map that has an ordered key, return a range of sorted keys
// (useful for sorted iteration of the values).
func SortedKeys[T constraints.Ordered, U any](m map[T]U) []T {
	keys := make([]T, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}
