package main

import (
	"sort"
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	i, j := 0, 0
	n1Len := len(nums1)
	n2Len := len(nums2)
	for ;i < n1Len && j < n2Len; {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}

func main() {
	res := intersect([]int{1,2,2,1}, []int{2,2})
	fmt.Println(res)
}