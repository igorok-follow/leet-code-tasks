package main

func main() {
	merge(
		[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
		6,
		[]int{1, 2, 2},
		3,
	)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
