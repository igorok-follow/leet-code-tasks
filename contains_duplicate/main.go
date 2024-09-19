package main

import "log"

func main() {
	log.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		l   = 0
		r   = 0
		set = make(map[int]struct{})
	)

	for r = 0; r < len(nums); r++ {
		if _, ok := set[nums[r]]; ok {
			return true
		}

		set[nums[r]] = struct{}{}

		if len(set) > k {
			delete(set, nums[l])
			l++
		}
	}

	return false
}
