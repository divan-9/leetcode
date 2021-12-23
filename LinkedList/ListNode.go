package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	currentNode := &ListNode{
		Val: slice[0],
	}
	head := currentNode
	for i := 1; i < len(slice); i++ {
		currentNode.Next = &ListNode{
			Val: slice[i],
		}
		currentNode = currentNode.Next
	}
	return head
}

func createCycledList(slice []int, pos int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	var prev, curr, head, cycled *ListNode
	for i := 0; i < len(slice); i++ {
		prev = curr
		curr = &ListNode{Val: slice[i]}
		if i == 0 {
			head = curr
		}
		if prev != nil {
			prev.Next = curr
		}
		if i == pos {
			cycled = curr
		}
	}

	curr.Next = cycled

	return head
}

func (list *ListNode) Slice(limit int) []int {
	res := make([]int, 0, limit)

	current := list
	for i := 0; i < limit && current != nil; i++ {
		res = append(res, current.Val)
		current = current.Next
	}

	return res
}
