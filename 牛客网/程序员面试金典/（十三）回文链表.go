/*
题目描述
请编写一个函数，检查链表是否为回文。

给定一个链表ListNode* pHead，请返回一个bool，代表链表是否为回文。

测试样例：
{1,2,3,2,1}
返回：true
{1,2,3,2,3}
返回：false

分析：
思路一：
可以先遍历一遍链表，使用数组来存放，然后从首和末尾开始在数组中遍历，判断是否为回文

思路二：栈思想
可以先用两个指针，第一个指针每次移动2个，第二个指针每次移动1个，当第一个指针移到末尾时，第二个指针就到中间。
有以下几种情况：
1）链表为空，或长度为1，返回true
2）链表长度为奇数个，这种情况中间值不需要加入到栈中
3）链表长度为偶数个，这种情况中间值要加入到栈中
然后从中间向两边比较，看是否为回文
 */
package main

import "fmt"

type ListNode struct {
	val int
	next *ListNode
}

func isPalindrome(pHead *ListNode) bool {
	if pHead == nil || pHead.next == nil {
		return true
	}
	stack := make([]int, 0, 1024)
	fast, show := pHead, pHead

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		stack = append(stack, show.val)
		show = show.next
	}
	if fast != nil {	// 链表长度为奇数情况
		show = show.next
	}
	stackLen := len(stack) - 1
	fmt.Println(stack)
	for show != nil {
		if stack[stackLen] != show.val {
			break
		}
		stackLen--
		show = show.next
	}
	if stackLen != -1 {
		return false
	}

	return true
}

func main() {
	pHead := testData()
	fmt.Println(isPalindrome(pHead))
}

func testData() *ListNode {
	var pHead ListNode

	ln := &pHead
	ln.val = 0
	for i := 1;i <= 5;i++ {
		ln.next = &ListNode{val:i}
		ln = ln.next
	}
	for i := 5;i >= 0;i-- {
		ln.next = &ListNode{val:i}
		ln = ln.next
	}
	ln = &pHead
	for ln != nil {
		fmt.Print(ln.val, " ")
		ln = ln.next
	}
	fmt.Println()
	return &pHead
}