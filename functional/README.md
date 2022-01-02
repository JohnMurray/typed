# Functional Methods

The `functional` package contains functional methods over `map[K]V` and
`[]T` types. A limitation of Go generics currently is that receiver methods
may not have additional type parameters. This means functions such as `Map`
and `FlatMap` that may return types different from the types they operate on
must be defined as standalone functions.

## Examples and Docs

pkg.go.dev doesn't currently support documentation generation for generics. So please refer
to the below examples for reference. Also feel free to read the source if that's your jam.

### `Map`

Map takes a slice or map and a function. It then calls the function over each item
in the collection. The return value is the collection containing the results of the
function call. The type returned by the function does not have to match the type
of the input slice/map.

```go
ints := []int{1, 2, 3, 4}
strs := Map(ints, strconv.Itoa)
// strs -> ["1", "2", "3", "4"]
```