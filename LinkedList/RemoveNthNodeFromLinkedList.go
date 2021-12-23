// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/603/

package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	var prevDel *ListNode
	counter := 0

	for current != nil {
		counter = counter + 1
		current = current.Next
		if prevDel != nil {
			prevDel = prevDel.Next
		}
		if counter == n+1 {
			prevDel = head
		}
	}

	if counter == n {
		return head.Next
	}

	if prevDel != nil {
		prevDel.Next = prevDel.Next.Next
	}

	return head
}
