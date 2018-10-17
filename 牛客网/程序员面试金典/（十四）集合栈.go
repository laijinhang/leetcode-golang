/*
题目描述
请实现一种数据结构SetOfStacks，由多个栈组成，其中每个栈的大小为size，当前一个栈填满时，新建一个栈。该数据结构应支持与普通栈相同的push和pop操作。

给定一个操作序列int[][2] ope(C++为vector&ltvector&ltint>>)，每个操作的第一个数代表操作类型，若为1，则为push操作，后一个数为应push的数字；若为2，则为pop操作，后一个数无意义。请返回一个int[][](C++为vector&ltvector&ltint>>)，为完成所有操作后的SetOfStacks，顺序应为从下到上，默认初始的SetOfStacks为空。保证数据合法。

 */
package main

import "fmt"

type SetOfStatck struct {}

func (SetOfStatck) setOfStatck(ope [][2]int, size int) [][]int {
	al := make([][]int, 0)
	st := make([]int, 0)

	k := 0
	len1 := len(ope)
	for i := 0;i < len1;i++ {
		if ope[i][0] == 1 {
			st = append(st, ope[i][1])
			k++
		} else {
			if k > 0 {
				k--
				st = st[:k]
			} else {
				if len(al) > 0 {
					k = size - 1
					st = al[len(al)-1][:k]
					al = al[:len(al)-1]
				} else {
					k = 0
				}
			}
		}
		if k == size {
			al = append(al, st)
			st = make([]int, 0)
			k = 0
		}
	}
	if k != 0 {
		al = append(al, st)
	}
	return al
}

func main() {
	//ope := [][2]int{{1, 2}, {2, 2}, {1, 3}, {1, 4}, {1, 2}, {2, 2}, {2, 3}, {1, 4}, {1, 12}}
	ope := [][2]int{{2, 2}, {2, 2}, {2, 3}, {1, 4}, {1, 2}, {2, 2}, {2, 3}, {1, 4}, {1, 12}}
	al := SetOfStatck{}.setOfStatck(ope, 2)
	fmt.Println(al)
}
