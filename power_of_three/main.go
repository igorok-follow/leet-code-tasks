package main

import "log"

func main() {
	log.Println(isPowerOfThree(9))
}

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
