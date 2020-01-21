package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	jw := 0
	p := l1
	for ;p.Next != nil && l2 != nil;p, l2 = p.Next, l2.Next {
		p.Val, jw = (p.Val + l2.Val + jw) % 10, (p.Val + l2.Val + jw) / 10
	}
	if l2 != nil {
		p.Val, jw = (p.Val + l2.Val + jw) % 10, (p.Val + l2.Val + jw) / 10
		p.Next = l2.Next
		if p.Next != nil {
			for p = p.Next;p.Next != nil;p = p.Next {
				p.Val, jw = (p.Val + jw) % 10, (p.Val + jw) / 10
			}
		}
		if p.Next != l2.Next {
			p.Val, jw = (p.Val + jw) % 10, (p.Val + jw) / 10
		}
	} else{
		for ;p.Next != nil;p = p.Next {
			p.Val, jw = (p.Val + jw) % 10, (p.Val + jw) / 10
		}
		p.Val, jw = (p.Val + jw) % 10, (p.Val + jw) / 10
	}
	if jw != 0 {
		p.Next = &ListNode{Val:jw}
	}
	return l1
}
