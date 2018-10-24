package main

import "fmt"

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = [][]int{}
	ctsum([]int{}, 1, k, n)
	fmt.Println(res)
	return res
}

func ctsum(a []int, i, k, n int) {
	if len(a) == k {
		for i := 0;i < len(a);i++ {
			n -= a[i]
		}
		if n == 0 {
			res = append(res, a)
		}
		return
	}
	if i > 9 {
		return
	}
	ctsum(a, i + 1, k, n)
	ctsum(append(append([]int{}, a...), i), i + 1, k, n)
}

func main()  {
	combinationSum3(3, 9)
}