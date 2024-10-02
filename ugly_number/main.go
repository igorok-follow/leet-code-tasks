package main

import "log"

func main() {
	log.Println(isUgly(6))
}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 0 || n == 1 || n == 2 || n == 3 || n == 5 {
		return true
	}

	if n%5 == 0 {
		return isUgly(n / 5)
	}

	if n%3 == 0 {
		return isUgly(n / 3)
	}

	if n%2 == 0 {
		return isUgly(n / 2)
	}

	return false
}
