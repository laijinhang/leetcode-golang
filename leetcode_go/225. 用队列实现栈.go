package main

type MyStack struct {
	v []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.v = append(this.v, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.v[len(this.v)-1]
	this.v = this.v[:len(this.v)-1]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	res := this.v[len(this.v)-1]
	return res
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.v) != 0 {
		return false
	}
	return true
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */