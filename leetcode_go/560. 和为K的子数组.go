package main

func subarraySum(nums []int, k int) int {
	num := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	for i := 0;i < len(nums);i++ {
		sum += nums[i]
		num += m[sum-k]
		m[sum]++
	}
	return num
}