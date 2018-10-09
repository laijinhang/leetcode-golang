/*
1的个数
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
小南刚学了二进制，他想知道一个数的二进制表示中有多少个1，你能帮他写一个程序来完成这个任务吗？
输入
第一行输入一个整数N，表示测试数据的组数(1<N<1000)
每组测试数据只有一行，是一个整数M(0=<M<=10000)
输出
每组测试输出占一行，输出M的二进制表示中1的个数
样例输入
3
4
6
7
样例输出
1
2
3
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&M)
		b := strconv.FormatInt(int64(M), 2)
		fmt.Println(strings.Count(b, "1"))
	}
}
