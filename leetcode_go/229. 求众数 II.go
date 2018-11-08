package main

func majorityElement(nums []int) []int {
	res := []int{}
	m := make(map[int]int)
	for i := 0;i < len(nums);i++ {
		m[nums[i]]++
		if float64(m[nums[i]]) > float64(len(nums)) / 3.0 {
			res = append(res, nums[i])
			m[nums[i]] = -len(nums)
		}
	}
	return res
}