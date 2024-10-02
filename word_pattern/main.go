package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(wordPattern("abbac", "dog cat cat dog rabbit"))
}

func wordPattern(pattern string, s string) bool {
	dict := make(map[string]uint8)
	used := make([]bool, 26, 26)

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i, word := range words {
		if code, ok := dict[word]; ok {
			if code != pattern[i] {
				return false
			}
		} else {
			if used[pattern[i]-'a'] {
				return false
			}

			used[pattern[i]-'a'] = true
			dict[word] = pattern[i]
		}
	}

	return true
}
