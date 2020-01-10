package main

import "math"

func maxSubArray(nums []int) int {
	ans := 0
	max := math.MinInt32
	for i := 0;i < len(nums);i++ {
		if ans < 0 {	// 如果前面的值加上为0，则重新计算
			ans = 0
		}
		ans += nums[i]
		if max < ans {
			max = ans
		}
	}
	return max
}