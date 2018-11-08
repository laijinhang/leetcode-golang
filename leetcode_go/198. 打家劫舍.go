package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 0 {
		return nums[0]
	} else if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	res := make([]int, len(nums) + 1)
	res[0] = nums[0];
	res[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < len(res); i++ {
		res[i] = int(math.Max(float64(res[i-2] + nums[i]), float64(nums[i-1])))
	}
	return res[len(res)-1]
}

func main() {
	fmt.Println(rob([]int{3,1,2,1}))
}