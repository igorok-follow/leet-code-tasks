package main

import "log"

func main() {
	log.Println(removeElement(
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
	))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}

	return i
}

func removeElement1(nums []int, val int) int {
	var (
		found = true
	)
	for found {
		found = false
		for i, num := range nums {
			if num == val {
				nums = append(nums[:i], nums[i+1:]...)
				found = true
				break
			}
		}

	}

	return len(nums)
}
