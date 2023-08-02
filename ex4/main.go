package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := []string{"a123bc34d8ef34", "A1b01c001", "A1b01c101"}
	for _, w := range word {
		count := numDifferentIntegers(w)
		fmt.Println(count)
	}
}

func numDifferentIntegers(word string) int {
	count := make(map[string]bool)

	for i := 0; i < len(word); i++ {
		if unicode.IsDigit(rune(word[i])) {
			j := i
			for j < len(word) && unicode.IsDigit(rune(word[j])) {
				j++
			}
			for i < j-1 && word[i] == '0' {
				i++
			}

			intStr := word[i:j]
			count[intStr] = true
			i = j
		}
	}

	return len(count)
}
