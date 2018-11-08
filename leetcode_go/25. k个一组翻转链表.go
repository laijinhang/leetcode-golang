package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	// 计算可以反转的次数
	p := head
	i := 0
	for ;p != nil;p = p.Next {
		i++
	}
	count := i / k
	if count == 0 {
		return head
	}
	// 反转头部
	p1, p2, f := head, head.Next, head.Next
	for i := 1;i < k;i++ {
		if f != nil {
			p2 = f
			f = f.Next
		}
		p2.Next = p1
		p1 = p2
	}
	h := head
	head.Next = f
	head = p1
	count--
	// 其他部位翻转
	for count > 0 {
		p1 = h.Next
		if p1 == nil || p1.Next == nil {
			return head
		}
		p2 = p1.Next
		f = p1.Next
		for i := 1;i < k;i++ {
			if f != nil {
				p2 = f
				f = f.Next
			}
			p2.Next = p1
			p1 = p2
		}
		tp := h.Next
		h.Next.Next = f
		h.Next = p1
		h = tp
		count--
	}
	return head
}