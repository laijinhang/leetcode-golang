package main

import (
	"fmt"
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	for i := 0;i < len(nums);i++ {
		nums[int(math.Abs(float64(nums[i])))-1] = -int(math.Abs(float64(nums[int(math.Abs(float64(nums[i])))-1])))
	}
	res := []int{}
	for i := 0;i < len(nums);i++ {
		if nums[i] > 0 {
			res = append(res, i + 1)
		}
	}
	return res
}

func main()  {
	fmt.Println(findDisappearedNumbers([]int{1,2,4,5,5}))
}