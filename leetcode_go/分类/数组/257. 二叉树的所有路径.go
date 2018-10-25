package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res = []string{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	res = []string{}
	goThrough := make(map[*TreeNode]bool)
	depthErgodic(root, []int{}, goThrough)
	fmt.Println(res)
	return res
}

func depthErgodic(t *TreeNode, temp []int, gt map[*TreeNode]bool) {
	temp = append(temp, t.Val)
	gt[t] = true
	if t.Left == nil && t.Right == nil {
		temps := strconv.Itoa(temp[0])
		for i := 1;i < len(temp);i++ {
			temps += "->" + strconv.Itoa(temp[i])
		}
		res = append(res, temps)
		return
	}
	if !gt[t.Left] && t.Left != nil {
		depthErgodic(t.Left, append([]int{}, temp...), gt) // 左子树
	}
	if !gt[t.Right] && t.Right != nil {
		depthErgodic(t.Right, append([]int{}, temp...), gt) // 右子树
	}
}

func main() {
}
