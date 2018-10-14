/*
素数距离问题
时间限制：3000 ms  |  内存限制：65535 KB
难度：2
描述
现在给出你一些数，要求你写出一个程序，输出这些整数相邻最近的素数，并输出其相距长度。如果左右有等距离长度素数，则输出左侧的值及相应距离。
如果输入的整数本身就是素数，则输出该素数本身，距离输出0
输入
第一行给出测试数据组数N(0<N<=10000)
接下来的N行每行有一个整数M(0<M<1000000)，
输出
每行输出两个整数 A B.
其中A表示离相应测试数据最近的素数，B表示其间的距离。
样例输入
3
6
8
10
样例输出
5 1
7 1
11 1
 */
package main

import "fmt"

var prime [1000000]bool

func createPrimeTable() {
	prime[1] = true
	for i := 2;i * i <= 1000000;i++ {
		if !prime[i] {
			for j := i + i;j < 1000000;j += i {
				prime[j] = true
			}
		}
	}
}

func main() {
	createPrimeTable()
	var N, M, leftP, rightP int
	fmt.Scan(&N)
	for ;N > 0;N-- {
		fmt.Scan(&M)
		if !prime[M] {
			fmt.Println(M, 0)
		} else {
			for leftP = M - 1;leftP >= 0;leftP-- {
				if !prime[leftP] {
					break
				}
			}
			for rightP = M + 1;rightP < 1000000;rightP++ {
				if !prime[rightP] {
					break
				}
			}
			if leftP >= 0 && rightP < 1000000 {
				if M - leftP <= rightP - M {
					fmt.Println(leftP, M - leftP)
				} else {
					fmt.Println(rightP, rightP - M)
				}
			} else if leftP < 0 {
				fmt.Println(rightP, M - rightP)
			} else {
				fmt.Println(leftP, M - leftP)
			}
		}
	}
}
