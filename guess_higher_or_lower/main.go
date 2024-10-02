package main

import "log"

func main() {
	log.Println(guessNumber(2))
}

func guessNumber(n int) int {
	l, r := 0, n

	for l < r {
		m := (l + r) / 2

		g := guess(m)
		if g == -1 {
			r = m
		} else if g == 1 {
			l = m + 1
		} else {
			return m
		}
	}

	return l
}

func guess(num int) int {
	if num == 1 {
		return 0
	} else if num > 1 {
		return -1
	} else {
		return 1
	}
}
