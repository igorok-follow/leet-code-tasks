package main

import "log"

func main() {
	log.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	lfreq := make(map[int32]int, len(magazine))

	for _, v := range magazine {
		lfreq[v]++
	}

	for _, n := range ransomNote {
		v := lfreq[n]
		if v == 0 {
			return false
		}
		lfreq[n]--
	}

	return true
}
