/*
题目描述
编写代码，以给定值x为基准将链表分割成两部分，所有小于x的结点排在大于或等于x的结点之前

给定一个链表的头指针 ListNode* pHead，请返回重新排列后的链表的头指针。注意：分割以后保持原来的数据顺序不变。

 */
package main

type ListNode struct {
	val int
	next *ListNode
}

type Partition struct{}

func (Partition) partition(pHead *ListNode, x int) *ListNode {
	if pHead == nil || pHead.next == nil {
		return pHead
	}
	p := pHead
	Shead := &ListNode{}
	Bhead := &ListNode{}
	stemp, btemp := Shead, Bhead
	// 拷贝
	for p != nil {
		if p.val < x {
			stemp.next = &ListNode{val:p.val}
			stemp = stemp.next
		} else {
			btemp.next = &ListNode{val:p.val}
			btemp = btemp.next
		}
		p = p.next
	}
	// 合并
	h := Shead
	if Shead.next != nil {
		p = stemp
	} else {
		p = Bhead
	}
	h.next = p
	p.next = Bhead.next

	return h
}