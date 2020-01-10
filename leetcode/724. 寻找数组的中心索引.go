package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	leftSum, rightSum := 0, 0
	for i := 1;i < len(nums);i++ {
		rightSum += nums[i]
	}
	i := 0
	for leftSum != rightSum {
		if i == len(nums) -1 {
			i = - 1
			break
		}
		leftSum += nums[i]
		i++
		rightSum -= nums[i]
	}

	if leftSum != rightSum {
		i = -1
	}
	return i
}

func main()  {
	fmt.Println(pivotIndex([]int{-1,-1,0,1,0,-1}))
}