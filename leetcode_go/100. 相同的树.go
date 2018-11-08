package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var res bool = true
	sameTree(p, q, &res)
	return res
}

func sameTree(p *TreeNode, q *TreeNode, t *bool) {
	if !*t || (p == nil && p == nil) {
		return
	}
	if (p.Left == nil && q.Left != nil) || (p.Right != nil && q.Right == nil) ||
		(p.Left != nil && q.Left != nil && p.Val != q.Val) ||
		(p.Right != nil && q.Right != nil && q.Val != q.Val){
		*t = false
		return
	}
	if p.Left != nil && q.Left != nil && p.Val == q.Val {
		sameTree(p.Left, q.Right, t)
	}
	if p.Right != nil && q.Right != nil && p.Val == q.Val {
		sameTree(p.Right, q.Right, t)
	}
}