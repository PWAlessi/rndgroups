package main

import (
	"testing"
)

func Test_shuffle(t *testing.T) {
	srcData := []string{"e", "c", "b", "d", "f", "a"}
	data := []string{"a", "b", "c", "d", "e", "f"}

	shuffle(data, 0)

	if !equal(srcData, data) {
		t.Errorf("Shuffled slice is not shuffled")
	}
}

func Test_convertByteSliceToList(t *testing.T) {
	srcData := []byte("a\nb\nc\nd")

	response := convertByteSliceToList(srcData)

	if len(response) != 4 {
		t.Errorf("Expected 4 elements, got %d", len(response))
	}

}

func Test_printGroups(t *testing.T) {
	// No checks here, just print the groups
	data := []string{"a", "b", "c", "d", "e", "f", "g"}

	printGroups(data, 4)
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
