package main

func searchInsert(nums []int, target int) int {
	res := 0
	if target <= nums[0] {
		return res
	}
	for i := 1;i < len(nums);i++ {
		if target <= nums[i] {
			res = i
			break
		}
	}
	if res == 0 {
		res = len(nums)
	}
	return res
}