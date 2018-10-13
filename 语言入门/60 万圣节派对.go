/*
万圣节派对
时间限制：1000 ms  |  内存限制：65535 KB
难度：1
描述
万圣节有一个Party，XadillaX显然也要去凑热闹了。因为去凑热闹的人数非常庞大，几十W的数量级吧，自然要进场就需要有门票了。很幸运的，XadillaX竟然拿到了一张真·门票！这真·门票的排列规则有些奇怪：

门票号是由0~6组成的六位数（0~6这几个数字可重用）

每一个门票号的每一位不能有三个连续相同的数字（如123335是不行的）

每一个门票号相邻的两位相差必须在四以下（≤4）（如016245是不行的）

输入
第一行一个n，代表输入个数
接下去n行，每行两个数字x,y(x <= y)
输出
对于每个测试，输出x到y之间的门票编号。每个测试结尾之间空行。
样例输入
2
001001 001002
001011 001012
样例输出
001001
001002

001011
001012
 */
package main

import "fmt"

func main() {
	var n int
	var usable bool
	var x, y string
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&x, &y)
		for a := x[0];a <= y[0];a++ {
			for b := x[1];b <= y[1];b++ {
				for c := x[2];c <= y[2];c++ {
					for d := x[3];d <= y[3];d++ {
						for e := x[4];e <= y[4];e++ {
							for f := x[5];f <= y[5];f++ {
								// 规则一
								usable = true
								if a == b && b == c {
									usable = false
									continue
								}
								if b == c && c == d {
									usable = false
									continue
								}
								if c == d && d == e {
									usable = false
									continue
								}
								if d == e && e == f {
									usable = false
									continue
								}
								// 规则2
								if int(a) - int(b) > 4 || int(b) - int(a) > 4 {
									usable = false
									continue
								}
								if int(b) - int(c) > 4 || int(c) - int(b) > 4 {
									usable = false
									continue
								}
								if int(c) - int(d) > 4 || int(d) - int(c) > 4 {
									usable = false
									continue
								}
								if int(d) - int(e) > 4 || int(e) - int(d) > 4 {
									usable = false
									continue
								}
								if int(e) - int(f) > 4 || int(f) - int(e) > 4 {
									usable = false
									continue
								}
								if usable {
									fmt.Println(string(a) + string(b) + string(c) + string(d) + string(e) + string(f))
								}
							}
						}
					}
				}
			}
		}
	}
}
