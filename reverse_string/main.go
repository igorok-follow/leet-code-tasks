package main

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
