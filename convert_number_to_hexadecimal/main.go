package main

import (
	"log"
	"math"
	"strconv"
)

func main() {
	log.Println(toHex(-1))
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num = int(math.Pow(2, 32)) + num
	}

	res := ""
	for num != 0 {
		tetrad := num & 15
		num >>= 4

		if tetrad <= 9 {
			res = strconv.Itoa(tetrad) + res
		} else {
			res = string(rune('a'-10+tetrad)) + res
		}
	}

	return res
}
