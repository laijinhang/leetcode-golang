package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	nLen := len(nums)
	max, tempMax := 0, 0
	if nums[0] == 1 {
		tempMax++
		max++
	}
	for i := 1;i < nLen;i++ {
		if nums[i] == nums[i-1] && nums[i] == 1 {
			tempMax++
			fmt.Println(tempMax)
			if max < tempMax {
				max = tempMax
			}
		} else if nums[i] == 1 {
			tempMax = 1
			if max < tempMax {
				max = tempMax
			}
		} else {
			tempMax = 0
		}
	}
	return max
}