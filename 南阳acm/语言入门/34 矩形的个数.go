/*
矩形的个数
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
在一个3*2的矩形中，可以找到6个1*1的矩形，4个2*1的矩形3个1*2的矩形，2个2*2的矩形，2个3*1的矩形和1个3*2的矩形，总共18个矩形。

给出A，B,计算可以从中找到多少个矩形。
输入
本题有多组输入数据（<10000），你必须处理到EOF为止

输入2个整数A,B(1<=A,B<=1000)

输出
输出找到的矩形数。
样例输入
1 2
3 2
样例输出
3
18
 */
package main

import "fmt"

func main() {
	var A, B, count int
	for {
		_, err := fmt.Scan(&A, &B)
		if err != nil {
			break
		}
		count = 0
		for i := 1;i <= A;i++ {
			for j := 1;j <= B;j++ {				// 组成坐标
				for k := 0;i + k <= A;k++ {		// 纵向走
					for f := 0;j + f <= B;f++ {	// 横向走
						count++
					}
				}
			}
		}
		fmt.Println(count)
	}
}