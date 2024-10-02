package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	log.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}
func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	var start int

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		if start == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
		}

		start = i + 1
	}

	return res
}
