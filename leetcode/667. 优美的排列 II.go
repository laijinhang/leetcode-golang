package main

func constructArray(n int, k int) []int {
	res := []int{1}
	//
	for i := 1;i <= k;i++ {
		if i % 2 == 0 {
			res = append(res, res[len(res)-1] - (k - i + 1))
		} else {
			res = append(res, (k - i + 1) + res[len(res)-1])
		}
	}
	//
	for i := k + 2;i <= n;i++ {
		res = append(res, i)
	}
	return res
}
/*
规律分析：
k = 1
1 2 3 4 5 6

k = 2
1 3 2 4 5 6

1 2 3 4 5 6   k = 3
1 2 3
1 4 2 3 5 6

k = 4
1 2 3 4
1 5 2 4 3 6

k = 5
1 2 3 4 5
1 6 2 5 3 4
*/