package main

import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	tn := []*TreeNode{root}
	tv := []int{}
	for len(tn) != 0 {
		if tn[0].Left != nil {
			tn = append(tn, tn[0].Left)
		}
		if tn[0].Right != nil {
			tn = append(tn, tn[0].Right)
		}
		tv = append(tv, tn[0].Val)
		tn = tn[1:]
	}
	sort.Ints(tv)
	minV := math.MaxInt32
	for i := 1;i < len(tv);i++ {
		if minV > tv[i] - tv[i-1] {
			minV = tv[i] - tv[i-1]
		}
	}
	return minV
}