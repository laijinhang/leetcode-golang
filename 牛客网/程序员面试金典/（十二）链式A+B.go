/*
题目描述
有两个用链表表示的整数，每个结点包含一个数位。这些数位是反向存放的，也就是个位排在链表的首部。编写函数对这两个整数求和，并用链表形式返回结果。

给定两个链表ListNode* A，ListNode* B，请返回A+B的结果(ListNode*)。

测试样例：
{1,2,3},{3,2,1}
返回：{4,4,4}

 */
package main

type ListNode struct {
	val int
	next *ListNode
}

type Plus struct{}

func plusAB(a, b *ListNode) *ListNode {
	var t1, t2 int
	var at *ListNode

	h := a
	for a != nil && b != nil {
		t1 = (a.val + b.val + t2) / 10
		a.val = (a.val + b.val + t2) % 10
		t2 = t1
		at = a
		a = a.next
		b = b.next
	}
	for a != nil {
		t1 = (a.val + t2) / 10
		a.val = (a.val + t2) % 10
		t2 = t1
		a = a.next
	}
	for b != nil {
		t1 = (b.val + t2) / 10
		b.val = (b.val + t2) % 10
		t2 = t1
		at.next = b
		at = at.next
		b = b.next
	}
	return h
}