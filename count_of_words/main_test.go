package main

import (
	"testing"
)

func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valueA := range a {
		valueB, exists := b[key]
		if !exists || valueA != valueB {
			return false
		}
	}
	return true
}

func TestCountOfWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"", map[string]int{}},
		{"hello world hello", map[string]int{"hello": 2, "world": 1}},                              // Simple case
		{"hello  world    hello", map[string]int{"hello": 2, "world": 1}},                          // Multiple spaces
		{"apple banana orange apple banana", map[string]int{"apple": 2, "banana": 2, "orange": 1}}, // Duplicate words
		{"one one two two three three", map[string]int{"one": 2, "two": 2, "three": 2}},            // Words with repetitions
	}

	for _, test := range tests {
		actual := CountOfWords(test.input)
		if !mapsEqual(actual, test.expected) {
			t.Errorf("Input: %s, Expected: %v, Actual: %v", test.input, test.expected, actual)
		}
	}
}
