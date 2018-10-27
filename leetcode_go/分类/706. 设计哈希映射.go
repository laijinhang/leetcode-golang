package main

type MyHashMap struct {
	v []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	var hmap MyHashMap
	hmap.v = make([]int, 1000001)
	for i := 0; i < 1000001; i++ {
		hmap.v[i] = -1
	}
	return hmap
}

/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int) {
	this.v[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.v[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.v[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
