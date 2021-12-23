// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/772/
package linkedlist

func isPalindrome(head *ListNode) bool {
	count := 0
	current := head
	for current != nil {
		current = current.Next
		count = count + 1
	}

	current = head
	for index := 0; index <= count/2-1; index++ {
		current = current.Next
	}

	var prev *ListNode
	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	left := head
	right := prev

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}
