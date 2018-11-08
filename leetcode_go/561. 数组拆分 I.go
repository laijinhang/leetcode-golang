package main

import "sort"

func arrayPairSum(nums []int) int {
	nLen := len(nums)
	sort.Ints(nums)
	max := 0
	for i := 0;i < nLen;i += 2 {
		max += nums[i]
	}
	return max
}