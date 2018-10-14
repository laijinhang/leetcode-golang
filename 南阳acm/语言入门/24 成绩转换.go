/*
成绩转换
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
输入一个百分制的成绩M，将其转换成对应的等级，具体转换规则如下：
90~100为A;
80~89为B;
70~79为C;
60~69为D;
0~59为E;
输入
第一行是一个整数N，表示测试数据的组数(N<10)
每组测试数据占一行，由一个整数M组成(0<=M<=100)。
输出
对于每组输入数据，输出一行。
样例输入
2
97
80
样例输出
A
B
 */
package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&M)
		if M >= 90 && M <= 100 {
			fmt.Println("A")
		} else if M >= 80 && M <= 89 {
			fmt.Println("B")
		} else if M >= 70 && M <= 79 {
			fmt.Println("C")
		} else if M >= 60 && M <= 69 {
			fmt.Println("D")
		} else if M >= 0 && M <= 59 {
			fmt.Println("E")
		}
	}
}
