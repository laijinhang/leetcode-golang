package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 首部去重
	p1, p2 := head, head.Next
	for {
		p1, p2 := head, head.Next
		for ;p2 != nil && p1.Val == p2.Val;p2 = p2.Next {}
		if p1.Next != p2 && p1.Next != nil {
			head = p2
			if head == nil || head.Next == nil {
				return head
			}
		} else {
			break
		}
	}
	p1, p2 = head, head.Next
	n := 0
	for p1 != nil {
		p2 = p1.Next
		n = 0
		for ;p2 != nil && p1.Next.Val == p2.Val;p2 = p2.Next {n++}
		if n > 1 {
			p1.Next = p2
		} else {
			p1 = p1.Next
		}
	}
	return head
}