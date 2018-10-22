package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rLen := len(nums)
	cLen := len(nums[0])
	if rLen * cLen != r * c {
		return nums
	}
	rnums := make([][]int, r)
	for i := 0;i < r;i++ {
		rnums[i] = make([]int, c)
		for j := 0;j < c;j++ {
			rnums[i][j] = nums[(i*c+j)/cLen][(i*c+j)%cLen]
		}
	}
	return rnums
}

func main()  {
	matrixReshape([][]int{{1,2},{3,4}}, 4, 1)
}