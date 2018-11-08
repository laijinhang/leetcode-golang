package main

import "fmt"

func missingNumber(nums []int) int {
	nLen := len(nums)
	n := (nLen + 1) * (0 + nLen) / 2
	for i := 0;i < nLen;i++ {
		n -= nums[i]
	}
	return n
}

func main()  {
	fmt.Println(missingNumber([]int{0}))
}