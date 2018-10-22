package main

import (
	"unicode"
	"math"
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
/*
一开始是用队列思路，层层遍历，但是不能判断隔层时能不能满足，然后看了好几篇博客，最后找到一篇用区间的来解，思考了下，区间这种方法行的通
*/
type TNode struct{
	node *TreeNode
	section [2]int	//闭区间
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	tn := []TNode{TNode{node:root,section:[2]int{math.MinInt32,math.MaxInt32}}}
	for len(tn) != 0 {
		t := tn[0]
		if t.node.Left != nil {
			if t.node.Left.Val >= t.node.Val || t.node.Left.Val < t.section[0]  {
				return false
			}
			tn = append(tn, TNode{t.node.Left,[2]int{t.section[0],t.node.Val-1}})
		}
		if t.node.Right != nil {
			if t.node.Right.Val <= t.node.Val || t.node.Right.Val > t.section[1] {
				return false
			}
			tn = append(tn, TNode{t.node.Right,[2]int{t.node.Val+1,t.section[1]}})
		}
		tn = tn[1:]
	}
	return true
}