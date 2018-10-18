package main

import (
	"sort"
	"math"
	"fmt"
)

/*
设字符串长度为L，可以分成 往上取整(L/(n+1))块，设字符串中字符出现次数最多的那个字符的个数为x，x <= 往上取整值(L/(n+1))时，可以满足这种。

那么 x 的取值范围就会有两种：
(1) x < 往上取整(L/(n+1))，此时字符串长度最短为L
(2) x >= 往上取整(L/(n+1))，此时字符串只能分x块，因为原来的 (L/(n+1))填不下
	前x-1块的话，长度为(x - 1) * (n + 1）
	第x块的话，需要看与那个字符长度最长的字符个数
 */

func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	tLen := len(tasks)
	for i := 0;i < tLen;i++ {
		count[tasks[i] - 'A']++
	}
	sort.Ints(count)
	num := 0
	if count[25] >= int(math.Ceil(float64(tLen) / float64(n + 1))) {
		num = (count[25] - 1) * (n + 1) + 1
		for i := 24;i >= 0 && count[i] == count[25] && count[i] != 0;i-- {
			num++
		}
	} else {
		num = tLen
	}
	return num
}

func main()  {
	n := leastInterval([]byte{'A','A','B','B'}, 2)
	fmt.Println(n)
}