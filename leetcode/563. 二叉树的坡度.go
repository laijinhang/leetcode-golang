package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type nodeS struct {
	leftSum int
	rightSum int
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	mTNode := make([]nodeS, 1)
	mTNode[0].leftSum = nodeSum(root.Left, &mTNode)
	mTNode[0].rightSum = nodeSum(root.Right, &mTNode)
	tiltSum := 0
	for i := 0;i < len(mTNode);i++ {	// 求出每个节点坡度，然后求和
		tiltSum += int(math.Abs(float64(mTNode[i].leftSum - mTNode[i].rightSum)))
	}
	return tiltSum
}

func nodeSum(t *TreeNode, mtn *[]nodeS) int {
	if t == nil {
		return 0
	}
	tn := nodeS{}
	tn.leftSum = nodeSum(t.Left, mtn)
	tn.rightSum = nodeSum(t.Right, mtn)
	*mtn = append(*mtn, tn)
	return t.Val + tn.leftSum + tn.rightSum
}