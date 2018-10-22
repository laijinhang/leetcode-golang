package main

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

type tNode struct {
	node *TreeNode
	h int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tn := []tNode{}
	tn = append(tn, tNode{root,1})
	h := 1
	for len(tn) > 0 {
		t := tn[0]
		if h < t.h {
			h = t.h
		}
		if t.node.Left != nil {
			tn = append(tn, tNode{t.node.Left,t.h+1})
		}
		if t.node.Right != nil {
			tn = append(tn, tNode{t.node.Right,t.h+1})
		}
		tn = tn[1:]
	}
	return h
}
