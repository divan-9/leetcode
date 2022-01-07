// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/
package trees

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt32, math.MaxInt32)
}

func isValid(root *TreeNode, l int, u int) bool {
	if root == nil {
		return true
	}
	if root.Val < l || root.Val > u {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return isValid(root.Left, l, root.Val-1) &&
		isValid(root.Right, root.Val+1, u)
}
