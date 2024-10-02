package main

import "log"

func main() {
	log.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	fc := 0
	sc := 0
	c := 0

	for fc != len(s) && sc != len(t) {
		if s[fc] == t[sc] {
			c++
			fc++
		}

		sc++
	}

	if c < len(s) {
		return false
	}

	return true
}
