package main

func main() {

}

func isAnagram(s string, t string) bool {
	c1 := [26]int{}
	c2 := [26]int{}

	for _, v := range s {
		c1[v-'a'] += 1
	}

	for _, v := range t {
		c2[v-'a'] += 1
	}

	if c1 == c2 {
		return true
	}

	return false
}
