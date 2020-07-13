＃ 1、解法一
package main

import (
	"math"
	"fmt"
)

type MinStack struct {
	v []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min:math.MaxInt32}
}

func (this *MinStack) Push(x int)  {
	this.v = append(this.v, x)
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop()  {
	if this.min == this.v[len(this.v)-1] {

		this.min = math.MaxInt32
		if len(this.v) > 1 {
			this.min = this.v[0]
			fmt.Println(this.min, this.v[0])
			for i := 1; i < len(this.v); i++ {
				if this.min > this.v[i] {
					this.min = this.v[i]
				}
			}
		} else {
			this.min = math.MaxInt32
		}
	}
	this.v = this.v[:len(this.v)-1]
}

func (this *MinStack) Top() int {
	return this.v[len(this.v)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
＃ 2、最优解
使用一个辅助栈：push、pop、GetMin都可以以时间复杂度取到
算法过程：

入栈

如果栈为空，入栈，辅助栈也入该值
如果栈不为空，入栈，辅助栈入 min(该值,辅助栈栈顶值)
出栈
栈和辅助栈都出值
