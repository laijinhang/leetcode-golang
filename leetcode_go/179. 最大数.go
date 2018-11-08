package main

import (
	"fmt"
)

func largestNumber(nums []int) string {
	resS := make([]string, len(nums))
	for i := 0;i < len(nums);i++ {
		resS[i] = strconv.Itoa(nums[i])
	}
	// 重排
	for i := 0;i < len(resS);i++ {
		for j := i + 1;j < len(resS);j++ {
			if !compare(resS[i], resS[j]) {	// 如果a[i] < a[j]，交换
				resS[i], resS[j] = resS[j], resS[i]
			}
		}
	}
	// 拼接
	res := ""
	for i := 0;i < len(resS);i++ {
		res += resS[i]
	}
	if strings.Count(res, "0") == len(res) {
		res = "0"
	}
	return res
}

func compare(a, b string) bool {
	// 三种比较情况：
	//		len(a) == len(b)
	//		len(a) < len(b)
	// 		len(a) > len(b)
	if len(a) == len(b) && a >= b {
		return true
	}
	if (len(a) < len(b) && a > b[:len(a)]) || (len(a) < len(b) && a == b[:len(a)] && compare(a, b[len(a):])) {
		return true
	}
	if (len(a) > len(b) && a[:len(b)] > b) || (len(a) > len(b) && a[:len(b)] == b && compare(a[len(b):], b)) {
		return true
	}
	return false
}