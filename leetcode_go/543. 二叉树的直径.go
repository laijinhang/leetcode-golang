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
 // 动归问题
type THight struct {
	lh, rh int
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tNode := make(map[*TreeNode]THight)
	tNode[nil] = THight{-1,-1}
	comTreeHight(root, tNode)
	maxH := tNode[nil].lh + tNode[nil].rh
	fmt.Println(tNode)
	for _, v := range tNode {
		if maxH < v.lh + v.rh {
			maxH = v.lh + v.rh
		}
	}
	return maxH
}

func comTreeHight(t *TreeNode, tn map[*TreeNode]THight) {
	if t == nil {
		return
	}
	if _, ok := tn[t.Left];!ok {
		comTreeHight(t.Left, tn)
	}
	if _, ok := tn[t.Right];!ok {
		comTreeHight(t.Right, tn)
	}
	tn[t] = THight{max(tn[t.Left].lh, tn[t.Left].rh) + 1,
	max(tn[t.Right].lh, tn[t.Right].rh) + 1}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
