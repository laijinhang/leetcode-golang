package main

func findPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := 0
	for i := 0;i < len(nums);i++ {
		if i == 0 && i != len(nums) - 1 {
			if nums[i] > nums[i+1] {
				res = i
				break
			}
		} else if i == len(nums) - 1 {
			if nums[i] > nums[i-1] {
				res = i
				break
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				res = i
				break
			}
		}
	}
	return res
}