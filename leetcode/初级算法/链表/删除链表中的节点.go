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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := []*ListNode{}
	p := head
	for p != nil {
		node = append(node, p)
		p = p.Next
	}
	if n == len(node) {	// 删除头节点
		if len(node) == 1 {
			return nil
		}
		return  node[1]
	}
	node[len(node)-n-1].Next = node[len(node)-n].Next
	return head
}

func main()  {
	h := &ListNode{Val:1}
	p := h
	for i := 2;i <= 5;i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}
	p = h
	for p != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
	h = removeNthFromEnd(h, 2)
	p = h
	for p != nil {
		fmt.Printf("%d -> ", p.Val)
		p = p.Next
	}
}