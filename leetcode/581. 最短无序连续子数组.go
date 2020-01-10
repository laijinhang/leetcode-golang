package main

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	tnums := append([]int{}, nums...)
	sort.Ints(tnums)
	i1, i2 := 0, 0
	// 起始位置
	for i1 = 0;i1 < len(nums);i1++ {
		if nums[i1] != tnums[i1] {
			break
		}
	}
	// 结束位置
	for i2 = len(nums) - 1;i2 > 0;i2-- {
		if nums[i2] != tnums[i2] {
			break
		}
	}
	// 已经排序好的情况
	if i1 == len(nums) {
		return 0
	}
	return i2 - i1 + 1
}

func main() {
	findUnsortedSubarray([]int{1,23,5,2,13})
}