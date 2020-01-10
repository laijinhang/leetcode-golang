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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	res1, res2 := []int{}, []int{}
	getLeafSimilarVal(root1, &res1)
	getLeafSimilarVal(root2, &res2)
	if len(res1) != len(res2) {
		return false
	}
	for i := 0;i < len(res1);i++ {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true
}

func getLeafSimilarVal(t *TreeNode, r *[]int) {
	if t.Left == nil && t.Right == nil {
		*r = append(*r, t.Val)
		return
	}
	if t.Left != nil {
		getLeafSimilarVal(t.Left, r)
	}
	if t.Right != nil {
		getLeafSimilarVal(t.Right, r)
	}
}

func main() {
}