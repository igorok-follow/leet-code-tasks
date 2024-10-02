package main

import "log"

func main() {
	log.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	set := map[uint8]struct{}{
		'a': {},
		'A': {},
		'e': {},
		'E': {},
		'o': {},
		'O': {},
		'i': {},
		'I': {},
		'u': {},
		'U': {},
	}
	res := []byte(s)

	l, r := 0, len(s)-1

	for l < r {
		if _, ok := set[s[l]]; !ok {
			l++
			continue
		}

		if _, ok := set[s[r]]; !ok {
			r--
			continue
		}

		res[l], res[r] = res[r], res[l]
		l++
		r--
	}

	return string(res)
}
