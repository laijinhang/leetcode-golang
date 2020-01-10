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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minV := 0
	tn := []*TreeNode{root}
	tVal := []int{}
	for len(tn) != 0 {
		if tn[0].Left != nil {
			tn = append(tn, tn[0].Left)
		}
		if tn[0].Right != nil {
			tn = append(tn, tn[0].Right)
		}
		tVal = append(tVal, tn[0].Val)
		tn = tn[1:]
	}
	if len(tVal) == 1 {
		minV = 0
	} else {
		minV = math.MaxInt32
		sort.Ints(tVal)
		for i := 1;i < len(tVal);i++ {
			if minV > tVal[i] - tVal[i-1] {
				minV = tVal[i] - tVal[i-1]
			}
		}
	}
	return minV
}