/*
变态最大值
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
Yougth讲课的时候考察了一下求三个数最大值这个问题，没想到大家掌握的这么烂，幸好在他的帮助下大家算是解决了这个问题，但是问题又来了。

他想在一组数中找一个数，这个数可以不是这组数中的最大的，但是要是相对比较大的，但是满足这个条件的数太多了，怎么办呢？他想到了一个办法，把这一组数从开始把每相邻三个数分成一组（组数是从1开始），奇数组的求最大值，偶数组的求最小值，然后找出这些值中的最大值。

输入
有多组测试数据，以文件结束符为标志。
每组测试数据首先一个N,是数组中数的个数。（0<N<10000，为降低题目难度，N是3的倍数）
然后是数组中的这些数。
输出
输出包括一行，就是其中的最大值。
样例输入
3
4 5 6
6
1 2 3 7 9 5
样例输出
6
5
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, max, min, maxL, a, b, c, i int
	var err error
	for {
		_, err = fmt.Scan(&n)
		if err != nil {
			break
		}
		max = int(math.MinInt32)
		min = int(math.MaxInt32)
		n /= 3
		for i = 1;i <= n;i++ {
			fmt.Scan(&a, &b, &c)
			if i % 2 == 0 {	// 偶数求最小
				if a <= b && a <= c {
					min = a
				} else if b <= a && b <= c {
					min = b
				} else if c <= a && c <= b {
					min = c
				}
			} else { // 奇数求最大
				if a >= b && a >= c {
					max = a
				} else if b >= a && b >= c {
					max = b
				} else if c >= a && c >= b {
					max = c
				}
			}
		}
		if n == 1 {
			maxL = max
		} else {
			if max >= min {
				maxL = max
			} else {
				maxL = min
			}
		}
		fmt.Println(maxL)
	}
}