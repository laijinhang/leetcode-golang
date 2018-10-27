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
func hasPathSum(root *TreeNode, sum int) bool {
	num := 0
	pathSum(root, 0, sum, &num)
	if num == 0 {
		return false
	}
	return true
}

func pathSum(t *TreeNode, sum, target int, num *int) {
	fmt.Println(sum, target, t)
	if t == nil {
		return
	}
	if t.Left == nil && t.Right == nil && sum == target {	// 到达了叶子节点
		*num++
	}
	pathSum(t.Left, sum + t.Val, target, num)
	pathSum(t.Right, sum + t.Val, target, num)
}