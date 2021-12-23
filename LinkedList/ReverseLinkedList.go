// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/560/

package linkedlist

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	current := head
	for current != nil {
		result = &ListNode{Val: current.Val, Next: result}
		current = current.Next
	}
	return result
}

func reverseListR(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseListR(next)
	next.Next = head
	head.Next = nil

	return newHead
}
