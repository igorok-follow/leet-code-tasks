package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	log.Println(climbStairs(37))
	log.Println(time.Since(t))
}

var cache = map[int]int{}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if v, ok := cache[n]; ok {
		return v
	}

	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}
