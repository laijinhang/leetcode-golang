package main

import "math"

func findDuplicate(nums []int) int {
	res := 0
	for i := 0;i < len(nums);i++ {
		nums[int(math.Abs(float64(nums[i])))] = -nums[int(math.Abs(float64(nums[i])))]
		if nums[int(math.Abs(float64(nums[i])))] > 0 {
			res = int(math.Abs(float64(nums[i])))
			break
		}
	}
	return res
}