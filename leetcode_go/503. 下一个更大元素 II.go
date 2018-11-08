package main

import (
	"fmt"
	"strconv"
)

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	max := nums[0]
	for i := 1;i < len(nums);i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	res := make([]int, len(nums))
	for i := 0;i < len(nums);i++ {
		if nums[i] == max {
			res[i] = -1
		} else {
			j := (i + 1) % len(nums)
			for ;nums[i] >= nums[j];j = (j + 1) % len(nums) {}
			res[i] = nums[j]
		}
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{1,2,1}))
}