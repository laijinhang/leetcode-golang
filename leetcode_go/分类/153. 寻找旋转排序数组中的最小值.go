package main

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 0;i < len(nums) - 1;i++ {
		if nums[i] > nums[i+1] {
			res = nums[i+1]
			break
		}
	}
	return res
}
