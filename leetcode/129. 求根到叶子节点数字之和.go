package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	allSum := 0
	getPathVals(root, 0, &allSum)
	return allSum
}

func getPathVals(t *TreeNode, sum int, allSum *int) {
	if t == nil {
		return
	}
	sum = sum * 10 + t.Val
	getPathVals(t.Left, sum, allSum)
	getPathVals(t.Right, sum, allSum)
	if t.Left == nil && t.Right == nil {
		*allSum += sum
	}
}