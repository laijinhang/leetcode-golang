package main

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	for i := 0;i < len(nums1);i++ {
		m1[nums1[i]] = true
	}
	res := []int{}
	for i := 0;i < len(nums2);i++ {
		if m1[nums2[i]] {
			res = append(res, nums2[i])
			delete(m1, nums2[i])
		}
	}
	return res
}