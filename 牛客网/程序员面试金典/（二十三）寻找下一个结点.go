/*
题目描述
请设计一个算法，寻找二叉树中指定结点的下一个结点（即中序遍历的后继）。

给定树的根结点指针TreeNode* root和结点的值int p，请返回值为p的结点的后继结点的值。保证结点的值大于等于零小于等于100000且没有重复值，若不存在后继返回-1。

 */
 package main

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func findSucc(root *TreeNode, p int) (res int) {
	if root.val == p {
		if root.right != nil {
			return root.right.val
		}
		if root.left != nil {
			return root.left.val
		}
		return -1
	}
	if root.right != nil {	// 遍历右子树
		r := findSucc(root.right, p)
		if r != -1 {
			return r
		}
	}
	if root.left != nil {	// 遍历左子树
		r := findSucc(root.left, p)
		if r != -1 {
			return r
		}
	}
	if root.right == nil && root.left == nil {	// 找不到结果
		return -1
	}
	return
}

func main() {

}