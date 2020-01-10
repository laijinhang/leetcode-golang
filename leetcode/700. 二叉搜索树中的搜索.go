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

func searchBST(root *TreeNode, val int) *TreeNode {
	tn := []*TreeNode{root}
	var rt *TreeNode = nil
	for len(tn) != 0 {
		if tn[0].Val == val {
			rt = tn[0]
			break
		}
		if tn[0].Left != nil {
			tn = append(tn, tn[0].Left)
		}
		if tn[0].Right != nil {
			tn = append(tn, tn[0].Right)
		}
		tn = tn[1:]
	}
	return rt
}