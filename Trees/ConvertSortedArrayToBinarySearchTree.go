// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/
package trees

import (
	"fmt"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	med := len(nums) / 2
	left := nums[:med]
	right := nums[med+1:]

	fmt.Printf("%v <- %v -> %v\n", left, nums[med], right)
	return &TreeNode{
		Val:   nums[med],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right)}
}
