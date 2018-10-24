package main

import (
	"sort"
	"math"
)

func thirdMax(nums []int) int {
	res := 0
	if len(nums) < 3 {
		sort.Ints(nums)
		res = nums[len(nums)-1]
	} else {
		m1, m2, m3 := math.MinInt32, math.MinInt32, math.MinInt32
		for i := 0;i < len(nums);i++ {
			if nums[i] > m1 {
				m3 = m2
				m2 = m1
				m1 = nums[i]
			} else if nums[i] > m2 && nums[i] != m1 {
				m3 = m2
				m2 = nums[i]
			} else if nums[i] > m3 && nums[i] != m2 {
				m3 = nums[i]
			}
			res = m3
			if m3 == math.MinInt32 {
				res = m1
			}
		}
	}
	return res
}