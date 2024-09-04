package main

import (
	"log"
)

func main() {
	log.Println(mySqrt(8))

}

func mySqrt(x int) int {
	l := 0
	r := x

	for l <= r {
		m := (l + r) / 2

		pow := m * m
		if pow > x {
			r = m - 1
		} else if pow < x {
			l = m + 1
		} else {
			return m
		}
	}

	return r
}
