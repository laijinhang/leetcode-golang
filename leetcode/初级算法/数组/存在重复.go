package main

import "sort"

func containsDuplicate(nums []int) bool {
	nLen := len(nums)
	sort.Ints(nums)
	for i := 0;i < nLen - 1;i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}