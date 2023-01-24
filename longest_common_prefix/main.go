package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(longestCommonPrefix(
		[]string{"flower", "flow", "flight"},
	))
}

func longestCommonPrefix(strs []string) string {
	var (
		min     int
		minElem string
		strLen  int
	)

	min = 201
	for i := 0; i < len(strs); i++ {
		strLen = len(strs[i])
		if strLen < min {
			min = strLen
			minElem = strs[i]
		}
	}

	var (
		longestPrefix string
		isContains    bool
	)

	isContains = true
	for i := 0; i < min; i++ {
		if isContains {
			longestPrefix += string([]rune(minElem)[i])
			for j := 0; j < len(strs); j++ {
				if strings.Index(strs[j], longestPrefix) != 0 {
					longestPrefix = longestPrefix[:len(longestPrefix)-1]
					isContains = false
					break
				}
			}
		} else {
			break
		}
	}

	return longestPrefix
}
