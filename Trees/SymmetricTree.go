// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/
package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil {
		return left.Val == right.Val &&
			isMirror(left.Left, right.Right) &&
			isMirror(left.Right, right.Left)
	}

	return false
}
