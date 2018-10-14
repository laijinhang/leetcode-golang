/*
5个数求最值
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
设计一个从5个整数中取最小数和最大数的程序
输入
输入只有一组测试数据，为五个不大于1万的正整数
输出
输出两个数，第一个为这五个数中的最小值，第二个为这五个数中的最大值，两个数字以空格格开。
样例输入
1 2 3 4 5
样例输出
1 5
 */
package main

import (
	"math"
	"fmt"
)

func main() {
	numbers := []int{1,2,3,4,5}
	max, min := int(math.MinInt32), int(math.MaxInt32)

	for _, val := range numbers {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Println(min, max)
}
