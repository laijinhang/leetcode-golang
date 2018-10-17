/*
题目描述
实现一个算法，删除单向链表中间的某个结点，假定你只能访问该结点。

给定待删除的节点，请执行删除操作，若该节点为尾节点，返回false，否则返回true

分析：从题目第一句来看，似乎无解。。。
节点位置有以下两种情况：
1）不在末尾，既然没有办法走向该节点的上一个节点，但是可以找到下一节点位置。
  那么就把下一节点的值赋值到当前节点，然后当前节点的执行指向下一个节点的下一个位置，返回true
2）在末尾（无解）
      返回false

 */
package main

type ListNode struct {
	val int
	next *ListNode
}

type Remove struct {}

func (*Remove) removeNode(pNode *ListNode) bool {
	if pNode == nil || pNode.next == nil {
		return false
	}
	pNode.val = pNode.next.val
	pNode.next = pNode.next.next
	return true
}

