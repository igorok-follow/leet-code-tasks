package main

import "log"

func main() {
	log.Println(isHappy(2))
}

func isHappy(n int) bool {
	set := make(map[int]struct{})

	for {
		if _, ok := set[n]; ok {
			return false
		}

		set[n] = struct{}{}

		sum := 0
		for n != 0 {
			digit := n % 10

			sum += digit * digit
			n /= 10
		}

		if sum == 1 {
			return true
		}

		n = sum
	}
}
