package main

import "log"

func main() {
	log.Println(hammingWeight(11))
}

func hammingWeight(n int) int {
	res := 0

	for i := 0; i < 64; i++ {
		res += n & 1
		n >>= 1
	}

	return res
}
