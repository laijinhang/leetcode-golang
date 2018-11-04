package main

import (
	"sort"
	"fmt"
)

func fourSum(nums []int, target int) [][]int {
	var i, j, left, right int
	sort.Ints(nums)
	res := [][]int{}
	for i = 0;i < len(nums);i++ {
		for j = len(nums) - 1;j > i;j-- {
			left = i + 1
			right = j - 1
			for left < right {
				if nums[i] + nums[left] + nums[right] + nums[j] == target {
					flag := true
					for n := 0;n < len(res);n++ {
						if res[n][0] == nums[i] && res[n][1] == nums[left] && res[n][2] == nums[right] && res[n][3] == nums[j] {
							flag = false
							break
						}
					}
					if flag {
						res = append(res, []int{nums[i],nums[left],nums[right],nums[j]})
					}
					left++
				} else if nums[i] + nums[left] + nums[right] + nums[j] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{-5,5,4,-3,0,0,4,-2}, 4))
}