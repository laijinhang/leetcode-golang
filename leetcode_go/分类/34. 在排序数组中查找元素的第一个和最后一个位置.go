package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1,-1}
	}
	f, l := 0, len(nums) - 1
	for f < l {
		if nums[(f+l)/2] == target {
			break
		} else if nums[(f+l)/2] < target {
			f = (f + l) / 2 + 1
		} else {
			l = (f + l) / 2 - 1
		}
	}
	if f >= l && nums[f] != target {
		return []int{-1,-1}
	}
	l = (f + l) / 2
	for f = l;f - 1 >= 0 && nums[f-1] == target;f-- {}
	for ;l + 1 < len(nums) && nums[l+1] == target;l++ {}
	return []int{f,l}
}