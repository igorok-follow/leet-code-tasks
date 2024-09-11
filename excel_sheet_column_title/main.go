package main

import "log"

func main() {
	log.Println(convertToTitle(701))
}

func convertToTitle(columnNumber int) string {
	var res string
	for columnNumber != 0 {
		quotient := (columnNumber - 1) / 26
		remainder := (columnNumber - 1) % 26
		columnNumber = quotient

		res = string(rune(remainder+65)) + res
	}

	return res
}
