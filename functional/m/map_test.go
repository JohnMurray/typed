package m

import (
	"reflect"
	"strconv"
	"testing"
)

func mapTest[K comparable, V any, KK comparable, VV any](
	t *testing.T,
	input map[K]V,
	f func(K, V) (KK, VV),
	output map[KK]VV,
) {
	out := Map(input, f)
	if !reflect.DeepEqual(out, output) {
		t.Fatalf("expected %v, got %v", output, out)
	}
}

func TestMap(t *testing.T) {
	keyToStr := func(k string, v int) (string, string) { return k, strconv.Itoa(v) }
	swap := func(k string, v int) (int, string) { return v, k }

	mapTest(t,
		map[string]int{
			"one": 1,
			"two": 2,
		},
		keyToStr,
		map[string]string{
			"one": "1",
			"two": "2",
		},
	)
	mapTest(t,
		map[string]int{
			"one": 1,
			"two": 2,
		},
		swap,
		map[int]string{
			1: "one",
			2: "two",
		},
	)
}

func TestFlatMap(t *testing.T) {

}