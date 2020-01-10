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

type TNode struct {
	node *TreeNode
	h int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	tn := []TNode{TNode{node:root,h:1}}
	for len(tn) > 0 {
		t := tn[0]
		tn = tn[1:]
		if t.h > len(res) {
			res = append(res, []int{})
		}
		res[t.h-1] = append(res[t.h-1], t.node.Val)
		if t.node.Left != nil {
			tn = append(tn, TNode{node:t.node.Left,h:t.h+1})
		}
		if t.node.Right != nil {
			tn = append(tn, TNode{node:t.node.Right,h:t.h+1})
		}
	}
	return res
}

func main()  {

}