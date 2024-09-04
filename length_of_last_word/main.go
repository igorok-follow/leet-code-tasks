package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(lengthOfLastWord("     fly me   to   the moon     "))
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")

	words := strings.Split(s, " ")

	return len([]rune(words[len(words)-1]))
}
