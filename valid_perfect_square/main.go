package main

import "log"

func main() {
	log.Println(isPerfectSquare(15))
}

func isPerfectSquare(num int) bool {
	l, r := 0, num

	for l < r {
		m := (l + r) / 2

		if m*m > num {
			r = m
		} else if m*m < num {
			l = m + 1
		} else {
			return true
		}
	}

	if l*l == num {
		return true
	}

	return false
}
