/*
题目描述
请实现一个函数，检查一棵二叉树是否为二叉查找树。

给定树的根结点指针TreeNode* root，请返回一个bool，代表该树是否为二叉查找树。

 */
package main

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func checkBST(root *TreeNode) bool {
	t := []*TreeNode{}
	t = append(t, root)
	for len(t) != 0 {
		t1 := t[0]
		t = t[1:]
		if t1.left != nil {
			if t1.val > t1.left.val {
				t = append(t, t1.left)
			} else {
				return false
			}
		}
		if t1.right != nil {
			if t1.val < t1.right.val {
				t = append(t, t1.right)
			} else {
				return false
			}
		}
	}
	return true
}

func main() {

}