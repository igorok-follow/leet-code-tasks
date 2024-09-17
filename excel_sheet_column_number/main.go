package main

import (
	"log"
	"math"
)

func main() {
	log.Println(titleToNumber("AAAA"))
}

func titleToNumber(columnTitle string) int {
	var res int
	for i := 0; i < len(columnTitle)-1; i++ {
		symbol := int(rune(columnTitle[i] - 65 + 1))
		res += int(math.Pow(26, float64(len(columnTitle)-i-1))) * symbol

		if res > 2147483647 {
			return 2147483647
		}
	}

	if len(columnTitle) > 1 {
		res += int(rune(columnTitle[len(columnTitle)-1] - 65 + 1))
	} else {
		res += int(rune(columnTitle[0] - 65 + 1))
	}

	return res
}
