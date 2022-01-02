// The functional package provides functional operations over types defined
// in the `typed` package as well as more types oriented toward functional
// style programming.
//
// Because methods are not allowed to define additional type paramters,
// many functions that ideally would exist on types such as Map, FlatMap,
// ForEach, etc. are defined as free functions over the types. See the
// sub-packages for working with maps and slices.
//   - `m`     -> Package for working with maps
//   - `slice` -> Package for working with slices
package functional
