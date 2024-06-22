package main

import "log"

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 6, 6}

	log.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	prev := nums[0]
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[l] = nums[i]
			l++
		}
		prev = nums[i]
	}
	return l
}
