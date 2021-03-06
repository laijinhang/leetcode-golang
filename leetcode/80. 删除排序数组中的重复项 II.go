package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, j, n, tempN, nLen := 0, 0, 0, 0, len(nums)
	for n < nLen {
		for ;n < nLen && nums[i] == nums[n];n++ {}
		if n - i > 2 { // 需要向前移动
			nLen -= (n - i - 2)
			j = i + 2
			i = j
			tempN = n
			for ;tempN < len(nums);j, tempN = j + 1, tempN + 1 {
				nums[j] = nums[tempN]
			}
			n = i
			if n >= nLen {
				break
			}
		} else {
			i = n
		}
	}
	return nLen
}

func main() {
	fmt.Println(removeDuplicates([]int{-49,-49,-46,-46,-46,-46,-46,-45,-45,-45,-44,-42,-42,-41,-39,-39,-39,-38,-38,-38,-38,-38,-37,-36,-36,-35,-35,-34,-34,-34,-34,-34,-33,-32,-32,-32,-31,-31,-31,-31,-30,-30,-30,-28,-28,-28,-27,-27,-27,-27,-25,-23,-23,-23,-22,-22,-21,-21,-20,-20,-20,-19,-19,-17,-15,-15,-14,-14,-13,-12,-11,-11,-10,-10,-9,-8,-8,-7,-7,-6,-6,-6,-6,-5,-5,-5,-4,-4,-4,-3,-3,-2,-2,0,0,0,0,1,1,2,3,4,4,5,5,7,7,7,8,9,9,10,11,11,11,12,12,12,12,13,15,15,16,16,17,17,18,19,20,20,21,21,22,22,22,22,24,24,25,25,25,28,28,28,28,28,29,29,29,30,30,30,31,31,32,33,33,33,34,34,36,36,37,37,37,37,37,38,38,39,40,40,41,41,41,41,42,44,44,44,44,45,45,46,47,47,47,48,48,48,48,48,48,50,50,50}))
}