/*
题目描述
有家动物收容所只收留猫和狗，但有特殊的收养规则，收养人有两种收养方式，第一种为直接收养所有动物中最早进入收容所的，第二种为选择收养的动物类型（猫或狗），并收养该种动物中最早进入收容所的。

给定一个操作序列int[][2] ope(C++中为vector<vector<int>>)代表所有事件。若第一个元素为1，则代表有动物进入收容所，第二个元素为动物的编号，正数代表狗，负数代表猫；若第一个元素为2，则代表有人收养动物，第二个元素若为0，则采取第一种收养方式，若为1，则指定收养狗，若为-1则指定收养猫。请按顺序返回收养的序列。若出现不合法的操作，即没有可以符合领养要求的动物，则将这次领养操作忽略。

测试样例：
[[1,1],[1,-1],[2,0],[2,-1]]
返回：[1,-1]

分析：
根据题意，可以得出以下信息
1）操作系列数据我们可以用 [][2]int  类型存储
2）取出最早到的，说明用队列
3）操作系列
  第一个元素值为1;第二个元素 为正数，取出狗;第二个元素 为负数，取出猫
  第二个元素值为2;第二个元素 为0，取出最早到的;第二个元素 为1，取出狗;第二个元素 为2，取出猫
  其他情况忽略
4）狗为正数，猫为负数

根据上面的分析，我们需要创建两个队列，分别是收养队列，领养队列


然后就开始编写代码：
 */
package main

import "fmt"

type CatDogAsylum struct {}

func (CatDogAsylum) asylum(ope [][2]int) []int {
	animal := make([]int, 0)
	list := make([]int, 0)

	for i := 0;i < len(ope);i++ {
		if ope[i][0] == 1 {
			if ope[i][1] != 0 {
				animal = append(animal, ope[i][1])
			}
		} else if ope[i][0] == 2 {
			if len(animal) == 0 {
				continue
			}
			if ope[i][1] == 0 {
				temp := animal[0]
				list = append(list, temp)
				animal = animal[1:]
			} else if ope[i][1] == 1 {
				i := 0
				for ;i < len(animal) && animal[i] < 0;i-- {}
				if i >= len(animal) || animal[i] == 0 {
					continue
				}
				temp := animal[i]
				list = append(list, temp)
				animal = append(animal[:i], animal[i+1:]...)
			} else if ope[i][1] == -1 {
				i := 0
				for ;i < len(animal) && animal[i] > 0;i-- {}
				if i >= len(animal) || animal[i] == 0 {
					continue
				}
				temp := animal[i]
				list = append(list, temp)
				animal = append(animal[:i], animal[i+1:]...)
			}
		}
	}
	return list
}

func main() {
	list := CatDogAsylum{}.asylum([][2]int{{1,1},{1,-1},{2,0},{2,-1},{1,1},{2,0},{2,0}})
	fmt.Print(list)
}
