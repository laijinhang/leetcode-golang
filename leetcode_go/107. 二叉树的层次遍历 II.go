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

type TNode struct {
	t *TreeNode
	h int
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	tn := []TNode{TNode{root,1}}

	for len(tn) != 0 {
		if tn[0].h > len(res) {
			res = append(res, []int{})
		}
		res[tn[0].h-1] = append(res[tn[0].h-1], tn[0].t.Val)
		if tn[0].t.Left != nil {
			tn = append(tn, TNode{tn[0].t.Left,tn[0].h+1})
		}
		if tn[0].t.Right != nil {
			tn = append(tn, TNode{tn[0].t.Right,tn[0].h+1})
		}
		tn = tn[1:]
	}
	r := [][]int{}
	for i := len(res) - 1;i >= 0;i-- {
		r = append(r, res[i])
	}
	return r
}