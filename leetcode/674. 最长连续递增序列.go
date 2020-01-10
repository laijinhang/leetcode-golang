package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := 1
	max := 1
	for i := 1;i < len(nums);i++ {
		if nums[i] - nums[i-1] > 0{
			n++
			if n > max {
				max = n
			}
		} else {
			n = 1
		}
	}
	return max
}