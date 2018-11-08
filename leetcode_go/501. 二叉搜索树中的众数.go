package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	nodeValm := make(map[int]int)
	inorderTraversal(root, nodeValm)
	res := []int{}
	max := 0
	for _, nums := range nodeValm {
		if nums > max {
			max = nums
		}
	}
	for v, nums := range nodeValm {
		if nums == max {
			res = append(res, v)
		}
	}
	return res
}

func inorderTraversal(t *TreeNode, nodeValm map[int]int) {
	if t == nil {
		return
	}
	inorderTraversal(t.Left, nodeValm)
	nodeValm[t.Val]++
	inorderTraversal(t.Right, nodeValm)
}