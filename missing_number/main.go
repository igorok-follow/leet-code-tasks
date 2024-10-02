package main

import "log"

func main() {
	log.Println(missingNumber([]int{3, 0, 1}))
}

func missingNumber(nums []int) int {
	var (
		sum1 int
		sum2 int
	)

	for i, num := range nums {
		sum1 += i
		sum2 += num
	}

	return sum1 + len(nums) - sum2
}
