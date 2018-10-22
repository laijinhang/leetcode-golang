package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TNode struct {
	node *TreeNode
	val int
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := []*TreeNode{root.Left}
	r := []*TreeNode{root.Right}
	leftTraversal(l)
	rightTraversal(r)
	fmt.Println(l)
	fmt.Println(r)
	if len(l) != len(r) {
		return false
	}
	for i := 0;i < len(l);i++ {
		if l[i] == nil && r[i] == nil {
			continue
		}
		if l[i] != nil && r[i] != nil {
			if l[i].Val != r[i].Val {
				return false
			}
		}
		return false
	}
	return true
}

func leftTraversal(l []*TreeNode) {	// 左子树遍历
	for i := 0;i < len(l);i++ {
		t := l[i]
		if t != nil {
			l = append(l, l[i].Left)
			l = append(l, l[i].Right)
		}
	}
}

func rightTraversal(r []*TreeNode) {	// 右子树遍历
	for i := 0;i < len(r);i++ {
		t := r[i]
		if t != nil {
			r = append(r, r[i].Right)
			r = append(r, r[i].Left)
		}
	}
}