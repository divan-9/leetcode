// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/

package trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}
