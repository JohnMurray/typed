package m

// Map iterates over a map, calling a function of (K, V) -> (KK, VV)
// on each key-value pair. The resulting set of key-value pairs from
// each function call are collected in a return map.
func Map[K comparable, V any, KK comparable, VV any](m map[K]V, f func(k K, v V) (KK, VV)) map[KK]VV {
	out := make(map[KK]VV, len(m))
	for k, v := range m {
		kk, vv := f(k, v)
		out[kk] = vv
	}
	return out
}

// FlatMap iterates over a map, calling a function of (K, V) -> map[KK]VV
// on each key-value pair. The resulting set of maps from each function call
// is flattened into a single return map.
//
// If duplicate keys are found while flattening the return map, no guarantee
// is made about the ordering in which keys are written.
func FlaMap[K comparable, V any, KK comparable, VV any](m map[K]V, f func(k K, v V) map[KK]VV) map[KK]VV {
	out := make(map[KK]VV, len(m))
	for k, v := range m {
		for kk, vv := range f(k, v) {
			out[kk] = vv
		}
	}
	return out
}