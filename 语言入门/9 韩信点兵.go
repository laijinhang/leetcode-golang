package main

import "fmt"

/*
韩信点兵
时间限制：3000 ms  |  内存限制：65535 KB
难度：1
描述
相传韩信才智过人，从不直接清点自己军队的人数，只要让士兵先后以三人一排、五人一排、七人一排地变换队形，而他每次只掠一眼队伍的排尾就知道总人数了。输入3个非负整数a,b,c ，表示每种队形排尾的人数（a<3,b<5,c<7），输出总人数的最小值（或报告无解）。已知总人数不小于10，不超过100 。
输入
输入3个非负整数a,b,c ，表示每种队形排尾的人数（a<3,b<5,c<7）。例如,输入：2 4 5
输出
输出总人数的最小值（或报告无解，即输出No answer）。实例，输出：89
样例输入
2 1 6
样例输出
41
 */
 func main() {
	var tool, a, b, c int
	fmt.Scan(&a, &b, &c)
	i := 0
	if c + 7 >= 10 {
		i = 1
	} else {
		i = 2
	}

	for ;tool < 100;i++ {
		tool = i * 7 + c
		if tool % 3 == a && tool % 5 == b {
			break
		}
	}
	if tool < 100 {
		fmt.Println(tool)
	} else {
		fmt.Println("No answer")
	}
 }