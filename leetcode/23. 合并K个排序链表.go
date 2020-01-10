package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// 两两合并
	ls := append([]*ListNode{}, lists...)
	head := &ListNode{}
	p := head
	for {
		i := 0
		for ;i < len(ls) && ls[i] == nil;i++{}
		if i == len(ls) {
			break
		}
		minV := ls[i].Val
		minI := i
		for ;i < len(ls);i++ {
			if ls[i] != nil && ls[i].Val < minV {
				minV = ls[i].Val
				minI = i
			}
		}
		p.Next = ls[minI]
		ls[minI] = ls[minI].Next
		p = p.Next
	}
	return head.Next
}