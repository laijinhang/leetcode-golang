/*
题目描述
输入一个链表，输出该链表中倒数第k个结点。

 */
package main

import (
	"container/list"
	"fmt"
)

func FindKthToTail(l *list.List, k uint32) *list.Element {
	h := l.Front()
	k = uint32(l.Len()) - k
	if k < 0 {return nil}
	for i := uint32(0);i < k;i++ {
		h = h.Next()
	}
	return h
}

func main() {
	l := list.List{}
	for i := 0;i < 10;i++ {
		l.PushBack(i)
	}
	f := FindKthToTail(&l, 9)
	fmt.Println(f.Value)
}
