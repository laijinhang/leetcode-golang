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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	sumOfLeft(root, &sum)
	return sum
}

func sumOfLeft(t *TreeNode, sum *int) {
	if t.Left != nil {
		if t.Left.Left == nil && t.Left.Right == nil {
			*sum += t.Val
		} else {
			sumOfLeft(t.Left, sum)
		}
	}
	if t.Right != nil {
		sumOfLeft(t, sum)
	}
}