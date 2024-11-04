package main

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := "asd"
	for i := 0; i < len(s); i++ {
		asd := strings.ToUpper(string(s[i]))
	}

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

func constructRectangle(area int) []int {
	strconv.Itoa()
	i := int(math.Sqrt(float64(area)))
	for ; i <= area; i++ {
		if area%i == 0 {
			return max(i, area/i)
		}
	}

	return nil
}

func max(x, y int) []int {
	if x > y {
		return []int{x, y}
	}
	return []int{y, x}
}
