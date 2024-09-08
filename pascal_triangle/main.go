package main

import "log"

func main() {
	log.Println(generate(5))
	log.Println(getRow(4))
}

func generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		line := make([]int, i+1, i+1)
		line[0] = 1
		line[len(line)-1] = 1

		for j := 1; j < len(line)-1; j++ {
			line[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, line)
	}

	return res
}

func getRow(rowIndex int) []int {
	res := make([][]int, 0)
	res = append(res, []int{1})

	for i := 1; i < rowIndex+1; i++ {
		line := make([]int, i+1, i+1)
		line[0] = 1
		line[len(line)-1] = 1

		for j := 1; j < len(line)-1; j++ {
			line[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, line)
	}

	return res[rowIndex]
}
