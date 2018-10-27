package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	tn := []*TreeNode{root}
	pathNum := 0
	for len(tn) != 0 {
		if tn[0].Left != nil {
			tn = append(tn, tn[0].Left)
		}
		if tn[0].Right != nil {
			tn = append(tn, tn[0].Right)
		}
		findPath(tn[0], tn[0].Val, sum, &pathNum)
		tn = tn[1:]
	}
	return pathNum
}

func findPath(t *TreeNode, nodeSum, sum int, pathNum *int) {
	if nodeSum == sum {
		*pathNum++
	}
	if t == nil {
		return
	}
	if t.Left != nil {
		findPath(t.Left, nodeSum + t.Left.Val, sum, pathNum)
	}
	if t.Right != nil {
		findPath(t.Right, nodeSum + t.Right.Val, sum, pathNum)
	}
}