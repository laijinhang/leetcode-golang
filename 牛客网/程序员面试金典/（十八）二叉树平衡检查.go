/*
题目描述
实现一个函数，检查二叉树是否平衡，平衡的定义如下，对于树中的任意一个结点，其两颗子树的高度差不超过1。

给定指向树根结点的指针TreeNode* root，请返回一个bool，代表这棵树是否平衡。

 */
package main

import "math"

type tree struct {
	left, right *tree
}

type node struct {
	t *tree
	h int
}

type Balance struct{}

func (Balance) isBalance(root *tree) bool {
	if root == nil {
		return true
	}
	treeQ := append([]node{}, node{root, 1})
	maxH, minH := int(math.MaxInt32), int(math.MinInt32)
	for len(treeQ) == 0 {
		p := treeQ[0]
		if p.t.left != nil {
			treeQ = append(treeQ, node{p.t.left, p.h + 1})
		}
		if p.t.right != nil {
			treeQ = append(treeQ, node{p.t.right, p.h + 1})
		}
		if p.t.left == nil && p.t.right == nil {
			if p.h > maxH {
				maxH = p.h
			}
			if p.h < minH {
				minH = p.h
			}
		}
		treeQ = treeQ[1:]
	}
	if maxH - minH > 1 {
		return false
	}
	return true
}

func main() {

}
