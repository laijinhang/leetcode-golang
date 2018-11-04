package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TLNode struct {
	t *TreeNode
	h int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	tn := []TLNode{TLNode{root,1}}
	res := [][]int{}
	for len(tn) != 0 {
		if tn[0].t.Left != nil {
			tn = append(tn, TLNode{tn[0].t.Left,tn[0].h+1})
		}
		if tn[0].t.Right != nil {
			tn = append(tn, TLNode{tn[0].t.Right,tn[0].h+1})
		}
		if len(res) < tn[0].h {
			res = append(res, []int{})
		}
		res[tn[0].h-1] = append(res[tn[0].h-1], tn[0].t.Val)
		tn = tn[1:]
	}
	for i := 1;i < len(res);i += 2 {
		for ti, tj := 0, len(res[i]) - 1;ti < tj;ti, tj = ti + 1, tj - 1 {
			res[i][ti], res[i][tj] = res[i][tj], res[i][ti]
		}
	}
	return res
}