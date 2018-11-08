package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	maxCount := 0
	res := []int{}
	nodeSumM := make(map[int]int)
	nodeSum(root, nodeSumM, &maxCount)
	for k, v := range nodeSumM {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res
}

func nodeSum(t *TreeNode, nodeSumM map[int]int, max *int) int {
	if t == nil {
		return 0
	}
	sum := nodeSum(t.Left, nodeSumM, max) + nodeSum(t.Right, nodeSumM, max) + t.Val
	nodeSumM[sum]++
	if *max < nodeSumM[sum] {
		*max = nodeSumM[sum]
	}
	strings.
	return sum
}