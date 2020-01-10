package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head
	for i := 2;i < m;i++ {
		p1 = p1.Next
	}
	tp := p1.Next
	p2 = p1.Next
	f := p2.Next
	tm := m
	if tm < 2 {
		tm = 2
	}
	for i := tm;i < n;i++ {
		if f != nil {
			tp = f
			f = f.Next
			tp.Next = p2
		}
		p2 = tp
	}
	if m == 1 {
		p1.Next.Next = p1
		p1.Next = f
		head = p2
	} else {
		p1.Next.Next = f
		p1.Next = p2
	}
	return head
}