// https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/587/
package MergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int
	for i < m && j < n {
		a := nums1[m-i-1]
		b := nums2[n-j-1]
		if a > b {
			nums1[m+n-1-i-j] = a
			i++
		} else {
			nums1[m+n-1-i-j] = b
			j++
		}
	}
	for ; j < n; j++ {
		nums1[m+n-1-i-j] = nums2[n-j-1]
	}
}
