package main

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	log.Println(rune(-1))
	log.Println(isPalindrome(" A"))
}

func isPalindrome(s string) bool {
	s = strings.Map(stringMap, s)

	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func stringMap(s rune) rune {
	if !unicode.IsLetter(s) && !unicode.IsNumber(s) {
		return -1
	}

	return unicode.ToLower(s)
}
