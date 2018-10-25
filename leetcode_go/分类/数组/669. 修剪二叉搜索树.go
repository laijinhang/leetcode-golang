package main
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

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	t := root
	if t.Val < L {
		t = trimBST(t.Right, L, R)
	} else if t.Val > R {
		t = trimBST(t.Left, L, R)
	} else {
		t.Left = trimBST(t.Left, L, R)
		t.Right = trimBST(t.Right, L, R)
	}
	return t
}

