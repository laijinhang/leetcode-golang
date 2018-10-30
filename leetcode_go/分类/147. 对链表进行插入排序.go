package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	n := p.Next
	for ;n != nil; {
		// 1）大于尾部不移动
		if p.Val <= n.Val {
			p = p.Next
			n = n.Next
			continue
		}
		p.Next = n.Next
		// 2）小于首部，移到首部
		if head.Val > n.Val {
			p.Next = n.Next
			n.Next = head
			head = n

			if p.Next == nil {
				break
			}
			n = p.Next
			continue
		}
		// 3）中间比较
		f := head
		for ;f.Next.Val < n.Val;f = f.Next {}
		n.Next = f.Next
		f.Next = n

		if p.Next == nil {
			break
		}
		n = p.Next
	}
	return head
}
/*
输入: 4->2->1->3
输出: 1->2->3->4

4->2->1->3
2->4->1->3
1->2->4->3
1->2->3->4
*/