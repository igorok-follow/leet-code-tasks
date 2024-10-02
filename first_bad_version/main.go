package main

func main() {
	firstBadVersion(50)
}

func firstBadVersion(n int) int {
	return lbinsearch(n)
}

func lbinsearch(n int) int {
	l := 0
	r := n

	for l < r {
		m := (l + r) / 2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func isBadVersion(version int) bool { return false }
