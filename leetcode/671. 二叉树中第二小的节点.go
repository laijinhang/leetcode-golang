package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {	// 根不存在或者不存在子节点
		return -1
	}
	lv, rv := root.Val, root.Val
	finMinVal(root.Left, root.Val, &lv)
	finMinVal(root.Right, root.Val, &rv)
	if root.Val != lv && root.Val != rv {
		return min(lv, rv)
	} else if root.Val == lv && root.Val != rv {
		return rv
	} else if root.Val != lv && root.Val == rv {
		return lv
	}
	return -1
}

func finMinVal(t *TreeNode, m int, v *int) {
	if t == nil {
		return
	}
	if t.Val != m {
		if *v == m {
			*v = t.Val
		} else {
			*v = min(*v, t.Val)
		}
		return
	}
	if t.Left != nil {
		finMinVal(t.Left, m, v)
		finMinVal(t.Right, m, v)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}