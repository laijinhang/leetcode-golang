package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	res := -1
	if target >= nums[0] {
		for i := 0;i < len(nums) && target >= nums[0];i++ {
			if nums[i] == target {
				res = i
			}
		}
	} else {
		for i := len(nums) - 1;i >= 0 && target <= nums[i];i-- {
			if nums[i] == target {
				res = i
			}
		}
	}
	return res
}