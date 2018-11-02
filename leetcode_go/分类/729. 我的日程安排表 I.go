package main

import "math"

type MyCalendar struct {
	val []struct{
		s, e int
	}
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	flag := true
	for i := 0;i < len(this.val);i++ {
		if (this.val[i].s <= start && start < this.val[i].e) ||
			(this.val[i].s < end && end < this.val[i].e) ||
			(this.val[i].s >= start && end >= this.val[i].e) {
			flag = false
			break
		}
	}
	if flag {
		this.val = append(this.val, struct{ s, e int }{s: start, e: end})
	}
	return flag
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */