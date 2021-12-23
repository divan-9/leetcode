// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/628/
package trees

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	levelOrderInternal(root, 0, &res)
	return res
}

func levelOrderInternal(node *TreeNode, level int, levels *[][]int) {
	if node == nil {
		return
	}
	if len(*levels) <= level {
		*levels = append(*levels, []int{})
	}
	tmp := *levels
	tmp[level] = append(tmp[level], node.Val)

	levelOrderInternal(node.Left, level+1, levels)
	levelOrderInternal(node.Right, level+1, levels)
}
