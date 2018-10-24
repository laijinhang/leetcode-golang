package main

import (
	"sort"
	"math"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return int(math.Max(float64(nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]), float64(nums[0] * nums[1] * nums[len(nums)-1])))
}