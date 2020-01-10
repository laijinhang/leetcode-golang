package main
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for p := head;p != nil;p = p.Next {
		arr = append(arr, p.Val)
	}
	for i, j := 0, len(arr) - 1;i <= j;i, j = i + 1, j - 1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}