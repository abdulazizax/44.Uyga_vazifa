package main

import (
	"strings"
)

func CountOfWords(str string) map[string]int {
	count := make(map[string]int)
	for _, word := range strings.Fields(str) {
		count[word]++
	}
	return count
}
