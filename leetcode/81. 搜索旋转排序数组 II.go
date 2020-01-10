package main

func search(nums []int, target int) bool {
	for i := 0;i < len(nums);i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}