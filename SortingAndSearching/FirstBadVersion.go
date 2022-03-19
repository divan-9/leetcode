package SortingAndSearching

func firstBadVersion(n int) int {
	return findBadVersion(0, n)
}

func findBadVersion(min int, max int) int {
	if min == max {
		return min
	}
	x := (max - min) / 2
	if isBadVersion(x) {
		return findBadVersion(min, x)
	} else {
		return findBadVersion(x, max)
	}
}

func isBadVersion(version int) bool {
	return version > 5
}
