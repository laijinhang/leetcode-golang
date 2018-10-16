package main

/*
url：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 */
import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	medianV := 0.0
	if len(nums) % 2 == 0 {
		medianV = (float64(nums[len(nums) / 2 -1]) + float64(nums[len(nums) / 2])) / 2.0
	} else {
		medianV = float64(nums[len(nums) / 2])
	}
	return medianV
}

func main() {
	findMedianSortedArrays([]int{1,2,3,8}, []int{2,4,6})
}