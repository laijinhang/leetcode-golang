package main

import "fmt"

type MyCalendarTwo struct {
	val []struct {
		s, e, num int
	}
}


func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}


func (this *MyCalendarTwo) Book(start int, end int) bool {
	flag := true
	itIndex := []int{}
	for i := 0;i < len(this.val);i++ {
		if (this.val[i].s <= start && start < this.val[i].e) ||
			(this.val[i].s < end && end < this.val[i].e) ||
			(this.val[i].s >= start && end >= this.val[i].e) {
			if this.val[i].num == 2 {
				flag = false
				break
			}
			itIndex = append(itIndex, i)
		}
	}
	if flag {
		// 提取出公共部分
		for i := 0;i < len(itIndex);i++ {
			// 画图分析一共可以得出7种情况
			if start < this.val[itIndex[i]].s && end > this.val[itIndex[i]].s && end < this.val[itIndex[i]].e {	// 第一种
				this.val = append(this.val, struct{ s, e, num int }{
					s: end, e: this.val[itIndex[i]].e, num: 1})
				this.val[itIndex[i]].e = end
				this.val[itIndex[i]].num++
				this.val = append(this.val, struct{ s, e, num int }{
					s: start, e: this.val[itIndex[i]].s, num: 1})
			} else if start < this.val[itIndex[i]].s && this.val[itIndex[i]].e == end { 	// 第二种
				this.val[itIndex[i]].num++
				this.val = append(this.val, struct{ s, e, num int }{
					s: start, e: this.val[itIndex[i]].s, num: 1})
			} else if start < this.val[itIndex[i]].s && end > this.val[itIndex[i]].e {	// 第三种
				this.val[itIndex[i]].num++
				this.val = append(this.val, struct{ s, e, num int }{
					s: start, e: this.val[itIndex[i]].s, num: 1})
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].e, e: end, num: 1})
			} else if start == this.val[itIndex[i]].s && end < this.val[itIndex[i]].e {		// 第四种
				this.val = append(this.val, struct{ s, e, num int }{
					s: start, e: this.val[itIndex[i]].e, num: 1})
				this.val[itIndex[i]].e = end
				this.val[itIndex[i]].num++
			} else if start == this.val[itIndex[i]].s && end == this.val[itIndex[i]].e {	// 第五种
				this.val[itIndex[i]].num++
			} else if start == this.val[itIndex[i]].s && end > this.val[itIndex[i]].e {		// 第六种
				this.val[itIndex[i]].num++
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].e, e: end, num: 1})
			} else if start > this.val[itIndex[i]].s && end < this.val[itIndex[i]].e {		// 第七种
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].s, e: start, num: 1})
				this.val = append(this.val, struct{ s, e, num int }{
					s: end, e: this.val[itIndex[i]].e, num: 1})
				this.val[itIndex[i]].s = start
				this.val[itIndex[i]].e = end
				this.val[itIndex[i]].num++
			} else if start > this.val[itIndex[i]].s && start < this.val[itIndex[i]].e && end > this.val[itIndex[i]].e {
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].s, e: start, num: 1})
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].e, e: end, num: 1})
				this.val[itIndex[i]].s = start
				this.val[itIndex[i]].num++
			} else if start > this.val[itIndex[i]].s && end == this.val[itIndex[i]].e {
				this.val = append(this.val, struct{ s, e, num int }{
					s: this.val[itIndex[i]].s, e: start, num: 1})
				this.val[itIndex[i]].s = start
				this.val[itIndex[i]].num++
			}
		}
		if len(itIndex) == 0 {
			this.val = append(this.val, struct{ s, e, num int }{
				s: start, e: end, num: 1})
		}
	}
	return flag
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */