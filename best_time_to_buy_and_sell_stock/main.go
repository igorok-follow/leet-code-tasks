package main

import "log"

func main() {
	log.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		min     = prices[0]
		maxProf = 0
	)
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if (prices[i] - min) > maxProf {
			maxProf = prices[i] - min
		}
	}

	return maxProf
}
