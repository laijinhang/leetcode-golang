package main

func arrayNesting(nums []int) int {
	res := 0
	numm := make(map[int]int)
	for i := 0;i < len(nums);i++ {
		numm[i] = nums[i]
	}
	for len(numm) != 0 {
		s := make(map[int]bool)
		i := 0
		for k, _ := range numm {
			i = k
			break
		}
		for {
			if _, ok := numm[i];!ok {
				break
			}
			s[i] = true
			tempI := numm[i]
			delete(numm, i)
			i = tempI
		}
		if len(s) > res {
			res = len(s)
		}
	}
	return res
}