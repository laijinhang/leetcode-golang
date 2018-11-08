package main

import (
	"sort"
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	csum(candidates, 0, 0, target, &res, []int{})
	return res
}

func csum(candidates []int, i, sum, target int, res *[][]int, r []int) {
	if sum == target {	// 如果已经加入，则不加入
		for i := 0;i < len(*res);i++ {
			if len((*res)[i]) == len(r) {
				j := 0
				for ;j < len(r);j++ {
					if (*res)[i][j] != r[j] {
						break
					}
				}
				if j == len(r) {	// 已经出现过
					return
				}
			}
		}
		*res = append(*res, r)
		return
	}
	if sum > target || i == len(candidates) {
		return
	}
	csum(candidates, i + 1, sum, target, res, append([]int{}, r...))				// 不加
	csum(candidates, i + 1, sum + candidates[i], target, res, append(r, candidates[i]))// 加
}

func main() {
	res := combinationSum2([]int{2,5,2,1,2}, 5)
	fmt.Println(res)
}