[1325. 删除给定值的叶子节点](https://leetcode-cn.com/problems/delete-leaves-with-a-given-value/)
```goalng
二叉树后序遍历

时间复杂度：O(n)，n为节点数
空间复杂度：O(1)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Left != nil {
        root.Left = removeLeafNodes(root.Left, target)
    }
    if root.Right != nil {
        root.Right = removeLeafNodes(root.Right, target)
    }
    if root.Left == nil && root.Right == nil && root.Val == target {
        return nil
    }
    return root
}
```
