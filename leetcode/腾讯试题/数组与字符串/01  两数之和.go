/*
url：https://leetcode-cn.com/problems/two-sum/description/
 */
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0;i < length;i++ {
		for j := i + 1;j < length;j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main()  {
	fmt.Println(twoSum([]int{0,4,3,0}, 0))
}