/*
题目描述
用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
 */
package main

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"fmt"
)

type Solution struct {
	stack1 []int
	stack2 []int
}

func (s *Solution) push(node int) {
	s.stack1 = append(s.stack1, node)
}

func (s *Solution) pop() (int, error) {
	if len(s.stack1) == 0 && len(s.stack2) == 0 {
		return 0, errors.New("Queue is empty!")
	} else if len(s.stack2) == 0 {
		for len(s.stack1) != 0 {
			s.stack2 = append(s.stack2, s.stack1[len(s.stack1)-1])
			s.stack1 = s.stack1[:len(s.stack1)-1]
		}
	}
	res := s.stack2[len(s.stack2)-1]
	s.stack2 = s.stack2[:len(s.stack2)-1]
	return res, nil
}

func main() {
	q := Solution{}
	q.push(123)
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)

	for val, err := q.pop();err == nil; val, err = q.pop(){
		fmt.Print(val, " ")
	}
	fmt.Println()
}
