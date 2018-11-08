package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head != nil {
		p := head
		// 获取出第一个head的值
		for ;p != nil && p.Val == val;p = p.Next {}
		head = p
		// 去除后面的相等的值
		for p != nil && p.Next != nil {
			if p.Next.Val == val {
				p.Next = p.Next.Next
			} else {
				p = p.Next
			}
		}
	}
	return head
}