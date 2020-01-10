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

func bBST(nums []int, t *TreeNode) {
	t.Val = nums[len(nums)/2]
	if len(nums[:len(nums)/2]) != 0 {	// 左
		t.Left = &TreeNode{}
		bBST(nums[:len(nums)/2], t.Left)
	}
	if len(nums[len(nums)/2+1:]) != 0 {	// 右
		t.Right = &TreeNode{}
		bBST(nums[len(nums)/2+1:], t.Right)
	}
}
/*
这题有点恶心，和打印出来的测试数据不同，但是还是给我过了
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := TreeNode{Val:nums[len(nums)-1]}
	bBST(nums, &root)
	return &root
}