package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var (
		current int
		next    int
		chars   []string
		res     int
	)

	numbers := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res = 0
	for _, r := range s {
		chars = append(chars, string(r))
	}

	for i := 0; i < len(chars); i++ {
		if i == len(chars)-1 {
			res += numbers[chars[i]]
			break
		}

		current = numbers[chars[i]]
		next = numbers[chars[i+1]]
		if current < next {
			res -= current
		} else {
			res += current
		}
	}

	return res
}
