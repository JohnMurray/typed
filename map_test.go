package typed

import "testing"

func TestMapKeys(t *testing.T) {
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	keys := Map[string, int](m).Keys()
	if len(keys) != 2 {
		t.Fatalf("keys should contain two items, found %d:  %v", len(keys), keys)
	}
	if keys[0] != "one" && keys[1] != "one" {
		t.Fatalf("missing 'one' key: %v", keys)
	}
	if keys[0] != "two" && keys[1] != "two" {
		t.Fatalf("missing 'two' key: %v", keys)
	}
}

func TestMapKeySet(t *testing.T) {
	m := Map[string, int]{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	set := m.KeySet()
	if len(set) != 3 {
		t.Fatalf("key set should equal map length")
	}
	if !set.Has("one") || !set.Has("two") || !set.Has("three") {
		t.Fatalf("key set is missing data")
	}
}