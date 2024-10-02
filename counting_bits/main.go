package main

import "log"

func main() {
	log.Println(countBits(5))
}

func countBits(n int) []int {
	res := make([]int, n+1, n+1)

	for i := 0; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}

	return res
}
