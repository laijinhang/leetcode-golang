package main

import "fmt"

func removeElement(nums []int, val int) int {
	nLen := len(nums)
	count := 0
	for i := 0;i < nLen - count;i++ {
		if nums[i] == val {	// 整体往前移动
			count++
			for j := i + 1;j < nLen;j++ {
				nums[j-1] = nums[j]
			}
			i--
		}
	}
	fmt.Println(nums)
	return nLen - count
}

func main()  {
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}