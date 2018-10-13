/*
A+B Problem III
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
求A+B是否与C相等。
输入
T组测试数据。
每组数据中有三个实数A,B,C(-10000.0<=A,B<=10000.0,-20000.0<=C<=20000.0)
数据保证小数点后不超过4位。

输出
如果相等则输出Yes
不相等则输出No
样例输入
3
-11.1 +11.1 0
11 -11.25 -0.25
1 2 +4
样例输出
Yes
Yes
No
 */
package main

import "fmt"

func main() {
	var n int
	var a, b, c float64
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&a, &b, &c)
		if int(a * 10000) + int(b * 10000) == int(c * 10000) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}