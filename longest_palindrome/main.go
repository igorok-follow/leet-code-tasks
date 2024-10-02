package main

import "log"

func main() {
	log.Println(longestPalindrome("ccc"))
}

func longestPalindrome(s string) int {
	res := 0
	count := make(map[int32]int, 52)

	for _, v := range s {
		count[v] += 1
	}

	var isOne bool
	for k, v := range count {
		if v/2 >= 1 {
			add := v / 2 * 2
			count[k] -= add
			res += add
		}

		if count[k] == 1 {
			isOne = true
		}
	}

	if isOne {
		return res + 1
	}

	return res
}
