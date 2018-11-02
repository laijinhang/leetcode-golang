package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	res := nums[0] + nums[1] + nums[2]
	for i := 0;i < len(nums);i++ {
		for j := i + 1;j < len(nums);j++ {
			for k := j + 1;k < len(nums);k++ {
				if int(math.Abs(float64((nums[i] + nums[j] + nums[k]) - target))) < int(math.Abs(float64(res - target))) {
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSumClosest([]int{46,-80,39,81,-44,22,76,-65,-11,71,-93,0,22,-85,-6,-6,34,-22,33,80,-2,93,-58,-89,-72,19,-47,-30,-14,92,25,5,16,98,62,-26,-11,7,8,4,22,-11,9,-3,55,78,-24,-38,-43,-57,-46,-35,34,93,4,23,-86,61,-74,-18,91,8,-10,89,15,24,94,84,67,37,4,-30,-26,85,-21,95,64,26,49,27,25,91,42,-79,-83,-12,-41,9,0,-100,62,66,96,66,-72,-7,-76,82,38,-58}, -89))
}