package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TNode struct {
	t *TreeNode
	h int
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	tn := []TNode{TNode{root,1}}
	nowL, nowMax := 1, root.Val    // 当前层，当前最大值
	for len(tn) != 0 {
		if tn[0].t.Left != nil {
			tn = append(tn, TNode{tn[0].t.Left,tn[0].h+1})
		}
		if tn[0].t.Right != nil {
			tn = append(tn, TNode{tn[0].t.Right,tn[0].h+1})
		}
		if nowL == tn[0].h {
			if tn[0].t.Val > nowMax {
				nowMax = tn[0].t.Val
			}
		} else {
			res = append(res, nowMax)
			nowL = tn[0].h
			nowMax = tn[0].t.Val
		}
		if len(tn) == 1 {
			res = append(res, nowMax)
		}
		tn = tn[1:]
	}
	return res
}