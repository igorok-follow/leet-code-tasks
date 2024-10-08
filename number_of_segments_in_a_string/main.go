package main

func main() {

}

func countSegments(s string) int {
	var res int
	var counter int
	for i, v := range s {
		if v == ' ' {
			if counter != 0 {
				res++
				counter = 0
			}
		} else if i == len(s)-1 {
			res++
		} else if v != ' ' {
			counter++
		}
	}

	return res
}
