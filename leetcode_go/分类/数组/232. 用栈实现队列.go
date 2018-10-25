package main

import "fmt"

type MyQueue struct {
	s []int
	t []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.s = append(this.s, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.s) != 0 {
		this.t = append(this.t, this.s[len(this.s)-1])
		this.s = this.s[:len(this.s)-1]
	}
	res := this.t[len(this.t)-1]
	this.t = this.t[:len(this.t)-1]
	for len(this.t) == 0 {
		this.s = append(this.s, this.t[len(this.t)-1])
		this.t = this.t[:len(this.t)-1]
	}
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	for len(this.s) != 0 {
		this.t = append(this.t, this.s[len(this.s)-1])
		this.s = this.s[:len(this.s)-1]
	}
	res := this.t[len(this.t)-1]
	for len(this.t) != 0 {
		this.s = append(this.s, this.t[len(this.t)-1])
		this.t = this.t[:len(this.t)-1]
	}
	return res
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	fmt.Println(len(this.s))
	if len(this.s) != 0 {
		return false
	}
	return true
}

