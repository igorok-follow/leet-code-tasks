package main

func main() {
}

func firstUniqChar(s string) int {
	count := make([]int, 26)

	for _, v := range s {
		count[v-'a']++
	}
	for i, l := range s {
		if count[l-'a'] == 1 {
			return i
		}
	}

	return -1
}
