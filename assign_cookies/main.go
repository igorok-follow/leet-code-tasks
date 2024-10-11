package main

import (
	"log"
	"sort"
)

func main() {
	log.Println(findContentChildren([]int{1, 2, 3}, []int{3}))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var c1, c2, count int
	for c1 < len(g) && c2 < len(s) {
		if g[c1] <= s[c2] {
			c1++
			c2++
			count++
		} else {
			c2++
		}
	}

	return count
}
