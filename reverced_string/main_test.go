package main

import "testing"

func TestRevercedString(t *testing.T) {
	test := []struct {
		str  string
		rstr string
	}{
		{"", ""},
		{"salom", "molas"},
		{"abdulaziz", "zizaludba"},
		{"asal", "lasa"},
		{"aziza", "aziza"},
	}

	for _, v := range test {
		expected := v.rstr
		actual := RevercedString(v.str)
		if expected != actual {
			t.Errorf("Expected %s do not match actual %s", expected, actual)
		}
	}
}
