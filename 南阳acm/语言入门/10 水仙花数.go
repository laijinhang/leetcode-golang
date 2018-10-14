package main

import "fmt"

/*
水仙花数
时间限制：1000 ms  |  内存限制：65535 KB
难度：0
描述
请判断一个数是不是水仙花数。
其中水仙花数定义各个位数立方和等于它本身的三位数。
输入
有多组测试数据，每组测试数据以包含一个整数n(100<=n<1000)
输入0表示程序输入结束。
输出
如果n是水仙花数就输出Yes
否则输出No
样例输入
153
154
0
样例输出
Yes
No
 */
func main()  {
	var n int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n == 153 || n == 370 || n == 371 || n == 407 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}