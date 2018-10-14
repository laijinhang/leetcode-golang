package main

import "fmt"

/*
素数求和问题
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
现在给你N个数（0<N<1000），现在要求你写出一个程序，找出这N个数中的所有素数，并求和。
输入
第一行给出整数M(0<M<10)代表多少组测试数据
每组测试数据第一行给你N，代表该组测试数据的数量。
接下来的N个数为要测试的数据，每个数小于1000
输出
每组测试数据结果占一行，输出给出的测试数据的所有素数和
样例输入
3
5
1 2 3 4 5
8
11 12 13 14 15 16 17 18
10
21 22 23 24 25 26 27 28 29 30
样例输出
10
41
52
 */
var prime [1001]bool

func creatPrimeTable() {	// false中存放的是素数
	prime[1] = true
	for i, j := 2, 0;i * i <= 1000;i++ {
		if !prime[i] {
			for j = i * i;j <= 1000;j += i {
				prime[j] = true
			}
		}
	}
}

func main() {
	creatPrimeTable()
	for i := 1;i < 1001;i++ {
		if !prime[i] {
			fmt.Println(i)
		}
	}
	var N, M int
	var n, sum int
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&M)
		sum = 0
		for ;M > 0;M-- {
			fmt.Scan(&n)
			if !prime[n] {
				fmt.Println(n)
				sum += n
			}
		}
		fmt.Println(sum)
	}
}