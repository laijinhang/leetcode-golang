/*
画图
时间限制：1000 ms  |  内存限制：65535 KB
难度：0
描述
计算机画图也挺有趣的哈！那我们就来用计算机画幅图吧。。。
输入
输入一个正整数N（0<N<=10），表示要输出的正方形的边上*的个数
输出
输出一个满足题意的正方形
样例输入
4
样例输出
****
****
****
****
 */
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		for j := 0;j < n;j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
