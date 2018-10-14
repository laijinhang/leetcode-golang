/*
喷水装置（一）
时间限制：3000 ms  |  内存限制：65535 KB
难度：3
描述
现有一块草坪，长为20米，宽为2米，要在横中心线上放置半径为Ri的喷水装置，每个喷水装置的效果都会让以它为中心的半径为实数Ri(0<Ri<15)的圆被湿润，这有充足的喷水装置i（1<i<600)个，并且一定能把草坪全部湿润，你要做的是：选择尽量少的喷水装置，把整个草坪的全部湿润。
输入
第一行m表示有m组测试数据
每一组测试数据的第一行有一个整数数n，n表示共有n个喷水装置，随后的一行，有n个实数ri，ri表示该喷水装置能覆盖的圆的半径。
输出
输出所用装置的个数
样例输入
2
5
2 3.2 4 4.5 6
10
1 2 3 1 2 1.2 3 1.1 1 2
样例输出
2
5

假设 半径大于1的圆所组成的能够全部覆盖到
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

type NewRi []float64

func main() {
	var m, n int
	var tempRi float64
	length := 20.0
	fmt.Scan(&m)
	for i := 0;i < m;i++ {
		fmt.Scan(&n)
		ri1 := []float64{}
		ri2 := []float64{}
		for j := 0;j < n;j++ {
			fmt.Scan(&tempRi)
			if tempRi > 1 {
				ri1 = append(ri1, 2 * math.Sqrt(tempRi*tempRi-1))
			} else {
				ri2 = append(ri2, tempRi)
			}
		}
		// 从小到大排序
		sort.Float64s(ri1)
		// 反转
		for i, j := 0, len(ri1) - 1;i < j;i, j = i + 1, j - 1 {
			ri1[i], ri1[j] = ri1[j], ri1[i]
		}
		//sort.Float64s(ri2)
		//for i, j := 0, len(ri2) - 1;i < j;i, j = i + 1, j - 1 {
		//	ri2[i], ri2[j] = ri2[j], ri2[i]
		//}
		num := 0
		for tempL, i := length, 0;tempL > 0 && i < len(ri1); i++ {
			tempL -= ri1[i]
			num++
		}
		fmt.Println(num)
	}
}

func (r NewRi) Len() int {
	return len(r)
}

func (r NewRi) Less(i, j int) {

}