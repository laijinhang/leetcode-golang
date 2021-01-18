[面试题 02.02. 返回倒数第 k 个节点](https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/)
```golang
思路1: 双指针,时间复杂度O(n),空间复杂度O(1)
思路2: 使用数组,时间复杂度O(n),空间复杂度O(n)

时间复杂度:O(n)
空间复杂度:O(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
    p1, p2 := head, head
    for i := 1;i < k;i++ {
        p2 = p2.Next
    }
    for ;p2.Next != nil;p2 = p2.Next {
        p1 = p1.Next
    }
    return p1.Val
}
```
