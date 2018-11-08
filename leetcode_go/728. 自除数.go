package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for ;left <= right;left++ {
		if left < 10 {
			res = append(res, left)
		} else if left < 100 {
			if (left / 10 != 0 && left % (left / 10) == 0) &&
				(left % 10 != 0 && left % (left % 10) == 0) {	// 不等于要不等于0
				res = append(res, left)
			}
		} else if left < 1000 {
			if (left / 100 != 0 && left % (left / 100) == 0) &&
				(left % 100 / 10 != 0 && left % (left % 100 / 10) == 0) &&
				(left % 10 != 0 && left % (left % 10) == 0) {
				res = append(res, left)
			}
		} else if left < 10000 {
			if (left / 1000 != 0 && left % (left / 1000) == 0) &&
				(left % 1000 / 100 != 0 && left % (left % 1000 / 100) == 0) &&
				(left % 100 / 10 != 0 && left % (left % 100 / 10) == 0) &&
				(left % 10 != 0 && left % (left % 10) == 0) {
					res = append(res, left)
			}
		} else if left == 10000 {
			res = append(res, left)
		}
	}
	return res
}

func main()  {
	fmt.Println(selfDividingNumbers(1, 22))
}