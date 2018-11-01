package main

import (
	"math/rand"
	"fmt"
)

type RandomizedSet struct {
	valm map[int]int
	vals []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{valm:make(map[int]int)}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.valm[val]
	if !ok {
		this.vals = append(this.vals, val)
		this.valm[val] = len(this.vals)-1
	}
	return !ok
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.valm[val]
	if ok {
		this.vals[this.valm[val]], this.vals[len(this.vals)-1] = this.vals[len(this.vals)-1], this.vals[this.valm[val]]
		this.valm[this.vals[this.valm[val]]] = this.valm[val]
		delete(this.valm, this.vals[len(this.vals)-1])
		this.vals = this.vals[:len(this.vals)-1]
	}
	return ok
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main()  {
	obj := Constructor();
	fmt.Println(obj.Insert(3))
	fmt.Println(obj.Remove(3))
	fmt.Println(obj.Remove(3))
	fmt.Println(obj.Insert(0))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(0))
}