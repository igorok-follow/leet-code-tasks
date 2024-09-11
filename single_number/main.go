package main

import "log"

func main() {
	log.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

// xor bit manipulation solution
func singleNumberXor(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

// dictionary solution
func singleNumber(nums []int) int {
	dict := make(map[int]int)

	for _, num := range nums {
		dict[num] += 1
	}

	for k, v := range dict {
		if v == 1 {
			return k
		}
	}

	return 0
}
