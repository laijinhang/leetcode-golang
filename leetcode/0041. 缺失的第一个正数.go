package main

func firstMissingPositive(nums []int) int {
	n := make(map[int]bool)
	min, max := 1, 0
	for i := 0;i < len(nums);i++ {
		if nums[i] > 0 {
			if nums[i] > max {
				max = nums[i]
			}
			n[nums[i]] = true
		}
	}
	for ;min <= max + 1 && n[min];min++ {}
	return min
}