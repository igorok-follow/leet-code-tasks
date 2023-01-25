package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(isValid("()))"))
}

func isValid(s string) bool {
	for {
		if strings.Contains(s, "{}") || strings.Contains(s, "()") || strings.Contains(s, "[]") {
			s = strings.Replace(s, "{}", "", -1)
			s = strings.Replace(s, "()", "", -1)
			s = strings.Replace(s, "[]", "", -1)
		} else {
			break
		}
	}

	if len(s) != 0 {
		return false
	}

	return true
}
