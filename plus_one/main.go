package main

import "log"

func main() {
	log.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1] += 1
	} else {
		digits[len(digits)-1] = 0
		if len(digits) > 1 {
			digits[len(digits)-2] += 1
		} else {
			return append([]int{1}, 0)
		}
	}

	return digits
}
