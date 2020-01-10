package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Next.Next == nil {
		p := head.Next
		head.Next.Next = head
		head.Next = nil
		head = p
		return head
	}
	var f, p, h *ListNode
	f = nil
	h = head
	p = head.Next
	for p.Next != nil {
		h.Next = f
		f = h
		h = p
		p = p.Next
	}
	h.Next = f
	p.Next = h
	h = p
	return h
}

func main()  {
	head := &ListNode{}
	p := head
	p.Val = 1
	for i := 2;i < 3;i++ {
		p.Next = &ListNode{Val:i}
		p = p.Next
	}
	p = reverseList(head)
	for p != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	fmt.Println()
}