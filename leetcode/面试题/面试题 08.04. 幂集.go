[面试题 08.04. 幂集](https://leetcode-cn.com/problems/power-set-lcci/)
```go
规律：[3,1,2]
（1）初始
[]

（2）加入第一位 3
[]

[3]

（2）加入第二位 1
[]
[3]

[1]
[1,3]

（3）加入第三位 2
[]
[3]
[1]
[1,3]

[2]
[2,3]
[1,2,3]

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for i := 0;i < len(nums);i++ {
		l := len(res)
		for j := 0;j < l;j++ {			
            arr := append([]int{}, res[j]...)
			k := 0
			for ;k < len(arr) && arr[k] <= nums[i] ;k++ {}
			if k >= len(arr) {
				arr = append(arr, nums[i])
			} else {
				arr = append(arr, 0)
				copy(arr[k+1:], arr[k:])
				arr[k] = nums[i]
			}
			res = append(res, arr)
		}
	}
	return res
}
```