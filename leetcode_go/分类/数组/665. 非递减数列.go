package main

func checkPossibility(nums []int) bool {
	if len(nums) == 2 {	// 不管怎么处理都是有救的
		return true
	}
	change := false
	for i := 0;i < len(nums) - 1;i++ {
		if nums[i] > nums[i+1] && change {
			return false
		}
		if nums[i] > nums[i+1] && !change {	// 矫正一次
			if i == 0{
				if nums[i+1] <= nums[i+2] {	// 有救，救一次
					nums[0] = nums[1]
				} else {	// 没救了
					return false
				}
			} else if i + 2 == len(nums) {	// 有救
				return true
			} else {
				//确定
				//n-1 n+2	，去掉首尾情况
				//n-1 > n+2：3 3 2 2		// 不管怎么救都是没救了
				if nums[i-1] > nums[i+2] {
					return false
				}
				//n-1 = n+2：3 4 2 3		// 如果n!=n-1且n+1!=n+2，没救了
				if nums[i-1] == nums[i+2] {
					if nums[i-1] != nums[i] && nums[i+1] != nums[i+2] {
						return false
					} else if nums[i-1] == nums[i] {
						nums[i+1] =nums[i]
					} else {
						nums[i] = nums[i+1]
					}
					continue
				}
				//n-1 < n+2：2  3 4 3		// 如果是n>n+2且n+1>n+2，没救了
				if nums[i-1] < nums[i+2] {
					if (nums[i] > nums[i+2] && nums[i+1] > nums[i+2]) || (nums[i] > nums[i-1] && nums[i+1] < nums[i+2] ) {
						return false
					} else if nums[i] == nums[i+2] {
						nums[i+1] = nums[i+2]
					} else {
						nums[i] = nums[i+1]
					}
				}
			}
			change = true
		}
	}
	return true
}