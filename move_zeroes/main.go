package main

func main() {
	moveZeroes([]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0})
}

func moveZeroes(nums []int) {
	cursor := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[cursor] = nums[cursor], nums[i]
			cursor++
		}
	}
}
