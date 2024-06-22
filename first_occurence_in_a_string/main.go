package main

import "log"

func main() {
	log.Println(strStr("eblaneblan", "eblanee"))
}

func strStr(haystack string, needle string) int {
	for i, v := range haystack {
		if v == []rune(needle)[0] && i+len(needle) <= len(haystack) {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}

	return -1
}
