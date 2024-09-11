package main

import "log"

func main() {
	log.Println(majorityElement([]int{2, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	majorityEl, majorityElFrequency := 0, 0
	for _, num := range nums {
		switch {
		case majorityElFrequency == 0:
			majorityEl, majorityElFrequency = num, 1
		case majorityEl == num:
			majorityElFrequency++
		default:
			majorityElFrequency--
		}
	}

	return majorityEl
}
