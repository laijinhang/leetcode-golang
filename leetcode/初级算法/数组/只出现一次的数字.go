package main

import "sort"

func singleNumber(nums []int) int {
	nLen := len(nums)
	sort.Ints(nums)
	if nLen == 1 || nums[0] != nums[1] {
		return nums[0]
	}
	for i := 1;i < nLen - 1;i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[nLen-1]
}