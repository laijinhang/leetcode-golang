package main

type RLEIterator struct {
	vals []int
}


func Constructor(A []int) RLEIterator {
	var rlei RLEIterator
	rlei.vals = A
	return rlei
}


func (this *RLEIterator) Next(n int) int {
	res := -1
	for i := 0;i < len(this.vals);i += 2 {
		if this.vals[i] >= n {
			res = this.vals[i+1]
			this.vals[i] -= n
			n = 0
			break
		} else {
			n -= this.vals[i]
			this.vals[i] = 0
		}
	}
	return res
}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */