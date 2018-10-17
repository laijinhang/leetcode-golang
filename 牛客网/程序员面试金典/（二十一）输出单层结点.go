/*
题目描述
对于一棵二叉树，请设计一个算法，创建含有某一深度上所有结点的链表。

给定二叉树的根结点指针TreeNode* root，以及链表上结点的深度，请返回一个链表ListNode，代表该深度上所有结点的值，请按树上从左往右的顺序链接，保证深度不超过树的高度，树上结点的值为非负整数且不超过100000。

分析：根据题目，可以知道，就是层次遍历，利用队列可以解决这题

代码实现：

 */
package main

type TreeNode struct {
	left *TreeNode
	right *TreeNode
}

type treeNode struct {
	node *TreeNode
	h int
}

type ListNode struct {
	node *TreeNode
	next *ListNode
}

func getTreeLevel(root *TreeNode, dep int) *ListNode {
	tn := []treeNode{}
	ln := &ListNode{}
	lh := ln
	tn = append(tn, treeNode{node:root, h:1})
	for len(tn) > 0 {
		t := tn[0]
		tn = tn[1:]
		if t.node.left != nil {
			if t.h + 1 == dep {
				ln.next = &ListNode{}
				ln.next.node = t.node.left
				ln = ln.next
			} else {
				tn = append(tn, treeNode{node:t.node.left, h:t.h + 1})
			}
		}
		if t.node.right != nil {
			if t.h + 1 == dep {
				ln.next = &ListNode{}
				ln.next.node = t.node.right
				ln = ln.next
			} else {
				tn = append(tn, treeNode{node:t.node.left, h:t.h + 1})
			}
		}
	}
	return lh
}

func main() {

}