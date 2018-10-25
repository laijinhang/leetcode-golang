package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	f, p := head, head
	// 前一个指针每次移动一位，后面一个指针每次移动两位，
	// 链表长度为奇数时，p到末尾时f刚好移动到中间，
	// 链表长度为偶数时，p到末尾时f刚好移动到两个中间值得第二个值位置
	for p.Next != nil {
		f = f.Next
		p = p.Next
		if p.Next != nil {
			p = p.Next
		}
	}
	return f
}