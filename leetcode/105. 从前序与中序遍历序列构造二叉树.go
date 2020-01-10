package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val:preorder[0]}
	m := make(map[int]int)
	for i := 0;i < len(inorder);i++ {
		m[inorder[i]] = i
	}
	for i := 1;i < len(preorder);i++ {
		t := root
		for {
			for ;t.Left != nil && m[t.Val] > m[preorder[i]];t = t.Left {}
			if t.Left == nil && m[t.Val] > m[preorder[i]] {
				t.Left = &TreeNode{Val:preorder[i]}
				break
			}
			for ;t.Right != nil && m[t.Val] < m[preorder[i]];t = t.Right {}
			if t.Right == nil && m[t.Val] < m[preorder[i]] {
				t.Right = &TreeNode{Val:preorder[i]}
				break
			}
		}
	}
	return root
}
