package main

func minMoves(nums []int) int {
	min := nums[0]
	for i := 1;i < len(nums);i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}
	count := 0
	for i := 0;i < len(nums);i++ {
		count += nums[i] - min
	}
	return count
}