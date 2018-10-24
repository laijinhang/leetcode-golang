package main

import "fmt"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	cbsum(candidates, []int{}, 0, 0, target)
	fmt.Println(res)
	return res
}

func cbsum(c, a []int, i, sum, target int) {
	if sum == target {
		res = append(res, a)
		return
	}
	if sum > target || i >= len(c) {
		return
	}
	cbsum(c, a, i+1, sum, target) // 不加入
	cbsum(c, append(append([]int{}, a...), c[i], c[i]), i, sum + c[i] * 2, target) // 加入，不向下移动位置
	cbsum(c, append(append([]int{}, a...), c[i]), i + 1, sum + c[i], target)	// 加入，向下移动位置
}

func main() {
	combinationSum([]int{2,3,5}, 8)
}