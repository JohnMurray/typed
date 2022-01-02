package slice

// Map iterates over a slice, calling a function of T -> U on each
// item and returning a slice of the results. The returned array
// should be of equal length to the input array.
// 
//   ints := []int{1, 2, 3, 4}
//   double = slice.Map(ints, func(i int) int { return i * 2 })
//   // double [2, 4, 6, 8]
func Map[T any, U any](ts []T, f func(t T) U) []U {
	us := make([]U, 0, len(ts))
	for _, t := range ts {
		us = append(us, f(t))
	}
	return us
}

// FlatMap iterates over a slice, calling a function of T -> []U on each
// item and collects the results. Unlike map however, the results of the
// call to 'f' are "flattened" so that the output array is not
// multi-dimensional
//
//   func duplicate(i int) []int { return []int{i, i} }
//   arr1 := []int{ 1, 2, 3, 4 }
//   arr2 := slice.FlatMap(arr1, duplicate)
//   // arr2 [1, 1, 2, 2, 3, 3, 4, 4]
func FlatMap[T any, U any](ts []T, f func(t T) []U) []U {
	// init capacity of the return array assuming each call to 'f'
	// produces at least one output.
	us := make([]U, 0, len(ts))

	for _, t := range ts {
		for _, u := range f(t) {
			us = append(us, u)
		}
	}
	return us
}

// ForEach
// FoldLeft
// FoldRight
// Reduce / Collect
