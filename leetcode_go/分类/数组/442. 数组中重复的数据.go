package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
	for i := 0;i < len(nums);i++ {
		nums[int(math.Abs(float64(nums[i])))-1] = -nums[int(math.Abs(float64(nums[i])))-1]
	}
	res := []int{}
	for i := 0;i < len(nums);i++ {
		if nums[int(math.Abs(float64(nums[i]))) - 1] > 0 {
			res = append(res, int(math.Abs(float64(nums[i]))))
		}
	}
	fmt.Println(res)
	res = res[:len(res)/2]
	return res
}

func main() {
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}