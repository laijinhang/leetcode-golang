package main

import "fmt"

type number struct {
	count int
	i1 int
	i2 int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]number)
	for i := 0;i < len(nums);i++ {
		if _, ok := m[nums[i]];ok == false {
			m[nums[i]] = number{1,i,i}
		} else {
			m[nums[i]] = number{m[nums[i]].count+1,m[nums[i]].i1,i}
		}
	}
	l := number{}
	for _, n := range m {
		if (n.count == l.count && n.i2 - n.i1 < l.i2 - l.i1) || n.count > l.count {
			l = n
		}
	}
	return l.i2 - l.i1 + 1
}

func main() {
	n := findShortestSubArray([]int{1,1,3,3,4,3})
	fmt.Println(n)
}