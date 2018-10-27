package main

import (
	"sort"
	"fmt"
)

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
	t *TreeNode
	h int
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	tn := []*TreeNode{root}
	hl, hr := 0, 0
	for len(tn) != 0 {
		if tn[0].Left != nil {
			tn = append(tn, tn[0].Left)	// 计算左子树高度
			hl = treeH(tn[0].Left, 1)
		} else {
			hl = 0
		}
		if tn[0].Right != nil {
			tn = append(tn, tn[0].Right)
			hr = treeH(tn[0].Right, 1)
		} else {
			hr = 0
		}
		if hl - hr > 1 || hr - hl > 1 {
			return false
		}
		tn = tn[1:]
	}
	return true
}

func treeH(t *TreeNode, h int) int {
	tn := []TNode{TNode{t:t,h:h}}
	maxH := h
	for len(tn) != 0 {
		if tn[0].t.Left != nil {
			tn = append(tn, TNode{tn[0].t.Left,tn[0].h+1})
		}
		if tn[0].t.Right != nil {
			tn = append(tn, TNode{tn[0].t.Right,tn[0].h+1})
		}
		if maxH < tn[0].h {
			maxH = tn[0].h
		}
		tn = tn[1:]
	}
	return maxH
}