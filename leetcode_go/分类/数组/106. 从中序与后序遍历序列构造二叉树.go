package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	} else if len(inorder) == 1 {
		return &TreeNode{Val:inorder[0]}
	}
	m := make(map[int]int)
	root := &TreeNode{Val:postorder[len(postorder)-1]}
	for i := 0;i < len(postorder);i++ {
		m[inorder[i]] = i
	}
	for i := len(postorder) - 2;i >= 0;i-- {
		t := root
		for {
			for ; t.Left != nil && m[t.Val] > m[postorder[i]]; t = t.Left {}
			if t.Left == nil && m[t.Val] > m[postorder[i]] {
				t.Left = &TreeNode{Val: postorder[i]}
				break
			}
			for ; t.Right != nil && m[t.Val] < m[postorder[i]]; t = t.Right {}
			if t.Right == nil && m[t.Val] < m[postorder[i]] {
				t.Right = &TreeNode{Val: postorder[i]}
				break
			}
		}
	}
	return root
}
