package main

func containsNearbyDuplicate(nums []int, k int) bool {
	res := false
	if k == 0 {
		res = true
	}
	for i := 0;i < len(nums) - k + 1;i++ {
		if nums[i] == nums[i+k-1] {
			res = true
			break
		}
	}
	return res
}