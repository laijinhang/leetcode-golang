package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}
	// 1、遍历s，找出和t中值相等的所有节点
	nodes := []*TreeNode{}
	sn := []*TreeNode{s}
	for len(sn) != 0 {
		if sn[0].Left != nil {
			sn = append(sn, sn[0].Left)
		}
		if sn[0].Right != nil {
			sn = append(sn, sn[0].Right)
		}
		if sn[0].Val == t.Val {
			nodes = append(nodes, sn[0])
		}
		sn = sn[1:]
	}
	// 2、将节点进行比较
	flag := true
	for i := 0;i < len(nodes);i++ {
		compare(nodes[i], t, &flag)
		if flag == true {
			fmt.Println(i)
			break
		}
		flag = true
	}
	return flag
}

func compare(t1, t2 *TreeNode, flag *bool) {
	if *flag {
		if t1 == t2 {
			fmt.Println(t1, t2)
			return
		}
		if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) || t1.Val != t2.Val {
			*flag = false
			return
		}
		compare(t1.Left, t2.Left, flag)
		compare(t1.Right, t2.Right, flag)
	}
}