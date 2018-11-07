package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 计算链表长度
	lLen := 0
	for p := root;p != nil;p = p.Next {
		lLen++
	}
	// 分割
	// 长度多1的链表的个数
	l1 := lLen % k
	resL := []*ListNode{}
	tempL := &ListNode{}
	p := root
	for i := 0;i < l1;i++ {
		resL = append(resL, p)
		for i := 1;i < lLen / k + 1;i++ {
			p = p.Next
		}
		tempL = p
		p = p.Next
		tempL.Next = nil
	}
	// 长度少1的链表
	l2 := k - l1
	for i := 0;i < l2;i++ {
		resL = append(resL, p)
		if p != nil {
			for i := 1;i < lLen / k && p != nil;i++ {
				p = p.Next
			}
			if p != nil {
				tempL = p
				p = p.Next
				tempL.Next = nil
			}
		}
	}
	return resL
}