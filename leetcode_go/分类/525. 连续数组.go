package main

import (
	"fmt"
	"time"
	"math/rand"
)

func findMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	zeroS := make([]int, len(nums)+1)
	oneS := make([]int, len(nums)+1)
	for i := 0;i < len(nums);i++ {
		zeroS[i+1] = zeroS[i]
		oneS[i+1] = oneS[i]
		if nums[i] == 0 {
			zeroS[i+1]++
		} else {
			oneS[i+1]++
		}
	}
	i, j := 0, 0
	maxL := 0
	for ;i <= len(nums) && len(nums) - maxL > i;i++ {
		for j = len(nums);j - i > maxL;j-- {
			if zeroS[j] - zeroS[i] == oneS[j] - oneS[i] {
				if maxL < j - i {
					maxL = j - i
				}
				break
			}
		}
	}
	return maxL
}

func main() {
	t := time.Now()
	a := make([]int, 100001)
	for i := 0;i < 100001;i++ {
		a[i] = rand.Intn(2)
	}
	fmt.Println(time.Since(t))
	fmt.Println(findMaxLength(a))
}