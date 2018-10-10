/*
两点距离
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
输入两点坐标（X1,Y1）,（X2,Y2）(0<=x1,x2,y1,y2<=1000),计算并输出两点间的距离。
输入
第一行输入一个整数n（0<n<=1000）,表示有n组测试数据;
随后每组占一行，由4个实数组成，分别表示x1,y1,x2,y2,数据之间用空格隔开。
输出
对于每组输入数据，输出一行，结果保留两位小数。
样例输入
2
0 0 0 1
0 1 1 0
样例输出
1.00
1.41
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var x1, x2, y1, y2 float64
	fmt.Scan(&n)
	for ;n > 0;n-- {
		fmt.Scan(&x1, &y1, &x2, &y2)
		distance := math.Sqrt((x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2))
		fmt.Printf("%.2f\n", distance)
	}
}