package main

import (
	"fmt"
	"strings"
)

func main() {
	const input = "Hello World!"
	words := strings.Fields(input)

	m := make(map[string]int8)
	for i := range words {
		word := string(words[i])

		elem, ok := m[word]
		if ok {
			m[word] = elem + 1
		} else {
			m[word] = 1
		}
	}

	fmt.Println("Word Count: ", m)
}
