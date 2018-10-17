/*
题目描述
输入一颗二叉树的跟节点和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。路径定义为从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。(注意: 在返回值的list中，数组长度大的数组靠前)

分析：其实就是树的深度遍历

代码实现：
 */
 package main

import "fmt"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

var goThrough map[*TreeNode]bool = make(map[*TreeNode]bool)

type TreeList []*TreeNode

func FindPath(root *TreeNode, expectNumber int) [][]*TreeNode {
	nl := 0
	lt := [][]*TreeNode{}
	list := []*TreeNode{root}
	goThrough[root] = true
	nl += root.val
	for len(list) != 0 {
		l := list[len(list)-1]
		if l.left != nil && !goThrough[l.left] {	// 遍历左子树
			goThrough[l.left] = true
			if nl + l.left.val == expectNumber {
				lt = append(lt, append(list, l.left))	// 路径保存
			} else if nl + l.left.val < expectNumber {
				list = append(list, l.left)
				nl += l.left.val
				l = l.left
			}
		} else if l.right != nil && !goThrough[l.right] {	// 遍历右子树
			goThrough[l.right] = true
			if l.right.val + nl == expectNumber {
				lt = append(lt, append(list, l.right))	// 路径保存
			} else if l.right.val + nl < expectNumber {
				list = append(list, l.right)
				nl += l.right.val
				l = l.right
			}
		} else {	// 左右子树都遍历了，往上走一步
			nl -= l.val
			list = list[:len(list)-1]
		}
	}
	return lt
}

func main() {
	var root TreeNode = TreeNode{val:1}
	t := &root
	t.left = &TreeNode{val:2}
	t.right = &TreeNode{val:3}
	t.left.left = &TreeNode{val:3}
	t.left.right = &TreeNode{val:2}
	t.right.left = &TreeNode{val:1}
	t.right.right = &TreeNode{val:1}
	t.left.left.left = &TreeNode{val:1}
	t.left.left.right = &TreeNode{val:2}
	/*
				1
			2		3
	     3    2	  1   1
	   1  2
	 */
	// 测试
	lists := FindPath(t, 5)
	for _, list := range lists {
		for _, l := range list {
			fmt.Printf("%d ", l.val)
		}
		fmt.Println()
	}
}