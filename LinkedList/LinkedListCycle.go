// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/773/
package linkedlist

func hasCycle(head *ListNode) bool {
	cur := head
	for cur != nil {
		if cur.Next == cur {
			return true
		}
		tmp := cur
		cur = cur.Next
		tmp.Next = tmp
	}
	return false
}
