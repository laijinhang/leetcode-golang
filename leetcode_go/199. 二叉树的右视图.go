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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 广度遍历，每层右边
	res := []int{}
	tn := []TNode{TNode{root,1}}
	for len(tn) != 0 {
		if tn[0].t.Left != nil {
			tn = append(tn, TNode{tn[0].t.Left,tn[0].h+1})
		}
		if tn[0].t.Right != nil {
			tn = append(tn, TNode{tn[0].t.Right,tn[0].h+1})
		}
		if len(res) < tn[0].h {
			res = append(res, tn[0].t.Val)
		} else {
			res[tn[0].h-1] = tn[0].t.Val
		}
		tn = tn[1:]
	}
	return res
}