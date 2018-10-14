/*
寻找最大数
时间限制：1000 ms  |  内存限制：65535 KB
难度：2
描述
请在整数 n 中删除m个数字, 使得余下的数字按原次序组成的新数最大，

比如当n=92081346718538，m=10时，则新的最大数是9888



输入
第一行输入一个正整数T，表示有T组测试数据
每组测试数据占一行，每行有两个数n,m（n可能是一个很大的整数，但其位数不超过100位，并且保证数据首位非0，m小于整数n的位数）
输出
每组测试数据的输出占一行，输出剩余的数字按原次序组成的最大新数
样例输入
2
92081346718538 10
1008908 5
样例输出
9888
98
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

var maxNum int

func findMax(s []byte, res []byte, n, m, maxM int) {
	if m == maxM {	// 最大允许去掉的数的个数
		tempR := append([]byte{}, res...)	// 深度拷贝，不深度拷贝的话，会把原来的覆盖掉
		if n < len(s) {
			tempR = append(tempR, s[n:]...)
		}
		tempV, _ := strconv.Atoi(string(tempR))
		if tempV > maxNum {
			maxNum = tempV
		}
		return
	}
	if n == len(s) {	// 走到末尾
		return
	}
	// 不去掉，go切片是引用，所以需要深度拷贝
	findMax(s, append(append([]byte{}, res...), s[n]), n+1, m, maxM)
	// 去掉
	findMax(s, append([]byte{}, res...), n+1, m+1, maxM)
}

func main() {
	var T, m int
	var n string
	fmt.Scan(&T)
	for i := 0;i < T;i++ {
		fmt.Scan(&n, &m)
		maxNum = 0
		findMax([]byte(n), []byte{}, 0, 0, m)
		fmt.Println(maxNum)
	}
}
