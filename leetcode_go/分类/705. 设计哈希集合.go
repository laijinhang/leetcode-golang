package main

type MyHashSet struct {
	v []bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	hs := MyHashSet{}
	hs.v = make([]bool, 1000001)
	return hs
}


func (this *MyHashSet) Add(key int)  {
	this.v[key] = true
}


func (this *MyHashSet) Remove(key int)  {
	this.v[key] = false
}


/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.v[key]
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */