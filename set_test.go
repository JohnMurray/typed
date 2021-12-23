package typed

import "testing"

func testSetBasicOperationsByType[T comparable](t *testing.T, value1, value2 T) {
	set := MakeSet[T]()
	if set == nil {
		t.Fatalf("Could not initialize Set")
	}

	// add values to the set
	set.Add(value1)
	set.Add(value2)

	// check both values are in the set
	if !set.Has(value1) || !set.Has(value2) {
		t.Fatalf("Failed to find inserted keys in set")
	}

	// Remove a value
	set.Remove(value2)
	if set.Has(value2) {
		t.Fatalf("Failed to remove item from the set")
	}
	if set.Length() != 1 {
		t.Fatalf("Expected to have only one item in set")
	}

	// Re-insert same item
	set.Add(value1)
	if set.Length() != 1 {
		t.Fatalf("Expected to have only one item in set")
	}
}

func TestSetBasicOperations(t *testing.T) {
	testSetBasicOperationsByType(t, "one", "two")
	testSetBasicOperationsByType(t, 1, 2)
	testSetBasicOperationsByType(t, true, false)

	type testStruct struct {
		a string;
		b int;
		c float64;
	}
	testSetBasicOperationsByType(t, &testStruct{}, &testStruct{ b: 5})
}

func TestSetVariadicConstruction(t *testing.T) {
	s1 := MakeSetValues(1, 2, 3, 4)
	if s1.Length() != 4 {
		t.Fatalf("Expected to construct set with 4 values")
	}
	if !s1.Has(1) || !s1.Has(2) || !s1.Has(3) || !s1.Has(4) {
		t.Fatalf("Expected to contain values initialized with")
	}

	// Create set with many duplicates
	s2 := MakeSetValues(1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4)
	if s2.Length() != 4 {
		t.Fatalf("Expected to construct set with 4 values")
	}
	if !s2.Has(1) || !s2.Has(2) || !s2.Has(3) || !s2.Has(4) {
		t.Fatalf("Expected to contain values initialized with")
	}
}

func TestSetIteration(t *testing.T) {
	set := MakeSetValues(1, 2, 3, 4)
	sum := 0
	set.ForEach(func(x int) {
		sum += x
	})

	if sum != (1 + 2 + 3 + 4) {
		t.Fatalf("Failed to iterate over all elements in the set")
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := MakeSetValues(1, 2, 3, 4, 5, 6)
	s2 := MakeSetValues(5, 6, 7, 8, 9, 10)

	s3 := s1.Intersection(s2)

	// Validate that s1 and s2 weren't modified by the operation
	if s1.Length() != 6 || s2.Length() != 6 {
		t.Fatalf("Original Set objects modified during intersection")
	}
	if !s1.Has(1) || !s1.Has(2) || !s1.Has(3) || !s1.Has(4) || !s1.Has(5) || !s1.Has(6) {
		t.Fatalf("Original set 's1' does not have expected contents after intersection")
	}
	if !s2.Has(5) || !s2.Has(6) || !s2.Has(7) || !s2.Has(8) || !s2.Has(9) || !s2.Has(10) {
		t.Fatalf("Original set 's2' does not have expected contents after intersection")
	}

	// Validate the intersection
	if s3.Length() != 2 || !s3.Has(5) || !s3.Has(6) {
		t.Fatalf("Intersection set 's3' is the wrong size or contains wrong data")
	}
}

func TestSetUnion(t *testing.T) {
	s1 := MakeSetValues(1, 2, 3, 4, 5, 6)
	s2 := MakeSetValues(5, 6, 7, 8, 9, 10)

	s3 := s1.Union(s2)

	// Validate that s1 and s2 weren't modified by the operation
	if s1.Length() != 6 || s2.Length() != 6 {
		t.Fatalf("Original Set objects modified during union")
	}
	if !s1.Has(1) || !s1.Has(2) || !s1.Has(3) || !s1.Has(4) || !s1.Has(5) || !s1.Has(6) {
		t.Fatalf("Original set 's1' does not have expected contents after union")
	}
	if !s2.Has(5) || !s2.Has(6) || !s2.Has(7) || !s2.Has(8) || !s2.Has(9) || !s2.Has(10) {
		t.Fatalf("Original set 's2' does not have expected contents after union")
	}

	// Validate the intersection
	if s3.Length() != 10 {
		t.Fatalf("Union set 's3' is the wrong size")
	}
	if !s3.Has(1) || !s3.Has(2) || !s3.Has(3) || !s3.Has(4) || !s3.Has(5) || !s3.Has(6) || !s3.Has(7) || !s3.Has(8) || !s3.Has(9) || !s3.Has(10) {
		t.Fatalf("Union set 's3' does not contain correct data")
	}
}

func TestSetSubtraction(t *testing.T) {
	s1 := MakeSetValues(1, 2, 3, 4)
	s2 := MakeSetValues(3, 4)

	s3 := s1.Subtract(s2)

	// Validate that s1 and s2 weren't modified by the operation
	if s1.Length() != 4 || !s1.Has(1) || !s1.Has(2) || !s1.Has(3) || !s1.Has(4) {
		t.Fatalf("Original set 's1' modified during subtraction")
	}
	if s2.Length() != 2 || !s2.Has(3) || !s2.Has(4) {
		t.Fatalf("Original set 's2' modified during subtraction")
	}

	// validate complement
	if s3.Length() != 2 || !s3.Has(1) || !s3.Has(2) {
		t.Fatalf("Complement set 's3' doesnot contain correct data")
	}
}