package main

// 时间复杂度T(n * (k % n))
// 空间复杂度O(1)
func rotate(nums []int, k int)  {
	nLen := len(nums)
	temp, n1 := 0, nums[0]
	k %= nLen
	for i := 0;i < k;i++ {
		n1 = nums[i]
		for j := i + 1;j <= nLen + i + 1;j += 1 {
			temp = nums[j % nLen]
			nums[j % nLen] = n1
			n1 = temp
		}
	}
}

func main() {
	rotate([]int{1,2,3,4,5,6,7}, 3)
}