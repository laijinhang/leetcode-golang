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

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Right != nil {
		root.Val += convertBST(root.Right).Val
	}
	if root.Left != nil {
		convertBST(root.Left).Val += root.Val
	}
	return root
}