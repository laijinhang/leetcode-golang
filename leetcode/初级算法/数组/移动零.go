package main

func moveZeroes(nums []int)  {
	nLen := len(nums)
	for i, j := 0, 0;i < nLen;i++ {
		for ;j < nLen && nums[j] != 0;j++ {}
		if j == nLen {
			break
		}
		for i = j + 1;i < nLen && nums[i] == 0;i++ {}
		if i == nLen {
			break
		}
		nums[j] = nums[i]
		nums[i] = 0
	}
}

func main() {
	moveZeroes([]int{0,1,0,3,12})
}